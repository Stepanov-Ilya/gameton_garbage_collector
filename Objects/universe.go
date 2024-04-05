package Objects

type Planet struct {
	Name    string             `json:"name"`
	Garbage map[string][]int64 `json:"garbage"`
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
