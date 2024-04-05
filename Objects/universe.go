package Objects

type Planet struct {
	Name    string                   `json:"name"`
	Garbage map[string][]interface{} `json:"garbage"`
}

type UniverseResponse struct {
	Name       string          `json:"name"`
	RoundName  string          `json:"roundName"`
	RoundEndIn int             `json:"roundEndIn"`
	Ship       Ship            `json:"ship"`
	Universe   [][]interface{} `json:"universe"`
	Attempt    int64           `json:"attempt"`
}

type Edge struct {
	From   string
	To     string
	Weight int
}

type ResponseCollect struct {
	Garbage map[string][]interface{} `json:"garbage"`
	Leaved  []string                 `json:"leaved"`
}

type ResponseReset struct {
	Success bool `json:"success"`
}

type Round struct {
	StartAt     string `json:"startAt"`
	EndAt       string `json:"endAt"`
	IsCurrent   bool   `json:"isCurrent"`
	Name        string `json:"name"`
	PlanetCount int64  `json:"planetCount"`
}

type ResponseRound struct {
	Rounds []Round `json:"rounds"`
}
