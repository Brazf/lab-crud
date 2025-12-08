package models

type user struct {
	ID    uint32 `json: "id"`
	name  string `json: "name"`
	email string `json: "email`
}
