package src

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseYaml(t *testing.T) {
	confTest := &Config{"3000", APIConfig{Prefix: "/api"}}
	config, err := ParseYaml("../config.yml")
	assert.NoError(t, err)
	if !reflect.DeepEqual(config, confTest) {
		t.Errorf("TestParseYaml %v got %v", confTest, config)
	}
}
