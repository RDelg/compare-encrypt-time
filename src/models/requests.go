package models

type Message struct {
	RequestId string            `json:"requestId,omitempty"`
	Caller    string            `json:"caller"`
	Session   string            `json:"sessionUser"`
	Context   map[string]string `json:"userDefinedContext"`
	Calls     [][]string        `json:"calls"`
}
