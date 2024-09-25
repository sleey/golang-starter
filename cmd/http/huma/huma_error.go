package hc

type CustomHumaError struct {
	Status  int
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func (e *CustomHumaError) Error() string {
	return e.Message
}

func (e *CustomHumaError) GetStatus() int {
	return e.Status
}
