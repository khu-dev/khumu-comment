package test

import (
	"github.com/khu-dev/khumu-comment/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

var (
	testData *TestData = &TestData{}
)

type User struct {
	Username string `yaml:"Username"`
}
type TestData struct {
	//Users []*User `yaml:"Users"`
	Users []*model.KhumuUserSimple `yaml:"Users"`

	Comments []*model.Comment `yaml:"Comments"`
}

func Test_testdata를_잘_인식하는지(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata.yaml")
	assert.NoError(t, err)
	//t.Log(string(buf))
	err = yaml.Unmarshal(buf, testData)
	assert.NoError(t, err)
	assert.NotNil(t, testData)
	//t.Log(testData)
	assert.Greater(t, len(testData.Comments), 1)
}
