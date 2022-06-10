package afterrace

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jszwec/csvutil"
	"github.com/sclevine/agouti"

	"main/internal/domain"
	"main/pkg/util"
)

type RaceDetail struct {
	RACE_ID string `csv:"レースID(新/馬番無)"`
	WIN     string `csv:"馬番"`
	TANSHO  string `csv:"単勝配当"`
	UMAREN  string `csv:"馬連"`
	UMATAN  string `csv:"馬単"`
	RENPUKU string `csv:"３連複"`
	RENTAN  string `csv:"３連単"`
}

type SpeedUrl struct {
	RaceId string
	UrlId  string
}

type Rap struct {
	NEWID              string `csv:"レースID(新)"`
	Weather            string `csv:"天候"`
	Condition          string `csv:"馬場状態"`
	Count              string `csv:"頭数"`
	Last5F             string `csv:"後5F"`
	Start5F            string `csv:"前5F"`
	Last4F             string `csv:"前4F"`
	Start4F            string `csv:"後4F"`
	Last3F             string `csv:"前3F"`
	Start3F            string `csv:"後3F"`
	Last2F             string `csv:"後2F"`
	Start2F            string `csv:"前2F"`
	Last1F             string `csv:"後1F"`
	Start1F            string `csv:"前1F"`
	Fast3F             string `csv:"最速3F"`
	RPCI               string `csv:"RPCI"`
	PCI3               string `csv:"PCI3"`
	FirstFact          string `csv:"1着決手"`
	SecondFact         string `csv:"2着決手"`
	ThirdFact          string `csv:"3着決手"`
	FirstFirstCorner   string `csv:"1着1角"`
	FirstSecondCorner  string `csv:"1着2角"`
	FirstThirdCorner   string `csv:"1着3角"`
	FirstFourthCorner  string `csv:"1着4角"`
	SecondFirstCorner  string `csv:"2着1角"`
	SecondSecondCorner string `csv:"2着2角"`
	SecondThirdCorner  string `csv:"2着3角"`
	SecondFourthCorner string `csv:"2着4角"`
	ThirdFirstCorner   string `csv:"3着1角"`
	ThirdSecondCorner  string `csv:"3着2角"`
	ThirdThirdCorner   string `csv:"3着3角"`
	ThirdFourthCorner  string `csv:"3着4角"`
	Ave3F              string `csv:"Ave-3F"`
	Ave3F3             string `csv:"A-3F*3"`
	Ave3FAll           string `csv:"A-3F全"`
	Ave3FDiff          string `csv:"A-3F差"`
	Ave3FO3            string `csv:"3F位*3"`
	Rap1               string `csv:"Lap1"`
	Rap2               string `csv:"Lap2"`
	Rap3               string `csv:"Lap3"`
	Rap4               string `csv:"Lap4"`
	Rap5               string `csv:"Lap5"`
	Rap6               string `csv:"Lap6"`
	Rap7               string `csv:"Lap7"`
	Rap8               string `csv:"Lap8"`
	Rap9               string `csv:"Lap9"`
	Rap10              string `csv:"Lap10"`
	Rap11              string `csv:"Lap11"`
	Rap12              string `csv:"Lap12"`
	Rap13              string `csv:"Lap13"`
	Rap14              string `csv:"Lap14"`
	Rap15              string `csv:"Lap15"`
	Rap16              string `csv:"Lap16"`
	Rap17              string `csv:"Lap17"`
	Rap18              string `csv:"Lap18"`
	Rap19              string `csv:"Lap19"`
	Rap20              string `csv:"Lap20"`
}

type Haitou struct {
	RACE_ID  string `csv:"レースID(新/馬番無)"`
	UMABAN1  string `csv:"馬番_1"`
	UMABAN2  string `csv:"馬番_2"`
	UMABAN3  string `csv:"馬番_3"`
	UMABAN4  string `csv:"馬番_4"`
	FUKUSHO1 string `csv:"複勝配当_1"`
	FUKUSHO2 string `csv:"複勝配当_2"`
	FUKUSHO3 string `csv:"複勝配当_3"`
	FUKUSHO4 string `csv:"複勝配当_4"`
}

var races []domain.Race
var raps []Rap
var details []RaceDetail
var ht []Haitou

var rapMap = make(map[string]map[string]interface{})

func SpeedIndex(date string, db *sql.DB) {
	rows, _ := db.Query(`SELECT id, CONCAT(SUBSTRING(id, 3, 2), SUBSTRING(id, 9, 8)) AS url_id FROM races where id AND date >= '` + date + `'`)
	defer rows.Close()

	var urls = make([]SpeedUrl, 0)
	for rows.Next() {
		var url SpeedUrl
		var id string
		var url_id string
		rows.Scan(&id, &url_id)

		url.RaceId = id
		url.UrlId = url_id

		urls = append(urls, url)
	}

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",             // headlessモードの指定
			"--window-size=1280,800", // ウィンドウサイズの指定
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Abend SpeedIndex Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Abend SpeedIndex Failed to open page:%v", err)
	}

	for _, url := range urls {
		u := "http://jiro8.sakura.ne.jp/index.php?code=" + url.UrlId
		fmt.Println(u)
		page.Navigate(u)

		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		contentDom, _ := goquery.NewDocumentFromReader(readerCurContents)

		h_num := contentDom.Find(".c1 > tbody > tr:nth-child(14) > td").Length() - 1

		for i := h_num; i > 0; i-- {
			num := h_num - i + 1
			str := fmt.Sprintf("%02d", num)

			preceding_index := ""
			pace_index := ""
			last3f_index := ""
			speed_index := ""
			id := url.RaceId + str
			for j := 14; j <= 17; j++ {
				elem := contentDom.Find(".c1 > tbody > tr:nth-child(" + strconv.Itoa(j) + ") > td:nth-child(" + strconv.Itoa(i) + ")")
				idx := elem.Text()
				if j == 14 {
					preceding_index = idx
				} else if j == 15 {
					pace_index = idx
				} else if j == 16 {
					last3f_index = idx
				} else {
					speed_index = idx
				}
			}

			_, err := db.Exec(
				`UPDATE horse_results set preceding_index = ?, pace_index = ?, last3f_index = ?, speed_index = ? where id = ?`,
				preceding_index,
				pace_index,
				last3f_index,
				speed_index,
				id,
			)

			if err != nil {
				log.Fatalf("Abend SpeedIndex:%v", err)
			}
		}

		time.Sleep(1 * time.Second)
	}

	defer db.Close()
}

func UpdateHorseRace(filename string, db *sql.DB) {
	var err error

	race_csv, _ := ioutil.ReadFile("./csv/" + filename)
	_ = csvutil.Unmarshal(race_csv, &races)

	gender := map[string]int{"牡": 1, "牝": 2, "セ": 3}
	affiliation := map[string]int{"(美)": 1, "(栗)": 2, "[外]": 3, "[地]": 4}
	race_type := map[string]int{"逃げ": 1, "先行": 2, "中団": 3, "後方": 4, "ﾏｸﾘ": 5}
	weathers := map[string]int{"晴": 1, "曇": 2, "雨": 3, "小雨": 4, "小雪": 5, "雪": 6}
	conditions := map[string]int{"良": 1, "稍": 2, "重": 3, "不": 4}

	for _, race := range races[0:] {
		db.Exec(`UPDATE races SET weather = ?, cond = ?, PCI =? WHERE id = ?`,
			weathers[strings.TrimSpace(race.Weather)],
			conditions[race.Condition],
			race.PCI,
			race.RACE_ID,
		)
	}

	for _, race := range races[0:] {
		w, b := util.Weight(race.WEIGHT)
		multi := race.IS_MULTIPLE

		r := race.RESULT
		var re *int = nil
		rb := 1
		if r == "消" || r == "外" {
			re = nil
			rb = 0
		} else if r == "止" {
			re = nil
		} else {
			res := util.ZenNumTohan(r)
			re = &res
		}

		m := false
		if multi != "" {
			m = true
		}

		var rd string = ""
		if race.DIFFERENCE != "----" {
			rd = race.DIFFERENCE
		}

		var rt *int = nil
		if race.RaceType != "" {
			rte := race_type[race.RaceType]
			rt = &rte
		}

		borns := strings.Split(race.BORN, ".")
		born := borns[0] + "-" + strings.TrimSpace(borns[1]) + "-" + strings.TrimSpace(borns[2])
		fmt.Println(race.HORSE_ID)
		_, err = db.Exec(
			`UPDATE horses SET born = ? WHERE id = ?`,
			born,
			race.HORSE_ID,
		)

		_, err = db.Exec(`
				UPDATE horse_results SET
				prev_id= ?, race_id= ?, result= ?, result_except= ?, jockey_id= ?, multiple= ?, affiliation_id= ?, 
				trainer_id= ?, gender= ?, age= ?, weight= ?, is_loss_jockey= ?, horse_id= ?, 
				body_weight= ?, body_weight_in_de= ?, popular= ?, odds= ?, speed = ?, fukusho_share = ?,
				waku= ?, h_num= ?, time= ?, difference= ?, courner1= ?, courner2= ?, courner3= ?, courner4= ?, last3f= ?, last3f_num= ?,
				prize= ?, add_prize= ?, career= ?, interval_week= ?, after_break= ?, ave_last3f_speed= ?, ave_first3f_speed= ?, ave_speed= ?,
				ave_1f_time= ?, ave3f= ?, race_type= ?, owner = ?, correct_time = ? WHERE id = ?`,
			util.NilOrString(race.PASSED_NEW_ID),
			race.RACE_ID,
			re,
			rb,
			race.JOCKEY_ID,
			m,
			affiliation[race.AFFILIATION],
			race.TRAINER_ID,
			gender[race.SEX],
			race.AGE,
			w,
			b,
			race.HORSE_ID,
			util.NilOrString(race.BODY_WEIGHT),
			util.NilOrInt(race.INCREASE_DECREASE),
			util.NilOrInt(race.POPULAR),
			util.NilOrFloat(race.ODDS),
			// util.NilOrInt(race.COMPI),
			// util.NilOrInt(race.CompiNum),
			util.NilOrInt(race.Zi),
			util.NilOrFloat(race.FUKUSHO),
			race.WAKU,
			race.UMABAN,
			util.TimeToSeconds(race.TIME),
			util.NilOrString(rd),
			util.NilOrString(race.COURNER_1),
			util.NilOrString(race.COURNER_2),
			util.NilOrString(race.COURNER_3),
			util.NilOrString(race.COURNER_4),
			util.NilOrFloat(race.LAST_3F),
			util.NilOrFloat(race.LAST_3F_NUM),
			util.ZenNumTohan(race.PRIZE),
			util.NilOrFloat(strings.TrimSpace(race.ADD_PRIZE)),
			race.CAREER,
			util.NilOrFloat(race.WEEK),
			util.NilOrFloat(race.INTERVAL_WEEK),
			util.NilOrFloat(race.AvgLast3FSpeed),
			util.NilOrFloat(race.AvgFirst3FSpeed),
			util.NilOrFloat(race.AvgSpeed),
			util.NilOrFloat(race.Avg1FTime),
			util.NilOrFloat(race.Avg3F),
			rt,
			race.OWNER,
			util.NilOrInt(race.CorrectTime),
			race.NEWID,
		)

		if err != nil {
			fmt.Println(race)
			log.Fatalf("insertUser db.Exec error err:%v", err)
		}
	}

	defer db.Close()
}

// ラップ
func UpdateRap(filename string, db *sql.DB) {
	rapcsv, _ := ioutil.ReadFile("./csv/" + filename)
	_ = csvutil.Unmarshal(rapcsv, &raps)

	for _, rap := range raps {
		rapMap[rap.NEWID] = util.StructToMap(rap)
	}

	weathers := map[string]int{"晴": 1, "曇": 2, "雨": 3, "小雨": 4, "小雪": 5, "雪": 6}
	conditions := map[string]int{"良": 1, "稍": 2, "重": 3, "不": 4}
	facts := map[string]int{"逃": 1, "先": 2, "差": 3, "追": 4, "マ": 5}

	for _, rap := range raps[1:] {
		r := util.StructToMap(rap)
		_, err := db.Exec(`UPDATE races SET weather= ?,cond= ?,h_num= ?,
							last_5f= ?, last_4f= ?, last_3f= ?, last_2f= ?, last_1f= ?,
							start_5f= ?, start_4f= ?, start_3f= ?, start_2f= ?, start_1f= ?,
							fast_3f= ?, RPCI= ?, PCI3= ?, first_fact= ?, second_fact= ?, third_fact= ?,
							ave_3f= ?, a_3f_3= ?, a_3f_all= ?, a_3f_diff= ?, a_3f_o3= ?
							WHERE id = ?
		`,
			weathers[strings.TrimSpace(rap.Weather)], conditions[rap.Condition], rap.Count,
			rap.Last5F, rap.Start5F, rap.Last4F, rap.Start4F, rap.Last3F,
			rap.Start3F, rap.Last2F, rap.Start2F, rap.Last1F, rap.Start1F, rap.Fast3F,
			rap.RPCI, rap.PCI3, facts[rap.FirstFact], facts[rap.SecondFact], facts[rap.ThirdFact],
			rap.Ave3F, rap.Ave3F3, rap.Ave3FAll, rap.Ave3FDiff,
			util.NilOrFloat(rap.Ave3FO3),
			rap.NEWID,
		)

		for i := 1; i <= 20; i++ {
			s := strconv.Itoa(i)
			if s == "" {
				db.Exec("UPDATE races SET rap" + s + "= null")
			} else {
				db.Exec("UPDATE races SET rap" + s + "=" + r["Rap"+s].(string) + ` where id = "` + rap.NEWID + `"`)
			}
		}

		if err != nil {
			fmt.Println(rap.NEWID)
			log.Fatalf("insertUser db.Exec error err:%v", err)
		}
	}
	defer db.Close()
}

// 配当結果
func UpdateHaito(filename string, db *sql.DB) {
	race_d, _ := ioutil.ReadFile("./csv/" + filename)
	_ = csvutil.Unmarshal(race_d, &details)

	for _, race := range details[0:] {
		_, err := db.Exec(
			`UPDATE races SET tansho = ?, umaren = ?, umatan = ?, sanren_fuku = ?, sanren_tan = ? WHERE id = ?`,
			race.TANSHO, race.UMAREN, race.UMATAN, race.RENPUKU,
			util.NilOrString(race.RENTAN), race.RACE_ID,
		)
		if err != nil {
			fmt.Println(race)
			log.Fatalf("Abend haitou:%v", err)
		}
	}

	fukusho, _ := ioutil.ReadFile("./fukusho20220904_fukusho.csv")
	_ = csvutil.Unmarshal(fukusho, &ht)
	for _, h := range ht[0:] {
		_, err := db.Exec(
			`UPDATE races set first = ?, second = ?, third = ?, 
			fourth = ?, fukusho_1 = ?, fukusho_2 = ?, fukusho_3 = ?, fukusho_4 = ?
			where id = ?`,
			h.UMABAN1, h.UMABAN2, util.NilOrString(h.UMABAN3), util.NilOrString(h.UMABAN4),
			h.FUKUSHO1, h.FUKUSHO2, util.NilOrString(h.FUKUSHO3), util.NilOrString(h.FUKUSHO4),
			h.RACE_ID,
		)

		if err != nil {
			fmt.Println(h)
			log.Fatalf("insertUser db.Exec error err:%v", err)
		}
	}

	defer db.Close()
}

func UpdateFukusho(filename string, db *sql.DB) {
	fukusho, _ := ioutil.ReadFile("./" + filename)
	_ = csvutil.Unmarshal(fukusho, &ht)

	for _, h := range ht[0:] {
		_, err := db.Exec(`
			UPDATE races set first = ?, second = ?, third = ?, 
			fourth = ?, fukusho_1 = ?, fukusho_2 = ?, fukusho_3 = ?, fukusho_4 = ?
			where id = ?`,
			h.UMABAN1, h.UMABAN2, util.NilOrString(h.UMABAN3), util.NilOrString(h.UMABAN4),
			h.FUKUSHO1, h.FUKUSHO2, util.NilOrString(h.FUKUSHO3), util.NilOrString(h.FUKUSHO4),
			h.RACE_ID,
		)

		if err != nil {
			fmt.Println(h)
			log.Fatalf("Abend fukusho_update:%v", err)
		}
	}

	defer db.Close()
}

func WeatherSql(target_date string, db *sql.DB) {
	blocks := map[int]string{1: "47412", 2: "47430", 3: "47595", 4: "47604", 5: "1133", 6: "1236", 7: "47636", 8: "47759", 9: "47770", 10: "47807"}
	precs := map[int]string{1: "14", 2: "23", 3: "36", 4: "54", 5: "44", 6: "45", 7: "51", 8: "61", 9: "63", 10: "82"}
	places := map[string]int{"14": 1, "23": 2, "36": 3, "54": 4, "44": 5, "45": 6, "51": 7, "61": 8, "63": 9, "82": 10}
	types := map[int]string{1: "s", 2: "s", 3: "s", 4: "s", 5: "a", 6: "a", 7: "s", 8: "s", 9: "s", 10: "s"}
	wind_directions := map[string]int{"北": 1, "北北東": 2, "北北西": 3, "北東": 4, "北西": 5, "南": 6, "南南東": 7, "南南西": 8, "南東": 9, "南西": 10, "東": 11, "東北東": 12, "東南東": 13, "西": 14, "西北西": 15, "西南西": 16, "静穏": 17}

	rows, _ := db.Query(`
		SELECT place_id, date FROM races
		WHERE place_id IS NOT NULL AND date >= '` + target_date + `'
		GROUP BY date, place_id
	`)
	defer rows.Close()

	var urls = make([]string, 0)
	for rows.Next() {
		var place_id int
		var date string
		rows.Scan(&place_id, &date)

		url := "https://www.data.jma.go.jp/obd/stats/etrn/view/10min_" + types[place_id] + "1.php?prec_no=" + precs[place_id] + "&block_no=" + blocks[place_id] + "&year=" + date[:4] + "&month=" + date[5:7] + "&day=" + date[8:10] + "&view=p1"
		urls = append(urls, url)
	}

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",             // headlessモードの指定
			"--window-size=1280,800", // ウィンドウサイズの指定
		}),
		agouti.Debug,
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	for _, u := range urls {
		fmt.Println(u)
		ur, _ := url.Parse(u)
		q := ur.Query()

		page.Navigate(u)
		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		contentDom, _ := goquery.NewDocumentFromReader(readerCurContents)

		contentDom.Find("#tablefix1 > tbody > tr:nth-child(n+60):nth-child(-n+105)").Each(func(idx int, s *goquery.Selection) {
			ti := ""
			var temperature string
			var wind string
			var wind_direction string
			if s.Find("td").Length() == 11 {
				ti = s.Find("td:nth-child(1)").Text()
				temperature = s.Find("td:nth-child(5)").Text()
				wind = s.Find("td:nth-child(7)").Text()
				wind_direction = s.Find("td:nth-child(8)").Text()
			} else {
				ti = s.Find("td:nth-child(1)").Text()
				temperature = s.Find("td:nth-child(3)").Text()
				wind = s.Find("td:nth-child(5)").Text()
				wind_direction = s.Find("td:nth-child(6)").Text()
			}

			var direction *int = nil
			d := nilOrWeatherString(wind_direction)
			if d != nil {
				a := *d
				b := wind_directions[a]
				direction = &b
			}

			_, err := db.Exec(
				`INSERT INTO weathers (date, place_id, observation_time, temperature, wind, wind_direction_id) VALUES (?, ?, ?, ?, ?, ?)`,
				q["year"][0]+"-"+q["month"][0]+"-"+q["day"][0],
				places[q["prec_no"][0]],
				ti,
				nilOrWeatherString(temperature),
				nilOrWeatherString(wind),
				direction,
			)

			if err != nil {
				log.Fatalf("insertUser db.Exec error err:%v", err)
			}
		})

		time.Sleep(2 * time.Second)
	}

	defer db.Close()
}

func RaceType(filename string, racetype int, db *sql.DB) {
	file, _ := os.Open("./csv/" + filename)
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	for _, race := range records[0:] {
		if racetype == 1 {
			db.Exec(
				`UPDATE races SET is_only_mare = 1 WHERE id = ?`,
				race[0],
			)
		} else if racetype == 2 {
			db.Exec(
				`UPDATE races SET is_handi = 1 WHERE id = ?`,
				race[0],
			)
		} else {
			db.Exec(
				`UPDATE races SET is_only_age3 = 1 WHERE id = ?`,
				race[0],
			)
		}
	}
}

func nilOrWeatherString(s string) *string {
	if s == "///" || s == "#" || s == "×" {
		return nil
	}

	s = strings.Replace(s, ")", "", 1)
	s = strings.Replace(s, "]", "", 1)
	s = strings.TrimSpace(s)

	return &s
}
