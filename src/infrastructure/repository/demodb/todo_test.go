package demodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	NewTodo()
	t.Skip("skipping test")
}
func TestGetTodoList(t *testing.T) {
	todo := NewTodo()
	_, err := todo.GetTodoList()
	assert.Nil(t, err)
}
