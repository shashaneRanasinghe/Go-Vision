package responses

type ClassifyResponse struct {
	Labels []string `json:"labels,omitempty"`
	Error  string   `json:"error,omitempty"`
}
