package protocol

type Status struct {
	Success bool `json:"success"`
}

type Auth struct {
	Auth string `json:"auth"`
}
