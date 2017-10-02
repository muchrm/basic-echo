package src

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseYaml(t *testing.T) {
	confTest := &Config{Port: "3000", Api: ApiConfig{Prefix: "/api"}}
	config, err := ParseYaml()
	assert.NoError(t, err)
	if !reflect.DeepEqual(config, confTest) {
		t.Errorf("TestParseYaml %v got %v", confTest, config)
	}
}
