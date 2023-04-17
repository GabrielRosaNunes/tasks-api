package main

type Pessoa struct {
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
}

type Tasks struct {
	Id string `json:"id"`
	Descricao string `json:"descricao"`
	Status string `json:"status"`
}

type HttpResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}