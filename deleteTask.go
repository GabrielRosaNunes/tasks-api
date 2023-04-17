package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	delete,err := db.Query("DELETE FROM tasks WHERE id = ?",id)
	defer delete.Close()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,HttpResponse{
			Status: "error",
			Message: err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK,HttpResponse{
		Status: "success",
		Message: "Registro deletado com sucesso",
	})
}