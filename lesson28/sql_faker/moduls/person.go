package moduls

type Person struct {
	ID        string `json:"id"`
	Fist_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
}
