package newrace

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	"github.com/sclevine/agouti"

	"main/pkg/config"
	"main/pkg/util"
)

type Training struct {
	NEWID     string `csv:"レースID(新)"`
	Six       string `csv:"6F"`
	SixFive   string `csv:"6F-5F"`
	Five      string `csv:"5F"`
	FiveFour  string `csv:"5F-4F"`
	Four      string `csv:"4F"`
	FourThree string `csv:"4F-3F"`
	Three     string `csv:"3F"`
	ThreeTwo  string `csv:"3F-2F"`
	TwoOne    string `csv:"2F-1F"`
	One       string `csv:"1F"`
	Course    string `csv:"training_course"`
	Load      string `csv:"training_load"`
	Rank      string `csv:"training_rank"`
	Condition string `csv:"調教馬場"`
	Place     string `csv:"位置"`
}

var trainings []Training

func TrainingData(date string, config config.ConfigList, dbConn *sql.DB) {
	training(date, dbConn, config)
	training_sql(dbConn)
}

// 調教データ取得
func training(date string, db *sql.DB, config config.ConfigList) {
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

	f, _ := os.OpenFile("training.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(f)

	rows, _ := db.Query(`select id from races where date = '` + date + `'`)
	defer rows.Close()

	for rows.Next() {
		var id string
		rows.Scan(&id)

		time.Sleep(1 * time.Second)
		fmt.Println(id)
		page_id := id[0:4] + id[8:]
		if err := page.Navigate("https://race.netkeiba.com/race/oikiri.html?race_id=" + page_id + "&type=2&rf=shutuba_submenu"); err != nil {
			log.Fatalf("Failed to compi page:%v", err)
		}

		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		contentDom, _ := goquery.NewDocumentFromReader(readerCurContents)

		contentDom.Find(".HorseList").Each(func(idx int, s *goquery.Selection) {
			var training = make([]string, 0)
			num := fmt.Sprintf("%s%02d", id, idx+1)
			training = append(training, num)

			s.Find(".TrainingTimeDataList").Each(func(idx int, s *goquery.Selection) {
				s.Find("li").Each(func(idx int, s *goquery.Selection) {
					text := s.Text()

					if utf8.RuneCountInString(text) > 3 {
						splits := strings.Split(text, "(")
						time := splits[0]
						if util.IsNumber(time) {
							training = append(training, time)
						} else {
							training = append(training, "")
						}

						if len(splits) == 2 {
							f := strings.Replace(splits[1], ")", "", 1)
							if util.IsNumber(f) {
								training = append(training, f)
							} else {
								training = append(training, "")
							}
						} else {
							training = append(training, "")
						}
					} else {
						training = append(training, "")
						training = append(training, "")
					}
				})
			})

			training = append(training, s.Find(".Training_Day+td").Text())
			training = append(training, s.Find(".TrainingLoad").Text())
			training = append(training, s.Find("td[class*='Rank_']").Text())
			training = append(training, s.Find(".Training_Day+td+td").Text())
			training = append(training, s.Find(".TrainingTimeData+td").Text())

			if err := w.Write(training); err != nil {
				fmt.Println("error writing record to csv:", err)
			}
		})
	}
	defer w.Flush()
}

func training_sql(db *sql.DB) {
	training_csv, _ := ioutil.ReadFile("./training.csv")
	_ = csvutil.Unmarshal(training_csv, &trainings)

	load := map[string]int{
		"馬也": 1,
		"一杯": 2,
		"強め": 3,
		"Ｇ強": 4,
		"仕掛": 5,
		"Ｇ一": 6,
		"直一": 7,
		"直強": 8,
	}
	conditions := map[string]int{"良": 1, "稍": 2, "重": 3, "不": 4}

	for _, training := range trainings[0:] {
		var course *string = nil
		if training.Course != "" {
			course = &training.Course
		}
		var l *int = nil
		if training.Load != "" {
			ld := load[training.Load]
			l = &ld
		}
		var rank *string = nil
		if training.Load != "" {
			rank = &training.Rank
		}
		var cd *int = nil
		if training.Load != "" {
			cdn := conditions[training.Condition]
			cd = &cdn
		}
		var p *string = nil
		if training.Place != "" {
			p = &training.Place
		}

		_, err := db.Exec(`INSERT IGNORE INTO trainings (horse_race_id,furlong_6,furlong_5,furlong_4,furlong_3,furlong_1,furlong_6_5,furlong_5_4,furlong_4_3,furlong_3_2,furlong_2_1,training_course,load_id,rank,training_cond,run_place)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			training.NEWID,
			util.NilOrFloat(training.Six),
			util.NilOrFloat(training.Five),
			util.NilOrFloat(training.Four),
			util.NilOrFloat(training.Three),
			util.NilOrFloat(training.One),
			util.NilOrFloat(training.SixFive),
			util.NilOrFloat(training.FiveFour),
			util.NilOrFloat(training.FourThree),
			util.NilOrFloat(training.ThreeTwo),
			util.NilOrFloat(training.TwoOne),
			course,
			l,
			rank,
			cd,
			p,
		)
		if err != nil {
			log.Fatalf("training_sql db.Exec error err:%v", err)
		}
	}
}
