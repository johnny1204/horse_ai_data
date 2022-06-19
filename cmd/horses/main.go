package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	ini "gopkg.in/ini.v1"

	"main/internal/data"
	"main/internal/netkeiba"
	"main/internal/newrace"
	"main/pkg/config"
	"main/pkg/util"
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
	update_horse(db)
	get_model_date("2022-06-03", db)
}

// 新規レース登録 → 血統 → 調教
func create_new_race(db *sql.DB) {
	newrace.CreateNewRace("test.csv", db)
	newrace.CreateHorseData(db)
	newrace.TrainingData("2022-06-03", Config, db)

	defer db.Close()
}

func update_horse(db *sql.DB) {
	file, _ := os.Open("./csv/horses.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	for _, r := range records[:] {
		var color_id int
		db.QueryRow("SELECT id FROM horse_colors WHERE name = ?", r[3]).Scan(&color_id)
		var producer_id string
		db.QueryRow("SELECT id FROM producers WHERE name = ?", r[5]).Scan(&producer_id)

		if producer_id == "" {
			stmt, err := db.Prepare(`INSERT INTO producers(name) VALUES(?) RETURNING id;`)
			defer stmt.Close()

			if err != nil {
				log.Fatalf("Abend create producers :%v", err)
			}

			stmt.QueryRow(r).Scan(&producer_id)
		}

		b := strings.Split(r[2], ".")
		born := b[0] + "-" + fmt.Sprintf("%02s", b[1]) + "-" + fmt.Sprintf("%02s", b[2])
		_, err := db.Exec(
			`UPDATE horses SET name = ?, born = ?, color_id = ?, sireline = ?, producer_id = ? WHERE id = ?`,
			r[1],
			born,
			color_id,
			util.NilOrString(r[5]),
			util.NilOrString(producer_id),
			r[0],
		)

		if err != nil {
			fmt.Println(r[0])
			log.Fatalf("Abend horse_detail_sql:%v", err)
		}
	}
}

func get_model_date(date string, db *sql.DB) {
	data.GenerateData(date, db)
}

// レース後データ
func after_race(db *sql.DB) {
	// afterrace.SpeedIndex("2022-06-03", db)
	// afterrace.UpdateHorseRace("test.csv", db)
	// afterrace.UpdateRap("raptest.csv", db)
	// afterrace.WeatherSql("2022-06-03", db)
	// afterrace.RaceType("hande.csv", 1, db)
	// netkeiba.GetNetKeibaData("test", db)
	netkeiba.CreateNetkeibaDetail(db)
	netkeiba.Start(db)
}

func compi_update() {
	com, _ := ioutil.ReadFile("./csv/horse_compi.csv")
	_ = csvutil.Unmarshal(com, &compis)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Config.DbUser, Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName))
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	for _, h := range compis[0:] {
		_, err := db.Exec(
			`UPDATE horse_results set compi = ?, compi_num = ? where id = ?`,
			h.COMPI,
			h.CompiNum,
			h.NEWID,
		)

		if err != nil {
			fmt.Println(h)
			log.Fatalf("compi_update db.Exec error err:%v", err)
		}
	}

	defer db.Close()
}
