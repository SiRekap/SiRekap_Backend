package controllers

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"sirekap/SiRekap_Backend/forms"
	"sirekap/SiRekap_Backend/models"

	"github.com/gin-gonic/gin"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type FormcImageController struct{}

func (f FormcImageController) SendFormcImagePayload(c *gin.Context) {
	var formcImagePayload models.FormcImagePayload

	if err := c.BindJSON(&formcImagePayload); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcImagePayload.SendFormcImagePayload()
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

	if err := c.BindJSON(&formcImageRaw); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := models.SendFormcImageRaw(formcImageRaw)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// TODO: Async
	formcImageVisionResponse, err := SendFormcImageVisionRequest(*form)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = SendFormcResultStreamProcessingRequest(formcImageVisionResponse)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
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

func SendFormcImageVisionRequest(form forms.FormcImageRawResponse) (forms.FormcImageVisionResponse, error) {
	// formcImageVisionRequest := forms.FormcImageVisionRequest{
	// 	IdImageList:  form.IdImageList,
	// 	PayloadList:  form.FormcImageRaw.PayloadList,
	// 	IdPaslonList: form.FormcImageRaw.IdPaslonList,
	// }

	resp := forms.FormcImageVisionResponse{}

	formcAdministrasiHlmSatuProses := models.FormcAdministrasiHlmSatuProses{
		IdImage: resp.IdImageList[0],

		PemilihDptL:    resp.PemilihDptL,
		PemilihDptP:    resp.PemilihDptP,
		PemilihDptJ:    resp.PemilihDptJ,
		PemilihDpphL:   resp.PemilihDpphL,
		PemilihDpphP:   resp.PemilihDpphP,
		PemilihDpphJ:   resp.PemilihDpphJ,
		PemilihDptbL:   resp.PemilihDptbL,
		PemilihDptbP:   resp.PemilihDptbP,
		PemilihDptbJ:   resp.PemilihDptbJ,
		PemilihTotalL:  resp.PemilihTotalL,
		PemilihTotalP:  resp.PemilihTotalP,
		PemilihTotalJ:  resp.PemilihTotalJ,
		PenggunaDptL:   resp.PenggunaDptL,
		PenggunaDptP:   resp.PenggunaDptP,
		PenggunaDptJ:   resp.PenggunaDptJ,
		PenggunaDpphL:  resp.PenggunaDpphL,
		PenggunaDpphP:  resp.PenggunaDpphP,
		PenggunaDpphJ:  resp.PenggunaDpphJ,
		PenggunaDptbL:  resp.PenggunaDptbL,
		PenggunaDptbP:  resp.PenggunaDptbP,
		PenggunaDptbJ:  resp.PenggunaDptbJ,
		PenggunaTotalL: resp.PenggunaTotalL,
		PenggunaTotalP: resp.PenggunaTotalP,
		PenggunaTotalJ: resp.PenggunaTotalJ,

		PemilihDisabilitasL:  resp.PemilihDisabilitasL,
		PemilihDisabilitasP:  resp.PemilihDisabilitasP,
		PemilihDisabilitasJ:  resp.PemilihDisabilitasJ,
		PenggunaDisabilitasL: resp.PenggunaDisabilitasL,
		PenggunaDisabilitasP: resp.PenggunaDisabilitasP,
		PenggunaDisabilitasJ: resp.PenggunaDisabilitasJ,

		SuratDiterima:       resp.SuratDiterima,
		SuratDikembalikan:   resp.SuratDikembalikan,
		SuratTidakDigunakan: resp.SuratTidakDigunakan,
		SuratDigunakan:      resp.SuratDigunakan,
	}

	_, err := models.SendFormcAdministrasiHlmSatuProses(formcAdministrasiHlmSatuProses)
	if err != nil {
		return resp, err
	}

	formcAdministrasiHlmSatuFinal := models.FormcAdministrasiHlmSatuFinal{
		IdImage: resp.IdImageList[0],
		IsBenar: true,
	}

	_, err = models.SendFormcAdministrasiHlmSatuFinal(formcAdministrasiHlmSatuFinal)
	if err != nil {
		return resp, err
	}

	formcAdministrasiHlmDuaProses := models.FormcAdministrasiHlmDuaProses{
		IdImage: resp.IdImageList[1],

		SuaraSah:            resp.SuaraSah,
		SuaraTidakSah:       resp.SuaraTidakSah,
		SuaraTotal:          resp.SuaraTotal,
		PenggunaHakPilih:    resp.PenggunaHakPilih,
		SuratSuaraDigunakan: resp.SuratSuaraDigunakan,
	}

	_, err = models.SendFormcAdministrasiHlmDuaProses(formcAdministrasiHlmDuaProses)
	if err != nil {
		return resp, err
	}

	formcAdministrasiHlmDuaFinal := models.FormcAdministrasiHlmDuaFinal{
		IdImage: resp.IdImageList[1],
		IsBenar: true,
	}

	_, err = models.SendFormcAdministrasiHlmDuaFinal(formcAdministrasiHlmDuaFinal)
	if err != nil {
		return resp, err
	}

	var formcImage models.FormcImage
	err = formcImage.GetFormcImage(form.IdImageList[0])
	if err != nil {
		return resp, err
	}

	formcImageGroup := models.FormcImageGroup{
		IdTps:          formcImage.IdTps,
		JenisPemilihan: formcImage.JenisPemilihan,
		IdImageHlm1:    resp.IdImageList[0],
		IdImageHlm2:    resp.IdImageList[1],
		IdImageHlm3:    resp.IdImageList[2],
	}

	_, err = formcImageGroup.SendFormcImageGroup()
	if err != nil {
		return resp, err
	}

	for i := 0; i < len(resp.IdPaslonList); i++ {
		suaracProses := models.SuaraCProses{
			IdImage:     resp.IdImageList[2],
			IdPaslon:    resp.IdPaslonList[i],
			JmlSuaraOcr: resp.JmlSuaraOcrList[i],
			JmlSuaraOmr: resp.JmlSuaraOmrList[i],
		}

		_, err = models.SendSuaraCProses(suaracProses)
		if err != nil {
			return resp, err
		}

		suaracFinal := models.SuaraCFinal{
			IdImage:  resp.IdImageList[2],
			IdPaslon: resp.IdPaslonList[i],
			JmlSuara: resp.JmlSuaraOcrList[i],
		}

		_, err = models.SendSuaraCFinal(suaracFinal)
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func SendFormcResultStreamProcessingRequest(form forms.FormcImageVisionResponse) error {
	var formcImage models.FormcImage

	for i := 0; i < len(form.IdPaslonList); i++ {
		err := formcImage.GetFormcImage(form.IdImageList[i])
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

		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "sulky.srvs.cloudkafka.com:9094",
			"sasl.jaas.config":  "org.apache.kafka.common.security.plain.PlainLoginModule required username=\"jetffmjg\" password=\"7LXhVBsEndIGSHeBmDComfQaajZVpWPZ\";",
		})
		if err != nil {
			panic(err)
		}

		defer p.Close()

		// Delivery report handler for produced messages
		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
					} else {
						fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}()

		var formcResultStreamProcessingRequestBuffer bytes.Buffer
		enc := gob.NewEncoder(&formcResultStreamProcessingRequestBuffer)
		err = enc.Encode(formcResultStreamProcessingRequest)
		if err != nil {
			return err
		}

		// Produce messages to topic (asynchronously)
		topic := "jetffmjg-sirekap-vote"
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(formcResultStreamProcessingRequestBuffer.Bytes()),
		}, nil)

		// Wait for message deliveries before shutting down
		p.Flush(15 * 1000)
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

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "sulky.srvs.cloudkafka.com:9094",
		"sasl.jaas.config":  "org.apache.kafka.common.security.plain.PlainLoginModule required username=\"jetffmjg\" password=\"7LXhVBsEndIGSHeBmDComfQaajZVpWPZ\";",
	})
	if err != nil {
		return err
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	var formcResultStreamProcessingRequestBuffer bytes.Buffer
	enc := gob.NewEncoder(&formcResultStreamProcessingRequestBuffer)
	err = enc.Encode(formcResultStreamProcessingRequest)
	if err != nil {
		return err
	}

	// Produce messages to topic (asynchronously)
	topic := "jetffmjg-sirekap-vote"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(formcResultStreamProcessingRequestBuffer.Bytes()),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	return nil
}
