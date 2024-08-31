package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquipartition(t *testing.T) {
	num, cnt := 100, 5
	assert.Equal(t, &[]int{20, 20, 20, 20, 20}, Equipartition(num, cnt))
	num, cnt = 101, 5
	assert.Equal(t, &[]int{21, 20, 20, 20, 20}, Equipartition(num, cnt))
	num, cnt = 4, 1
	assert.Equal(t, &[]int{4}, Equipartition(num, cnt))
	num, cnt = 4, 0
	assert.Nil(t, Equipartition(num, cnt))
}
