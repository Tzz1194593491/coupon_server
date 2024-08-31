package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRedisKey(t *testing.T) {
	assert.Equal(t, PreKey+"abc", GetRedisKey("abc"))
	assert.Equal(t, PreKey+"abc:efg", GetRedisKey("abc", "efg"))
}

func TestShareKey(t *testing.T) {
	assert.Equal(t, &[]string{"a:0", "a:1", "a:2", "a:3"}, ShareKey("a", 4))
}
