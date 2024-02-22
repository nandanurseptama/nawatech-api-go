package responses

type Response[T comparable] struct {
	RequestId string  `json:"request_id"`
	Status    string  `json:"status"`
	Message   string  `json:"message,omitempty"`
	Data      *T      `json:"data,omitempty"`
	Error     *string `json:"error,omitempty"`
}
