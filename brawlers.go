package courseProject

type BrawlersList struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	CurrentLevel string `json:"currentLevel"`
	NewLevel     string `json:"newLevel"`
	AvailablePP  string `json:"availablePP"`
}

type BrawlersListCalc struct {
	Id               int    `json:"id" db:"id"`
	Name             string `json:"name" db:"brawlers_name"`
	CurrentLevel     int    `json:"currentLevel" db:"current_level"`
	AvailablePP      int    `json:"availablePP" db:"available_pp"`
	NewLevel         int    `json:"newLevel" db:"new_level"`
	Gold             int    `json:"gold" db:"gold"`
	Pp               int    `json:"powerPoints" db:"pp"`
	CpForGold        int    `json:"cpForGold" db:"cp_for_gold"`
	CpGold           int    `json:"cpGold" db:"cp_gold"`
	CpForPowerPoints int    `json:"cpForPowerPoints" db:"cp_for_pp"`
	CpPowerPoints    int    `json:"cpPowerPoints" db:"cp_pp"`
	CpTotal          int    `json:"cpTotal" db:"cp_total"`
}
