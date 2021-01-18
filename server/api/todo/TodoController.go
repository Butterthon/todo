package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo 構造体
type Todo struct {
	ID   string
	Text string
}

// GetAll TODOリストを取得
func GetAll(context *gin.Context) {
	response := make([]Todo, 0, 10)
	todo := Todo{}
	todo.ID = "1"
	todo.Text = "ご飯を食べる"
	response = append(response, todo)
	context.JSON(http.StatusOK, response)
}
