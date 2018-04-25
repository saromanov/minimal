package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// All of the data is optional.
type Error struct {
	// Id of this particular occurrence of the problem.
	ID string `json:"id,omitempty"`
	// HTTP Status code of response, as a string.
	Status string `json:"status,omitempty"`
	// Application-specific status code.
	Code string `json:"code,omitempty"`
	// Title of the error
	Title string `json:"title,omitempty"`
	// Advanced details about error. Optional
	Detail string `json:"detail,omitempty"`
	// Source of the error. Optional
	Source map[string]string `json:"source,omitempty"`
	// Meta of the error. Optional
	Meta map[string]string `json:"meta,omitempty"`
}

func WriteBasicError(w http.ResponseWriter, status, title, detail string) {
	e := Error{Status: status, Title: title, Detail: detail}
	marshal(w, e)
}

// bardRequest provides setting of 400 error on response
func badRequest(w http.ResponseWriter, text, description string) {
	w.WriteHeader(http.StatusBadRequest)
	WriteBasicError(w, "400", text, description)
}

func marshal(w http.ResponseWriter, obj interface{}) {
	res, _ := json.Marshal(obj)
	fmt.Fprintln(w, string(res))
}
