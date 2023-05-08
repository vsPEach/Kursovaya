package main

import (
	"context"
	"fmt"
	"github.com/vsPEach/Kursovaya/internal/queue/rabbit"
	"log"
	"sync"
)

func main() {
	mq, err := rabbit.NewSubscriber()
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	err = mq.Subscribe(context.Background())
	if err != nil {
		log.Println(err)
	}
}
