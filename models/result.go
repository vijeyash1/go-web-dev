package models

type Result struct {
	ID         string `json:"id"`
	RunnerId   string `json:"runner_id"`
	RaceResult string `json:"race_result"`
	Location   string `json:"location"`
	Position   int    `json:"position,omitempty"`
	Year       int    `json:"year"`
}
