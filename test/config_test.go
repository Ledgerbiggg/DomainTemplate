package test

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestConfig test config initialization
func TestConfig(t *testing.T) {
	c := gConfig
	val := reflect.ValueOf(c).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name
		assert.Falsef(t, field.IsZero(), "Field %s is not initialized", fieldName)
	}
}
