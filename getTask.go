package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTask(c *gin.Context) {
	var task Tasks
	id := c.Param("id")

	err := db.QueryRow("SELECT id, descricao,status FROM tasks WHERE id = ?",id).Scan(&task.Id,&task.Descricao,&task.Status)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,HttpResponse{
			Status: "error",
			Message:err.Error() ,
		})
		return
	}

	c.IndentedJSON(http.StatusOK,task)
}