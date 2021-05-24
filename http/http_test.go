package http

import (
	"github.com/khu-dev/khumu-comment/ent"
	"github.com/khu-dev/khumu-comment/ent/enttest"
	"github.com/khu-dev/khumu-comment/test"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

var (
	repo               *ent.Client
)

// B는 Before each의 acronym
func BeforeMiddlewareTest(tb testing.TB) {
	repo = enttest.Open(tb, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	test.SetUpUsers(repo)
}

// A는 After each의 acronym
func AfterMiddlewareTest(tb testing.TB) {
	repo.Close()
}