package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTasks(c *gin.Context) {
	var tasks []Tasks

	results, err := db.Query("SELECT id, descricao, status FROM tasks")
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,tasks)
		return
	}

	for results.Next() {
		var task Tasks

		err = results.Scan(
			&task.Id,
			&task.Descricao,
			&task.Status,
		)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,tasks)
		}

		tasks = append(tasks,task)
	}

	c.IndentedJSON(http.StatusOK,tasks)
}