package entities

type Item struct {
	ChrtId      int32  `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       uint   `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  uint   `json:"total_price"`
	NmId        uint   `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      uint   `json:"status"`
}
