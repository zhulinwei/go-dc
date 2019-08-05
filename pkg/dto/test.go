package dto

type Test struct {
	Age  int    `json:"age" binding:"required"`
	Name string `json:"name" binding:"required"`
}
