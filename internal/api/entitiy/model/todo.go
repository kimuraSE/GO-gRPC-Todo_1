package model

type Todo struct {
	ID     uint
	Title  string
	UserID uint
}

type TodoResponse struct {
	Id    uint
	Title string
}
