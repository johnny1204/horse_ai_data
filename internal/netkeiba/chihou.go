package netkeiba

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	"github.com/sclevine/agouti"

	"main/pkg/config"
	"main/pkg/util"
)

type ChihouRace struct {
	RACE_ID     string `csv:"レースID"`
	RACE_NAME   string `csv:"レース名"`
	DATE        string `csv:"日付"`
	BABA        string `csv:"馬場指数"`
	RAP         string `csv:"ラップ"`
	FIRST_TIME  string `csv:"1着タイム"`
	SECOND_TIME string `csv:"2着タイム"`
	RACE_INFO   string `csv:"レース情報"`
	COUNT       string `csv:"頭数"`
}

type Chihou struct {
	RACE_ID     string `csv:"レースID"`
	RESULT      string `csv:"着順"`
	HORSE_NAME  string `csv:"馬名"`
	HORSE_ID    string `csv:"馬ID"`
	WAKU        string `csv:"枠番"`
	UMABAN      string `csv:"馬番"`
	AGE         string `csv:"性齢"`
	WEIGHT      string `csv:"斤量"`
	JOCKEY      string `csv:"騎手"`
	TIME        string `csv:"タイム"`
	RAP         string `csv:"通過"`
	LAST_3F     string `csv:"上がり"`
	ODDS        string `csv:"オッズ"`
	POPULAR     string `csv:"人気"`
	BODY_WEIGHT string `csv:"馬体重"`
	AFFILIATION string `csv:"所属"`
	TRAINER     string `csv:"調教師"`
	OWNER       string `csv:"馬主"`
	PRIZE       string `csv:"賞金`
}

var chihou_races []ChihouRace
var chihou_info []Chihou

var seasons = map[string]int{"01": 4, "02": 4, "03": 1, "04": 1, "05": 1, "06": 2, "07": 2, "08": 2, "09": 3, "10": 3, "11": 3, "12": 4}

func NetkeibaChihou(db *sql.DB) {
	create_race(db)
	// create_race_info(db)
	// update_race_id(db)
	// update_horse(db)
}

func NetkeibaChihouRace(config config.ConfigList) {
	f, _ := os.OpenFile("chihou_race.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(f)
	fr, _ := os.OpenFile("chihou_race_detail.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	wr := csv.NewWriter(fr)

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

	page.Navigate("https://regist.netkeiba.com/account/?pid=login")

	page.FindByXPath("/html/body/div[1]/div/div/form/div/ul/li[1]/input").Fill(config.NetKeibaEmail)
	page.FindByXPath("/html/body/div[1]/div/div/form/div/ul/li[2]/input").Fill(config.NetKeibaPass)

	if err := page.FindByXPath("/html/body/div[1]/div/div/form/div/div[1]/input").Click(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}

	file, _ := os.Open("target_ids.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	header_r := [9]string{
		"レースID",
		"レース名",
		"日付",
		"馬場指数", "ラップ",
		"1着タイム", "2着タイム", "レース情報", "頭数",
	}
	w.Write(header_r[:])

	header := [20]string{
		"レースID",
		"着順", "枠番",
		"馬番", "馬名", "馬ID",
		"性齢", "斤量", "騎手",
		"タイム", "タイム指数", "通過", "上がり", "オッズ", "人気",
		"馬体重", "所属", "調教師", "馬主", "賞金",
	}
	wr.Write(header[:])

	for _, i := range records[:] {
		time.Sleep(1 * time.Second)

		// page_id := i[0]
		id := i[0]
		page_id := id[0:4] + id[8:10] + id[4:8] + id[14:]
		fmt.Println(page_id)
		if err := page.Navigate("https://db.netkeiba.com/race/" + page_id); err != nil {
			log.Fatalf("Failed access page:%v", err)
		}

		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		s, _ := goquery.NewDocumentFromReader(readerCurContents)

		var hr = make([]string, 0)
		hr = append(hr, id)
		hr = append(hr, s.Find("#main > div > div > div > diary_snap > div > div > dl > dd > h1").Text())
		hr = append(hr, s.Find("#main > div > div > div > diary_snap > div > div > p").Text())
		hr = append(hr, s.Find("#contents > div.result_info.box_left > table:nth-child(2) > tbody > tr:nth-child(1) > td").Text())
		hr = append(hr, s.Find("#contents > div.result_info.box_left > table:nth-child(4) > tbody > tr:nth-child(1) > td").Text())
		hr = append(hr, strings.TrimSpace(s.Find("#contents_liquid > table > tbody > tr:nth-child(2) > td:nth-child(8)").Text()))
		hr = append(hr, strings.TrimSpace(s.Find("#contents_liquid > table > tbody > tr:nth-child(3) > td:nth-child(8)").Text()))
		hr = append(hr, s.Find("#main > div > div > div > diary_snap > div > div > dl > dd > p > diary_snap_cut > span").Text())
		hr = append(hr, strconv.Itoa(s.Find("#contents_liquid > table > tbody > tr").Length()-1))

		if err := w.Write(hr); err != nil {
			fmt.Println("error writing record to race csv:", err)
		}

		s.Find("#contents_liquid > table > tbody > tr:nth-child(n+2)").Each(func(idx int, s *goquery.Selection) {
			var h = make([]string, 0)
			// str := fmt.Sprintf("%s%02d", id, num)
			h = append(h, page_id)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(1)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(2)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(3)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(4)").Text()))
			i, _ := s.Find("td:nth-child(4) > a").Attr("href")
			h = append(h, i)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(5)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(6)").Text()))

			jhref, _ := s.Find("td:nth-child(7) > a").Attr("href")
			h = append(h, jhref)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(8)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(10)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(11)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(12)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(13)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(14)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(15)").Text()))

			h = append(h, strings.TrimRight(strings.TrimSpace(s.Find("td:nth-child(19)").Text()), "\n"))
			thref, _ := s.Find("td:nth-child(19) > a").Attr("href")
			h = append(h, thref)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(20)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(21)").Text()))

			if err := wr.Write(h); err != nil {
				fmt.Println("error writing record to detail csv:", err)
			}
		})
	}

	defer w.Flush()
	defer wr.Flush()
}

func NetkeibaChihouDetail(config config.ConfigList) {
	f, _ := os.OpenFile("chihou_new.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(f)

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

	page.Navigate("https://regist.netkeiba.com/account/?pid=login")

	page.FindByXPath("/html/body/div[1]/div/div/form/div/ul/li[1]/input").Fill(config.NetKeibaEmail)
	page.FindByXPath("/html/body/div[1]/div/div/form/div/ul/li[2]/input").Fill(config.NetKeibaPass)

	if err := page.FindByXPath("/html/body/div[1]/div/div/form/div/div[1]/input").Click(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}

	file, _ := os.Open("prev_id.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	header := [17]string{
		"レースID", "着順",
		"枠番", "馬番", "性齢", "斤量", "騎手",
		"タイム", "通過", "上がり", "オッズ", "人気",
		"馬体重", "所属", "調教師", "馬主", "賞金",
	}
	w.Write(header[:])

	for _, i := range records {
		id := i[0]
		time.Sleep(1 * time.Second)
		fmt.Println(id)
		page_id := id[0:4] + id[8:10] + id[4:8] + id[14:16]
		if err := page.Navigate("https://db.netkeiba.com/race/" + page_id); err != nil {
			log.Fatalf("Failed access page:%v", err)
		}

		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		contentDom, _ := goquery.NewDocumentFromReader(readerCurContents)

		contentDom.Find("#contents_liquid > table > tbody > tr").Each(func(idx int, s *goquery.Selection) {
			var h = make([]string, 0)
			// str := fmt.Sprintf("%s%02d", id, num)
			h = append(h, id)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(1)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(2)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(3)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(5)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(6)").Text()))

			jhref, _ := s.Find("td:nth-child(7) > a").Attr("href")
			h = append(h, jhref)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(8)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(11)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(12)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(13)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(14)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(15)").Text()))

			h = append(h, strings.TrimRight(strings.TrimSpace(s.Find("td:nth-child(19)").Text()), "\n"))
			thref, _ := s.Find("td:nth-child(19) > a").Attr("href")
			h = append(h, thref)
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(20)").Text()))
			h = append(h, strings.TrimSpace(s.Find("td:nth-child(21)").Text()))

			if err := w.Write(h); err != nil {
				fmt.Println("error writing record to csv:", err)
			}
		})
	}

	defer w.Flush()
}

func Horse() {
	f, _ := os.OpenFile("chihou_race4.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(f)

	file, _ := os.Open("./chihou_horse_id.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

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

	header_r := [4]string{
		"horse_id", "date", "レースID", "馬番",
	}
	w.Write(header_r[:])

	records, _ := r.ReadAll()
	for _, record := range records[:] {
		fmt.Println(record[0])
		if err := page.Navigate("https://db.netkeiba.com/horse/" + record[0]); err != nil {
			log.Fatalf("Failed access page:%v", err)
		}

		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		contentDom, _ := goquery.NewDocumentFromReader(readerCurContents)

		contentDom.Find("#contents > div.db_main_race.fc > div > table > tbody > tr").Each(func(idx int, s *goquery.Selection) {
			var h = make([]string, 0)
			href, _ := s.Find("td:nth-child(5) > a").Attr("href")
			h = append(h, record[0])
			h = append(h, s.Find("td:nth-child(1)").Text())
			h = append(h, href)
			h = append(h, s.Find("td:nth-child(9)").Text())

			if err := w.Write(h); err != nil {
				fmt.Println("error writing record to csv:", err)
			}
		})

		// time.Sleep(1 * time.Second)
	}

	defer w.Flush()
}

func create_race(db *sql.DB) {
	newraces, _ := ioutil.ReadFile("./chihou_race10.csv")
	_ = csvutil.Unmarshal(newraces, &chihou_races)

	fmt.Println(len(chihou_races))
	for _, c_race := range chihou_races[0:] {
		if c_race.COUNT == "-1" {
			continue
		}

		// pid := c_race.RACE_ID[4:6]

		re1 := regexp.MustCompile("([0-9]+)回")
		re2 := regexp.MustCompile("([0-9]+)日目")
		d := strings.Split(c_race.DATE, " ")
		fs1 := re1.FindStringSubmatch(d[1])
		fs2 := re2.FindStringSubmatch(d[1])

		race_id := c_race.RACE_ID[0:4] + c_race.RACE_ID[6:10] + c_race.RACE_ID[4:6] + fmt.Sprintf("%02s", fs1[1]) + fmt.Sprintf("%02s", fs2[1]) + c_race.RACE_ID[10:]
		fmt.Println(c_race.RACE_ID)

		var place_id string
		db.QueryRow("SELECT id FROM places WHERE id = ?", c_race.RACE_ID[8:10]).Scan(&place_id)
		if place_id == "" {
			continue
		}

		if c_race.FIRST_TIME == "" {
			continue
		}

		date := c_race.RACE_ID[0:4] + "-" + c_race.RACE_ID[6:8] + "-" + c_race.RACE_ID[8:10]
		season := seasons[c_race.RACE_ID[6:8]]
		info := splitRaceInfo(c_race.RACE_INFO)
		r := rap(c_race.RAP)
		names := strings.Split(c_race.RACE_NAME, "(")

		grade := 12
		if len(names) > 1 {
			if strings.Contains(names[1], "G1") {
				grade = 11
			} else if strings.Contains(names[1], "G2") {
				grade = 10
			} else if strings.Contains(names[1], "G3") {
				grade = 9
			} else if strings.Contains(names[1], "G") {
				grade = 8
			}
		}

		_, err := db.Exec(`INSERT INTO races (id, baba, weather, cond,
				place_id, h_num, course, distance, grade, date,
				season_id, start_time,
				rap1, rap2, rap3, rap4, rap5, rap6, rap7, rap8, rap9, rap10,
				rap11, rap12, rap13, rap14, rap15, rap16, rap17, rap18, rap19, rap20
			) VALUES (
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?
			) ON DUPLICATE KEY UPDATE id = ?, grade = ?
		`,
			c_race.RACE_ID, util.NilOrString(c_race.BABA), info["weather"], info["condition"],
			place_id, c_race.COUNT, info["course"], info["distance"], grade, date,
			season, util.NilOrString(info["start_time"]),
			r["rap1"], r["rap2"], r["rap3"], r["rap4"], r["rap5"], r["rap6"], r["rap7"], r["rap8"], r["rap9"], r["rap10"],
			r["rap11"], r["rap12"], r["rap13"], r["rap14"], r["rap15"], r["rap16"], r["rap17"], r["rap18"], r["rap19"], r["rap20"],
			c_race.RACE_ID, grade,
		)

		if err != nil {
			log.Fatalf("Abend create race: %v %v", race_id, err)
		}
	}
}

func create_race_info(db *sql.DB) {
	newinfo, _ := ioutil.ReadFile("./chihou_race6.csv")
	_ = csvutil.Unmarshal(newinfo, &chihou_info)

	gender := map[string]int{"牡": 1, "牝": 2, "セ": 3}

	wt := 0
	f_id := "0"
	c_id := "0"

	for _, race := range chihou_info[0:] {
		if race.WAKU == "" {
			continue
		}

		id := race.RACE_ID
		umaban := fmt.Sprintf("%02s", race.UMABAN)
		id = id[0:4] + id[6:10] + id[4:6] + "%" + id[10:]

		var race_id string
		db.QueryRow("SELECT id FROM races WHERE id like ?", id).Scan(&race_id)
		if race_id == "" {
			continue
		}

		h_race_id := race_id + umaban
		// fmt.Println(h_race_id)

		age := race.AGE

		j := race.JOCKEY
		js := strings.Split(j, "/")

		t := race.TRAINER
		ts := strings.Split(t, "/")

		var time *int
		if race.TIME != "" {
			time = util.TimeToSeconds(race.TIME)
		} else {
			time = nil
		}

		if race.RESULT == "1" {
			wt = *time
			f_id = h_race_id
			if c_id != race.RACE_ID {
				c_id = race.RACE_ID
			}
		}

		var d *float64
		if wt != 0 {
			if time == nil {
				d = nil
			} else {
				ft := float64(*time)
				fwt := float64(wt)
				dt := ((ft - fwt) / 10)
				d = &dt
			}
		}

		if race.RESULT == "2" && c_id == race.RACE_ID {
			_, err := db.Exec(`UPDATE horse_results SET difference = ? WHERE id = ?`, (*d)*-1.0, f_id)
			if err != nil {
				log.Fatalf("Abend update race info:%v %v", race_id, err)
			}
		}

		raps := strings.Split(race.RAP, "-")
		var rap2 *string
		var rap3 *string
		var rap4 *string
		if len(raps) == 2 {
			rap2 = &raps[1]
			rap3 = nil
			rap4 = nil
		} else if len(raps) == 3 {
			rap3 = &raps[2]
			rap4 = nil
		} else if len(raps) == 4 {
			rap2 = &raps[1]
			rap3 = &raps[2]
			rap4 = &raps[3]
		} else {
			rap2 = nil
			rap3 = nil
			rap4 = nil
		}

		var a int
		if strings.Contains(race.AFFILIATION, "西") {
			a = 2
		} else if strings.Contains(race.AFFILIATION, "東") {
			a = 1
		} else {
			a = 4
		}

		var bw *string
		var bwid *string
		if race.BODY_WEIGHT == "計不" {
			bw = nil
			bwid = nil
		} else {
			b := strings.Split(race.BODY_WEIGHT, "(")
			bid := strings.Split(b[1], ")")

			bw = &b[0]
			bwid = &bid[0]
		}

		var last3f *string
		if race.LAST_3F == "" {
			last3f = nil
		} else {
			last3f = &race.LAST_3F
		}

		var prize *string
		if race.PRIZE == "" {
			prize = nil
		} else {
			prize = &race.PRIZE
		}

		var re *string
		var rb int
		if race.RESULT == "取" || race.RESULT == "除" {
			re = nil
			rb = 0
		} else if race.RESULT == "中" || race.RESULT == "失" {
			rd := "99"
			re = &rd
			rb = 0
		} else if strings.Contains(race.RESULT, "降") {
			rs := strings.Split(race.RESULT, "(")
			re = &rs[0]
			rb = 0
		} else {
			re = &race.RESULT
			rb = 1
		}

		hs := strings.Split(race.HORSE_ID, "/")
		horse_id := hs[2]
		_, err := db.Exec(`
					INSERT INTO horses (
						id, name, peds4, peds5, peds6, peds7, peds8, peds9, peds10, 
						peds11, peds12, peds13, peds14, peds15, peds16, peds17, peds18, peds19, peds20,
						peds21, peds22, peds23, peds24, peds25, peds26, peds27, peds28, peds29, peds30,
						peds31, peds32, peds33, peds34, peds35, peds36, peds37, peds38, peds39, peds40,
						peds41, peds42, peds43, peds44, peds45, peds46, peds47, peds48, peds49, peds50,
						peds51, peds52, peds53, peds54, peds55, peds56, peds57, peds58, peds59, peds60,
						peds61, peds62
					)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,
						?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
						?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
						?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
						?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
						?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
						?, ?
					) ON DUPLICATE KEY UPDATE id = ?, name = ?`,
			horse_id, race.HORSE_NAME, "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", horse_id, race.HORSE_NAME,
		)

		fmt.Println(horse_id)
		if err != nil {
			log.Fatalf("Abend create race info:%v %v", horse_id, err)
		}

		_, err = db.Exec(`
			INSERT INTO horse_results(
				id, horse_id, race_id, age, gender, result, result_except, jockey_id,
				trainer_id, affiliation_id, weight, multiple,
				body_weight, body_weight_in_de,
				popular, odds, waku, h_num, time, difference,
				courner1, courner2, courner3, courner4,
				last3f, prize, owner
			)
			VALUES (
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?
			) ON DUPLICATE KEY UPDATE difference = ?`,
			h_race_id, horse_id, race_id, age[3:], gender[age[:3]], re, rb,
			js[4], ts[4], a, race.WEIGHT, 0,
			bw, bwid, util.NilOrInt(race.POPULAR), util.NilOrFloat(race.ODDS), race.WAKU, race.UMABAN,
			time, d, util.NilOrInt(raps[0]), rap2, rap3, rap4,
			last3f, prize, race.OWNER, d,
		)

		if err != nil {
			fmt.Println(race_id)
			log.Fatalf("Abend create race info:%v %v", h_race_id, err)
		}
	}
}

func update_horse(db *sql.DB) {
	file, _ := os.Open("./chihou_race_detail.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	for _, r := range records[1:] {
		rs := strings.Split(r[3], "/")
		h_id := rs[2]

		umaban := fmt.Sprintf("%02s", r[1])
		id := r[0]
		h_race_id := id[0:4] + id[6:10] + id[4:6] + "%" + id[10:] + umaban

		fmt.Println(h_race_id)

		_, err := db.Exec(`UPDATE horse_results SET horse_id = ? WHERE id like ?`, h_id, h_race_id)
		if err != nil {
			log.Fatalf("Abend update horse:%v %v", h_race_id, err)
		}

		_, err1 := db.Exec(`
			INSERT IGNORE INTO horses (
				id, name, peds4, peds5, peds6, peds7, peds8, peds9, peds10, 
				peds11, peds12, peds13, peds14, peds15, peds16, peds17, peds18, peds19, peds20,
				peds21, peds22, peds23, peds24, peds25, peds26, peds27, peds28, peds29, peds30,
				peds31, peds32, peds33, peds34, peds35, peds36, peds37, peds38, peds39, peds40,
				peds41, peds42, peds43, peds44, peds45, peds46, peds47, peds48, peds49, peds50,
				peds51, peds52, peds53, peds54, peds55, peds56, peds57, peds58, peds59, peds60,
				peds61, peds62
			)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
				?, ?
			) ON DUPLICATE KEY UPDATE id = ?`,
			h_id, r[2], "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", "", "", "", "", "", "", "", "",
			"", "", h_id,
		)

		if err1 != nil {
			log.Fatalf("Abend create race info:%v %v", h_id, err)
		}
	}
}

func update_race_id(db *sql.DB) {
	file, _ := os.Open("chihou_race.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	re1 := regexp.MustCompile("([0-9]+)回")
	re2 := regexp.MustCompile("([0-9]+)日目")
	for _, r := range records[1:] {
		id := r[0]
		race_id := id[0:4] + id[6:10] + id[4:6] + "0000" + id[10:]

		d := strings.Split(r[1], " ")
		fs1 := re1.FindStringSubmatch(d[1])
		fs2 := re2.FindStringSubmatch(d[1])

		new_id := id[0:4] + id[6:10] + id[4:6] + fmt.Sprintf("%02s", fs1[1]) + fmt.Sprintf("%02s", fs2[1]) + id[10:]

		fmt.Println(new_id)
		db.Exec(`UPDATE races SET id = ? WHERE id = ?`, new_id, race_id)
	}
}

func splitRaceInfo(info string) map[string]string {
	s := strings.Split(info, " / ")

	courses := map[string]int{"芝": 1, "ダ": 2}
	weathers := map[string]int{"晴": 1, "曇": 2, "雨": 3, "小雨": 4, "小雪": 5, "雪": 6}
	conditions := map[string]int{"良": 1, "稍重": 2, "重": 3, "不良": 4}

	r := regexp.MustCompile("([0-9]+)")
	fs := r.FindString(s[0])

	rc := regexp.MustCompile("(右|左)")
	fcs := rc.FindString(s[0])

	w := strings.Split(s[1], " : ")
	wt := "0"
	if len(w) > 1 {
		wt = strconv.Itoa(weathers[w[1]])
	}

	c := strings.Split(s[2], " : ")
	cd := "0"
	if len(c) > 1 {
		cd = strconv.Itoa(conditions[c[1]])
	}

	t := strings.Split(s[3], " : ")

	time := ""
	if len(t) > 1 {
		time = t[1]
	}

	return map[string]string{
		"course":     strconv.Itoa(courses[s[0][0:3]]),
		"distance":   fs,
		"courner":    fcs,
		"weather":    wt,
		"condition":  cd,
		"start_time": time,
	}
}

func rap(info string) map[string]string {
	raps := map[string]string{
		"rap1": "0", "rap2": "0", "rap3": "0", "rap4": "0", "rap5": "0", "rap6": "0", "rap7": "0", "rap8": "0", "rap9": "0", "rap10": "0",
		"rap11": "0", "rap12": "0", "rap13": "0", "rap14": "0", "rap15": "0", "rap16": "0", "rap17": "0", "rap18": "0", "rap19": "0", "rap20": "0",
	}
	rap := strings.Split(info, " - ")

	for i, r := range rap {
		raps["rap"+strconv.Itoa(i)] = r
	}

	return raps
}

func otherIds() {
	file, _ := os.Open("other_race_id.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	f, _ := os.OpenFile("target_ids.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(f)

	records, _ := r.ReadAll()

	ids := make([]string, 0)
	for _, r := range records[1:] {
		id := r[0]
		ids = append(ids, id[0:10]+id[14:])
	}

	var date string
	var place_id string
	for _, r := range records[1:] {
		if date == r[1] && place_id == r[2] {
			continue
		}

		date = r[1]
		place_id = r[2]

		for i := 1; i <= 12; i++ {
			var h = make([]string, 0)
			s := date[0:4] + date[5:7] + date[8:10] + place_id
			str := fmt.Sprintf("%s%02s", s, strconv.Itoa(i))
			if util.ArrayContains(ids, str) {
				continue
			}

			h = append(h, str)
			if err := w.Write(h); err != nil {
				fmt.Println("error writing record to csv:", err)
			}
		}

	}

	defer w.Flush()
}
