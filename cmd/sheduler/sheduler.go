package main

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/vsPEach/Kursovaya/internal/entity"
	"github.com/vsPEach/Kursovaya/internal/queue/rabbit"
	"log"
	"time"
)

func main() {
	mq, err := rabbit.NewPublisher()
	if err != nil {
		log.Fatalln(err)
	}
	marshal, err := json.Marshal(entity.Notice{
		ID:     uuid.New(),
		Title:  "Hello world",
		Date:   time.Now(),
		UserID: uuid.New(),
	})
	if err != nil {
		log.Println(err)
	}
	mq.Publish(context.Background(), marshal)
}
