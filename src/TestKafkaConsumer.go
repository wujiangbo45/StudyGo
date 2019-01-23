package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	wg     sync.WaitGroup
	logger = log.New(os.Stderr, "[srama]", log.LstdFlags)
)

func main() {
	sarama.Logger = logger
	config := sarama.NewConfig()
	config.ClientID = "test"
	consumer, err := sarama.NewConsumer(strings.Split("172.16.1.158:9092", ","), config)
	if err != nil {
		logger.Printf("Failed to start consumer: %s", err)
	}
	partitionList, err := consumer.Partitions("wujiangbo")
	if err != nil {
		logger.Printf("Failed to get the list of partitions: %s", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("wujiangbo", int32(partition), sarama.OffsetNewest)
		if err != nil {
			logger.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
		}(pc)
	}

	wg.Wait()

	logger.Println("Done consuming topic hello")
	consumer.Close()
}
