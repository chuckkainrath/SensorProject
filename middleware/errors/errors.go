package errors

type AppError struct {
	Code    int    `json:",omitempty"` //makes empty values not show
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
