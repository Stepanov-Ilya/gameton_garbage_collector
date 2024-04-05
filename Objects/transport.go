package Objects

type Ship struct {
	FuelUsed  int64              `json:"fuelUsed"`
	Planet    Planet             `json:"planet"`
	CapacityX int64              `json:"capacityX"`
	CapacityY int64              `json:"capacityY"`
	Garbage   map[string][]int64 `json:"garbage"`
}
