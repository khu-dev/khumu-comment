package khumu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/khu-dev/khumu-comment/config"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	API_TIMEOUT = 10 * time.Second
)

// KhumuAPIAdapter 는 khumu의 API를 바탕으로한 마이크로서비스간의 통신을 위한 struct
type KhumuAPIAdapter interface {
	IsAuthor(articleID int, authorID string) <-chan bool
}
type KhumuAPIAdapterImpl struct {
	CommandCenterRootURL string
}

func NewKhumuAPIAdapter() KhumuAPIAdapter {
	adapter := &KhumuAPIAdapterImpl{
		CommandCenterRootURL: config.Config.Khumu.CommandCenter.RootURL,
	}
	return adapter
}

// IsAuthor 는 concurrent하게 isAuthor 를 이용할 수 있게 해줌.
func (a *KhumuAPIAdapterImpl) IsAuthor(articleID int, authorID string) <-chan bool {
	resultChan := make(chan bool)
	go a.isAuthor(articleID, authorID, resultChan)
	return resultChan
}

// isAuthor 는 synchorous하게 command-center microservice에게 게시글의 글쓴이 정보를 체크함.
func (a *KhumuAPIAdapterImpl) isAuthor(articleID int, authorID string, resultChan chan<- bool) {
	client := http.Client{Timeout: API_TIMEOUT}
	body := &IsAuthorReq{
		Author: authorID,
	}
	data, err := json.Marshal(body)
	if err != nil {
		log.Error(err)
		resultChan <- false
	}
	log.Info(string(data))

	resp, err := client.Post(fmt.Sprintf("%s/articles/%d/is-author", a.CommandCenterRootURL, articleID), "application/json", bytes.NewReader(data))

	if err != nil {
		log.Error(err)
		resultChan <- false
		return
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		resultChan <- false
		return
	}
	log.Info(string(data))
	result := new(IsAuthorResp)
	err = json.Unmarshal(data, result)
	if err != nil {
		log.Error(err)
		resultChan <- false
		return
	}

	resultChan <- result.Data.IsAuthor
}
