package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocarina/gocsv"
	ini "gopkg.in/ini.v1"

	"main/internal/afterrace"
	"main/internal/netkeiba"
	"main/internal/newrace"
	"main/pkg/config"
)

var Config config.ConfigList

type Compi struct {
	NEWID    string `csv:"レースID(新)"`
	COMPI    string `csv:"コンピ指数"`
	CompiNum string `csv:"コンピ順位"`
}

var compis []Compi

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Config = config.ConfigList{
		DbHost:        cfg.Section("db").Key("host").String(),
		DbPort:        cfg.Section("db").Key("port").MustInt(),
		DbName:        cfg.Section("db").Key("name").String(),
		DbUser:        cfg.Section("db").Key("user").String(),
		DbPassword:    cfg.Section("db").Key("password").String(),
		NetKeibaEmail: cfg.Section("netkeiba").Key("email").String(),
		NetKeibaPass:  cfg.Section("netkeiba").Key("password").String(),
	}
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.DbUser, Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName))
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	create_new_race(db)
	after_race(db)
}

// 新規レース登録 → 血統 → 調教
func create_new_race(db *sql.DB) {
	newrace.CreateNewRace("test.csv", db)
	newrace.CreateHorseData(db)
	newrace.TrainingData("2022-06-03", Config, db)

	defer db.Close()
}

// レース後データ
func after_race(db *sql.DB) {
	afterrace.SpeedIndex("2022-06-03", db)
	afterrace.UpdateHorseRace("all20220603.csv", db)
	afterrace.UpdateRap("rap20220603.csv", db)
	afterrace.WeatherSql("2022-06-03", db)
	afterrace.RaceType("hande.csv", 1, db)
	netkeiba.GetNetKeibaData("2022-06-03", db)
	netkeiba.CreateNetkeibaDetail(db)
	netkeiba.Start(db)
}
