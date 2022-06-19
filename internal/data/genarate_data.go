package data

import (
	"database/sql"
	"encoding/csv"
	"os"
)

type NullString struct {
	sql.NullString
}

type HorsePillar struct {
	id                       NullString
	date                     NullString
	horse_id                 NullString
	race_id                  NullString
	age                      NullString
	course                   NullString
	distance                 NullString
	result                   NullString
	jockey_id                NullString
	multiple                 NullString
	affiliation_id           NullString
	trainer_id               NullString
	gender                   NullString
	weight                   NullString
	is_loss_jockey           NullString
	body_weight              NullString
	body_weight_in_de        NullString
	popular                  NullString
	odds                     NullString
	compi                    NullString
	compi_num                NullString
	speed                    NullString
	waku                     NullString
	h_num                    NullString
	interval_week            NullString
	after_break              NullString
	producer                 NullString
	owner                    NullString
	time_odds                NullString
	time_popular             NullString
	born_month               NullString
	stallion_id              NullString
	color_id                 NullString
	training_urlong_6        NullString
	training_furlong_5       NullString
	training_furlong_4       NullString
	training_furlong_3       NullString
	training_furlong_1       NullString
	training_furlong_6_5     NullString
	training_furlong_5_4     NullString
	training_furlong_4_3     NullString
	training_furlong_3_2     NullString
	training_furlong_2_1     NullString
	training_training_course NullString
	training_load_id         NullString
	training_rank            NullString
	training_training_cond   NullString
	training_run_place       NullString
	one_result               NullString
	one_relative_result      NullString
	one_jockey_id            NullString
	one_weight               NullString
	one_is_loss_jockey       NullString
	one_body_weight_inde     NullString
	one_popular              NullString
	one_relative_popular     NullString
	one_odds                 NullString
	one_compi                NullString
	one_compi_num            NullString
	one_zi                   NullString
	one_waku                 NullString
	one_horse_num            NullString
	one_time                 NullString
	one_correct_time         NullString
	one_difference           NullString
	one_one_courner          NullString
	one_two_courner          NullString
	one_three_courner        NullString
	one_fourth_courner       NullString
	one_is_brinker           NullString
	one_deokure              NullString
	one_last3f               NullString
	one_last3f_num           NullString
	one_last_3f_speed        NullString
	one_first_3f_speed       NullString
	one_speed_avg            NullString
	one_avg_1f_time          NullString
	one_avg_3f               NullString
	one_race_type            NullString
	one_preceding_index      NullString
	one_pace_index           NullString
	one_last3f_index         NullString
	one_speed_index          NullString
	one_unfavorable          NullString
	one_baba                 NullString
	one_h_num                NullString
	one_distance             NullString
	one_place_id             NullString
	one_course               NullString
	one_grade                NullString
	one_first_fact           NullString
	one_ave_3f               NullString
	one_last_5f              NullString
	one_last_4f              NullString
	one_last_3f              NullString
	one_last_2f              NullString
	one_last_1f              NullString
	one_first_5f             NullString
	one_first_4f             NullString
	one_first_3f             NullString
	one_first_2f             NullString
	one_first_1f             NullString
	one_rap1                 NullString
	one_rap2                 NullString
	one_rap3                 NullString
	one_rap4                 NullString
	one_rap5                 NullString
	one_rap6                 NullString
	one_rap7                 NullString
	one_rap8                 NullString
	one_rap9                 NullString
	one_rap10                NullString
	one_rap11                NullString
	one_rap12                NullString
	one_rap13                NullString
	one_rap14                NullString
	one_rap15                NullString
	one_rap16                NullString
	one_rap17                NullString
	one_rap18                NullString
	one_rap19                NullString
	one_rap20                NullString
	two_result               NullString
	two_relative_result      NullString
	two_jockey_id            NullString
	two_weight               NullString
	two_is_loss_jockey       NullString
	two_body_weight_inde     NullString
	two_popular              NullString
	two_relative_popular     NullString
	two_odds                 NullString
	two_compi                NullString
	two_compi_num            NullString
	two_zi                   NullString
	two_waku                 NullString
	two_horse_num            NullString
	two_time                 NullString
	two_correct_time         NullString
	two_difference           NullString
	two_one_courner          NullString
	two_two_courner          NullString
	two_three_courner        NullString
	two_fourth_courner       NullString
	two_is_brinker           NullString
	two_deokure              NullString
	two_last3f               NullString
	two_last3f_num           NullString
	two_last_3f_speed        NullString
	two_first_3f_speed       NullString
	two_speed_avg            NullString
	two_avg_1f_time          NullString
	two_avg_3f               NullString
	two_race_type            NullString
	two_preceding_index      NullString
	two_pace_index           NullString
	two_last3f_index         NullString
	two_speed_index          NullString
	two_unfavorable          NullString
	two_baba                 NullString
	two_h_num                NullString
	two_distance             NullString
	two_place_id             NullString
	two_course               NullString
	two_grade                NullString
	two_first_fact           NullString
	two_ave_3f               NullString
	two_last_5f              NullString
	two_last_4f              NullString
	two_last_3f              NullString
	two_last_2f              NullString
	two_last_1f              NullString
	two_first_5f             NullString
	two_first_4f             NullString
	two_first_3f             NullString
	two_first_2f             NullString
	two_first_1f             NullString
	two_rap1                 NullString
	two_rap2                 NullString
	two_rap3                 NullString
	two_rap4                 NullString
	two_rap5                 NullString
	two_rap6                 NullString
	two_rap7                 NullString
	two_rap8                 NullString
	two_rap9                 NullString
	two_rap10                NullString
	two_rap11                NullString
	two_rap12                NullString
	two_rap13                NullString
	two_rap14                NullString
	two_rap15                NullString
	two_rap16                NullString
	two_rap17                NullString
	two_rap18                NullString
	two_rap19                NullString
	two_rap20                NullString
	three_result             NullString
	three_relative_result    NullString
	three_jockey_id          NullString
	three_weight             NullString
	three_is_loss_jockey     NullString
	three_body_weight_inde   NullString
	three_popular            NullString
	three_relative_popular   NullString
	three_odds               NullString
	three_compi              NullString
	three_compi_num          NullString
	three_zi                 NullString
	three_waku               NullString
	three_horse_num          NullString
	three_time               NullString
	three_correct_time       NullString
	three_difference         NullString
	three_one_courner        NullString
	three_two_courner        NullString
	three_three_courner      NullString
	three_fourth_courner     NullString
	three_is_brinker         NullString
	three_deokure            NullString
	three_last3f             NullString
	three_last3f_num         NullString
	three_last_3f_speed      NullString
	three_first_3f_speed     NullString
	three_speed_avg          NullString
	three_avg_1f_time        NullString
	three_avg_3f             NullString
	three_race_type          NullString
	three_preceding_index    NullString
	three_pace_index         NullString
	three_last3f_index       NullString
	three_speed_index        NullString
	three_unfavorable        NullString
	three_baba               NullString
	three_h_num              NullString
	three_distance           NullString
	three_place_id           NullString
	three_course             NullString
	three_grade              NullString
	three_first_fact         NullString
	three_ave_3f             NullString
	three_last_5f            NullString
	three_last_4f            NullString
	three_last_3f            NullString
	three_last_2f            NullString
	three_last_1f            NullString
	three_first_5f           NullString
	three_first_4f           NullString
	three_first_3f           NullString
	three_first_2f           NullString
	three_first_1f           NullString
	three_rap1               NullString
	three_rap2               NullString
	three_rap3               NullString
	three_rap4               NullString
	three_rap5               NullString
	three_rap6               NullString
	three_rap7               NullString
	three_rap8               NullString
	three_rap9               NullString
	three_rap10              NullString
	three_rap11              NullString
	three_rap12              NullString
	three_rap13              NullString
	three_rap14              NullString
	three_rap15              NullString
	three_rap16              NullString
	three_rap17              NullString
	three_rap18              NullString
	three_rap19              NullString
	three_rap20              NullString
	four_result              NullString
	four_relative_result     NullString
	four_jockey_id           NullString
	four_weight              NullString
	four_is_loss_jockey      NullString
	four_body_weight_inde    NullString
	four_popular             NullString
	four_relative_popular    NullString
	four_odds                NullString
	four_compi               NullString
	four_compi_num           NullString
	four_zi                  NullString
	four_waku                NullString
	four_horse_num           NullString
	four_time                NullString
	four_correct_time        NullString
	four_difference          NullString
	four_one_courner         NullString
	four_two_courner         NullString
	four_three_courner       NullString
	four_fourth_courner      NullString
	four_is_brinker          NullString
	four_deokure             NullString
	four_last3f              NullString
	four_last3f_num          NullString
	four_last_3f_speed       NullString
	four_first_3f_speed      NullString
	four_speed_avg           NullString
	four_avg_1f_time         NullString
	four_avg_3f              NullString
	four_race_type           NullString
	four_preceding_index     NullString
	four_pace_index          NullString
	four_last3f_index        NullString
	four_speed_index         NullString
	four_unfavorable         NullString
	four_baba                NullString
	four_h_num               NullString
	four_distance            NullString
	four_place_id            NullString
	four_course              NullString
	four_grade               NullString
	four_first_fact          NullString
	four_ave_3f              NullString
	four_last_5f             NullString
	four_last_4f             NullString
	four_last_3f             NullString
	four_last_2f             NullString
	four_last_1f             NullString
	four_first_5f            NullString
	four_first_4f            NullString
	four_first_3f            NullString
	four_first_2f            NullString
	four_first_1f            NullString
	four_rap1                NullString
	four_rap2                NullString
	four_rap3                NullString
	four_rap4                NullString
	four_rap5                NullString
	four_rap6                NullString
	four_rap7                NullString
	four_rap8                NullString
	four_rap9                NullString
	four_rap10               NullString
	four_rap11               NullString
	four_rap12               NullString
	four_rap13               NullString
	four_rap14               NullString
	four_rap15               NullString
	four_rap16               NullString
	four_rap17               NullString
	four_rap18               NullString
	four_rap19               NullString
	four_rap20               NullString
	five_result              NullString
	five_relative_result     NullString
	five_jockey_id           NullString
	five_weight              NullString
	five_is_loss_jockey      NullString
	five_body_weight_inde    NullString
	five_popular             NullString
	five_relative_popular    NullString
	five_odds                NullString
	five_compi               NullString
	five_compi_num           NullString
	five_zi                  NullString
	five_waku                NullString
	five_horse_num           NullString
	five_time                NullString
	five_correct_time        NullString
	five_difference          NullString
	five_one_courner         NullString
	five_two_courner         NullString
	five_three_courner       NullString
	five_fourth_courner      NullString
	five_is_brinker          NullString
	five_deokure             NullString
	five_last3f              NullString
	five_last3f_num          NullString
	five_last_3f_speed       NullString
	five_first_3f_speed      NullString
	five_speed_avg           NullString
	five_avg_1f_time         NullString
	five_avg_3f              NullString
	five_race_type           NullString
	five_preceding_index     NullString
	five_pace_index          NullString
	five_last3f_index        NullString
	five_speed_index         NullString
	five_unfavorable         NullString
	five_baba                NullString
	five_h_num               NullString
	five_distance            NullString
	five_place_id            NullString
	five_course              NullString
	five_grade               NullString
	five_first_fact          NullString
	five_ave_3f              NullString
	five_last_5f             NullString
	five_last_4f             NullString
	five_last_3f             NullString
	five_last_2f             NullString
	five_last_1f             NullString
	five_first_5f            NullString
	five_first_4f            NullString
	five_first_3f            NullString
	five_first_2f            NullString
	five_first_1f            NullString
	five_rap1                NullString
	five_rap2                NullString
	five_rap3                NullString
	five_rap4                NullString
	five_rap5                NullString
	five_rap6                NullString
	five_rap7                NullString
	five_rap8                NullString
	five_rap9                NullString
	five_rap10               NullString
	five_rap11               NullString
	five_rap12               NullString
	five_rap13               NullString
	five_rap14               NullString
	five_rap15               NullString
	five_rap16               NullString
	five_rap17               NullString
	five_rap18               NullString
	five_rap19               NullString
	five_rap20               NullString
}

type HorsePillarShinba struct {
	id                       NullString
	date                     NullString
	horse_id                 NullString
	race_id                  NullString
	age                      NullString
	course                   NullString
	distance                 NullString
	result                   NullString
	result_except            NullString
	jockey_id                NullString
	multiple                 NullString
	affiliation_id           NullString
	trainer_id               NullString
	gender                   NullString
	weight                   NullString
	is_loss_jockey           NullString
	body_weight              NullString
	body_weight_in_de        NullString
	popular                  NullString
	odds                     NullString
	waku                     NullString
	h_num                    NullString
	producer                 NullString
	owner                    NullString
	time_odds                NullString
	time_popular             NullString
	born_month               NullString
	stallion_id              NullString
	sireline                 NullString
	color_id                 NullString
	brother_race_count       NullString
	brother_win_count        NullString
	brother_win_rate         NullString
	brother_rentai_rate      NullString
	brother_fukusho_rate     NullString
	training_urlong_6        NullString
	training_furlong_5       NullString
	training_furlong_4       NullString
	training_furlong_3       NullString
	training_furlong_1       NullString
	training_furlong_6_5     NullString
	training_furlong_5_4     NullString
	training_furlong_4_3     NullString
	training_furlong_3_2     NullString
	training_furlong_2_1     NullString
	training_training_course NullString
	training_load_id         NullString
	training_rank            NullString
	training_training_cond   NullString
	training_run_place       NullString
}

type JockeyRate struct {
	id                          NullString
	horse_race_id               NullString
	jockey_id                   NullString
	place_id                    NullString
	course                      NullString
	distance                    NullString
	date                        NullString
	count                       NullString
	win_rate                    NullString
	rentai_rate                 NullString
	fukusho_rate                NullString
	place_count                 NullString
	place_win_rate              NullString
	place_rentai_rate           NullString
	place_fukusho_rate          NullString
	course_count                NullString
	course_win_rate             NullString
	course_rentai_rate          NullString
	course_fukusho_rate         NullString
	distance_count              NullString
	distance_win_rate           NullString
	distance_rentai_rate        NullString
	distance_fukusho_rate       NullString
	place_distance_count        NullString
	place_distance_win_rate     NullString
	place_distance_rentai_rate  NullString
	place_distance_fukusho_rate NullString
}

type TrainerRate struct {
	id                          NullString
	horse_race_id               NullString
	trainer_id                  NullString
	place_id                    NullString
	course                      NullString
	distance                    NullString
	date                        NullString
	count                       NullString
	win_rate                    NullString
	rentai_rate                 NullString
	fukusho_rate                NullString
	place_count                 NullString
	place_win_rate              NullString
	place_rentai_rate           NullString
	place_fukusho_rate          NullString
	course_count                NullString
	course_win_rate             NullString
	course_rentai_rate          NullString
	course_fukusho_rate         NullString
	distance_count              NullString
	distance_win_rate           NullString
	distance_rentai_rate        NullString
	distance_fukusho_rate       NullString
	place_distance_count        NullString
	place_distance_win_rate     NullString
	place_distance_rentai_rate  NullString
	place_distance_fukusho_rate NullString
	age_year_win_rate           NullString
	age_year_rentai_rate        NullString
	age_year_fukusho_rate       NullString
	age_win_rate                NullString
	age_rentai_rate             NullString
	age_fukusho_rate            NullString
}

type StallionRate struct {
	race_id                       NullString
	name                          NullString
	courner_type                  NullString
	course                        NullString
	distance                      NullString
	stallion_id                   NullString
	date                          NullString
	start_time                    NullString
	count                         NullString
	win_rate                      NullString
	rentai_rate                   NullString
	fukushou_rate                 NullString
	place_count                   NullString
	place_win_rate                NullString
	place_rentai_rate             NullString
	place_fukushou_rate           NullString
	course_count                  NullString
	course_win_rate               NullString
	course_rentai_rate            NullString
	course_fukushou_rate          NullString
	distance_count                NullString
	distance_win_rate             NullString
	distance_rentai_rate          NullString
	distance_fukushou_rate        NullString
	course_distance_count         NullString
	course_distance_win_rate      NullString
	course_distance_rentai_rate   NullString
	course_distance_fukushou_rate NullString
	circle_count                  NullString
	circle_win_rate               NullString
	circle_rentai_rate            NullString
	circle_fukushou_rate          NullString
	waku_count                    NullString
	waku_win_rate                 NullString
	waku_rentai_rate              NullString
	waku_fukushou_rate            NullString
}

type HorseSireRate struct {
	id                            NullString
	race_id                       NullString
	name                          NullString
	place_id                      NullString
	course                        NullString
	distance                      NullString
	sire_id                       NullString
	date                          NullString
	start_time                    NullString
	count                         NullString
	win_rate                      NullString
	rentai_rate                   NullString
	fukushou_rate                 NullString
	place_count                   NullString
	place_win_rate                NullString
	place_rentai_rate             NullString
	place_fukushou_rate           NullString
	course_count                  NullString
	course_win_rate               NullString
	course_rentai_rate            NullString
	course_fukushou_rate          NullString
	distance_count                NullString
	distance_win_rate             NullString
	distance_rentai_rate          NullString
	distance_fukushou_rate        NullString
	course_distance_count         NullString
	course_distance_win_rate      NullString
	course_distance_rentai_rate   NullString
	course_distance_fukushou_rate NullString
}

type HorseResults struct {
	id                                 NullString
	date                               NullString
	race_id                            NullString
	horse_id                           NullString
	brinker                            NullString
	age                                NullString
	course                             NullString
	distance                           NullString
	career                             NullString
	nige                               NullString
	senkou                             NullString
	sashi                              NullString
	oikomi                             NullString
	makuri                             NullString
	ave_last_3f                        NullString
	win_rate                           NullString
	same_place_win_rate                NullString
	same_distance_win_rate             NullString
	same_place_disatance_win_place     NullString
	same_jockey_win_rate               NullString
	course_type_win_rate               NullString
	same_grade_win_rate                NullString
	sane_waku_win_rate                 NullString
	same_clockwise_win_rate            NullString
	rentai_rate                        NullString
	same_place_rentai_rate             NullString
	same_distance_rentai_rate          NullString
	same_place_disatance_rentai_place  NullString
	same_grade_distance_rentai_rate    NullString
	same_jockey_rentai_rate            NullString
	course_type_rentai_rate            NullString
	sane_waku_rentai_rate              NullString
	same_clockwise_rentai_rate         NullString
	fukusho_rate                       NullString
	same_place_fukusho_rate            NullString
	same_distance_fukusho_rate         NullString
	same_place_disatance_fukusho_place NullString
	same_jockey_fukusho_rate           NullString
	same_grade_distance_fukusho_rate   NullString
	course_type_fukusho_rate           NullString
	sane_waku_fukusho_rate             NullString
	start_late                         NullString
	jockey_start_late                  NullString
	total_prize                        NullString
	total_add_prize                    NullString
	same_clockwise_fukusho_rate        NullString
	first_brinker                      NullString
}

type SeasonRate struct {
	id             NullString
	date           NullString
	horse_id       NullString
	season_count   NullString
	season_win     NullString
	season_rentai  NullString
	season_fukusho NullString
}

func (ns NullString) MarshalJSON() string {
	if !ns.Valid {
		return ""
	}

	return ns.String
}

func GenerateData(date string, db *sql.DB) {
	get_horse_pillar(date, db)
	get_jockey_rate(date, db)
	get_horse_results(date, db)
	get_trainer_rate(date, db)
	get_stallion_rate(date, db)
	get_horse_sire_rate(date, db)
	get_season_rate(date, db)
	get_shinba(date, db)
}

func get_horse_results(date string, db *sql.DB) {
	rows, _ := db.Query("select * from horse_results_detail where date = ?", date)
	pillars := make([][]string, 0)
	pillars = append(pillars, []string{
		"id", "date", "race_id", "horse_id", "brinker", "age", "course", "distance", "career", "逃げ率", "先行率", "中団率", "追込率", "マクリ率", "上がり3F平均",
		"勝率", "同競馬場勝率", "同距離勝率", "同競馬場同距離勝率", "同騎手騎乗勝率", "コースタイプ勝率", "同距離同クラス勝率", "同枠タイプ生涯勝率", "同周り勝率",
		"連対率", "同競馬場連対率", "同距離連対率", "同競馬場同距離連対率", "同騎手騎乗連対率", "コースタイプ連対率", "同距離同クラス連対率", "同枠タイプ生涯連対率", "同周り連対率",
		"複勝率", "同競馬場複勝率", "同距離複勝率", "同競馬場同距離複勝率", "同騎手騎乗複勝率", "コースタイプ複勝率", "同距離同クラス複勝率", "同枠タイプ生涯複勝率",
		"生涯出遅れ率", "騎乗騎手年間出遅れ率", "獲得賞金合計", "付加賞金合計", "同周り複勝率", "初ブリンカー"})

	for rows.Next() {
		hr := HorseResults{}
		err := rows.Scan(
			&hr.id, &hr.date, &hr.race_id, &hr.horse_id, &hr.brinker, &hr.age, &hr.course, &hr.distance, &hr.career, &hr.nige, &hr.senkou, &hr.sashi, &hr.oikomi, &hr.makuri, &hr.ave_last_3f,
			&hr.win_rate, &hr.same_place_win_rate, &hr.same_distance_win_rate, &hr.same_place_disatance_win_place, &hr.same_jockey_win_rate, &hr.course_type_win_rate, &hr.same_grade_win_rate, &hr.sane_waku_win_rate, &hr.same_clockwise_win_rate,
			&hr.rentai_rate, &hr.same_place_rentai_rate, &hr.same_distance_rentai_rate, &hr.same_place_disatance_rentai_place, &hr.same_grade_distance_rentai_rate, &hr.same_jockey_rentai_rate, &hr.course_type_rentai_rate, &hr.sane_waku_rentai_rate, &hr.same_clockwise_rentai_rate,
			&hr.fukusho_rate, &hr.same_place_fukusho_rate, &hr.same_distance_fukusho_rate, &hr.same_place_disatance_fukusho_place, &hr.same_jockey_fukusho_rate, &hr.same_grade_distance_fukusho_rate, &hr.course_type_fukusho_rate, &hr.sane_waku_fukusho_rate,
			&hr.start_late, &hr.jockey_start_late, &hr.total_prize, &hr.total_add_prize, &hr.same_clockwise_fukusho_rate, &hr.first_brinker)
		if err != nil {
			panic(err)
		}

		pillars = append(pillars, []string{
			hr.id.MarshalJSON(), hr.date.MarshalJSON(), hr.race_id.MarshalJSON(), hr.horse_id.MarshalJSON(), hr.brinker.MarshalJSON(), hr.age.MarshalJSON(), hr.course.MarshalJSON(), hr.distance.MarshalJSON(), hr.career.MarshalJSON(), hr.nige.MarshalJSON(), hr.senkou.MarshalJSON(), hr.sashi.MarshalJSON(), hr.oikomi.MarshalJSON(), hr.makuri.MarshalJSON(), hr.ave_last_3f.MarshalJSON(),
			hr.win_rate.MarshalJSON(), hr.same_place_win_rate.MarshalJSON(), hr.same_distance_win_rate.MarshalJSON(), hr.same_place_disatance_win_place.MarshalJSON(), hr.same_jockey_win_rate.MarshalJSON(), hr.course_type_win_rate.MarshalJSON(), hr.same_grade_win_rate.MarshalJSON(), hr.sane_waku_win_rate.MarshalJSON(), hr.same_clockwise_win_rate.MarshalJSON(),
			hr.rentai_rate.MarshalJSON(), hr.same_place_rentai_rate.MarshalJSON(), hr.same_distance_rentai_rate.MarshalJSON(), hr.same_place_disatance_rentai_place.MarshalJSON(), hr.same_grade_distance_rentai_rate.MarshalJSON(), hr.same_jockey_rentai_rate.MarshalJSON(), hr.course_type_rentai_rate.MarshalJSON(), hr.sane_waku_rentai_rate.MarshalJSON(), hr.same_clockwise_rentai_rate.MarshalJSON(),
			hr.fukusho_rate.MarshalJSON(), hr.same_place_fukusho_rate.MarshalJSON(), hr.same_distance_fukusho_rate.MarshalJSON(), hr.same_place_disatance_fukusho_place.MarshalJSON(), hr.same_jockey_fukusho_rate.MarshalJSON(), hr.same_grade_distance_fukusho_rate.MarshalJSON(), hr.course_type_fukusho_rate.MarshalJSON(), hr.sane_waku_fukusho_rate.MarshalJSON(), hr.start_late.MarshalJSON(), hr.jockey_start_late.MarshalJSON(),
			hr.total_prize.MarshalJSON(), hr.total_add_prize.MarshalJSON(), hr.same_clockwise_fukusho_rate.MarshalJSON(), hr.first_brinker.MarshalJSON()})

		f, _ := os.OpenFile("./data/race_detail.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()

		w := csv.NewWriter(f)

		w.WriteAll(pillars)

		w.Flush()
	}
}

func get_horse_pillar(date string, db *sql.DB) {
	rows, _ := db.Query("select * from horse_pillar where date = ?", date)

	pillars := make([][]string, 0)
	pillars = append(pillars, []string{
		"horse_race_id", "date", "horse_id", "race_id", "age", "course", "distance", "result",
		"jockey_id", "multiple", "affiliation_id", "trainer_id", "gender", "weight",
		"減量騎手", "body_weight", "body_weight_in_de", "popular", "odds", "compi", "compi_num",
		"speed", "waku", "h_num", "間隔", "休み明け〜戦目", "producer", "owner",
		"time_odds", "time_popular", "born_month", "stallion_id", "color_id",
		"furlong_6", "furlong_5", "furlong_4", "furlong_3", "furlong_1",
		"furlong_6_5", "furlong_5_4", "furlong_4_3", "furlong_3_2", "furlong_2_1",
		"training_course", "load_id", "rank", "training_cond", "run_place",
		"1走前結果", "1走前相対着順", "1走前騎手ID", "1走前斤量", "1走前減量騎手か", "1走前体重増減", "1走前人気", "1走前相対人気", "1走前オッズ", "1走前コンピ指数", "1走前コンピ順位", "1走前スピードZI", "1走前枠番", "1走前馬番", "1走前走破タイム", "1走前補正タイム", "1走前着差", "1走前1コーナー順", "1走前2コーナー順", "1走前3コーナー順", "1走前4コーナー順", "1走前ブリンカー着用", "1走前出遅れ", "1走前上がり3F", "1走前上がり3F順", "1走前上がり3F速度", "1走前テン3F速度", "1走前平均速度", "1走前平均1Fタイム", "1走前Ave-3F", "1走前脚質",
		"1走前先行指数", "1走前ペース指数", "1走前上がり指数", "1走前スピード指数", "1走前不利", "1走前馬場指数", "1走前出走頭数", "1走前距離", "1走前場所", "1走前コース", "1走前出走クラス", "1走前1着脚質", "1走前3F平均", "1走前_後5F", "1走前_後4F", "1走前_後3F", "1走前_後2F", "1走前_後1F", "1走前_前5F", "1走前_前4F", "1走前_前3F", "1走前_前2F", "1走前_前1F",
		"1走前ラップ1", "1走前ラップ2", "1走前ラップ3", "1走前ラップ4", "1走前ラップ5", "1走前ラップ6", "1走前ラップ7", "1走前ラップ8", "1走前ラップ9", "1走前ラップ10", "1走前ラップ11", "1走前ラップ12", "1走前ラップ13", "1走前ラップ14", "1走前ラップ15", "1走前ラップ16", "1走前ラップ17", "1走前ラップ18", "1走前ラップ19", "1走前ラップ20",
		"2走前結果", "2走前相対着順", "2走前騎手ID", "2走前斤量", "2走前減量騎手か", "2走前体重増減", "2走前人気", "2走前相対人気", "2走前オッズ", "2走前コンピ指数", "2走前コンピ順位", "2走前スピードZI", "2走前枠番", "2走前馬番", "2走前走破タイム", "2走前補正タイム", "2走前着差", "2走前1コーナー順", "2走前2コーナー順", "2走前3コーナー順", "2走前4コーナー順", "2走前ブリンカー着用", "2走前出遅れ", "2走前上がり3F", "2走前上がり3F順", "2走前上がり3F速度", "2走前テン3F速度", "2走前平均速度", "2走前平均1Fタイム", "2走前Ave-3F", "2走前脚質",
		"2走前先行指数", "2走前ペース指数", "2走前上がり指数", "2走前スピード指数", "2走前不利",
		"2走前馬場指数", "2走前出走頭数", "2走前距離", "2走前場所", "2走前コース", "2走前出走クラス", "2走前1着脚質", "2走前3F平均", "2走前_後5F", "2走前_後4F", "2走前_後3F", "2走前_後2F", "2走前_後1F", "2走前_前5F", "2走前_前4F", "2走前_前3F", "2走前_前2F", "2走前_前1F",
		"2走前ラップ1", "2走前ラップ2", "2走前ラップ3", "2走前ラップ4", "2走前ラップ5", "2走前ラップ6", "2走前ラップ7", "2走前ラップ8", "2走前ラップ9", "2走前ラップ10", "2走前ラップ11", "2走前ラップ12", "2走前ラップ13", "2走前ラップ14", "2走前ラップ15", "2走前ラップ16", "2走前ラップ17", "2走前ラップ18", "2走前ラップ19", "2走前ラップ20",
		"3走前結果", "3走前相対着順", "3走前騎手ID", "3走前斤量", "3走前減量騎手か", "3走前体重増減", "3走前人気", "3走前相対人気", "3走前オッズ", "3走前コンピ指数", "3走前コンピ順位", "3走前スピードZI", "3走前枠番", "3走前馬番", "3走前走破タイム", "3走前補正タイム", "3走前着差", "3走前1コーナー順", "3走前2コーナー順", "3走前3コーナー順", "3走前4コーナー順", "3走前ブリンカー着用", "3走前出遅れ", "3走前上がり3F", "3走前上がり3F順", "3走前上がり3F速度", "3走前テン3F速度", "3走前平均速度", "3走前平均1Fタイム", "3走前Ave-3F", "3走前脚質",
		"3走前先行指数", "3走前ペース指数", "3走前上がり指数", "3走前スピード指数", "3走前不利",
		"3走前馬場指数", "3走前出走頭数", "3走前距離", "3走前場所", "3走前コース", "3走前出走クラス", "3走前1着脚質", "3走前3F平均", "3走前_後5F", "3走前_後4F", "3走前_後3F", "3走前_後2F", "3走前_後1F", "3走前_前5F", "3走前_前4F", "3走前_前3F", "3走前_前2F", "3走前_前1F",
		"3走前ラップ1", "3走前ラップ2", "3走前ラップ3", "3走前ラップ4", "3走前ラップ5", "3走前ラップ6", "3走前ラップ7", "3走前ラップ8", "3走前ラップ9", "3走前ラップ10", "3走前ラップ11", "3走前ラップ12", "3走前ラップ13", "3走前ラップ14", "3走前ラップ15", "3走前ラップ16", "3走前ラップ17", "3走前ラップ18", "3走前ラップ19", "3走前ラップ20",
		"4走前結果", "4走前相対着順", "4走前騎手ID", "4走前斤量", "4走前減量騎手か", "4走前体重増減", "4走前人気", "4走前相対人気", "4走前オッズ", "4走前コンピ指数", "4走前コンピ順位", "4走前スピードZI", "4走前枠番", "4走前馬番", "4走前走破タイム", "4走前補正タイム", "4走前着差", "4走前1コーナー順", "4走前2コーナー順", "4走前3コーナー順", "4走前4コーナー順", "4走前ブリンカー着用", "4走前出遅れ", "4走前上がり3F", "4走前上がり3F順", "4走前上がり3F速度", "4走前テン3F速度", "4走前平均速度", "4走前平均1Fタイム", "4走前Ave-3F", "4走前脚質",
		"4走前先行指数", "4走前ペース指数", "4走前上がり指数", "4走前スピード指数", "4走前不利",
		"4走前馬場指数", "4走前出走頭数", "4走前距離", "4走前場所", "4走前コース", "4走前出走クラス", "4走前1着脚質", "4走前3F平均", "4走前_後5F", "4走前_後4F", "4走前_後3F", "4走前_後2F", "4走前_後1F", "4走前_前5F", "4走前_前4F", "4走前_前3F", "4走前_前2F", "4走前_前1F",
		"4走前ラップ1", "4走前ラップ2", "4走前ラップ3", "4走前ラップ4", "4走前ラップ5", "4走前ラップ6", "4走前ラップ7", "4走前ラップ8", "4走前ラップ9", "4走前ラップ10", "4走前ラップ11", "4走前ラップ12", "4走前ラップ13", "4走前ラップ14", "4走前ラップ15", "4走前ラップ16", "4走前ラップ17", "4走前ラップ18", "4走前ラップ19", "4走前ラップ20",
		"5走前結果", "5走前相対着順", "5走前騎手ID", "5走前斤量", "5走前減量騎手か", "5走前体重増減", "5走前人気", "5走前相対人気", "5走前オッズ", "5走前コンピ指数", "5走前コンピ順位", "5走前スピードZI", "5走前枠番", "5走前馬番", "5走前走破タイム", "5走前補正タイム", "5走前着差", "5走前1コーナー順", "5走前2コーナー順", "5走前3コーナー順", "5走前4コーナー順", "5走前ブリンカー着用", "5走前出遅れ", "5走前上がり3F", "5走前上がり3F順", "5走前上がり3F速度", "5走前テン3F速度", "5走前平均速度", "5走前平均1Fタイム", "5走前Ave-3F", "5走前脚質",
		"5走前先行指数", "5走前ペース指数", "5走前上がり指数", "5走前スピード指数", "5走前不利",
		"5走前馬場指数", "5走前出走頭数", "5走前距離", "5走前場所", "5走前コース", "5走前出走クラス", "5走前1着脚質", "5走前3F平均", "5走前_後5F", "5走前_後4F", "5走前_後3F", "5走前_後2F", "5走前_後1F", "5走前_前5F", "5走前_前4F", "5走前_前3F", "5走前_前2F", "5走前_前1F",
		"5走前ラップ1", "5走前ラップ2", "5走前ラップ3", "5走前ラップ4", "5走前ラップ5", "5走前ラップ6", "5走前ラップ7", "5走前ラップ8", "5走前ラップ9", "5走前ラップ10", "5走前ラップ11", "5走前ラップ12", "5走前ラップ13", "5走前ラップ14", "5走前ラップ15", "5走前ラップ16", "5走前ラップ17", "5走前ラップ18", "5走前ラップ19", "5走前ラップ20"})

	for rows.Next() {
		hr := HorsePillar{}
		err := rows.Scan(
			&hr.id, &hr.date, &hr.horse_id, &hr.race_id, &hr.age, &hr.course, &hr.distance, &hr.result,
			&hr.jockey_id, &hr.multiple, &hr.affiliation_id, &hr.trainer_id, &hr.gender, &hr.weight, &hr.is_loss_jockey, &hr.body_weight, &hr.body_weight_in_de, &hr.popular, &hr.odds, &hr.compi, &hr.compi_num, &hr.speed, &hr.waku, &hr.h_num, &hr.interval_week, &hr.after_break,
			&hr.producer, &hr.owner, &hr.time_odds, &hr.time_popular,
			&hr.born_month, &hr.stallion_id, &hr.color_id,
			&hr.training_urlong_6, &hr.training_furlong_5, &hr.training_furlong_4, &hr.training_furlong_3, &hr.training_furlong_1,
			&hr.training_furlong_6_5, &hr.training_furlong_5_4, &hr.training_furlong_4_3, &hr.training_furlong_3_2, &hr.training_furlong_2_1,
			&hr.training_training_course, &hr.training_load_id, &hr.training_rank, &hr.training_training_cond, &hr.training_run_place,
			&hr.one_result, &hr.one_relative_result, &hr.one_jockey_id, &hr.one_weight, &hr.one_is_loss_jockey, &hr.one_body_weight_inde, &hr.one_popular, &hr.one_relative_popular, &hr.one_odds, &hr.one_compi, &hr.one_compi_num, &hr.one_zi, &hr.one_waku, &hr.one_horse_num, &hr.one_time, &hr.one_correct_time, &hr.one_difference, &hr.one_one_courner, &hr.one_two_courner, &hr.one_three_courner, &hr.one_fourth_courner, &hr.one_is_brinker, &hr.one_deokure, &hr.one_last3f, &hr.one_last3f_num, &hr.one_last_3f_speed, &hr.one_first_3f_speed, &hr.one_speed_avg, &hr.one_avg_1f_time, &hr.one_avg_3f, &hr.one_race_type, &hr.one_preceding_index, &hr.one_pace_index, &hr.one_last3f_index,
			&hr.one_speed_index, &hr.one_unfavorable, &hr.one_baba, &hr.one_h_num, &hr.one_distance, &hr.one_place_id, &hr.one_course, &hr.one_grade, &hr.one_first_fact, &hr.one_ave_3f, &hr.one_last_5f, &hr.one_last_4f, &hr.one_last_3f, &hr.one_last_2f, &hr.one_last_1f, &hr.one_first_5f, &hr.one_first_4f, &hr.one_first_3f, &hr.one_first_2f, &hr.one_first_1f,
			&hr.one_rap1, &hr.one_rap2, &hr.one_rap3, &hr.one_rap4, &hr.one_rap5, &hr.one_rap6, &hr.one_rap7, &hr.one_rap8, &hr.one_rap9, &hr.one_rap10, &hr.one_rap11, &hr.one_rap12, &hr.one_rap13, &hr.one_rap14, &hr.one_rap15, &hr.one_rap16, &hr.one_rap17, &hr.one_rap18, &hr.one_rap19, &hr.one_rap20,
			&hr.two_result, &hr.two_relative_result, &hr.two_jockey_id, &hr.two_weight, &hr.two_is_loss_jockey, &hr.two_body_weight_inde, &hr.two_popular, &hr.two_relative_popular, &hr.two_odds, &hr.two_compi, &hr.two_compi_num, &hr.two_zi, &hr.two_waku, &hr.two_horse_num, &hr.two_time, &hr.two_correct_time, &hr.two_difference, &hr.two_one_courner, &hr.two_two_courner, &hr.two_three_courner, &hr.two_fourth_courner, &hr.two_is_brinker, &hr.two_deokure, &hr.two_last3f, &hr.two_last3f_num, &hr.two_last_3f_speed, &hr.two_first_3f_speed, &hr.two_speed_avg, &hr.two_avg_1f_time, &hr.two_avg_3f, &hr.two_race_type, &hr.two_preceding_index, &hr.two_pace_index, &hr.two_last3f_index, &hr.two_speed_index, &hr.two_unfavorable,
			&hr.two_baba, &hr.two_h_num, &hr.two_distance, &hr.two_place_id, &hr.two_course, &hr.two_grade, &hr.two_first_fact, &hr.two_ave_3f, &hr.two_last_5f, &hr.two_last_4f, &hr.two_last_3f, &hr.two_last_2f, &hr.two_last_1f, &hr.two_first_5f, &hr.two_first_4f, &hr.two_first_3f, &hr.two_first_2f, &hr.two_first_1f,
			&hr.two_rap1, &hr.two_rap2, &hr.two_rap3, &hr.two_rap4, &hr.two_rap5, &hr.two_rap6, &hr.two_rap7, &hr.two_rap8, &hr.two_rap9, &hr.two_rap10, &hr.two_rap11, &hr.two_rap12, &hr.two_rap13, &hr.two_rap14, &hr.two_rap15, &hr.two_rap16, &hr.two_rap17, &hr.two_rap18, &hr.two_rap19, &hr.two_rap20,
			&hr.three_result, &hr.three_relative_result, &hr.three_jockey_id, &hr.three_weight, &hr.three_is_loss_jockey, &hr.three_body_weight_inde, &hr.three_popular, &hr.three_relative_popular, &hr.three_odds, &hr.three_compi, &hr.three_compi_num, &hr.three_zi, &hr.three_waku, &hr.three_horse_num, &hr.three_time, &hr.three_correct_time, &hr.three_difference, &hr.three_one_courner, &hr.three_two_courner, &hr.three_three_courner, &hr.three_fourth_courner, &hr.three_is_brinker, &hr.three_deokure, &hr.three_last3f, &hr.three_last3f_num, &hr.three_last_3f_speed, &hr.three_first_3f_speed, &hr.three_speed_avg, &hr.three_avg_1f_time, &hr.three_avg_3f, &hr.three_race_type,
			&hr.three_preceding_index, &hr.three_pace_index, &hr.three_last3f_index, &hr.three_speed_index, &hr.three_unfavorable,
			&hr.three_baba, &hr.three_h_num, &hr.three_distance, &hr.three_place_id, &hr.three_course, &hr.three_grade, &hr.three_first_fact, &hr.three_ave_3f, &hr.three_last_5f, &hr.three_last_4f, &hr.three_last_3f, &hr.three_last_2f, &hr.three_last_1f, &hr.three_first_5f, &hr.three_first_4f, &hr.three_first_3f, &hr.three_first_2f, &hr.three_first_1f,
			&hr.three_rap1, &hr.three_rap2, &hr.three_rap3, &hr.three_rap4, &hr.three_rap5, &hr.three_rap6, &hr.three_rap7, &hr.three_rap8, &hr.three_rap9, &hr.three_rap10, &hr.three_rap11, &hr.three_rap12, &hr.three_rap13, &hr.three_rap14, &hr.three_rap15, &hr.three_rap16, &hr.three_rap17, &hr.three_rap18, &hr.three_rap19, &hr.three_rap20,
			&hr.four_result, &hr.four_relative_result, &hr.four_jockey_id, &hr.four_weight, &hr.four_is_loss_jockey, &hr.four_body_weight_inde, &hr.four_popular, &hr.four_relative_popular, &hr.four_odds, &hr.four_compi, &hr.four_compi_num, &hr.four_zi, &hr.four_waku, &hr.four_horse_num, &hr.four_time, &hr.four_correct_time,
			&hr.four_difference, &hr.four_one_courner, &hr.four_two_courner, &hr.four_three_courner, &hr.four_fourth_courner, &hr.four_is_brinker, &hr.four_deokure, &hr.four_last3f, &hr.four_last3f_num, &hr.four_last_3f_speed, &hr.four_first_3f_speed, &hr.four_speed_avg, &hr.four_avg_1f_time, &hr.four_avg_3f, &hr.four_race_type, &hr.four_preceding_index, &hr.four_pace_index, &hr.four_last3f_index, &hr.four_speed_index, &hr.four_unfavorable,
			&hr.four_baba, &hr.four_h_num, &hr.four_distance, &hr.four_place_id, &hr.four_course, &hr.four_grade, &hr.four_first_fact, &hr.four_ave_3f, &hr.four_last_5f, &hr.four_last_4f, &hr.four_last_3f, &hr.four_last_2f, &hr.four_last_1f, &hr.four_first_5f, &hr.four_first_4f, &hr.four_first_3f, &hr.four_first_2f, &hr.four_first_1f,
			&hr.four_rap1, &hr.four_rap2, &hr.four_rap3, &hr.four_rap4, &hr.four_rap5, &hr.four_rap6, &hr.four_rap7, &hr.four_rap8, &hr.four_rap9, &hr.four_rap10, &hr.four_rap11, &hr.four_rap12, &hr.four_rap13, &hr.four_rap14, &hr.four_rap15, &hr.four_rap16, &hr.four_rap17, &hr.four_rap18, &hr.four_rap19, &hr.four_rap20,
			&hr.five_result, &hr.five_relative_result, &hr.five_jockey_id, &hr.five_weight, &hr.five_is_loss_jockey, &hr.five_body_weight_inde, &hr.five_popular, &hr.five_relative_popular, &hr.five_odds, &hr.five_compi, &hr.five_compi_num, &hr.five_zi, &hr.five_waku, &hr.five_horse_num, &hr.five_time, &hr.five_correct_time, &hr.five_difference,
			&hr.five_one_courner, &hr.five_two_courner, &hr.five_three_courner, &hr.five_fourth_courner, &hr.five_is_brinker, &hr.five_deokure, &hr.five_last3f, &hr.five_last3f_num, &hr.five_last_3f_speed, &hr.five_first_3f_speed, &hr.five_speed_avg, &hr.five_avg_1f_time, &hr.five_avg_3f, &hr.five_race_type, &hr.five_preceding_index, &hr.five_pace_index, &hr.five_last3f_index, &hr.five_speed_index, &hr.five_unfavorable,
			&hr.five_baba, &hr.five_h_num, &hr.five_distance, &hr.five_place_id, &hr.five_course, &hr.five_grade, &hr.five_first_fact, &hr.five_ave_3f, &hr.five_last_5f, &hr.five_last_4f, &hr.five_last_3f, &hr.five_last_2f, &hr.five_last_1f, &hr.five_first_5f, &hr.five_first_4f, &hr.five_first_3f, &hr.five_first_2f, &hr.five_first_1f,
			&hr.five_rap1, &hr.five_rap2, &hr.five_rap3, &hr.five_rap4, &hr.five_rap5, &hr.five_rap6, &hr.five_rap7, &hr.five_rap8, &hr.five_rap9, &hr.five_rap10, &hr.five_rap11, &hr.five_rap12, &hr.five_rap13, &hr.five_rap14, &hr.five_rap15, &hr.five_rap16, &hr.five_rap17, &hr.five_rap18, &hr.five_rap19, &hr.five_rap20)

		if err != nil {
			panic(err)
		}

		pillars = append(pillars, []string{
			hr.id.MarshalJSON(), hr.date.MarshalJSON(), hr.horse_id.MarshalJSON(), hr.race_id.MarshalJSON(), hr.age.MarshalJSON(), hr.course.MarshalJSON(), hr.distance.MarshalJSON(),
			hr.result.MarshalJSON(), hr.jockey_id.MarshalJSON(), hr.multiple.MarshalJSON(), hr.affiliation_id.MarshalJSON(), hr.trainer_id.MarshalJSON(), hr.gender.MarshalJSON(), hr.weight.MarshalJSON(), hr.is_loss_jockey.MarshalJSON(), hr.body_weight.MarshalJSON(), hr.body_weight_in_de.MarshalJSON(), hr.popular.MarshalJSON(), hr.odds.MarshalJSON(), hr.compi.MarshalJSON(), hr.compi_num.MarshalJSON(), hr.speed.MarshalJSON(), hr.waku.MarshalJSON(), hr.h_num.MarshalJSON(), hr.interval_week.MarshalJSON(), hr.after_break.MarshalJSON(), hr.producer.MarshalJSON(), hr.owner.MarshalJSON(),
			hr.time_odds.MarshalJSON(), hr.time_popular.MarshalJSON(), hr.born_month.MarshalJSON(), hr.stallion_id.MarshalJSON(), hr.color_id.MarshalJSON(),
			hr.training_urlong_6.MarshalJSON(), hr.training_furlong_5.MarshalJSON(), hr.training_furlong_4.MarshalJSON(), hr.training_furlong_3.MarshalJSON(), hr.training_furlong_1.MarshalJSON(),
			hr.training_furlong_6_5.MarshalJSON(), hr.training_furlong_5_4.MarshalJSON(), hr.training_furlong_4_3.MarshalJSON(), hr.training_furlong_3_2.MarshalJSON(), hr.training_furlong_2_1.MarshalJSON(),
			hr.training_training_course.MarshalJSON(), hr.training_load_id.MarshalJSON(), hr.training_rank.MarshalJSON(), hr.training_training_cond.MarshalJSON(), hr.training_run_place.MarshalJSON(),
			hr.one_result.MarshalJSON(), hr.one_relative_result.MarshalJSON(), hr.one_jockey_id.MarshalJSON(), hr.one_weight.MarshalJSON(), hr.one_is_loss_jockey.MarshalJSON(), hr.one_body_weight_inde.MarshalJSON(), hr.one_popular.MarshalJSON(), hr.one_relative_popular.MarshalJSON(), hr.one_odds.MarshalJSON(), hr.one_compi.MarshalJSON(), hr.one_compi_num.MarshalJSON(), hr.one_zi.MarshalJSON(), hr.one_waku.MarshalJSON(), hr.one_horse_num.MarshalJSON(), hr.one_time.MarshalJSON(), hr.one_correct_time.MarshalJSON(), hr.one_difference.MarshalJSON(), hr.one_one_courner.MarshalJSON(), hr.one_two_courner.MarshalJSON(), hr.one_three_courner.MarshalJSON(), hr.one_fourth_courner.MarshalJSON(), hr.one_is_brinker.MarshalJSON(), hr.one_deokure.MarshalJSON(), hr.one_last3f.MarshalJSON(), hr.one_last3f_num.MarshalJSON(), hr.one_last_3f_speed.MarshalJSON(), hr.one_first_3f_speed.MarshalJSON(), hr.one_speed_avg.MarshalJSON(),
			hr.one_avg_1f_time.MarshalJSON(), hr.one_avg_3f.MarshalJSON(), hr.one_race_type.MarshalJSON(), hr.one_preceding_index.MarshalJSON(), hr.one_pace_index.MarshalJSON(), hr.one_last3f_index.MarshalJSON(), hr.one_speed_index.MarshalJSON(), hr.one_unfavorable.MarshalJSON(),
			hr.one_baba.MarshalJSON(), hr.one_h_num.MarshalJSON(), hr.one_distance.MarshalJSON(), hr.one_place_id.MarshalJSON(), hr.one_course.MarshalJSON(), hr.one_grade.MarshalJSON(), hr.one_first_fact.MarshalJSON(), hr.one_ave_3f.MarshalJSON(), hr.one_last_5f.MarshalJSON(), hr.one_last_4f.MarshalJSON(), hr.one_last_3f.MarshalJSON(), hr.one_last_2f.MarshalJSON(), hr.one_last_1f.MarshalJSON(), hr.one_first_5f.MarshalJSON(), hr.one_first_4f.MarshalJSON(), hr.one_first_3f.MarshalJSON(), hr.one_first_2f.MarshalJSON(), hr.one_first_1f.MarshalJSON(),
			hr.one_rap1.MarshalJSON(), hr.one_rap2.MarshalJSON(), hr.one_rap3.MarshalJSON(), hr.one_rap4.MarshalJSON(), hr.one_rap5.MarshalJSON(), hr.one_rap6.MarshalJSON(), hr.one_rap7.MarshalJSON(), hr.one_rap8.MarshalJSON(), hr.one_rap9.MarshalJSON(), hr.one_rap10.MarshalJSON(), hr.one_rap11.MarshalJSON(), hr.one_rap12.MarshalJSON(), hr.one_rap13.MarshalJSON(), hr.one_rap14.MarshalJSON(), hr.one_rap15.MarshalJSON(), hr.one_rap16.MarshalJSON(), hr.one_rap17.MarshalJSON(), hr.one_rap18.MarshalJSON(), hr.one_rap19.MarshalJSON(), hr.one_rap20.MarshalJSON(),
			hr.two_result.MarshalJSON(), hr.two_relative_result.MarshalJSON(), hr.two_jockey_id.MarshalJSON(), hr.two_weight.MarshalJSON(), hr.two_is_loss_jockey.MarshalJSON(), hr.two_body_weight_inde.MarshalJSON(), hr.two_popular.MarshalJSON(), hr.two_relative_popular.MarshalJSON(), hr.two_odds.MarshalJSON(), hr.two_compi.MarshalJSON(), hr.two_compi_num.MarshalJSON(), hr.two_zi.MarshalJSON(), hr.two_waku.MarshalJSON(), hr.two_horse_num.MarshalJSON(), hr.two_time.MarshalJSON(), hr.two_correct_time.MarshalJSON(), hr.two_difference.MarshalJSON(), hr.two_one_courner.MarshalJSON(), hr.two_two_courner.MarshalJSON(), hr.two_three_courner.MarshalJSON(), hr.two_fourth_courner.MarshalJSON(), hr.two_is_brinker.MarshalJSON(), hr.two_deokure.MarshalJSON(), hr.two_last3f.MarshalJSON(), hr.two_last3f_num.MarshalJSON(), hr.two_last_3f_speed.MarshalJSON(), hr.two_first_3f_speed.MarshalJSON(), hr.two_speed_avg.MarshalJSON(),
			hr.two_avg_1f_time.MarshalJSON(), hr.two_avg_3f.MarshalJSON(), hr.two_race_type.MarshalJSON(), hr.two_preceding_index.MarshalJSON(), hr.two_pace_index.MarshalJSON(), hr.two_last3f_index.MarshalJSON(), hr.two_speed_index.MarshalJSON(), hr.two_unfavorable.MarshalJSON(),
			hr.two_baba.MarshalJSON(), hr.two_h_num.MarshalJSON(), hr.two_distance.MarshalJSON(), hr.two_place_id.MarshalJSON(), hr.two_course.MarshalJSON(), hr.two_grade.MarshalJSON(), hr.two_first_fact.MarshalJSON(), hr.two_ave_3f.MarshalJSON(), hr.two_last_5f.MarshalJSON(), hr.two_last_4f.MarshalJSON(), hr.two_last_3f.MarshalJSON(), hr.two_last_2f.MarshalJSON(), hr.two_last_1f.MarshalJSON(), hr.two_first_5f.MarshalJSON(), hr.two_first_4f.MarshalJSON(), hr.two_first_3f.MarshalJSON(), hr.two_first_2f.MarshalJSON(), hr.two_first_1f.MarshalJSON(),
			hr.two_rap1.MarshalJSON(), hr.two_rap2.MarshalJSON(), hr.two_rap3.MarshalJSON(), hr.two_rap4.MarshalJSON(), hr.two_rap5.MarshalJSON(), hr.two_rap6.MarshalJSON(), hr.two_rap7.MarshalJSON(), hr.two_rap8.MarshalJSON(), hr.two_rap9.MarshalJSON(), hr.two_rap10.MarshalJSON(), hr.two_rap11.MarshalJSON(), hr.two_rap12.MarshalJSON(), hr.two_rap13.MarshalJSON(), hr.two_rap14.MarshalJSON(), hr.two_rap15.MarshalJSON(), hr.two_rap16.MarshalJSON(), hr.two_rap17.MarshalJSON(), hr.two_rap18.MarshalJSON(), hr.two_rap19.MarshalJSON(), hr.two_rap20.MarshalJSON(),
			hr.three_result.MarshalJSON(), hr.three_relative_result.MarshalJSON(), hr.three_jockey_id.MarshalJSON(), hr.three_weight.MarshalJSON(), hr.three_is_loss_jockey.MarshalJSON(), hr.three_body_weight_inde.MarshalJSON(), hr.three_popular.MarshalJSON(), hr.three_relative_popular.MarshalJSON(), hr.three_odds.MarshalJSON(), hr.three_compi.MarshalJSON(), hr.three_compi_num.MarshalJSON(), hr.three_zi.MarshalJSON(), hr.three_waku.MarshalJSON(), hr.three_horse_num.MarshalJSON(), hr.three_time.MarshalJSON(), hr.three_correct_time.MarshalJSON(), hr.three_difference.MarshalJSON(), hr.three_one_courner.MarshalJSON(), hr.three_two_courner.MarshalJSON(), hr.three_three_courner.MarshalJSON(), hr.three_fourth_courner.MarshalJSON(), hr.three_is_brinker.MarshalJSON(), hr.three_deokure.MarshalJSON(), hr.three_last3f.MarshalJSON(), hr.three_last3f_num.MarshalJSON(), hr.three_last_3f_speed.MarshalJSON(), hr.three_first_3f_speed.MarshalJSON(),
			hr.three_speed_avg.MarshalJSON(), hr.three_avg_1f_time.MarshalJSON(), hr.three_avg_3f.MarshalJSON(), hr.three_race_type.MarshalJSON(), hr.three_preceding_index.MarshalJSON(), hr.three_pace_index.MarshalJSON(), hr.three_last3f_index.MarshalJSON(), hr.three_speed_index.MarshalJSON(), hr.three_unfavorable.MarshalJSON(),
			hr.three_baba.MarshalJSON(), hr.three_h_num.MarshalJSON(), hr.three_distance.MarshalJSON(), hr.three_place_id.MarshalJSON(), hr.three_course.MarshalJSON(), hr.three_grade.MarshalJSON(), hr.three_first_fact.MarshalJSON(), hr.three_ave_3f.MarshalJSON(), hr.three_last_5f.MarshalJSON(), hr.three_last_4f.MarshalJSON(), hr.three_last_3f.MarshalJSON(), hr.three_last_2f.MarshalJSON(), hr.three_last_1f.MarshalJSON(), hr.three_first_5f.MarshalJSON(), hr.three_first_4f.MarshalJSON(), hr.three_first_3f.MarshalJSON(), hr.three_first_2f.MarshalJSON(), hr.three_first_1f.MarshalJSON(),
			hr.three_rap1.MarshalJSON(), hr.three_rap2.MarshalJSON(), hr.three_rap3.MarshalJSON(), hr.three_rap4.MarshalJSON(), hr.three_rap5.MarshalJSON(), hr.three_rap6.MarshalJSON(), hr.three_rap7.MarshalJSON(), hr.three_rap8.MarshalJSON(), hr.three_rap9.MarshalJSON(), hr.three_rap10.MarshalJSON(), hr.three_rap11.MarshalJSON(), hr.three_rap12.MarshalJSON(), hr.three_rap13.MarshalJSON(), hr.three_rap14.MarshalJSON(), hr.three_rap15.MarshalJSON(), hr.three_rap16.MarshalJSON(), hr.three_rap17.MarshalJSON(), hr.three_rap18.MarshalJSON(), hr.three_rap19.MarshalJSON(), hr.three_rap20.MarshalJSON(),
			hr.four_result.MarshalJSON(), hr.four_relative_result.MarshalJSON(), hr.four_jockey_id.MarshalJSON(), hr.four_weight.MarshalJSON(), hr.four_is_loss_jockey.MarshalJSON(), hr.four_body_weight_inde.MarshalJSON(), hr.four_popular.MarshalJSON(), hr.four_relative_popular.MarshalJSON(), hr.four_odds.MarshalJSON(), hr.four_compi.MarshalJSON(), hr.four_compi_num.MarshalJSON(), hr.four_zi.MarshalJSON(), hr.four_waku.MarshalJSON(), hr.four_horse_num.MarshalJSON(), hr.four_time.MarshalJSON(), hr.four_correct_time.MarshalJSON(),
			hr.four_difference.MarshalJSON(), hr.four_one_courner.MarshalJSON(), hr.four_two_courner.MarshalJSON(), hr.four_three_courner.MarshalJSON(), hr.four_fourth_courner.MarshalJSON(), hr.four_is_brinker.MarshalJSON(), hr.four_deokure.MarshalJSON(), hr.four_last3f.MarshalJSON(), hr.four_last3f_num.MarshalJSON(), hr.four_last_3f_speed.MarshalJSON(), hr.four_first_3f_speed.MarshalJSON(), hr.four_speed_avg.MarshalJSON(), hr.four_avg_1f_time.MarshalJSON(), hr.four_avg_3f.MarshalJSON(), hr.four_race_type.MarshalJSON(), hr.four_preceding_index.MarshalJSON(), hr.four_pace_index.MarshalJSON(), hr.four_last3f_index.MarshalJSON(), hr.four_speed_index.MarshalJSON(), hr.four_unfavorable.MarshalJSON(),
			hr.four_baba.MarshalJSON(), hr.four_h_num.MarshalJSON(), hr.four_distance.MarshalJSON(), hr.four_place_id.MarshalJSON(), hr.four_course.MarshalJSON(), hr.four_grade.MarshalJSON(), hr.four_first_fact.MarshalJSON(), hr.four_ave_3f.MarshalJSON(), hr.four_last_5f.MarshalJSON(), hr.four_last_4f.MarshalJSON(), hr.four_last_3f.MarshalJSON(), hr.four_last_2f.MarshalJSON(), hr.four_last_1f.MarshalJSON(), hr.four_first_5f.MarshalJSON(), hr.four_first_4f.MarshalJSON(), hr.four_first_3f.MarshalJSON(), hr.four_first_2f.MarshalJSON(), hr.four_first_1f.MarshalJSON(),
			hr.four_rap1.MarshalJSON(), hr.four_rap2.MarshalJSON(), hr.four_rap3.MarshalJSON(), hr.four_rap4.MarshalJSON(), hr.four_rap5.MarshalJSON(), hr.four_rap6.MarshalJSON(), hr.four_rap7.MarshalJSON(), hr.four_rap8.MarshalJSON(), hr.four_rap9.MarshalJSON(), hr.four_rap10.MarshalJSON(), hr.four_rap11.MarshalJSON(), hr.four_rap12.MarshalJSON(), hr.four_rap13.MarshalJSON(), hr.four_rap14.MarshalJSON(), hr.four_rap15.MarshalJSON(), hr.four_rap16.MarshalJSON(), hr.four_rap17.MarshalJSON(), hr.four_rap18.MarshalJSON(), hr.four_rap19.MarshalJSON(), hr.four_rap20.MarshalJSON(),
			hr.five_result.MarshalJSON(), hr.five_relative_result.MarshalJSON(), hr.five_jockey_id.MarshalJSON(), hr.five_weight.MarshalJSON(), hr.five_is_loss_jockey.MarshalJSON(), hr.five_body_weight_inde.MarshalJSON(), hr.five_popular.MarshalJSON(), hr.five_relative_popular.MarshalJSON(), hr.five_odds.MarshalJSON(), hr.five_compi.MarshalJSON(), hr.five_compi_num.MarshalJSON(), hr.five_zi.MarshalJSON(), hr.five_waku.MarshalJSON(), hr.five_horse_num.MarshalJSON(), hr.five_time.MarshalJSON(), hr.five_correct_time.MarshalJSON(), hr.five_difference.MarshalJSON(),
			hr.five_one_courner.MarshalJSON(), hr.five_two_courner.MarshalJSON(), hr.five_three_courner.MarshalJSON(), hr.five_fourth_courner.MarshalJSON(), hr.five_is_brinker.MarshalJSON(), hr.five_deokure.MarshalJSON(), hr.five_last3f.MarshalJSON(), hr.five_last3f_num.MarshalJSON(), hr.five_last_3f_speed.MarshalJSON(), hr.five_first_3f_speed.MarshalJSON(), hr.five_speed_avg.MarshalJSON(), hr.five_avg_1f_time.MarshalJSON(), hr.five_avg_3f.MarshalJSON(), hr.five_race_type.MarshalJSON(), hr.five_preceding_index.MarshalJSON(), hr.five_pace_index.MarshalJSON(), hr.five_last3f_index.MarshalJSON(), hr.five_speed_index.MarshalJSON(), hr.five_unfavorable.MarshalJSON(),
			hr.five_baba.MarshalJSON(), hr.five_h_num.MarshalJSON(), hr.five_distance.MarshalJSON(), hr.five_place_id.MarshalJSON(), hr.five_course.MarshalJSON(), hr.five_grade.MarshalJSON(), hr.five_first_fact.MarshalJSON(), hr.five_ave_3f.MarshalJSON(), hr.five_last_5f.MarshalJSON(), hr.five_last_4f.MarshalJSON(), hr.five_last_3f.MarshalJSON(), hr.five_last_2f.MarshalJSON(), hr.five_last_1f.MarshalJSON(), hr.five_first_5f.MarshalJSON(), hr.five_first_4f.MarshalJSON(), hr.five_first_3f.MarshalJSON(), hr.five_first_2f.MarshalJSON(), hr.five_first_1f.MarshalJSON(),
			hr.five_rap1.MarshalJSON(), hr.five_rap2.MarshalJSON(), hr.five_rap3.MarshalJSON(), hr.five_rap4.MarshalJSON(), hr.five_rap5.MarshalJSON(), hr.five_rap6.MarshalJSON(), hr.five_rap7.MarshalJSON(), hr.five_rap8.MarshalJSON(), hr.five_rap9.MarshalJSON(), hr.five_rap10.MarshalJSON(), hr.five_rap11.MarshalJSON(), hr.five_rap12.MarshalJSON(), hr.five_rap13.MarshalJSON(), hr.five_rap14.MarshalJSON(), hr.five_rap15.MarshalJSON(), hr.five_rap16.MarshalJSON(), hr.five_rap17.MarshalJSON(), hr.five_rap18.MarshalJSON(), hr.five_rap19.MarshalJSON(), hr.five_rap20.MarshalJSON()})
	}

	f, _ := os.OpenFile("./data/horse_results.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(pillars)

	w.Flush()
}

func get_jockey_rate(date string, db *sql.DB) {
	rows, _ := db.Query("select * from jockey_rate where date = ?", date)

	jockies := make([][]string, 0)
	jockies = append(jockies, []string{
		"id", "horse_race_id", "jockey_id", "place_id", "course", "distance", "date", "騎手騎乗回数", "騎手全体勝率", "騎手全体連対率", "騎手全体複勝率", "騎手競馬場別騎乗回数", "騎手競馬場別勝率",
		"騎手競馬場別連対率", "騎手競馬場別複勝率", "騎手コース別騎乗回数", "騎手コース別勝率", "騎手コース別連対率", "騎手コース別複勝率", "騎手距離別騎乗回数", "騎手距離別勝率", "騎手距離別連対率", "騎手距離別複勝率", "騎手同コース同距離別騎乗回数", "騎手同コース同距離別勝率", "騎手同コース同距離別連対率", "騎手同コース同距離別複勝率"})

	for rows.Next() {
		j := JockeyRate{}
		err := rows.Scan(
			&j.id,
			&j.horse_race_id,
			&j.jockey_id,
			&j.place_id,
			&j.course,
			&j.distance,
			&j.date,
			&j.count,
			&j.win_rate,
			&j.rentai_rate,
			&j.fukusho_rate,
			&j.place_count,
			&j.place_win_rate,
			&j.place_rentai_rate,
			&j.place_fukusho_rate,
			&j.course_count,
			&j.course_win_rate,
			&j.course_rentai_rate,
			&j.course_fukusho_rate,
			&j.distance_count,
			&j.distance_win_rate,
			&j.distance_rentai_rate,
			&j.distance_fukusho_rate,
			&j.place_distance_count,
			&j.place_distance_win_rate,
			&j.place_distance_rentai_rate,
			&j.place_distance_fukusho_rate)

		if err != nil {
			panic(err)
		}

		jockies = append(jockies, []string{
			j.id.MarshalJSON(),
			j.horse_race_id.MarshalJSON(),
			j.jockey_id.MarshalJSON(),
			j.place_id.MarshalJSON(),
			j.course.MarshalJSON(),
			j.distance.MarshalJSON(),
			j.date.MarshalJSON(),
			j.count.MarshalJSON(),
			j.win_rate.MarshalJSON(),
			j.rentai_rate.MarshalJSON(),
			j.fukusho_rate.MarshalJSON(),
			j.place_count.MarshalJSON(),
			j.place_win_rate.MarshalJSON(),
			j.place_rentai_rate.MarshalJSON(),
			j.place_fukusho_rate.MarshalJSON(),
			j.course_count.MarshalJSON(),
			j.course_win_rate.MarshalJSON(),
			j.course_rentai_rate.MarshalJSON(),
			j.course_fukusho_rate.MarshalJSON(),
			j.distance_count.MarshalJSON(),
			j.distance_win_rate.MarshalJSON(),
			j.distance_rentai_rate.MarshalJSON(),
			j.distance_fukusho_rate.MarshalJSON(),
			j.place_distance_count.MarshalJSON(),
			j.place_distance_win_rate.MarshalJSON(),
			j.place_distance_rentai_rate.MarshalJSON(),
			j.place_distance_fukusho_rate.MarshalJSON()})
	}

	f, _ := os.OpenFile("./data/jockey_rate.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(jockies)

	w.Flush()
}

func get_trainer_rate(date string, db *sql.DB) {
	rows, _ := db.Query("select * from trainer_rate where date = ?", date)

	trainers := make([][]string, 0)
	trainers = append(trainers, []string{
		"id", "horse_race_id", "trainer_id", "place_id", "course", "distance", "date",
		"調教師出走回数", "調教師全体勝率", "調教師全体連対率", "調教師全体複勝率", "調教師競馬場別騎乗回数", "調教師競馬場別勝率",
		"調教師競馬場別連対率", "調教師競馬場別複勝率", "調教師コース別騎乗回数", "調教師コース別勝率", "調教師コース別連対率",
		"調教師コース別複勝率", "調教師距離別騎乗回数", "調教師距離別勝率", "調教師距離別連対率", "調教師距離別複勝率", "調教師同コース同距離別騎乗回数",
		"調教師同コース同距離別勝率", "調教師同コース同距離別連対率", "調教師同コース同距離別複勝率", "調教師年齢別年間勝率",
		"調教師年齢別年間連対率", "調教師年齢別年間複勝率", "調教師年齢別勝率", "調教師年齢別連対率", "調教師年齢別複勝率"})

	for rows.Next() {
		j := TrainerRate{}
		err := rows.Scan(
			&j.id,
			&j.horse_race_id,
			&j.trainer_id,
			&j.place_id,
			&j.course,
			&j.distance,
			&j.date,
			&j.count,
			&j.win_rate,
			&j.rentai_rate,
			&j.fukusho_rate,
			&j.place_count,
			&j.place_win_rate,
			&j.place_rentai_rate,
			&j.place_fukusho_rate,
			&j.course_count,
			&j.course_win_rate,
			&j.course_rentai_rate,
			&j.course_fukusho_rate,
			&j.distance_count,
			&j.distance_win_rate,
			&j.distance_rentai_rate,
			&j.distance_fukusho_rate,
			&j.place_distance_count,
			&j.place_distance_win_rate,
			&j.place_distance_rentai_rate,
			&j.place_distance_fukusho_rate,
			&j.age_year_win_rate,
			&j.age_year_rentai_rate,
			&j.age_year_fukusho_rate,
			&j.age_win_rate,
			&j.age_rentai_rate,
			&j.age_fukusho_rate,
		)

		if err != nil {
			panic(err)
		}

		trainers = append(trainers, []string{
			j.id.MarshalJSON(),
			j.horse_race_id.MarshalJSON(),
			j.trainer_id.MarshalJSON(),
			j.place_id.MarshalJSON(),
			j.course.MarshalJSON(),
			j.distance.MarshalJSON(),
			j.date.MarshalJSON(),
			j.count.MarshalJSON(),
			j.win_rate.MarshalJSON(),
			j.rentai_rate.MarshalJSON(),
			j.fukusho_rate.MarshalJSON(),
			j.place_count.MarshalJSON(),
			j.place_win_rate.MarshalJSON(),
			j.place_rentai_rate.MarshalJSON(),
			j.place_fukusho_rate.MarshalJSON(),
			j.course_count.MarshalJSON(),
			j.course_win_rate.MarshalJSON(),
			j.course_rentai_rate.MarshalJSON(),
			j.course_fukusho_rate.MarshalJSON(),
			j.distance_count.MarshalJSON(),
			j.distance_win_rate.MarshalJSON(),
			j.distance_rentai_rate.MarshalJSON(),
			j.distance_fukusho_rate.MarshalJSON(),
			j.place_distance_count.MarshalJSON(),
			j.place_distance_win_rate.MarshalJSON(),
			j.place_distance_rentai_rate.MarshalJSON(),
			j.age_fukusho_rate.MarshalJSON(),
			j.age_year_win_rate.MarshalJSON(),
			j.age_year_rentai_rate.MarshalJSON(),
			j.age_year_fukusho_rate.MarshalJSON(),
			j.age_win_rate.MarshalJSON(),
			j.age_rentai_rate.MarshalJSON(),
			j.age_fukusho_rate.MarshalJSON()})
	}

	f, _ := os.OpenFile("./data/trainer_rate.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(trainers)

	w.Flush()
}

func get_stallion_rate(date string, db *sql.DB) {
	rows, _ := db.Query("select * from stallion_rate where date = ?", date)
	stallions := make([][]string, 0)
	stallions = append(stallions, []string{
		"race_id", "name", "place_id", "courner_type", "course", "distance", "stallion_id", "date", "start_time",
		"種牡馬出走頭数", "種牡馬全体勝率", "種牡馬全体連対率", "種牡馬全体複勝率",
		"種牡馬競馬場別出走頭数", "種牡馬競馬場別勝率", "種牡馬競馬場別連対率", "種牡馬競馬場別複勝率",
		"種牡馬コース別出走頭数", "種牡馬コース別勝率", "種牡馬コース別連対率", "種牡馬コース別複勝率",
		"種牡馬距離別出走頭数", "種牡馬距離別勝率", "種牡馬距離別連対率", "種牡馬距離別複勝率",
		"種牡馬同コース同距離別出走頭数", "種牡馬同コース同距離別勝率", "種牡馬同コース同距離別連対率", "種牡馬同コース同距離別複勝率",
		"種牡馬同周り出走頭数", "種牡馬同周り勝率", "種牡馬同周り連対率", "種牡馬同周り複勝率",
		"種牡馬同枠勝率", "種牡馬同枠連対率", "種牡馬同枠複勝率"})

	for rows.Next() {
		s := StallionRate{}
		err := rows.Scan(
			&s.race_id,
			&s.name,
			&s.courner_type,
			&s.course,
			&s.distance,
			&s.stallion_id,
			&s.date,
			&s.start_time,
			&s.count,
			&s.win_rate,
			&s.rentai_rate,
			&s.fukushou_rate,
			&s.place_count,
			&s.place_win_rate,
			&s.place_rentai_rate,
			&s.place_fukushou_rate,
			&s.course_count,
			&s.course_win_rate,
			&s.course_rentai_rate,
			&s.course_fukushou_rate,
			&s.distance_count,
			&s.distance_win_rate,
			&s.distance_rentai_rate,
			&s.distance_fukushou_rate,
			&s.course_distance_count,
			&s.course_distance_win_rate,
			&s.course_distance_rentai_rate,
			&s.course_distance_fukushou_rate,
			&s.circle_count,
			&s.circle_win_rate,
			&s.circle_rentai_rate,
			&s.circle_fukushou_rate,
			&s.waku_count,
			&s.waku_win_rate,
			&s.waku_rentai_rate,
			&s.waku_fukushou_rate,
		)

		if err != nil {
			panic(err)
		}

		stallions = append(stallions, []string{
			s.race_id.MarshalJSON(),
			s.name.MarshalJSON(),
			s.courner_type.MarshalJSON(),
			s.course.MarshalJSON(),
			s.distance.MarshalJSON(),
			s.stallion_id.MarshalJSON(),
			s.date.MarshalJSON(),
			s.start_time.MarshalJSON(),
			s.count.MarshalJSON(),
			s.win_rate.MarshalJSON(),
			s.rentai_rate.MarshalJSON(),
			s.fukushou_rate.MarshalJSON(),
			s.place_count.MarshalJSON(),
			s.place_win_rate.MarshalJSON(),
			s.place_rentai_rate.MarshalJSON(),
			s.place_fukushou_rate.MarshalJSON(),
			s.course_count.MarshalJSON(),
			s.course_win_rate.MarshalJSON(),
			s.course_rentai_rate.MarshalJSON(),
			s.course_fukushou_rate.MarshalJSON(),
			s.distance_count.MarshalJSON(),
			s.distance_win_rate.MarshalJSON(),
			s.distance_rentai_rate.MarshalJSON(),
			s.distance_fukushou_rate.MarshalJSON(),
			s.course_count.MarshalJSON(),
			s.course_distance_win_rate.MarshalJSON(),
			s.course_distance_rentai_rate.MarshalJSON(),
			s.course_distance_fukushou_rate.MarshalJSON(),
			s.circle_count.MarshalJSON(),
			s.circle_win_rate.MarshalJSON(),
			s.circle_rentai_rate.MarshalJSON(),
			s.circle_fukushou_rate.MarshalJSON(),
			s.waku_count.MarshalJSON(),
			s.waku_win_rate.MarshalJSON(),
			s.waku_rentai_rate.MarshalJSON(),
			s.waku_fukushou_rate.MarshalJSON()})
	}

	f, _ := os.OpenFile("./data/stallion_rate.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(stallions)

	w.Flush()
}

func get_horse_sire_rate(date string, db *sql.DB) {
	rows, _ := db.Query("select * from horse_sire_rate where date = ?", date)
	sires := make([][]string, 0)
	sires = append(sires, []string{
		"id", "race_id", "name", "place_id", "course", "distance", "sire_id", "date", "start_time",
		"父系統出走頭数", "父系統全体勝率", "父系統全体連対率", "父系統全体複勝率",
		"父系統競馬場別出走頭数", "父系統競馬場別勝率", "父系統競馬場別連対率", "父系統競馬場別複勝率",
		"父系統コース別出走頭数", "父系統コース別勝率", "父系統コース別連対率", "父系統コース別複勝率",
		"父系統距離別出走頭数", "父系統距離別勝率", "父系統距離別連対率", "父系統距離別複勝率",
		"父系統同コース同距離別出走頭数", "父系統同コース同距離別勝率", "父系統同コース同距離別連対率", "父系統同コース同距離別複勝率"})

	for rows.Next() {
		s := HorseSireRate{}
		err := rows.Scan(
			&s.id,
			&s.race_id,
			&s.name,
			&s.place_id,
			&s.course,
			&s.distance,
			&s.sire_id,
			&s.date,
			&s.start_time,
			&s.count,
			&s.win_rate,
			&s.rentai_rate,
			&s.fukushou_rate,
			&s.place_count,
			&s.place_win_rate,
			&s.place_rentai_rate,
			&s.place_fukushou_rate,
			&s.course_count,
			&s.course_win_rate,
			&s.course_rentai_rate,
			&s.course_fukushou_rate,
			&s.distance_count,
			&s.distance_win_rate,
			&s.distance_rentai_rate,
			&s.distance_fukushou_rate,
			&s.course_distance_count,
			&s.course_distance_win_rate,
			&s.course_distance_rentai_rate,
			&s.course_distance_fukushou_rate,
		)

		if err != nil {
			panic(err)
		}

		sires = append(sires, []string{
			s.id.MarshalJSON(),
			s.race_id.MarshalJSON(),
			s.name.MarshalJSON(),
			s.place_id.MarshalJSON(),
			s.course.MarshalJSON(),
			s.distance.MarshalJSON(),
			s.sire_id.MarshalJSON(),
			s.date.MarshalJSON(),
			s.start_time.MarshalJSON(),
			s.count.MarshalJSON(),
			s.win_rate.MarshalJSON(),
			s.rentai_rate.MarshalJSON(),
			s.fukushou_rate.MarshalJSON(),
			s.place_count.MarshalJSON(),
			s.place_win_rate.MarshalJSON(),
			s.place_rentai_rate.MarshalJSON(),
			s.place_fukushou_rate.MarshalJSON(),
			s.course_count.MarshalJSON(),
			s.course_win_rate.MarshalJSON(),
			s.course_rentai_rate.MarshalJSON(),
			s.course_fukushou_rate.MarshalJSON(),
			s.distance_count.MarshalJSON(),
			s.distance_win_rate.MarshalJSON(),
			s.distance_rentai_rate.MarshalJSON(),
			s.distance_fukushou_rate.MarshalJSON(),
			s.course_count.MarshalJSON(),
			s.course_distance_win_rate.MarshalJSON(),
			s.course_distance_rentai_rate.MarshalJSON(),
			s.course_distance_fukushou_rate.MarshalJSON()})
	}

	f, _ := os.OpenFile("./data/horse_sire_rate.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(sires)

	w.Flush()
}

func get_season_rate(date string, db *sql.DB) {
	rows, _ := db.Query("select * from season_rate where date = ?", date)
	seasons := make([][]string, 0)
	seasons = append(seasons, []string{"id", "horse_id", "date", "季節出走回数", "季節勝率", "季節連対率", "季節複勝率"})

	for rows.Next() {
		s := SeasonRate{}
		err := rows.Scan(
			&s.id,
			&s.horse_id,
			&s.date,
			&s.season_count,
			&s.season_win,
			&s.season_rentai,
			&s.season_fukusho,
		)

		if err != nil {
			panic(err)
		}

		seasons = append(seasons, []string{
			s.id.MarshalJSON(),
			s.horse_id.MarshalJSON(),
			s.date.MarshalJSON(),
			s.season_count.MarshalJSON(),
			s.season_win.MarshalJSON(),
			s.season_rentai.MarshalJSON(),
			s.season_fukusho.MarshalJSON(),
		})

		f, _ := os.OpenFile("./data/season_rate.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()

		w := csv.NewWriter(f)

		w.WriteAll(seasons)

		w.Flush()
	}
}

func get_shinba(date string, db *sql.DB) {
	rows, _ := db.Query("select * from horse_pillar_shinba where date = ?", date)

	pillars := make([][]string, 0)
	pillars = append(pillars, []string{
		"horse_race_id", "date", "horse_id", "race_id", "age", "course", "distance", "result",
		"result_except", "jockey_id", "multiple", "affiliation_id", "trainer_id", "gender", "weight",
		"減量騎手", "body_weight", "body_weight_in_de", "popular", "odds", "waku", "h_num",
		"time_odds", "time_popular",
		"producer", "owner",
		"born_month", "stallion_id", "sireline", "color_id",
		"兄弟出走数", "兄弟勝利数", "兄弟勝率", "兄弟連対率", "兄弟複勝率",
		"furlong_6", "furlong_5", "furlong_4", "furlong_3", "furlong_1",
		"furlong_6_5", "furlong_5_4", "furlong_4_3", "furlong_3_2", "furlong_2_1",
		"training_course", "load_id", "rank", "training_cond", "run_place"})

	for rows.Next() {
		hr := HorsePillarShinba{}
		err := rows.Scan(
			&hr.id, &hr.date, &hr.horse_id, &hr.race_id, &hr.age, &hr.course, &hr.distance, &hr.result,
			&hr.result_except, &hr.jockey_id, &hr.multiple, &hr.affiliation_id, &hr.trainer_id, &hr.gender, &hr.weight,
			&hr.is_loss_jockey, &hr.body_weight, &hr.body_weight_in_de,
			&hr.popular, &hr.odds, &hr.waku, &hr.h_num,
			&hr.time_odds, &hr.time_popular,
			&hr.producer, &hr.owner,
			&hr.born_month, &hr.stallion_id, &hr.sireline, &hr.color_id,
			&hr.brother_race_count, &hr.brother_win_count, &hr.brother_win_rate, &hr.brother_rentai_rate, &hr.brother_fukusho_rate,
			&hr.training_urlong_6, &hr.training_furlong_5, &hr.training_furlong_4, &hr.training_furlong_3, &hr.training_furlong_1,
			&hr.training_furlong_6_5, &hr.training_furlong_5_4, &hr.training_furlong_4_3, &hr.training_furlong_3_2, &hr.training_furlong_2_1,
			&hr.training_training_course, &hr.training_load_id, &hr.training_rank, &hr.training_training_cond, &hr.training_run_place)
		if err != nil {
			panic(err)
		}

		pillars = append(pillars, []string{
			hr.id.MarshalJSON(), hr.date.MarshalJSON(), hr.horse_id.MarshalJSON(), hr.race_id.MarshalJSON(), hr.age.MarshalJSON(),
			hr.course.MarshalJSON(), hr.distance.MarshalJSON(), hr.result.MarshalJSON(), hr.result_except.MarshalJSON(),
			hr.jockey_id.MarshalJSON(), hr.multiple.MarshalJSON(), hr.affiliation_id.MarshalJSON(), hr.trainer_id.MarshalJSON(),
			hr.gender.MarshalJSON(), hr.weight.MarshalJSON(), hr.is_loss_jockey.MarshalJSON(), hr.body_weight.MarshalJSON(), hr.body_weight_in_de.MarshalJSON(),
			hr.popular.MarshalJSON(), hr.odds.MarshalJSON(), hr.waku.MarshalJSON(), hr.h_num.MarshalJSON(),
			hr.time_odds.MarshalJSON(), hr.time_popular.MarshalJSON(),
			hr.producer.MarshalJSON(), hr.owner.MarshalJSON(),
			hr.born_month.MarshalJSON(), hr.stallion_id.MarshalJSON(), hr.sireline.MarshalJSON(), hr.color_id.MarshalJSON(),
			hr.brother_race_count.MarshalJSON(), hr.brother_win_count.MarshalJSON(), hr.brother_win_rate.MarshalJSON(),
			hr.brother_rentai_rate.MarshalJSON(), hr.brother_fukusho_rate.MarshalJSON(),
			hr.training_urlong_6.MarshalJSON(), hr.training_furlong_5.MarshalJSON(), hr.training_furlong_4.MarshalJSON(), hr.training_furlong_3.MarshalJSON(), hr.training_furlong_1.MarshalJSON(),
			hr.training_furlong_6_5.MarshalJSON(), hr.training_furlong_5_4.MarshalJSON(), hr.training_furlong_4_3.MarshalJSON(), hr.training_furlong_3_2.MarshalJSON(), hr.training_furlong_2_1.MarshalJSON(),
			hr.training_training_course.MarshalJSON(), hr.training_load_id.MarshalJSON(), hr.training_rank.MarshalJSON(), hr.training_training_cond.MarshalJSON(), hr.training_run_place.MarshalJSON()})
	}

	f, _ := os.OpenFile("./data/race_detail_shinba.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()

	w := csv.NewWriter(f)

	w.WriteAll(pillars)

	w.Flush()
}
