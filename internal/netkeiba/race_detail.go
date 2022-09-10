package netkeiba

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jszwec/csvutil"
	"github.com/sclevine/agouti"

	"main/pkg/config"
	"main/pkg/util"
)

type RACE_NETKEIBA struct {
	RACEID       string `csv:"レースID"`
	BABA         string `csv:"馬場指数"`
	WIDE1        string `csv:"ワイド1馬番"`
	WIDE2        string `csv:"ワイド2馬番"`
	WIDE3        string `csv:"ワイド3馬番"`
	WIDE4        string `csv:"ワイド4馬番"`
	WIDE5        string `csv:"ワイド5馬番"`
	WIDE6        string `csv:"ワイド6馬番"`
	WIDE7        string `csv:"ワイド7馬番"`
	WIDE_HAITO_1 string `csv:"ワイド1"`
	WIDE_HAITO_2 string `csv:"ワイド2"`
	WIDE_HAITO_3 string `csv:"ワイド3"`
	WIDE_HAITO_4 string `csv:"ワイド4"`
	WIDE_HAITO_5 string `csv:"ワイド5"`
	WIDE_HAITO_6 string `csv:"ワイド6"`
	WIDE_HAITO_7 string `csv:"ワイド7"`
}

var netkeibas []RACE_NETKEIBA

func GetNetKeibaData(date string, config config.ConfigList, db *sql.DB) {
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

	f, _ := os.OpenFile("other_comment2.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	w := csv.NewWriter(f)

	n, _ := os.OpenFile("race_netkeiba2.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	wn := csv.NewWriter(n)

	header := [2]string{"レースID", "備考"}
	w.Write(header[:])

	header_w := [16]string{
		"レースID", "馬場指数",
		"ワイド1馬番", "ワイド2馬番", "ワイド3馬番", "ワイド4馬番", "ワイド5馬番", "ワイド6馬番", "ワイド7馬番",
		"ワイド1", "ワイド2", "ワイド3", "ワイド4", "ワイド5", "ワイド6", "ワイド7"}
	wn.Write(header_w[:])

	rows, _ := db.Query(`select id from races where date >= '` + date + `'`)
	defer rows.Close()

	for rows.Next() {
		var id string
		rows.Scan(&id)
		time.Sleep(1 * time.Second)

		fmt.Println(id)
		page_id := id[0:4] + id[8:]
		if err := page.Navigate("https://db.netkeiba.com/race/" + page_id); err != nil {
			log.Fatalf("Failed access page:%v", err)
		}

		html, _ := page.HTML()
		readerCurContents := strings.NewReader(html)
		contentDom, _ := goquery.NewDocumentFromReader(readerCurContents)

		contentDom.Find(".race_table_01 tr:nth-child(n+2)").Each(func(idx int, s *goquery.Selection) {
			var comment = make([]string, 0)
			num, _ := strconv.Atoi(s.Find("td:nth-child(3)").Text())
			str := fmt.Sprintf("%s%02d", id, num)
			comment = append(comment, str)
			comment = append(comment, strings.TrimSpace(s.Find("td:nth-child(18)").Text()))
			if err := w.Write(comment); err != nil {
				fmt.Println("error writing record to csv:", err)
			}
		})

		var rd = make([]string, 0)
		rd = append(rd, id)
		rd = append(rd, contentDom.Find(`#contents > div.result_info.box_left > table:nth-child(2) > tbody > tr:nth-child(1) > td`).Text())

		wide := contentDom.Find(`#contents > div.result_info.box_left > diary_snap > dl > dd > table:nth-child(2) > tbody > tr:nth-child(1) > td:nth-child(2)`).Contents()

		wc := 0
		wide.Each(func(i int, s *goquery.Selection) {
			if !s.Is("br") {
				rd = append(rd, strings.TrimSpace(s.Text()))
				wc++
			}
		})
		for i := wc; i < 7; i++ {
			rd = append(rd, "")
		}

		wide_haitou := contentDom.Find(`#contents > div.result_info.box_left > diary_snap > dl > dd > table:nth-child(2) > tbody > tr:nth-child(1) > td:nth-child(3)`).Contents()
		wide_haitou.Each(func(i int, s *goquery.Selection) {
			if !s.Is("br") {
				rd = append(rd, strings.TrimSpace(s.Text()))
			}
		})
		for i := wc; i < 7; i++ {
			rd = append(rd, "")
		}

		if err := wn.Write(rd); err != nil {
			fmt.Println("Abend GetNetKeibaData error writing record to csv:", err)
		}
	}

	defer w.Flush()
	defer wn.Flush()
}

func CreateNetkeibaDetail(db *sql.DB) {
	newrace, _ := ioutil.ReadFile("./race_netkeiba2.csv")
	_ = csvutil.Unmarshal(newrace, &netkeibas)

	for _, race := range netkeibas[0:] {
		fmt.Println(race.RACEID)

		wide1_1, wide1_2 := wide(race.WIDE1)
		wide2_1, wide2_2 := wide(race.WIDE2)
		wide3_1, wide3_2 := wide(race.WIDE3)
		wide4_1, wide4_2 := wide(race.WIDE4)
		wide5_1, wide5_2 := wide(race.WIDE5)
		wide6_1, wide6_2 := wide(race.WIDE6)
		wide7_1, wide7_2 := wide(race.WIDE7)
		_, err := db.Exec(
			`UPDATE races SET 
			baba = ?,
			wide1_uma1 = ?, wide1_uma2 = ?, wide2_uma1 = ?, wide2_uma2 = ?, wide3_uma1 = ?, wide3_uma2 = ?,
			wide4_uma1 = ?, wide4_uma2 = ?, wide5_uma1 = ?, wide5_uma2 = ?, wide6_uma1 = ?, wide6_uma2 = ?,
			wide7_uma1 = ?, wide7_uma2 = ?, wide_1 = ?, wide_2 = ?, wide_3 = ?, wide_4 = ?, wide_5 = ?, wide_6 = ?, wide_7 = ?
			WHERE id = ?`,
			util.NilOrInt(race.BABA),
			wide1_1, wide1_2, wide2_1, wide2_2, wide3_1, wide3_2,
			wide4_1, wide4_2, wide5_1, wide5_2, wide6_1, wide6_2, wide7_1, wide7_2,
			util.NilOrInt(race.WIDE_HAITO_1), util.NilOrInt(race.WIDE_HAITO_2), util.NilOrInt(race.WIDE_HAITO_3),
			util.NilOrInt(race.WIDE_HAITO_4), util.NilOrInt(race.WIDE_HAITO_5), util.NilOrInt(race.WIDE_HAITO_6), util.NilOrInt(race.WIDE_HAITO_7),
			race.RACEID,
		)

		if err != nil {
			fmt.Println(race)
			log.Fatalf("Abend CreateNetkeibaDetail:%v", err)
		}
	}
}

func Start(db *sql.DB) {
	file, _ := os.Open("other_comment2.csv")
	r := csv.NewReader(file)
	r.FieldsPerRecord = -1

	records, _ := r.ReadAll()
	for _, hr := range records[1:] {
		id := hr[0]
		fmt.Println(id)

		comment := hr[1]
		if strings.Index(comment, "不利") != -1 || strings.Index(comment, "躓く") != -1 || strings.Index(comment, "接触") != -1 {
			_, err := db.Exec(`UPDATE horse_results SET unfavorable = 1 WHERE id = ?`, id)

			if err != nil {
				fmt.Println(hr)
				log.Fatalf("Abend Start:%v", err)
			}
		}
	}
}

func wide(w string) (*string, *string) {
	if w != "" {
		arr := strings.Split(w, "-")

		return &arr[0], &arr[1]
	}

	return nil, nil
}
