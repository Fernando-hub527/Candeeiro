package utils_test

import (
	"testing"

	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestRemoveItemFromSlice(t *testing.T) {
	t.Run("If slice has item, item is removed", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		newList := utils.RemoveItemFromSlice2(list, 3, func(a, b int) bool { return a == b })
		assert.Equal(t, len(list)-1, len(newList))
	})
	t.Run("If slice has no item, slice is kept", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		newList := utils.RemoveItemFromSlice2(list, 13, func(a, b int) bool { return a == b })
		assert.Equal(t, len(list), len(newList))
	})
}
