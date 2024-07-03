package products

import (
	//import model product
	"context"
	"encoding/json"
	"log"
	"ms-go/app/kafka"
	"ms-go/app/models"
)

func sendProductMessage(product models.Product){
	kafkaWriter := kafka.Writer()
	
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		log.Printf("failed to json product: %s", err)
		return
	}
	
	err = kafkaWriter.WriteMessages(
		context.Background(),
		kafka.Message(jsonProduct),
	)
	
	if err != nil {
		log.Printf("failed to write messages: %s", err)
	}

	if err := kafkaWriter.Close(); err != nil {
		log.Printf("failed to close writer: %s", err)
	}
	
}