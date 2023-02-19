package models

const (
	IN_WORK   = 1
	ON_REVIEW = 2
	DONE      = 3
)

type Status struct {
	ID    int    `json:"id"`
	Names string `json:"name"`
}
