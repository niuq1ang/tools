package timestamp

import (
	"sync"
	"time"
)

/*
  用于生成在一个进程内唯一的微秒时间戳, 生成速度为1,000,000/s
*/
var (
	lastMicroSec int64
	workerID     int8 //取值范围：0~3
	lock         sync.Mutex
)

/*
* 取值范围: 0~3
 */
func SetWorkerID(id int8) {
	workerID = int8(id) & MASK
}

const (
	MASK_BITS = 4
	MASK      = 1<<MASK_BITS - 1
)

// 微秒(系统唯一)
func Next() int64 {
	lock.Lock()
	defer lock.Unlock()

	miscroSec := getMaskedMicroSec()
	for miscroSec <= lastMicroSec {
		miscroSec = getMaskedMicroSec()
	}
	lastMicroSec = miscroSec
	return miscroSec + int64(workerID)
}

func getMaskedMicroSec() int64 {
	return GetMicroSec() & (^int64(MASK))
}

// 微秒
func GetMicroSec() int64 {
	return time.Now().UnixNano() / Thousand
}

func GetMicroSecFromTime(t time.Time) int64 {
	return t.UnixNano() / Thousand
}

func GetSec() int64 {
	return time.Now().Unix()
}

func SecToDateString(i int64) string {
	tm := time.Unix(i, 0)
	return tm.Format("2006-01-02")
}

func SecToDateTimeString(i int64) string {
	tm := time.Unix(i, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func TimeStringToInt64(s string, loc *time.Location) int64 {
	tm, _ := time.ParseInLocation("2006-01-02", s, loc)
	return tm.Unix()
}

// 时间戳转换为以天为单位的时间戳, UTC 时区
func TodaySecondFromTimestamp(timestamp int64) int64 {
	str := SecToDateString(timestamp)
	return TimeStringToInt64(str, time.UTC)
}

// 当天时间的时间戳, local 时区
func TodaySecond() int64 {
	currentTime := GetSec()
	str := SecToDateString(currentTime)
	return TimeStringToInt64(str, time.Local)
}

func DateRange(f, t int64) (r []string) {
	r = make([]string, 0)
	const daySecond = 1 * 24 * 60 * 60
	for {
		if f > t {
			break
		}
		r = append(r, SecToDateString(f))
		f += daySecond
	}
	return
}
