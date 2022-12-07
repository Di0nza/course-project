package repository

import (
	courseProject "CourseProject"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

var goldArr [10]int = [10]int{20, 35, 75, 140, 290, 480, 800, 1250, 1875, 2800}
var powerPointArr [10]int = [10]int{20, 30, 50, 80, 130, 210, 340, 550, 890, 1440}
var limitsPP [10]int = [10]int{3740, 3720, 3690, 3640, 3560, 3430, 3220, 2880, 2330, 1440}

type BrawlerListPostgres struct {
	db *sqlx.DB
}

func NewBrawlerListPostgres(db *sqlx.DB) *BrawlerListPostgres {
	return &BrawlerListPostgres{db: db}
}

func (r *BrawlerListPostgres) Create(userId int, brawlersList courseProject.BrawlersList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int

	brawlersName := brawlersList.Name

	fmt.Println(brawlersList.CurrentLevel)
	level, err := strconv.Atoi(brawlersList.CurrentLevel)
	if err != nil {
		logrus.Printf("failed to parse int: %s", err.Error())
	}
	levelUp, err := strconv.Atoi(brawlersList.NewLevel)
	if err != nil {
		logrus.Printf("failed to parse int: %s", err.Error())
	}
	powerPoints, err := strconv.Atoi(brawlersList.AvailablePP)
	if err != nil {
		logrus.Printf("failed to parse int: %s", err.Error())
	}

	//brawlersName := brawlersList.Name
	//jsonDataLevel := []byte(string(brawlersList.CurrentLevel))
	//level := jsonDataLevel
	//jsonDataLevelUp := []byte(string(brawlersList.NewLevel))
	//levelUp := jsonDataLevelUp
	//jsonDataPP := []byte(string(brawlersList.AvailablePP))
	//powerPoints := jsonDataPP
	var flag = true

	if level < 1 || level > 10 {
		logrus.Printf("wrong  current level")
	}
	if level > levelUp || levelUp < 2 || levelUp > 11 {
		logrus.Printf("wrong new level")
	}

	switch level {
	case 1:
		if powerPoints > limitsPP[0] {
			flag = false
			break
		}
		flag = true
		break
	case 2:
		if powerPoints > limitsPP[1] {
			flag = false
			break
		}
		flag = true
		break
	case 3:
		if powerPoints > limitsPP[2] {
			flag = false
			break
		}
		flag = true
		break
	case 4:
		if powerPoints > limitsPP[3] {
			flag = false
			break
		}
		flag = true
		break
	case 5:
		if powerPoints > limitsPP[4] {
			flag = false
			break
		}
		flag = true
		break
	case 6:
		if powerPoints > limitsPP[5] {
			flag = false
			break
		}
		flag = true
		break
	case 7:
		if powerPoints > limitsPP[6] {
			flag = false
			break
		}
		flag = true
		break
	case 8:
		if powerPoints > limitsPP[7] {
			flag = false
			break
		}
		flag = true
		break
	case 9:
		if powerPoints > limitsPP[8] {
			flag = false
			break
		}
		flag = true
		break
	case 10:
		if powerPoints > limitsPP[9] {
			flag = false
			break
		}
		flag = true
		break
	default:
		flag = false
		break
	}
	if !flag {
		log.Printf("wrong amount of power points")
	}

	var goldSum int = 0
	var powerPointSum int = 0
	for i := level - 1; i < levelUp-1; i++ {
		goldSum += goldArr[i]
		powerPointSum += powerPointArr[i]
	}
	var redPointsP int
	var redPointsPCount int
	var redPointsG int
	var redPointsGCount int
	redPointsGCount = goldSum / 250
	if float64(goldSum%250) != 0 {
		redPointsGCount += 1
	}
	redPointsG = redPointsGCount * 100

	if powerPointSum > powerPoints {
		redPointsPCount = (powerPointSum - powerPoints) / 100
		if float64(powerPointSum%100) != 0 {
			redPointsPCount += 1
		}
		redPointsP = redPointsPCount * 60
		powerPointSum = powerPointSum - powerPoints
	} else {
		redPointsPCount = 0
		redPointsP = 0
		powerPointSum = 0
	}

	var allRedPoints = redPointsG + redPointsP

	createListQuery := fmt.Sprintf("INSERT INTO %s (brawlers_name, current_level, available_PP, new_level, gold, pp, cp_for_gold, cp_gold, cp_for_pp, cp_pp, cp_total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		brawlersTable)
	row := tx.QueryRow(createListQuery, brawlersName, level, powerPoints, levelUp, goldSum, powerPointSum, redPointsG, redPointsGCount*250, redPointsP, redPointsPCount*100, allRedPoints)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, lists_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BrawlerListPostgres) GetAll(userId int) ([]courseProject.BrawlersListCalc, error) {
	var lists []courseProject.BrawlersListCalc

	query := fmt.Sprintf("SELECT tl.id, tl.brawlers_name, tl.current_level, tl.available_PP, tl.new_level, tl.gold, tl.pp , tl.cp_for_gold, tl.cp_gold, tl.cp_for_pp, tl.cp_pp, tl.cp_total  FROM %s tl INNER JOIN %s ul on tl.id = ul.lists_id WHERE ul.user_id = $1",
		brawlersTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *BrawlerListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.lists_id AND ul.user_id=$1 AND ul.lists_id=$2",
		brawlersTable, usersListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
