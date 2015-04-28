package main

import (
	"encoding/json"
)

// Describes light status parameters.
type Status struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Queue  string `json:"queue"`
}

func (s Status) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{
		"id": s.Queue,
	}

	switch s.Status {
	case "red":
		raw["red"] = "0"
		raw["yellow"] = ""
		raw["green"] = ""
	case "yellow":
		raw["red"] = ""
		raw["yellow"] = "1"
		raw["green"] = ""
	case "green":
		raw["red"] = ""
		raw["yellow"] = ""
		raw["green"] = "2"
	case "":
		raw["red"] = ""
		raw["yellow"] = ""
		raw["green"] = ""
	}

	return json.Marshal(raw)
}
