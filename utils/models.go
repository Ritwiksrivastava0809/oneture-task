package utils

type Record struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

type Batch struct {
	Records []Record `json:"records"`
}
