package afterrace

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"

	"main/pkg/util"
)

func ChihouHaitou(db *sql.DB) {
	f, _ := os.OpenFile("chihou_haitou2.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	rows, _ := db.Query("select place_id, date from races inner join places on places.id = races.place_id where places.is_jra = 0 and date >= '2022-09-07' group by place_id, date order by date desc")
	defer rows.Close()

	place := map[int]int{
		30: 36, 34: 7, 35: 10, 36: 11, 42: 18, 43: 19, 44: 20, 45: 21,
		46: 22, 47: 23, 48: 24, 50: 27, 51: 28, 53: 30, 54: 31,
		55: 32, 56: 33, 58: 8}

	for rows.Next() {
		r := rand.Intn(1)
		time.Sleep(time.Duration(r+1) * time.Second)
		// time.Sleep(1 * time.Second)

		var place_id int
		var date string
		rows.Scan(&place_id, &date)

		d := strings.Split(date, "-")

		url := fmt.Sprintf("https://www.keiba.go.jp/KeibaWeb/TodayRaceInfo/RefundMoneyList?k_raceDate=%s/%s/%s&k_babaCode=%d", d[0], d[1], d[2], place[place_id])

		fmt.Println(url)
		page.Navigate(url)
		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		s, _ := goquery.NewDocumentFromReader(readerCurContents)

		t := ""
		s.Find("#mainContainer > article.raceList > div > section > div > div.priceArea.clearfix").Each(func(idx int, s *goquery.Selection) {
			var h = make([]string, 0)
			h = append(h, date)
			h = append(h, strconv.Itoa(place[place_id]))
			h = append(h, strconv.Itoa(idx+1))
			s.Find("table > tbody > tr").Each(func(idx int, s *goquery.Selection) {
				th := s.Find("th").Text()
				if th != "" {
					t = th
				}

				h = append(h, t+":"+s.Find("td:nth-child(2)").Text()+"、"+s.Find("td:nth-child(3)").Text())
			})
			w.Write(h)
		})
	}

	defer w.Flush()
}

func UpdateChihouHaitou(db *sql.DB) {
	file, _ := os.Open("chihou_haitou.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()

	place := map[string]string{
		"36": "30", "7": "34", "10": "35", "11": "36", "18": "42", "19": "43", "20": "44", "21": "45",
		"22": "46", "23": "47", "24": "48", "27": "50", "28": "51", "30": "53", "31": "54",
		"32": "55", "33": "56", "8": "58"}

	types := map[string]string{
		"単勝": "tansho", "複勝": "fukusho", "馬連複": "umaren", "馬連単": "umatan",
		"ワイド": "wide", "三連複": "sanren_fuku", "三連単": "sanren_tan",
	}
	for _, record := range records[:] {
		d := strings.Split(record[0], "-")

		id := d[0] + d[1] + d[2] + place[record[1]] + "%"

		c_type := ""
		haitou := map[string]string{
			"first": "", "second": "", "third": "", "fourth": "",
			"tansho": "", "fukusho_1": "", "fukusho_2": "", "fukusho_3": "", "fukusho_4": "",
			"wide_1": "", "wide_2": "", "wide_3": "", "wide_4": "",
			"wide_5": "", "wide_6": "", "wide_7": "",
			"wide1_uma1": "", "wide1_uma2": "", "wide2_uma1": "", "wide2_uma2": "", "wide3_uma1": "", "wide3_uma2": "",
			"wide4_uma1": "", "wide4_uma2": "", "wide5_uma1": "", "wide5_uma2": "", "wide6_uma1": "", "wide6_uma2": "",
			"wide7_uma1": "", "wide7_uma2": "",
			"umaren": "", "umatan": "", "sanren_fuku": "", "sanren_tan": "",
		}

		var i = 1
		for _, win := range record[:] {
			w := strings.Split(win, "、")
			if len(w) == 1 {
				continue
			}

			n := strings.Split(w[0], ":")
			ht := n[0]

			if ht == "枠連複" || ht == "枠連単" {
				continue
			}

			t := types[ht]
			c_name := ""

			if c_type == ht {
				c_type = ht
				i++
				c_name = t + "_" + strconv.Itoa(i)

				if ht == "馬連複" || ht == "馬連単" || ht == "三連複" || ht == "三連単" {
					continue
				}
			} else {
				c_type = n[0]
				if ht == "単勝" || ht == "馬連複" || ht == "馬連単" || ht == "三連複" || ht == "三連単" {
					i = 1
					c_name = t
				} else {
					c_name = t + "_" + strconv.Itoa(i)
					i = 1
				}
			}

			if ht == "複勝" {
				if i == 1 {
					haitou["first"] = n[1]
				} else if i == 2 {
					haitou["second"] = n[1]
				} else if i == 3 {
					haitou["third"] = n[1]
				} else {
					haitou["fourth"] = n[1]
				}
			}

			if ht == "ワイド" {
				um := strings.Split(n[1], "-")
				haitou["wide"+strconv.Itoa(i)+"_uma1"] = um[0]
				haitou["wide"+strconv.Itoa(i)+"_uma2"] = um[1]
			}

			haitou[c_name] = w[1]
		}

		ri := fmt.Sprintf("%s%02s", id, record[2])
		fmt.Println(ri)

		_, err := db.Exec(
			`update races set first = ?, second = ?, third = ?, fourth = ?,
			tansho = ?, fukusho_1 = ?, fukusho_2 = ?, fukusho_3 = ?, fukusho_4 = ?,
			wide_1 = ?, wide_2 = ?, wide_3 = ?, wide_4 = ?, wide_5 = ?, wide_6 = ?, wide_7 = ?,
			wide1_uma1 = ?, wide1_uma2 = ?, wide2_uma1 = ?, wide2_uma2 = ?, wide3_uma1 = ?, wide3_uma2 = ?,
			wide4_uma1 = ?, wide4_uma2 = ?, wide5_uma1 = ?, wide5_uma2 = ?, wide6_uma1 = ?, wide6_uma2 = ?,
			wide7_uma1 = ?, wide7_uma2 = ?, umaren = ?, umatan = ?, sanren_fuku = ?, sanren_tan = ?
			where id like ?`,
			util.NilOrInt(haitou["first"]),
			util.NilOrInt(haitou["second"]),
			util.NilOrInt(haitou["third"]),
			util.NilOrInt(haitou["fourth"]),
			util.NilOrInt(haitou["tansho"]),
			util.NilOrInt(haitou["fukusho_1"]),
			util.NilOrInt(haitou["fukusho_2"]),
			util.NilOrInt(haitou["fukusho_3"]),
			util.NilOrInt(haitou["fukusho_4"]),
			util.NilOrInt(haitou["wide_1"]),
			util.NilOrInt(haitou["wide_2"]),
			util.NilOrInt(haitou["wide_3"]),
			util.NilOrInt(haitou["wide_4"]),
			util.NilOrInt(haitou["wide_5"]),
			util.NilOrInt(haitou["wide_6"]),
			util.NilOrInt(haitou["wide_7"]),
			util.NilOrInt(haitou["wide1_uma1"]),
			util.NilOrInt(haitou["wide1_uma2"]),
			util.NilOrInt(haitou["wide2_uma1"]),
			util.NilOrInt(haitou["wide2_uma2"]),
			util.NilOrInt(haitou["wide3_uma1"]),
			util.NilOrInt(haitou["wide3_uma2"]),
			util.NilOrInt(haitou["wide4_uma1"]),
			util.NilOrInt(haitou["wide4_uma2"]),
			util.NilOrInt(haitou["wide5_uma1"]),
			util.NilOrInt(haitou["wide5_uma2"]),
			util.NilOrInt(haitou["wide6_uma1"]),
			util.NilOrInt(haitou["wide6_uma2"]),
			util.NilOrInt(haitou["wide7_uma1"]),
			util.NilOrInt(haitou["wide7_uma2"]),
			util.NilOrInt(haitou["umaren"]),
			util.NilOrInt(haitou["umatan"]),
			util.NilOrInt(haitou["sanren_fuku"]),
			util.NilOrInt(haitou["sanren_tan"]),
			ri)

		if err != nil {
			log.Fatalf("Abend update race: %v %v", ri, err)
		}
	}
}
