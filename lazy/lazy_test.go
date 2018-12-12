package lazy_test

import (
	"testing"

	"github.com/domipi/domitools/lazy"
	"github.com/stretchr/testify/assert"
)

func TestPoVToV(t *testing.T) {
	intValue := 2
	strValue := "coucou"
	var int64Value int64 = 5

	intPtr := &intValue
	strPtr := &strValue
	int64Ptr := &int64Value

	var nilStrPtr *string
	var nilInt64Ptr *int64
	var nilBoolPtr *bool

	assert.Equal(t, 2, lazy.PoVToV(intPtr).(int), "int ptr ok")
	assert.Equal(t, 2, lazy.PoVToV(intValue).(int), "int ptr ok")
	assert.Equal(t, "coucou", lazy.PoVToV(strPtr).(string), "str ptr ok")
	assert.Equal(t, "coucou", lazy.PoVToV(strValue).(string), "str ptr ok")
	assert.Equal(t, int64(5), lazy.PoVToV(int64Ptr).(int64), "int64 ptr ok")
	assert.Equal(t, int64(5), lazy.PoVToV(int64Value).(int64), "int64 ptr ok")
	assert.Equal(t, "", lazy.PoVToV(nilStrPtr).(string), "nil str ptr ok")
	assert.Equal(t, int64(0), lazy.PoVToV(nilInt64Ptr).(int64), "nil int64 ptr ok")
	assert.Equal(t, false, lazy.PoVToV(nilBoolPtr).(bool), "nil bool ptr ok")
}
