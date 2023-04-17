package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func updateTask(c *gin.Context) {
	var task Tasks
	id := c.Param("id")

	if err := c.BindJSON(&task); err != nil {
		return
	}

	if message,ok := validTask(task); !ok {
		response := HttpResponse{
			Status: "error",
			Message: message,
		}
		c.IndentedJSON(http.StatusNotAcceptable,response)
		return
	}

	updateReturn, err := db.Query("UPDATE tasks SET descricao = ?, status = ? WHERE id = ?",task.Descricao,task.Status,id)
	defer updateReturn.Close()

	if err != nil {
		response := HttpResponse{
			Status: "error",
			Message: err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	response := HttpResponse{
		Status: "success",
		Message: "Registro atualizado com sucesso",
	}
	c.IndentedJSON(http.StatusOK,response)
}