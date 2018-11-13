package models

type Request struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   string `json:"status"`
	Role     string `json:"role"`
}

type Response struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
	Data    data   `json:"data"`
}

type data struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}
