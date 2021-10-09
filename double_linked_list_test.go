package reverse_double_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValuesEmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	assert.Equal(t, list.Values(), []int{})
}

func TestValuesOneElement(t *testing.T) {
	list := NewDoubleLinkedList(1)

	assert.Equal(t, list.Values(), []int{1})
}

func TestValuesMultipleElements(t *testing.T) {
	list := NewDoubleLinkedList(1, 2, 3, 4, 5)

	assert.Equal(t, list.Values(), []int{1, 2, 3, 4, 5})
}

func TestReverseEmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	assert.Equal(t, list.Reverse().Values(), []int{})
}

func TestReverseOneElement(t *testing.T) {
	list := NewDoubleLinkedList(1)

	assert.Equal(t, list.Reverse().Values(), []int{1})
}

func TestReverseMultipleElements(t *testing.T) {
	list := NewDoubleLinkedList(1, 2, 3, 4, 5)

	assert.Equal(t, list.Reverse().Values(), []int{5, 4, 3, 2, 1})
}
