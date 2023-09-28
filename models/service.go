package models

import "net/http"

type Service struct {
	Key       string
	Name      string
	Version   string
	Host      string
	Port      string
	Namespace string
	WebURL    string
	Env       string
}

type HTTPRoute struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type ApiResponse struct {
	//ProcessTime int64       `json:"processTime"`
	Success bool        `json:"success"`
	Status  int         `json:"statusCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
