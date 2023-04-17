package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func insertTask(c *gin.Context) {
	var newTask Tasks

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	if message,ok := validTask(newTask); !ok {
		response := HttpResponse{
			Status: "error",
			Message: message,
		}
		c.IndentedJSON(http.StatusNotAcceptable,response)
		return
	}

	insert, err := db.Query("INSERT INTO tasks (descricao,status) VALUES (?,?)",newTask.Descricao,newTask.Status)
	defer insert.Close()

	if err != nil {
		fmt.Println("Erro para executar insert into ",err.Error())
		response := HttpResponse{
			Status: "error",
			Message: "Erro ao tentar inserir no banco: "+err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError,response)
		return
	}

	

	response := HttpResponse{
		Status: "success",
		Message: "Incluído com sucesso",
	}

	c.IndentedJSON(http.StatusOK,response)
}

func validTask(task Tasks) (message string, ok bool) {
	switch task.Status {
		case 
			"Não Iniciado",
			"Em andamento",
			"Aguardando",
			"Concluído":
			return "",true
		default:
			return "Opção de status não é aceita, somente é aceito: Não Iniciado, Em andamento, Aguardando ou Concluído", false
	}
}