package types

type Client struct {
	Email     string `json: "email"`
	Surname   string `json: "surname"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
	Password  string `json: "password"`
}
