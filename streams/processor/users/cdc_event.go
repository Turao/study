package users

type CDCEvent struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Before *User `json:"before"` // ! Whether or not this field is available is highly dependent on the REPLICA IDENTITY setting for each table
	After  *User `json:"after"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}
