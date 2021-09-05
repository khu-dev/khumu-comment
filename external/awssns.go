package external

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/data"
	log "github.com/sirupsen/logrus"
)

var (
	CommentCreateMessageAttribute = map[string]*sns.MessageAttributeValue{
		"resource_kind": {DataType: aws.String("String"), StringValue: aws.String("comment")},
		"event_kind":    {DataType: aws.String("String"), StringValue: aws.String("create")},
	}
)

type SnsClient interface {
	PublishMessage(comment *data.CommentOutput)
}

type SnsClientImpl struct {
	Sns *sns.SNS
}

func NewSnsClient() SnsClient {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})
	if err != nil {
		log.Error("NewSession error:", err)
		return nil
	}

	snsClient := sns.New(sess)

	return &SnsClientImpl{
		Sns: snsClient,
	}
}

func (client *SnsClientImpl) PublishMessage(comment *data.CommentOutput) {
	jsonData, err := json.Marshal(comment)
	if err != nil {
		log.Error(err)
		return
	}
	input := &sns.PublishInput{
		Message:           aws.String(string(jsonData)),
		TopicArn:          aws.String(config.Config.Sns.TopicArn),
		MessageAttributes: CommentCreateMessageAttribute,
	}

	result, err := client.Sns.Publish(input)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Publish SNS Message " + *input.Message)
	log.Info(input.MessageAttributes)

	log.Info(result)
}
