package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	consumer "github.com/harlow/kinesis-consumer"
	"github.com/jbvmio/awsgo"
)

func allKinesisStreams(aFlags AWSFlags) {
	streams := getKinesisStreams(aFlags)
	var shards []*string
	for _, s := range streams.StreamDescription.Shards {
		shards = append(shards, s.ShardId)
	}
	var itr []*string
	for _, s := range shards {
		input := kinesis.GetShardIteratorInput{
			ShardId:           s,
			ShardIteratorType: aws.String(`LATEST`),
			//ShardIteratorType: aws.String(`TRIM_HORIZON`),
			StreamName: aws.String(`etr-sandbox-container-logs`),
		}
		output, err := client.Kinesis().GetShardIterator(&input)
		if err != nil {
			Failf("error listing repositories: %v", err)
		}
		itr = append(itr, output.ShardIterator)
	}
	var wg sync.WaitGroup
	stopChan := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for _, i := range itr {
		wg.Add(1)
		go func(k *kinesis.Kinesis, iterator *string, sChan chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			in := kinesis.GetRecordsInput{
				Limit:         aws.Int64(10000),
				ShardIterator: iterator,
			}
			out, err := k.GetRecords(&in)
			if err != nil {
				Warnf("error retrieving records: %v", err)
				return
			}
			for {
				select {
				case <-stopChan:
					return
				default:
					in = kinesis.GetRecordsInput{
						Limit:         aws.Int64(10000),
						ShardIterator: out.NextShardIterator,
					}
					out, err = client.Kinesis().GetRecords(&in)
					if err != nil {
						Warnf("error retrieving records2: %v", err)
						return
					}
					for _, r := range out.Records {
						Infof("%s", r.Data)
					}
				}
			}

		}(client.Kinesis(), i, stopChan, &wg)
	}
waitLoop:
	for {
		select {
		case <-sigChan:
			close(stopChan)
			break waitLoop
		}
	}
	wg.Wait()
	Infof("Done.")
}

func consumeKinesis(aFlags AWSFlags) {
	client.AddConfig(awsgo.SvcTypeKinesis, aFlags)
	c, err := consumer.New(
		`etr-sandbox-container-logs`,
		consumer.WithClient(client.Kinesis()),
		consumer.WithShardIteratorType(`TRIM_HORIZON`),
	)
	if err != nil {
		Failf("consumer error: %v", err)
	}
	ctx := trap()
	err = c.Scan(ctx, func(r *consumer.Record) error {
		Infof("%s", r.Data)
		return nil // continue scanning
	})
	if err != nil {
		log.Fatalf("scan error: %v", err)
	}
}

func trap() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		log.Printf("received %s", sig)
		cancel()
	}()

	return ctx
}

func getKinesisRecords(aFlags AWSFlags) {
	client.AddConfig(awsgo.SvcTypeKinesis, aFlags)
	input := kinesis.GetRecordsInput{
		Limit:         aws.Int64(10000),
		ShardIterator: aws.String(`AAAAAAAAAAFeC6Z8JjBY7rQQsLvrngTZjQLruPyXkvtA/8eR2FTZG9PXNR0Q1QOFIVTTXikHma5mIGqLA4GMyHgo8pVCBGwmOCPmRy/E1Ub+ilOJsK7vxZJkpwlNCru/iftb0AozaKxqOztIOrWKp4c17+cKXUKGLK3tScPMp8XoB9wFrm1JqEygT9pZC8cMATGjYVzq28cYRhOhwX9wAyJfsTeaw38h4CjHOha0truAp1spkdczeA==`),
	}
	output, err := client.Kinesis().GetRecords(&input)
	if err != nil {
		Failf("error listing repositories: %v", err)
	}
	for len(output.Records) < 1 {
		input = kinesis.GetRecordsInput{
			Limit:         aws.Int64(10000),
			ShardIterator: output.NextShardIterator,
		}
		output, err = client.Kinesis().GetRecords(&input)
		if err != nil {
			Failf("error listing repositories: %v", err)
		}
	}
	Infof("\n\n%+v\n\n", output)
}

func getShardIterator(aFlags AWSFlags) {
	client.AddConfig(awsgo.SvcTypeKinesis, aFlags)
	input := kinesis.GetShardIteratorInput{
		ShardId: aws.String(`shardId-000000000000`),
		//ShardIteratorType: aws.String(`LATEST`),
		ShardIteratorType: aws.String(`TRIM_HORIZON`),
		StreamName:        aws.String(`etr-sandbox-container-logs`),
	}
	output, err := client.Kinesis().GetShardIterator(&input)
	if err != nil {
		Failf("error listing repositories: %v", err)
	}
	Infof("%s", output)
}

func getKinesisStreams(aFlags AWSFlags) *kinesis.DescribeStreamOutput {
	client.AddConfig(awsgo.SvcTypeKinesis, aFlags)
	input := kinesis.DescribeStreamInput{
		StreamName: aws.String(`etr-sandbox-container-logs`),
	}
	output, err := client.Kinesis().DescribeStream(&input)
	if err != nil {
		Failf("error listing repositories: %v", err)
	}
	return output
}

func listKinesisStreams(aFlags AWSFlags) {
	client.AddConfig(awsgo.SvcTypeKinesis, aFlags)
	output, err := client.Kinesis().ListStreams(nil)
	if err != nil {
		Failf("error listing repositories: %v", err)
	}
	Infof("%s", output)
}
