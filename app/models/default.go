package models

type ResponseWithData[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Errors  any    `json:"error"`
}

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type PaginateResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta    Meta   `json:"meta"`
}

type Meta struct {
	Total    int `json:"total"`
	PerPage  int `json:"per_page"`
	LastPage int `json:"last_page"`
	CurrPage int `json:"curr_page"`
}
