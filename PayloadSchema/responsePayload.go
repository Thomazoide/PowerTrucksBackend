package payloadschema

type ResponsePayload struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Error      bool   `json:"error"`
}
