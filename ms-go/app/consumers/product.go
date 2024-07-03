package consumers

import (
	"context"
	"encoding/json"
	"log"
	"ms-go/app/kafka"
	"ms-go/app/models"
	"ms-go/app/services/products"
)

func ListenProduct(){
	kafkaReader := kafka.Reader()
	
	var product models.Product
	for {
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		//fmt.Printf("rails-to-go received: %s \n", string(msg.Value))
		
		if err := json.Unmarshal(msg.Value, &product); err != nil {
			log.Printf("error to unmarshal product json: %s", err)
			continue
		}
		
		dbProduct, _ := products.Details(product)
		
		if dbProduct == nil {
			_, err := products.Create(product, false)
			if err != nil {
				log.Printf("error to create product on consumer: %s", err)
			}
		}else{
			_, err := products.Update(product, false)
			if err != nil {
				log.Printf("error to update product on consumer: %s", err)
			}
		}
		
	}
	
	if err := kafkaReader.Close(); err != nil {
		log.Printf("failed to close reader: %s", err)
	}
}