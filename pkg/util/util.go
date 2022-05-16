package util

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/ktnyt/go-moji"
)

func NilOrString(s string) *string {
	if s != "" {
		return &s
	}

	return nil
}

func NilOrInt(s string) *int {
	if s != "" {
		i, _ := strconv.Atoi(s)
		return &i
	}

	return nil
}

func NilOrFloat(s string) *float64 {
	if s != "" {
		f, _ := strconv.ParseFloat(s, 64)
		return &f
	}

	return nil
}

func Weight(s string) (float64, bool) {
	s = strings.TrimSpace(s)
	rex := regexp.MustCompile(`\d+\.?\d+`)
	sw := rex.FindString(s)
	f, _ := strconv.ParseFloat(sw, 64)

	rex = regexp.MustCompile(`▲|☆|△|◇|★`)

	return f, rex.MatchString(s)
}

func IsNumber(s string) bool {
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return true
	}

	return false
}

func ZenNumTohan(value string) int {
	r, _ := strconv.Atoi(moji.Convert(strings.TrimSpace(value), moji.ZE, moji.HE))

	return r
}

func TimeToSeconds(t string) *int {
	if t == "" {
		return nil
	}

	first, _ := strconv.Atoi(t[:1])
	middle, _ := strconv.Atoi(t[2:4])
	last, _ := strconv.Atoi(t[5:])

	time := (first * 60 * 10) + (middle * 10) + last

	return &time
}

func StructToMap(data interface{}) map[string]interface{} {
	typeof := reflect.TypeOf(data)
	valueof := reflect.ValueOf(data)

	race := make(map[string]interface{})
	for it := 0; it < typeof.NumField(); it++ {
		ff := typeof.Field(it)
		key := ff.Name
		value := valueof.FieldByName(key).Interface()
		if key == "RESULT" || key == "PRIZE" {
			r, _ := strconv.Atoi(moji.Convert(strings.TrimSpace(value.(string)), moji.ZE, moji.HE))
			value = r
		}

		if key == "TIME" {
			m, _ := strconv.Atoi(value.(string)[:1])
			s, _ := strconv.Atoi(value.(string)[3:4])
			n, _ := strconv.Atoi(value.(string)[6:])
			value = strconv.Itoa((m * 60 * 10) + (s * 10) + n)
		}

		if key == "DISTANCE" {
			race["COURSE"] = strings.TrimSpace(value.(string)[:3])
			value = value.(string)[3:]
		}

		race[key] = value
	}

	return race
}

func ArrayContains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
