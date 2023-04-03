package models

type BigQueryRemoteFunctionResponse struct {
	Replies      []string `json:"replies,omitempty"`
	ErrorMessage string   `json:"errorMessage,omitempty"`
}
