package Objects

type Ship struct {
	FuelUsed  int64                    `json:"fuelUsed"`
	Planet    Planet                   `json:"planet"`
	CapacityX int64                    `json:"capacityX"`
	CapacityY int64                    `json:"capacityY"`
	Garbage   map[string][]interface{} `json:"garbage"`
}

type TravelRequest struct {
	Planets []string `json:"planets"`
}

type TravelResponse struct {
	FuelDiff      int64                    `json:"fuelDiff"`
	PlanetDiffs   []PlanetDiffs            `json:"planetDiffs"`
	PlanetGarbage map[string][]interface{} `json:"planetGarbage"`
	ShipGarbage   map[string][]interface{} `json:"shipGarbage"`
}

type PlanetDiffs struct {
	From string `json:"from"`
	Fuel uint64 `json:"fuel"`
	To   string `json:"to"`
}
