package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"go-ops/config"
	"sync"
)

var (
	wg sync.WaitGroup
	host string
	port string
	url string
	addrs []string
)

func init() {
	kafka := config.GetConfig().Kafka
	host = kafka.Host
	port = kafka.Port
	url = fmt.Sprintf("%s:%s", host, port)
	addrs = []string{url}

}


func Consume(topic string) {
	//创建消费者
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	//设置分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)
	//循环分区
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}

		}(pc)
	}
	//time.Sleep(time.Hour)
	wg.Wait()
	consumer.Close()
}
