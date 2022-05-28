package newrace

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/jszwec/csvutil"
)

type Ped struct {
	Id    string `csv:"horse_id"`
	Ped1  string `csv:"peds_0"`
	Ped2  string `csv:"peds_1"`
	Ped3  string `csv:"peds_2"`
	Ped4  string `csv:"peds_3"`
	Ped5  string `csv:"peds_4"`
	Ped6  string `csv:"peds_5"`
	Ped7  string `csv:"peds_6"`
	Ped8  string `csv:"peds_7"`
	Ped9  string `csv:"peds_8"`
	Ped10 string `csv:"peds_9"`
	Ped11 string `csv:"peds_10"`
	Ped12 string `csv:"peds_11"`
	Ped13 string `csv:"peds_12"`
	Ped14 string `csv:"peds_13"`
	Ped15 string `csv:"peds_14"`
	Ped16 string `csv:"peds_15"`
	Ped17 string `csv:"peds_16"`
	Ped18 string `csv:"peds_17"`
	Ped19 string `csv:"peds_18"`
	Ped20 string `csv:"peds_19"`
	Ped21 string `csv:"peds_20"`
	Ped22 string `csv:"peds_21"`
	Ped23 string `csv:"peds_22"`
	Ped24 string `csv:"peds_23"`
	Ped25 string `csv:"peds_24"`
	Ped26 string `csv:"peds_25"`
	Ped27 string `csv:"peds_26"`
	Ped28 string `csv:"peds_27"`
	Ped29 string `csv:"peds_28"`
	Ped30 string `csv:"peds_29"`
	Ped31 string `csv:"peds_30"`
	Ped32 string `csv:"peds_31"`
	Ped33 string `csv:"peds_32"`
	Ped34 string `csv:"peds_33"`
	Ped35 string `csv:"peds_34"`
	Ped36 string `csv:"peds_35"`
	Ped37 string `csv:"peds_36"`
	Ped38 string `csv:"peds_37"`
	Ped39 string `csv:"peds_38"`
	Ped40 string `csv:"peds_39"`
	Ped41 string `csv:"peds_40"`
	Ped42 string `csv:"peds_41"`
	Ped43 string `csv:"peds_42"`
	Ped44 string `csv:"peds_43"`
	Ped45 string `csv:"peds_44"`
	Ped46 string `csv:"peds_45"`
	Ped47 string `csv:"peds_46"`
	Ped48 string `csv:"peds_47"`
	Ped49 string `csv:"peds_48"`
	Ped50 string `csv:"peds_49"`
	Ped51 string `csv:"peds_50"`
	Ped52 string `csv:"peds_51"`
	Ped53 string `csv:"peds_52"`
	Ped54 string `csv:"peds_53"`
	Ped55 string `csv:"peds_54"`
	Ped56 string `csv:"peds_55"`
	Ped57 string `csv:"peds_56"`
	Ped58 string `csv:"peds_57"`
	Ped59 string `csv:"peds_58"`
	Ped60 string `csv:"peds_59"`
	Ped61 string `csv:"peds_60"`
	Ped62 string `csv:"peds_61"`
}

var peds []Ped

func CreateHorseData(db *sql.DB) {
	horse_sql(db)
	create_stallion(db)
}

func horse_sql(db *sql.DB) {
	ped_csv, _ := ioutil.ReadFile("./peds_1.csv")
	_ = csvutil.Unmarshal(ped_csv, &peds)

	for _, ped := range peds[0:] {
		fmt.Println(ped.Id)
		_, err := db.Exec(`INSERT INTO horses (id,
			peds1, peds2, peds3, peds4, peds5, peds6, peds7, peds8, peds9, peds10,
			peds11, peds12, peds13, peds14, peds15, peds16, peds17, peds18, peds19, peds20,
			peds21, peds22, peds23, peds24, peds25, peds26, peds27, peds28, peds29, peds30,
			peds31, peds32, peds33, peds34, peds35, peds36, peds37, peds38, peds39, peds40,
			peds41, peds42, peds43, peds44, peds45, peds46, peds47, peds48, peds49, peds50,
			peds51, peds52, peds53, peds54, peds55, peds56, peds57, peds58, peds59, peds60,
			peds61, peds62
		) VALUES (
			?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			?, ?
		) ON DUPLICATE KEY UPDATE id = ?,
			peds1 = ?, peds2 = ?, peds3 = ?, peds4 = ?, peds5 = ?, peds6 = ?, peds7 = ?, peds8 = ?, peds9 = ?, peds10 = ?,
			peds11 = ?, peds12 = ?, peds13 = ?, peds14 = ?, peds15 = ?, peds16 = ?, peds17 = ?, peds18 = ?, peds19 = ?, peds20 = ?,
			peds21 = ?, peds22 = ?, peds23 = ?, peds24 = ?, peds25 = ?, peds26 = ?, peds27 = ?, peds28 = ?, peds29 = ?, peds30 = ?,
			peds31 = ?, peds32 = ?, peds33 = ?, peds34 = ?, peds35 = ?, peds36 = ?, peds37 = ?, peds38 = ?, peds39 = ?, peds40 = ?,
			peds41 = ?, peds42 = ?, peds43 = ?, peds44 = ?, peds45 = ?, peds46 = ?, peds47 = ?, peds48 = ?, peds49 = ?, peds50 = ?,
			peds51 = ?, peds52 = ?, peds53 = ?, peds54 = ?, peds55 = ?, peds56 = ?, peds57 = ?, peds58 = ?, peds59 = ?, peds60 = ?,
			peds61 = ?, peds62 = ?
		`,
			ped.Id, ped.Ped1, ped.Ped2, ped.Ped3, ped.Ped4, ped.Ped5, ped.Ped6, ped.Ped7, ped.Ped8, ped.Ped9, ped.Ped10, ped.Ped11, ped.Ped12, ped.Ped13, ped.Ped14, ped.Ped15, ped.Ped16, ped.Ped17, ped.Ped18, ped.Ped19, ped.Ped20, ped.Ped21, ped.Ped22, ped.Ped23, ped.Ped24, ped.Ped25, ped.Ped26, ped.Ped27,
			ped.Ped28, ped.Ped29, ped.Ped30, ped.Ped31, ped.Ped32, ped.Ped33, ped.Ped34, ped.Ped35, ped.Ped36, ped.Ped37, ped.Ped38, ped.Ped39, ped.Ped40, ped.Ped41, ped.Ped42, ped.Ped43, ped.Ped44, ped.Ped45, ped.Ped46, ped.Ped47, ped.Ped48, ped.Ped49, ped.Ped50, ped.Ped51, ped.Ped52, ped.Ped53, ped.Ped54, ped.Ped55, ped.Ped56, ped.Ped57, ped.Ped58, ped.Ped59, ped.Ped60, ped.Ped61, ped.Ped62,
			ped.Id, ped.Ped1, ped.Ped2, ped.Ped3, ped.Ped4, ped.Ped5, ped.Ped6, ped.Ped7, ped.Ped8, ped.Ped9, ped.Ped10, ped.Ped11, ped.Ped12, ped.Ped13, ped.Ped14, ped.Ped15, ped.Ped16, ped.Ped17, ped.Ped18, ped.Ped19, ped.Ped20, ped.Ped21, ped.Ped22, ped.Ped23, ped.Ped24, ped.Ped25, ped.Ped26, ped.Ped27,
			ped.Ped28, ped.Ped29, ped.Ped30, ped.Ped31, ped.Ped32, ped.Ped33, ped.Ped34, ped.Ped35, ped.Ped36, ped.Ped37, ped.Ped38, ped.Ped39, ped.Ped40, ped.Ped41, ped.Ped42, ped.Ped43, ped.Ped44, ped.Ped45, ped.Ped46, ped.Ped47, ped.Ped48, ped.Ped49, ped.Ped50, ped.Ped51, ped.Ped52, ped.Ped53, ped.Ped54, ped.Ped55, ped.Ped56, ped.Ped57, ped.Ped58, ped.Ped59, ped.Ped60, ped.Ped61, ped.Ped62,
		)

		if err != nil {
			log.Fatalf("Abend horse_sql:%v", err)
		}
	}
}

// 種牡馬データ作成
func create_stallion(db *sql.DB) {
	db.Exec(`
	UPDATE horses
	LEFT JOIN stallions ON stallions.base_name = horses.peds1 SET horses.stallion_id = stallions.id
	WHERE stallion_id IS NULL and peds1 is not null
	`)

	db.Exec(`UPDATE horses
	LEFT JOIN sires ON sires.name = horses.sireline SET horses.sire_id = sires.id
	WHERE sire_id IS NULL
	`)

	rows, _ := db.Query(`select peds1 from horses where stallion_id is null and peds3 != "" group by peds1`)
	defer rows.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name)

		rex := regexp.MustCompile(`(.*)\s([0-9]+).*\s(.*系)`)
		result := rex.FindStringSubmatch(name)

		if len(result) < 3 {
			continue
		}

		db.Exec(
			`INSERT INTO stallions (base_name, name, born_year, sire) VALUES (?, ?, ?, ?)`,
			name, result[1], result[2], result[3],
		)
	}

	db.Exec(`
	UPDATE horses
	LEFT JOIN stallions ON stallions.base_name = horses.peds1 SET horses.stallion_id = stallions.id
	WHERE stallion_id IS NULL and peds1 is not null
	`)

	db.Exec(`UPDATE horses
	LEFT JOIN sires ON sires.name = horses.sireline SET horses.sire_id = sires.id
	WHERE sire_id IS NULL
	`)
}
