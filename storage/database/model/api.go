package model

type ResponseHttp struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseLinkHttp struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Links   []*Link     `json:"links,omitempty"`
	Data    interface{} `json:"data"`
}

type Link struct {
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"`
	Link   string `json:"link,omitempty"`
}
