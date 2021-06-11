package sanyutil

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StrMapToInterface(fields map[string]string) map[string]interface{} {
	m := make(map[string]interface{})
	for k, v := range fields {
		if strings.Contains(v, ".") {
			m[k], _ = strconv.ParseFloat(v, 64)
		} else {
			m[k], _ = strconv.ParseInt(v, 10, 64)
		}
	}
	return m
}

// MsUnixToTime Unix毫秒转time
func MsUnixToTime(unix int64) time.Time {
	return time.Unix(unix/1000, (unix-unix/1000*1000)*1000000)
}

func Retry(fn func() error, attempts int, sleep time.Duration) error {
	err := fn()
	if err != nil {
		if attempts--; attempts > 0 {
			log.Warnf("Retry func error: %s. attempts #%d after %s.", err.Error(), attempts, sleep)
			time.Sleep(sleep)
			return Retry(fn, attempts, sleep)
		}
	}
	return err
}

func DefaultRetry(fn func() error) error {
	return Retry(fn, 2, 500*time.Millisecond)
}

func ArraysToMap(ks []string, vs []interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(ks); i++ {
		m[ks[i]] = vs[i]
	}
	return m
}

func StructToMap(obj interface{}, tagName string) map[string]interface{} {
	var m = make(map[string]interface{})
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	for i := 0; i < obj1.NumField(); i++ {
		tag := obj1.Field(i).Tag.Get(tagName)
		if tag != "" {
			m[tag] = obj2.Field(i).Interface()
		}
	}
	return m
}

func StructsToMaps(objs []interface{}, tagName string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, obj := range objs {
		result = append(result, StructToMap(obj, tagName))
	}
	return result
}

func StructToMapByJson(obj interface{}) (map[string]interface{}, error) {
	var m = make(map[string]interface{})
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// HumpToLine 驼峰转下划线 未完成
func HumpToLine(str string) string {
	result := strings.ReplaceAll(str, "(\\w)([A-Z])", "_$0")
	result = strings.ToLower(result)
	return result
}

// LowerStart 首字母小写
func LowerStart(str string) string {
	return strings.ToLower(str[0:1]) + str[1:len(str)]
}

func WeekInYear(t time.Time) int {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return week
}

func JoinError(err1 error, err2 error) error {
	if err1 == nil {
		return err2
	} else {
		return errors.Errorf("%s/n%s", err1, err2)
	}
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// StrArrayDifference 求差集 slice1-并集
func StrArrayDifference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

func StrArrayEqual(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// JoinMap 将m2的数据并到m1中，若有重复key则更新m1
func JoinMap(m1, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// CalculateSlope 计算斜率
func CalculateSlope(x1, y1, x2, y2, x3 float64) float64 {
	k := (y2 - y1) / (x2 - x1)
	v := y1 - k*x1
	y3 := k*x3 + v
	return y3
}
