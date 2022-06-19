package domain

type Race struct {
	NEWID             string `csv:"レースID(新)"`
	PASSED_NEW_ID     string `csv:"前走レースID(新)"`
	PLACE             string `csv:"場所"`
	COURSE            string `csv:"芝・ダ"`
	DISTANCE          string `csv:"距離"`
	GRADE             string `csv:"クラス名"`
	RESULT            string `csv:"着順"`
	IS_MULTIPLE       string `csv:"多頭出し"`
	H_COUNT           string `csv:"出走頭数"`
	AFFILIATION       string `csv:"所属"`
	NAME              string `csv:"馬名S"`
	SEX               string `csv:"性別"`
	AGE               string `csv:"年齢"`
	COLOR             string `csv:"毛色"`
	WEIGHT            string `csv:"斤量"`
	BODY_WEIGHT       string `csv:"馬体重"`
	INCREASE_DECREASE string `csv:"馬体重増減"`
	BRINKER           string `csv:"ブリンカー"`
	BORN              string `csv:"生年月日"`
	BORN_YEAR         string `csv:"生年"`
	BORN_MONTH        string `csv:"誕生日"`
	SIRE_TYPE         string `csv:"母父タイプ名"`
	POPULAR           string `csv:"人気"`
	ODDS              string `csv:"単勝オッズ"`
	WAKU              string `csv:"枠番"`
	UMABAN            string `csv:"馬番"`
	TIME              string `csv:"走破タイム"`
	DIFFERENCE        string `csv:"着差"`
	COURNER_1         string `csv:"1角"`
	COURNER_2         string `csv:"2角"`
	COURNER_3         string `csv:"3角"`
	COURNER_4         string `csv:"4角"`
	LAST_3F           string `csv:"上り3F"`
	LAST_3F_NUM       string `csv:"上り3F順"`
	PRIZE             string `csv:"賞金"`
	ADD_PRIZE         string `csv:"付加賞金"`
	RACE_ID           string `csv:"レースID(新/馬番無)"`
	Zi                string `csv:"指数"`
	Compi_ONE         string `csv:"コンピ一位指数"`
	Compi_ONE_DIFF    string `csv:"コンピ一位指数差"`
	HORSE_ID          string `csv:"血統登録番号"`
	COMPI_PREV        string `csv:"コンピ前指数"`
	COMPI_AFTER       string `csv:"コンピ後指数"`
	JOCKEY_ID         string `csv:"騎手コード"`
	TRAINER_ID        string `csv:"調教師コード"`
	PRODUCER          string `csv:"生産者"`
	OWNER             string `csv:"馬主(最新/仮想)"`
	INTERVAL_WEEK     string `csv:"休み明け〜戦目"`
	WEEK              string `csv:"間隔"`
	CAREER            string `csv:"キャリア"`
	FUKUSHO           string `csv:"複勝シェア"`
	AvgLast3FSpeed    string `csv:"上り3F平均速度"`
	AvgFirst3FSpeed   string `csv:"-3F平均速度"`
	PCI               string `csv:"PCI"`
	AvgSpeed          string `csv:"平均速度"`
	Avg1FTime         string `csv:"平均1Fタイム"`
	RaceType          string `csv:"脚質"`
	Avg3F             string `csv:"Ave-3F"`
	COMPI             string `csv:"コンピ指数"`
	CompiNum          string `csv:"コンピ順位"`
	CorrectTime       string `csv:"補正"`
	START_TIME        string `csv:"発走時刻"`
	Weather           string `csv:"天気"`
	Condition         string `csv:"馬場状態"`
}

type NewRace struct {
	NEWID             string `csv:"レースID(新)"`
	HORSE_ID          string `csv:"血統登録番号"`
	START_TIME        string `csv:"発走時刻"`
	PLACE             string `csv:"場所"`
	GRADE             string `csv:"クラス"`
	COURSE            string `csv:"芝ダ"`
	H_COUNT           string `csv:"頭"`
	DISTANCE          string `csv:"距離"`
	WAKU              string `csv:"枠番"`
	UMABAN            string `csv:"馬番"`
	NAME              string `csv:"馬名S"`
	RESULT            string `csv:"着"`
	DEL               string `csv:"取消しフラグ"`
	SEX               string `csv:"性別"`
	AGE               string `csv:"年齢"`
	WEIGHT            string `csv:"斤量"`
	JOCKEY_ID         string `csv:"騎手コード"`
	TRAINER_ID        string `csv:"調教師コード"`
	IS_MULTIPLE       string `csv:"調教師多頭出し数"`
	AFFILIATION       string `csv:"所属"`
	BODY_WEIGHT       string `csv:"馬体重"`
	BRINKER           string `csv:"B"`
	INCREASE_DECREASE string `csv:"増減"`
	ODDS              string `csv:"指時系1・単勝"`
	POPULAR           string `csv:"指時系1・人気"`
	CAREER            string `csv:"キャリア(最新)"`
	PRIZE             string `csv:"本賞金"`
	ADD_PRIZE         string `csv:"収得賞金"`
	OWNER             string `csv:"馬主"`
	PRODUCER          string `csv:"生産者"`
	BORN_YEAR         string `csv:"生年"`
	BORN_MONTH        string `csv:"誕生日"`
	SIRE_TYPE         string `csv:"母父タイプ名"`
	STALLION          string `csv:"種牡馬"`
	SIRE              string `csv:"母父名"`
	COLOR             string `csv:"毛色"`
	WEEK              string `csv:"間隔"`
	INTERVAL_WEEK     string `csv:"休明戦目"`
	PASSED_NEW_ID     string `csv:"前レースID(新)"`
}
