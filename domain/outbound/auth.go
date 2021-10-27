package outbound

type Authentication struct {
	UUID     string            `json:"uuid,omitempty"`
	Login    *ResponseLogin    `json:"login,omitempty"`
	Logout   *ResponseLogout   `json:"logout,omitempty"`
	Token    *ResponseToken    `json:"token,omitempty"`
	Register *ResponseRegister `json:"register,omitempty"`
}

type ResponseLogin struct {
	User *User `json:"user,omitempty"`
}

type ResponseLogout struct {
}

type ResponseToken struct {
	Jwt string `json:"jwt,omitempty"`
}

type ResponseRegister struct {
	User *User `json:"user,omitempty"`
}
