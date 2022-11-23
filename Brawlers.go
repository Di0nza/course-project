package courseProject

type BrawlersList struct {
	UserId       int    `json:"userId"`
	Name         string `json:"name"`
	CurrentLevel int    `json:"currentLevel"`
	AvailablePP  int    `json:"availablePP"`
}
