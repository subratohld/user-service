package server

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New(viper.New(), nil)
	assert.NotNil(t, s)
}
