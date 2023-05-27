package models

type Shards struct {
	Total     int `json:"total"`
	Succesful int `json:"succesful"`
	Skipped   int `json:"skipped"`
	Failed    int `json:"failed"`
}

type Total struct {
	Value int `json:"value"`
}

type Source struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	TimeStamp string  `json:"@timestamp"`
	Source    Email   `json:"_source"`
}

type Hits struct {
	Total    Total    `json:"total"`
	MaxScore float64  `json:"max_score"`
	Hits     []Source `json:"hits"`
}

type ZincSearch struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}
