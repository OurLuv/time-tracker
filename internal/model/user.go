package model

type User struct {
	Id             int    `json:"id"`
	PassportNumber string `json:"passport_number"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"addres"`
}
