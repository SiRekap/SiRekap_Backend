package controllers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sirekap/SiRekap_Backend/forms"
	"sirekap/SiRekap_Backend/models"
	"sirekap/SiRekap_Backend/notifications"
	"strconv"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"

	"github.com/signintech/gopdf"

	"cloud.google.com/go/storage"
)

type FormcImageController struct{}

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

const (
	projectID  = "sirekap-383605"
	bucketName = "staging-sirekap-form"
)

func (f FormcImageController) SendFormcImagePayload(c *gin.Context) {
	var formcImagePayload models.FormcImagePayload

	if err := c.BindJSON(&formcImagePayload); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcImagePayload.SendFormcImagePayload(formcImagePayload)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcImage(c *gin.Context) {
	var formcImage models.FormcImage

	if err := c.BindJSON(&formcImage); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcImage.SendFormcImage()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcImageRaw(c *gin.Context) {
	var formcImageRaw forms.FormcImageRaw

	// Mengambil body JSON dari request frontend
	if err := c.BindJSON(&formcImageRaw); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	// c.JSON(http.StatusOK, "Form C has been sent, please wait for notification from firebase")

	// Penyimpanan url dan informasi gambar
	form, err := models.SendFormcImageRaw(formcImageRaw)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Pengiriman gambar untuk dipindai ke sistem vision
	formcImageVisionResponse, err := SendFormcImageVisionRequest(form)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var wg sync.WaitGroup

	// Penyimpanan hasil pemindaian ke database secara asynchronous
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = SaveVisionResultToDatabase(formcImageVisionResponse)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}()

	// Pengiriman hasil pemindaian ke stream processing secara asynchronous
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = SendFormcResultStreamProcessingRequest(formcImageVisionResponse)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}()

	// Pengiriman hasil pemindaian ke tabel formc_result secara asynchronous
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = SendFormcResultByVisionResponse(formcImageVisionResponse)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}()

	wg.Wait()

	response := GenerateScanResultResponse(formcImageVisionResponse, formcImageRaw)
	c.JSON(http.StatusOK, response)

	// c.JSON(http.StatusOK, "Form C result has been saved to database and delivered to stream processor")
	// SendPushNotification(formcImageVisionResponse, formcImageRaw)
}

func (f FormcImageController) SendFormcStatusData(c *gin.Context) {
	var formcStatusData models.FormcStatusData

	if err := c.BindJSON(&formcStatusData); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcStatusData.SendFormcStatusData()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcStatusImage(c *gin.Context) {
	var formcStatusImage models.FormcStatusImage

	if err := c.BindJSON(&formcStatusImage); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcStatusImage.SendFormcStatusImage()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func SendFormcResultToKafkaTest(c *gin.Context) {
	err := SendFormcResultStreamProcessingRequestTest()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, "Test success")
	}
}

func SendFormcImageVisionRequest(form forms.FormcImageRawResponse) (forms.FormcImageVisionResponse, error) {
	formcImageVisionRequest := forms.FormcImageVisionRequest{
		IdImageList:  form.IdImageList,
		ImageUrlList: form.FormcImageRaw.PayloadList,
		IdPaslonList: form.FormcImageRaw.IdPaslonList,
	}

	requestURL := "http://35.184.140.99/v1/read"
	jsonBody, err := json.Marshal(formcImageVisionRequest)
	if err != nil {
		return forms.FormcImageVisionResponse{}, err
	}

	request, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return forms.FormcImageVisionResponse{}, err
	}

	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return forms.FormcImageVisionResponse{}, err
	}

	defer res.Body.Close()

	resp := &forms.FormcImageVisionResponse{}
	err = json.NewDecoder(res.Body).Decode(resp)
	if err != nil {
		return forms.FormcImageVisionResponse{}, err
	}

	return *resp, nil
}

func SaveVisionResultToDatabase(form forms.FormcImageVisionResponse) error {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		formcAdministrasiHlmSatuProses := models.FormcAdministrasiHlmSatuProses{
			IdImage: form.IdImageList[0],

			PemilihDptL:    form.PemilihDptL,
			PemilihDptP:    form.PemilihDptP,
			PemilihDptJ:    form.PemilihDptJ,
			PemilihDpphL:   form.PemilihDpphL,
			PemilihDpphP:   form.PemilihDpphP,
			PemilihDpphJ:   form.PemilihDpphJ,
			PemilihDptbL:   form.PemilihDptbL,
			PemilihDptbP:   form.PemilihDptbP,
			PemilihDptbJ:   form.PemilihDptbJ,
			PemilihTotalL:  form.PemilihTotalL,
			PemilihTotalP:  form.PemilihTotalP,
			PemilihTotalJ:  form.PemilihTotalJ,
			PenggunaDptL:   form.PenggunaDptL,
			PenggunaDptP:   form.PenggunaDptP,
			PenggunaDptJ:   form.PenggunaDptJ,
			PenggunaDpphL:  form.PenggunaDpphL,
			PenggunaDpphP:  form.PenggunaDpphP,
			PenggunaDpphJ:  form.PenggunaDpphJ,
			PenggunaDptbL:  form.PenggunaDptbL,
			PenggunaDptbP:  form.PenggunaDptbP,
			PenggunaDptbJ:  form.PenggunaDptbJ,
			PenggunaTotalL: form.PenggunaTotalL,
			PenggunaTotalP: form.PenggunaTotalP,
			PenggunaTotalJ: form.PenggunaTotalJ,

			PemilihDisabilitasL:  form.PemilihDisabilitasL,
			PemilihDisabilitasP:  form.PemilihDisabilitasP,
			PemilihDisabilitasJ:  form.PemilihDisabilitasJ,
			PenggunaDisabilitasL: form.PenggunaDisabilitasL,
			PenggunaDisabilitasP: form.PenggunaDisabilitasP,
			PenggunaDisabilitasJ: form.PenggunaDisabilitasJ,

			SuratDiterima:       form.SuratDiterima,
			SuratDikembalikan:   form.SuratDikembalikan,
			SuratTidakDigunakan: form.SuratTidakDigunakan,
			SuratDigunakan:      form.SuratDigunakan,
		}

		_, err := models.SendFormcAdministrasiHlmSatuProses(formcAdministrasiHlmSatuProses)
		if err != nil {
			panic(err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		formcAdministrasiHlmSatuFinal := models.FormcAdministrasiHlmSatuFinal{
			IdImage: form.IdImageList[0],
			IsBenar: true,
		}

		_, err := models.SendFormcAdministrasiHlmSatuFinal(formcAdministrasiHlmSatuFinal)
		if err != nil {
			panic(err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		formcAdministrasiHlmDuaProses := models.FormcAdministrasiHlmDuaProses{
			IdImage:             form.IdImageList[1],
			SuaraSah:            form.SuaraSah,
			SuaraTidakSah:       form.SuaraTidakSah,
			SuaraTotal:          form.SuaraTotal,
			PenggunaHakPilih:    form.PenggunaHakPilih,
			SuratSuaraDigunakan: form.SuratSuaraDigunakan,
		}

		_, err := models.SendFormcAdministrasiHlmDuaProses(formcAdministrasiHlmDuaProses)
		if err != nil {
			panic(err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		formcAdministrasiHlmDuaFinal := models.FormcAdministrasiHlmDuaFinal{
			IdImage: form.IdImageList[1],
			IsBenar: true,
		}

		_, err := models.SendFormcAdministrasiHlmDuaFinal(formcAdministrasiHlmDuaFinal)
		if err != nil {
			panic(err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var formcImageModel models.FormcImage
		formcImage, err := formcImageModel.GetFormcImage(form.IdImageList[0])
		if err != nil {
			panic(err)
		}

		formcImagePayloadModel := models.FormcImagePayload{}

		formcImagePayload, err := formcImagePayloadModel.GetFormcImagePayload(form.IdImageList[0])
		if err != nil {
			panic(err)
		}
		img1Url := formcImagePayload.Payload

		formcImagePayload, err = formcImagePayloadModel.GetFormcImagePayload(form.IdImageList[1])
		if err != nil {
			panic(err)
		}
		img2Url := formcImagePayload.Payload

		formcImagePayload, err = formcImagePayloadModel.GetFormcImagePayload(form.IdImageList[2])
		if err != nil {
			panic(err)
		}
		img3Url := formcImagePayload.Payload

		pdfFileName := "kesesuaian-" + strconv.Itoa(form.IdImageList[0]+form.IdImageList[1]+form.IdImageList[2]) + ".pdf"

		err = GeneratePdfAndSendToBucket(img1Url, img2Url, img3Url, pdfFileName)
		if err != nil {
			panic(err)
		}

		err = DeleteGeneratedLocalFiles(pdfFileName)
		if err != nil {
			panic(err)
		}

		formcImageGroup := models.FormcImageGroup{
			IdTps:          formcImage.IdTps,
			JenisPemilihan: formcImage.JenisPemilihan,
			IdImageHlm1:    form.IdImageList[0],
			IdImageHlm2:    form.IdImageList[1],
			IdImageHlm3:    form.IdImageList[2],
			PdfUrl:         "https://storage.googleapis.com/staging-sirekap-form/pdf/" + pdfFileName,
		}

		_, err = formcImageGroup.SendFormcImageGroup()
		if err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(form.IdPaslonList); i++ {
			suaracProses := models.SuaraCProses{
				IdImage:     form.IdImageList[2],
				IdPaslon:    form.IdPaslonList[i],
				JmlSuaraOcr: form.JmlSuaraOcrList[i],
				JmlSuaraOmr: form.JmlSuaraOmrList[i],
			}

			_, err := models.SendSuaraCProses(suaracProses)
			if err != nil {
				panic(err)
			}

			suaracFinal := models.SuaraCFinal{
				IdImage:  form.IdImageList[2],
				IdPaslon: form.IdPaslonList[i],
				JmlSuara: form.JmlSuaraOcrList[i],
			}

			_, err = models.SendSuaraCFinal(suaracFinal)
			if err != nil {
				panic(err)
			}
		}
	}()

	wg.Wait()

	return nil
}

func GeneratePdfAndSendToBucket(img1url string, img2url string, img3url string, pdfFileName string) error {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	if err := DownloadFile("img1url.jpg", img1url); err != nil {
		return err
	}
	pdf.AddPage()
	pdf.Image("img1url.jpg", 0, 0, &gopdf.Rect{W: 595.28, H: 841.89})

	if err := DownloadFile("img2url.jpg", img2url); err != nil {
		return err
	}
	pdf.AddPage()
	pdf.Image("img2url.jpg", 0, 0, &gopdf.Rect{W: 595.28, H: 841.89})

	if err := DownloadFile("img3url.jpg", img3url); err != nil {
		return err
	}
	pdf.AddPage()
	pdf.Image("img3url.jpg", 0, 0, &gopdf.Rect{W: 595.28, H: 841.89})

	pdf.WritePdf(pdfFileName)

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "sirekap-383605-473b98d7b061.json")

	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader := &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "pdf/",
	}

	file, err := os.Open(pdfFileName)
	if err != nil {
		return err
	}

	err = uploader.UploadFile(file, pdfFileName)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func DeleteGeneratedLocalFiles(pdfFileName string) error {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := os.Remove("img1url.jpg")
		if err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := os.Remove("img2url.jpg")
		if err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := os.Remove("img3url.jpg")
		if err != nil {
			panic(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := os.Remove(pdfFileName)
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()

	return nil
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func (c *ClientUploader) UploadFile(file *os.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func (f FormcImageController) GetFormcImageGroupByIdTpsAndJenisPemilihan(c *gin.Context) {
	params := c.Request.URL.Query()

	idTps := params.Get("id_tps")
	jenisPemilihan := params.Get("jenis_pemilihan")
	if idTps == "" || jenisPemilihan == "" {
		c.String(http.StatusBadRequest, "Id Tps is not provided or Jenis Pemilihan is not provided!")
		return
	}

	integerIdTps, err := strconv.Atoi(idTps)
	if err != nil {
		c.String(http.StatusBadRequest, "Id Tps is not valid!")
		return
	}

	integerJenisPemilihan, err := strconv.Atoi(jenisPemilihan)
	if err != nil {
		c.String(http.StatusBadRequest, "Jenis Pemilihan is not valid!")
		return
	}

	var formcImageGroup models.FormcImageGroup
	res, err := formcImageGroup.GetFormcImageGroupByIdTpsAndJenisPemilihan(integerIdTps, integerJenisPemilihan)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func SendFormcResultByVisionResponse(form forms.FormcImageVisionResponse) error {
	var formcImageModel models.FormcImage

	for i := 0; i < len(form.IdPaslonList); i++ {
		formcImage, err := formcImageModel.GetFormcImage(form.IdImageList[i])
		if err != nil {
			return err
		}

		tps, err := models.GetTpsById(formcImage.IdTps)
		if err != nil {
			return err
		}

		fullWilayahIdList, err := models.GetFullWilayahIdList(tps.IdWilayahDasar)

		formcResult := models.FormcResult{
			IdTps:           formcImage.IdTps,
			IdKelurahan:     fullWilayahIdList.IdKelurahan,
			IdKecamatan:     fullWilayahIdList.IdKecamatan,
			IdKabupatenKota: fullWilayahIdList.IdKabupatenKota,
			IdProvinsi:      fullWilayahIdList.IdProvinsi,
			IdPaslon:        form.IdPaslonList[i],
			JmlSuara:        form.JmlSuaraOcrList[i],
			JenisPemilihan:  formcImage.JenisPemilihan,
		}

		_, err = models.SendFormcResult(formcResult)
		if err != nil {
			return err
		}
	}

	return nil
}

func SendFormcResultStreamProcessingRequest(form forms.FormcImageVisionResponse) error {
	var formcImageModel models.FormcImage

	for i := 0; i < len(form.IdPaslonList); i++ {
		formcImage, err := formcImageModel.GetFormcImage(form.IdImageList[i])
		if err != nil {
			return err
		}

		tps, err := models.GetTpsById(formcImage.IdTps)
		if err != nil {
			return err
		}

		fullWilayahIdList, err := models.GetFullWilayahIdList(tps.IdWilayahDasar)

		formcResultStreamProcessingRequest := forms.FormcResultStreamProcessingRequest{
			IdTps:           formcImage.IdTps,
			IdKelurahan:     fullWilayahIdList.IdKelurahan,
			IdKecamatan:     fullWilayahIdList.IdKecamatan,
			IdKabupatenKota: fullWilayahIdList.IdKabupatenKota,
			IdProvinsi:      fullWilayahIdList.IdProvinsi,
			IdPaslon:        form.IdPaslonList[i],
			JmlSuara:        form.JmlSuaraOcrList[i],
			JenisPemilihan:  formcImage.JenisPemilihan,
		}

		var formcResultStreamProcessingRequestBuffer bytes.Buffer
		enc := gob.NewEncoder(&formcResultStreamProcessingRequestBuffer)
		err = enc.Encode(formcResultStreamProcessingRequest)
		if err != nil {
			return err
		}

		brokersUrl := []string{"34.70.240.187:9092"}
		producer, err := ConnectProducer(brokersUrl)
		if err != nil {
			return err
		}
		defer producer.Close()
		msg := &sarama.ProducerMessage{
			Topic: "vote",
			Value: sarama.StringEncoder(formcResultStreamProcessingRequestBuffer.Bytes()),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			return err
		}
		fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", "vote", partition, offset)
	}

	return nil
}

func SendFormcResultStreamProcessingRequestTest() error {

	formcResultStreamProcessingRequest := forms.FormcResultStreamProcessingRequest{
		IdTps:           1,
		IdKelurahan:     2,
		IdKecamatan:     3,
		IdKabupatenKota: 4,
		IdProvinsi:      5,
		IdPaslon:        6,
		JmlSuara:        100,
		JenisPemilihan:  2,
	}

	// p, err := kafka.NewProducer(&kafka.ConfigMap{
	// 	"bootstrap.servers": "35.238.208.74:9092",
	// })
	// if err != nil {
	// 	return err
	// }

	// defer p.Close()

	// // Delivery report handler for produced messages
	// go func() {
	// 	for e := range p.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()

	var formcResultStreamProcessingRequestBuffer bytes.Buffer
	enc := gob.NewEncoder(&formcResultStreamProcessingRequestBuffer)
	err := enc.Encode(formcResultStreamProcessingRequest)
	if err != nil {
		return err
	}

	// // Produce messages to topic (asynchronously)
	// topic := "vote"
	// p.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 	Value:          []byte(formcResultStreamProcessingRequestBuffer.Bytes()),
	// }, nil)

	// // Wait for message deliveries before shutting down
	// p.Flush(15 * 1000)

	brokersUrl := []string{"34.70.240.187:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: "vote",
		Value: sarama.StringEncoder(formcResultStreamProcessingRequestBuffer.Bytes()),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", "vote", partition, offset)

	return nil
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GenerateScanResultResponse(formcImageVisionResponse forms.FormcImageVisionResponse, formcImageRaw forms.FormcImageRaw) forms.FormcScanResponse {
	pdfFileName := "kesesuaian-" + strconv.Itoa(formcImageVisionResponse.IdImageList[0]+formcImageVisionResponse.IdImageList[1]+formcImageVisionResponse.IdImageList[2]) + ".pdf"

	formcScanNotificationResponse := forms.FormcScanResponse{
		PemilihDptL:    formcImageVisionResponse.PemilihDptL,
		PemilihDptP:    formcImageVisionResponse.PemilihDptP,
		PemilihDptJ:    formcImageVisionResponse.PemilihDptJ,
		PemilihDpphL:   formcImageVisionResponse.PemilihDpphL,
		PemilihDpphP:   formcImageVisionResponse.PemilihDpphP,
		PemilihDpphJ:   formcImageVisionResponse.PemilihDpphJ,
		PemilihDptbL:   formcImageVisionResponse.PemilihDptbL,
		PemilihDptbP:   formcImageVisionResponse.PemilihDptbP,
		PemilihDptbJ:   formcImageVisionResponse.PemilihDptbJ,
		PemilihTotalL:  formcImageVisionResponse.PemilihTotalL,
		PemilihTotalP:  formcImageVisionResponse.PemilihTotalP,
		PemilihTotalJ:  formcImageVisionResponse.PemilihTotalJ,
		PenggunaDptL:   formcImageVisionResponse.PenggunaDptL,
		PenggunaDptP:   formcImageVisionResponse.PenggunaDptP,
		PenggunaDptJ:   formcImageVisionResponse.PenggunaDptJ,
		PenggunaDpphL:  formcImageVisionResponse.PenggunaDpphL,
		PenggunaDpphP:  formcImageVisionResponse.PenggunaDpphP,
		PenggunaDpphJ:  formcImageVisionResponse.PenggunaDpphJ,
		PenggunaDptbL:  formcImageVisionResponse.PenggunaDptbL,
		PenggunaDptbP:  formcImageVisionResponse.PenggunaDptbP,
		PenggunaDptbJ:  formcImageVisionResponse.PenggunaDptbJ,
		PenggunaTotalL: formcImageVisionResponse.PenggunaTotalL,
		PenggunaTotalP: formcImageVisionResponse.PenggunaTotalP,
		PenggunaTotalJ: formcImageVisionResponse.PenggunaTotalJ,

		PemilihDisabilitasL:  formcImageVisionResponse.PemilihDisabilitasL,
		PemilihDisabilitasP:  formcImageVisionResponse.PemilihDisabilitasP,
		PemilihDisabilitasJ:  formcImageVisionResponse.PemilihDisabilitasJ,
		PenggunaDisabilitasL: formcImageVisionResponse.PenggunaDisabilitasL,
		PenggunaDisabilitasP: formcImageVisionResponse.PenggunaDisabilitasP,
		PenggunaDisabilitasJ: formcImageVisionResponse.PenggunaDisabilitasJ,

		SuratDiterima:       formcImageVisionResponse.SuratDiterima,
		SuratDikembalikan:   formcImageVisionResponse.SuratDikembalikan,
		SuratTidakDigunakan: formcImageVisionResponse.SuratTidakDigunakan,
		SuratDigunakan:      formcImageVisionResponse.SuratDigunakan,

		SuaraSah:            formcImageVisionResponse.SuaraSah,
		SuaraTidakSah:       formcImageVisionResponse.SuaraTidakSah,
		SuaraTotal:          formcImageVisionResponse.SuaraTotal,
		PenggunaHakPilih:    formcImageVisionResponse.PenggunaHakPilih,
		SuratSuaraDigunakan: formcImageVisionResponse.SuratSuaraDigunakan,

		JmlSuaraOcrList: formcImageVisionResponse.JmlSuaraOcrList,
		JmlSuaraOmrList: formcImageVisionResponse.JmlSuaraOmrList,

		IdTps:          formcImageRaw.IdTps,
		JenisPemilihan: formcImageRaw.JenisPemilihan,
		PayloadList:    formcImageRaw.PayloadList,
		IdPaslonList:   formcImageRaw.IdPaslonList,
		IdImageList:    formcImageVisionResponse.IdImageList,
		PdfUrl:         "https://storage.googleapis.com/staging-sirekap-form/pdf/" + pdfFileName,
	}

	return formcScanNotificationResponse
}

func SendPushNotification(formcImageVisionResponse forms.FormcImageVisionResponse, formcImageRaw forms.FormcImageRaw) error {
	formcScanNotificationResponse := GenerateScanResultResponse(formcImageVisionResponse, formcImageRaw)

	err := notifications.SendToToken("", formcScanNotificationResponse)
	if err != nil {
		return err
	}

	return nil
}
