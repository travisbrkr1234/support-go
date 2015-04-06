package main

// Describes light status parameters.
type Status struct {
	ID        int    `json:"id"`
	Status string `json:"status"`
	Queue  string `json:"queue"`
}
