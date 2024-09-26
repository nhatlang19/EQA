package sqs

import (
	cloud_aws "EQA/backend/app/pkg/cloud/aws"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

const DefaultTimeout = 5 * time.Second

type SQS struct {
	timeout time.Duration
	client  *sqs.SQS
}

func (s SQS) Send(ctx context.Context, req *SendRequest) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	attrs := make(map[string]*sqs.MessageAttributeValue, len(req.Attributes))
	for _, attr := range req.Attributes {
		attrs[attr.Key] = &sqs.MessageAttributeValue{
			StringValue: aws.String(attr.Value),
			DataType:    aws.String(attr.Type),
		}
	}
	uuid, _ := uuid.NewUUID()
	groupId := uuid.String()

	res, err := s.client.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		MessageAttributes:      attrs,
		MessageBody:            aws.String(req.Body),
		QueueUrl:               aws.String(req.QueueURL),
		MessageGroupId:         &groupId,
		MessageDeduplicationId: &groupId,
	})
	if err != nil {
		return "", fmt.Errorf("send: %w", err)
	}

	return *res.MessageId, nil
}

func (s SQS) Receive(ctx context.Context, queueURL string, maxMsg int64) ([]*sqs.Message, error) {
	var waitTimeSeconds int64 = 10
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(waitTimeSeconds+5))
	defer cancel()

	res, err := s.client.ReceiveMessageWithContext(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(queueURL),
		MaxNumberOfMessages:   aws.Int64(maxMsg),
		WaitTimeSeconds:       aws.Int64(waitTimeSeconds),
		MessageAttributeNames: aws.StringSlice([]string{"All"}),
	})
	if err != nil {
		return nil, fmt.Errorf("receive: %w", err)
	}

	return res.Messages, nil
}

func (s SQS) Delete(ctx context.Context, queueURL, rcvHandle string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.client.DeleteMessageWithContext(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(rcvHandle),
	}); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (s SQS) PurgeQueue(ctx context.Context, queueURL string) error {
	if _, err := s.client.PurgeQueue(&sqs.PurgeQueueInput{QueueUrl: &queueURL}); err != nil {
		return fmt.Errorf("purge: %w", err)
	}
	return nil
}

func NewSQS() *SQS {
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")

	awsSession := cloud_aws.NewClient(accessKey, secretKey)
	return &SQS{
		timeout: DefaultTimeout,
		client:  sqs.New(awsSession),
	}
}
