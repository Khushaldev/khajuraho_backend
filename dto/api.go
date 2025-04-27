package dto

type SendSuccess struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Successfull"`
	Data    interface{} `json:"data,omitempty"`
}

type SendError struct {
	Success    bool     `json:"success" example:"false"`
	Message    string   `json:"message" example:"Failed"`
	Errors     []string `json:"errors,omitempty" example:"Unexpected issue"`
	StackTrace string   `json:"stack_trace,omitempty" example:"Error stack trace"`
}
