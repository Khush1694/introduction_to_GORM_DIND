package structs

// Response is implemented in all routes
type Response struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validations,omitempty"`
	Message     string       `json:"message,omitempty"`
	Messages    []string     `json:"messages,omitempty"`
	Data        interface{}  `json:"data,omitempty"`
}

// Validation is part of the Reponse struct
type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
