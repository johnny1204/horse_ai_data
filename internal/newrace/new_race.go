package newrace

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jszwec/csvutil"

	"main/internal/domain"
	"main/pkg/util"
)

var races []domain.NewRace

func CreateNewRace(filename string, db *sql.DB) {
	new_race(filename, db)
	horse_detail_sql(filename, db)
}

func new_race(filename string, db *sql.DB) {
	newrace, _ := ioutil.ReadFile("./csv/" + filename)
	_ = csvutil.Unmarshal(newrace, &races)

	places := map[string]int{"札幌": 1, "函館": 2, "福島": 3, "新潟": 4, "東京": 5, "中山": 6, "中京": 7, "京都": 8, "阪神": 9, "小倉": 10}
	courses := map[string]int{"芝": 1, "ダ": 2}
	grades := map[string]int{"新馬": 1, "未勝利": 2, "1勝": 3, "2勝": 4, "3勝": 5, "ｵｰﾌﾟﾝ": 6, "OP(L)": 7, "重賞": 8, "Ｇ３": 9, "ＪＧ３": 9, "Ｇ２": 10, "ＪＧ２": 10, "Ｇ１": 11, "ＪＧ１": 11}
	seasons := map[string]int{"01": 4, "02": 4, "03": 1, "04": 1, "05": 1, "06": 2, "07": 2, "08": 2, "09": 3, "10": 3, "11": 3, "12": 4}

	for _, rap := range races[0:] {
		date := rap.NEWID[0:4] + "-" + rap.NEWID[4:6] + "-" + rap.NEWID[6:8]
		_, err := db.Exec(`INSERT IGNORE INTO races (
			id, place_id, h_num, course, distance, grade, date, season_id, start_time
			) VALUES (
				?, ?, ?, ?, ?, ?, ?, ?, ?
			)`,
			rap.NEWID[:16], places[rap.PLACE], rap.H_COUNT, courses[rap.COURSE], rap.DISTANCE, grades[rap.GRADE],
			date, seasons[rap.NEWID[4:6]], rap.START_TIME,
		)

		if err != nil {
			fmt.Println(rap)
			log.Fatalf("error err:%v", err)
		}
	}

	gender := map[string]int{"牡": 1, "牝": 2, "セ": 3}
	affiliation := map[string]int{"(美)": 1, "(栗)": 2, "[外]": 3, "[地]": 4}
	for _, race := range races[0:] {
		fmt.Println(race.NEWID)
		w, b := util.Weight(race.WEIGHT)
		multi := race.IS_MULTIPLE
		is_multi := false
		if multi != "" {
			is_multi = true
		}

		rb := true
		if race.DEL != "" {
			rb = false
		}

		var jockey_id string
		db.QueryRow("SELECT id FROM jockies WHERE origin_id = ?", race.JOCKEY_ID).Scan(&jockey_id)

		_, err := db.Exec(`INSERT INTO horse_results(
			id, prev_id, race_id, jockey_id, affiliation_id, trainer_id, gender, age,
			weight, is_loss_jockey, horse_id, multiple,
			waku, h_num, career, interval_week, after_break, owner, result_except
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?
		) ON DUPLICATE KEY UPDATE id = ?`,
			race.NEWID,
			util.NilOrString(race.PASSED_NEW_ID),
			race.NEWID[:16],
			jockey_id,
			affiliation[race.AFFILIATION],
			race.TRAINER_ID,
			gender[race.SEX],
			race.AGE,
			w,
			b,
			race.HORSE_ID,
			is_multi,
			race.WAKU,
			race.UMABAN,
			util.NilOrInt(race.CAREER),
			util.NilOrFloat(race.INTERVAL_WEEK),
			util.NilOrInt(race.WEEK),
			race.OWNER,
			rb,
			race.NEWID,
		)

		if err != nil {
			fmt.Println(race)
			log.Fatalf("race error err:%v", err)
		}
	}

	defer db.Close()
}

func horse_detail_sql(filename string, db *sql.DB) {
	race_csv, _ := ioutil.ReadFile("./csv/" + filename)
	_ = csvutil.Unmarshal(race_csv, &races)

	for _, race := range races[0:] {
		b := false
		if race.BRINKER == "B" {
			b = true
		}

		_, err := db.Exec(
			`UPDATE horse_results SET brinker = ? WHERE id = ?`,
			b,
			race.NEWID,
		)

		var horse_id string
		db.QueryRow("SELECT horse_id FROM horse_results WHERE id = ?", race.NEWID).Scan(&horse_id)
		var color_id int
		db.QueryRow("SELECT id FROM horse_colors WHERE name = ?", race.COLOR).Scan(&color_id)
		var producer_id string
		db.QueryRow("SELECT id FROM producers WHERE name = ?", race.PRODUCER).Scan(&producer_id)

		if producer_id == "" {
			stmt, err := db.Prepare(`INSERT INTO producers(name) VALUES(?) RETURNING id;`)
			defer stmt.Close()

			if err != nil {
				log.Fatalf("Abend create producers :%v", err)
			}

			stmt.QueryRow(race.PRODUCER).Scan(&producer_id)
		}

		born := race.BORN_YEAR + "-" + race.BORN_MONTH[:2] + "-" + race.BORN_MONTH[2:]
		_, err = db.Exec(
			`UPDATE horses SET born = ?, color_id = ?, sireline = ?, name = ?, producer_id = ? WHERE id = ?`,
			born,
			color_id,
			util.NilOrString(race.SIRE_TYPE),
			race.NAME,
			util.NilOrString(producer_id),
			horse_id,
		)

		if err != nil {
			fmt.Println(race)
			log.Fatalf("Abend horse_detail_sql:%v", err)
		}
	}

	db.Exec(`update horses set sire_id = (select id from sires where name = horses.sireline) where sire_id is null`)
}
