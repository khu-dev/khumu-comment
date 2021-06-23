package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/umi0410/ezconfig"
	"os"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	wd, _ := os.Getwd()

	t.Run("기본 unmarshal", func(t *testing.T) {
		config := &KhumuConfig{}

		os.Setenv("KHUMU_DB_KIND", "this is just test")
		ezconfig.LoadConfig("KHUMU", config, []string{wd, strings.ToLower(os.Getenv("KHUMU_CONFIG_PATH"))})
		assert.NotEqual(t, "", config.DB.Kind)
	})

	t.Run("env override", func(t *testing.T) {
		config := &KhumuConfig{}
		os.Setenv("KHUMU_DB.KIND", "this is just test")
		t.Log(os.Getenv("KHUMU_DB.KIND"))

		ezconfig.LoadConfig("KHUMU", config, []string{wd, strings.ToLower(os.Getenv("KHUMU_CONFIG_PATH"))})
		assert.Equal(t, "this is just test", config.DB.Kind)
	})

}
