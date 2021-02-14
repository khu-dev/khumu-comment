package repository

import (
    "github.com/khu-dev/khumu-comment/model"
    "github.com/khu-dev/khumu-comment/test"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestRedisEventMessageRepository_publishCommentEvent(t *testing.T) {
    B(t)
    defer A(t)
    tmpComment := test.CommentsData["JinsuAnonymousComment"]
    assert.NotNil(t, eventMessageRepository)
    assert.NotNil(t, tmpComment)
    err := eventMessageRepository.publishCommentEvent(&model.EventMessage{
        ResourceKind:"comment",
        EventKind: "create",
        Resource: tmpComment,
    })
    assert.NoError(t, err)
}
