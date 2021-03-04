package repository

import (
    "github.com/khu-dev/khumu-comment/model"
    "github.com/khu-dev/khumu-comment/test"
    "github.com/stretchr/testify/assert"
    "testing"
)

var (
    redisEventMessageRepository *RedisEventMessageRepository
)
func TestRedisEventMessageRepository_publishCommentEvent(t *testing.T) {
    test.SetUp()
    redisEventMessageRepository = NewRedisEventMessageRepository().(*RedisEventMessageRepository)
    tmpComment := test.Comment1JinsuAnnonymous
    assert.NotNil(t, redisEventMessageRepository)
    assert.NotNil(t, tmpComment)
    err := redisEventMessageRepository.publishCommentEvent(&model.EventMessage{
        ResourceKind:"comment",
        EventKind: "create",
        Resource: tmpComment,
    })
    assert.NoError(t, err)
}
