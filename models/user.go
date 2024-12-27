package models

type User struct {
	ID		int64	`json:"id"`
	Name 	string	`json:"name"`
	Age		int16	`json:"age"`
	Salary	float32	`json:"salary"`
} 