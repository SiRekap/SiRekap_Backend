package externalcomponent

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/Shopify/sarama"
)

type FormcResultStreamProcessingRequest struct {
	IdTps           int `json:"id_tps" binding:"required"`
	IdKelurahan     int `json:"id_kelurahan" binding:"required"`
	IdKecamatan     int `json:"id_kecamatan" binding:"required"`
	IdKabupatenKota int `json:"id_kabupaten_kota" binding:"required"`
	IdProvinsi      int `json:"id_provinsi" binding:"required"`
	IdPaslon        int `json:"id_paslon" binding:"required"`
	JmlSuara        int `json:"jml_suara" binding:"required"`
	JenisPemilihan  int `json:"jenis_pemilihan" binding:"required"`
}

func init() {
	functions.HTTP("ExternalComponent", ExternalComponent)
}

// ExternalComponent is an HTTP Cloud Function with a request parameter.
func ExternalComponent(w http.ResponseWriter, r *http.Request) {
	err := SendFormcResultStreamProcessingRequestTest()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println("Success")
}

func SendFormcResultStreamProcessingRequestTest() error {

	formcResultStreamProcessingRequest := FormcResultStreamProcessingRequest{
		IdTps:           1,
		IdKelurahan:     2,
		IdKecamatan:     3,
		IdKabupatenKota: 4,
		IdProvinsi:      5,
		IdPaslon:        6,
		JmlSuara:        100,
		JenisPemilihan:  2,
	}

	var formcResultStreamProcessingRequestBuffer bytes.Buffer
	enc := gob.NewEncoder(&formcResultStreamProcessingRequestBuffer)
	err := enc.Encode(formcResultStreamProcessingRequest)
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
