package repository

import (
    "github.com/khu-dev/khumu-comment/model"
    "github.com/khu-dev/khumu-comment/test"
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

var (
	redisEventMessageRepository *RedisEventMessageRepository
)

func TestRedisEventMessageRepository_publishCommentEvent(t *testing.T) {
	// 내가 실행시킬 떄에만 푸시 테스트 실행
	if os.Getenv("KHUMU_EXECUTOR") != "jinsu" {
		t.Skip("jinsu가 실행하지 않아서 푸시 알림 테스트는 스킵합니다.")
	} else {
		test.SetUp()
		redisEventMessageRepository = NewRedisEventMessageRepository().(*RedisEventMessageRepository)
		tmpComment := test.Comment1JinsuAnnonymous
		assert.NotNil(t, redisEventMessageRepository)
		assert.NotNil(t, tmpComment)
		err := redisEventMessageRepository.publishCommentEvent(&model.EventMessage{
			ResourceKind: "comment",
			EventKind:    "create",
			Resource:     tmpComment,
		})
		assert.NoError(t, err)
	}

}
