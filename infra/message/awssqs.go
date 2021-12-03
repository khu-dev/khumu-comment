package message

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/khu-dev/khumu-comment/config"
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type MessageHandler interface {
	Listen()
	OnUserCreated(body *data.CommandCenterKhumuUserDto) (*ent.KhumuUser, error)
	OnUserUpdated(body *data.CommandCenterKhumuUserDto) (*ent.KhumuUser, error)
	OnUserDeleted(body *data.CommandCenterKhumuUserDto) (*ent.KhumuUser, error)
}

type Body struct {
	Message           string
	MessageAttributes struct {
		ResourceKind MessageAttribute `json:"resource_kind"`
		EventKind    MessageAttribute `json:"event_kind"`
	}
}

type MessageAttribute struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}
type SqsMessageHandler struct {
	sqs        *sqs.SQS
	userDB     *ent.KhumuUserClient
	termSignal <-chan os.Signal
}

func NewSqsMessageHandler(termSignal <-chan os.Signal, db *ent.Client) MessageHandler {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})
	if err != nil {
		log.Error("NewSession error:", err)
		return nil
	}

	sqsClient := sqs.New(sess)

	return &SqsMessageHandler{
		sqs:        sqsClient,
		userDB:     db.KhumuUser,
		termSignal: termSignal,
	}
}

func (h *SqsMessageHandler) Listen() {
	for {
		select {
		case sig := <-h.termSignal:
			log.Infof("%#v 시그널에 의해 SqsMessageHandler의 Listen을 종료합니다.", sig)
			return
		default:
			output, err := h.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
				AttributeNames:          nil,
				MaxNumberOfMessages:     aws.Int64(10),
				MessageAttributeNames:   nil,
				QueueUrl:                aws.String(config.Config.Sqs.QueueURL),
				ReceiveRequestAttemptId: nil,
				VisibilityTimeout:       nil,
				WaitTimeSeconds:         aws.Int64(20),
			})
			if err != nil {
				log.Errorf("%+v", errors.Wrap(err, "SQS 메시지를 불러오던 중 에러 발생"))
			} else {
				for _, message := range output.Messages {
					parsedBody := new(Body)
					err := json.Unmarshal([]byte(*message.Body), parsedBody)
					if err != nil {
						log.Errorf("%+v", errors.Wrap(err, "SQS 메시지를 파싱하던 중 에러 발생"))
					}

					resourceKind := parsedBody.MessageAttributes.ResourceKind.Value
					eventKind := parsedBody.MessageAttributes.EventKind.Value
					if resourceKind == "" || eventKind == "" {
						log.Errorf("resource_kind와 event_kind가 올바르게 전달되지 못했습니다. %#v, %#v", resourceKind, eventKind)
						continue
					}
					switch resourceKind {
					case "user":
						body := new(data.CommandCenterKhumuUserDto)
						err := json.Unmarshal([]byte(parsedBody.Message), body)
						if err != nil {
							log.Errorf("%+v", errors.Wrap(err, "SQS 메시지를 파싱하던 중 에러 발생"))
							time.Sleep(5 * time.Second)
							continue
						}
						switch eventKind {
						case "create":
							if _, err := h.OnUserCreated(body); err != nil {
								log.Errorf("%+v", err)
							}
						case "update":
							if _, err := h.OnUserUpdated(body); err != nil {
								log.Errorf("%+v", err)
							}
						case "delete":
							if _, err := h.OnUserDeleted(body); err != nil {
								log.Errorf("%+v", err)
							}
						}
					}

					if _, err := h.sqs.DeleteMessage(&sqs.DeleteMessageInput{QueueUrl: aws.String(config.Config.Sqs.QueueURL), ReceiptHandle: message.ReceiptHandle}); err != nil {
						log.Errorf("%+v", err)
					}
				}
			}
		}
	}
}

func (h *SqsMessageHandler) OnUserCreated(body *data.CommandCenterKhumuUserDto) (*ent.KhumuUser, error) {
	log.Infof("유저 %#v 를 생성합니다.\n", body)
	user, err := h.userDB.Create().SetID(body.Username).SetStatus(body.Status).SetNickname(body.Nickname).Save(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "유저 생성 실패")
	}

	return user, nil
}

func (h *SqsMessageHandler) OnUserUpdated(body *data.CommandCenterKhumuUserDto) (*ent.KhumuUser, error) {
	log.Infof("유저 %#v 를 수정합니다.\n", body)

	user, err := h.userDB.UpdateOneID(body.Username).SetNickname(body.Nickname).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}

func (h *SqsMessageHandler) OnUserDeleted(body *data.CommandCenterKhumuUserDto) (*ent.KhumuUser, error) {
	log.Infof("유저 %#v 를 삭제(soft delete)합니다.\n", body)

	user, err := h.userDB.UpdateOneID(body.Username).SetStatus("deleted").Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return user, nil
}
