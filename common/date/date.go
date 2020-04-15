package date

import "time"

func DayStart(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func DayMiddle(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 12, 0, 0, 0, t.Location())
}

func DayEnd(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, 0, t.Location())
}

func MonthStart(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func MonthEnd(t time.Time) time.Time {
	return MonthStart(t).AddDate(0, 1, 0).Add(-time.Second)
}

func YearStart(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func YearEnd(t time.Time) time.Time {
	return YearStart(t).AddDate(1, 0, 0).Add(-time.Second)
}

func ISOWeekStart(t time.Time) time.Time {
	t = DayStart(t)
	wd := t.Weekday()
	if wd == time.Monday {
		return t
	}
	offset := int(time.Monday - wd)
	if offset > 0 {
		offset -= 7
	}
	return t.AddDate(0, 0, offset)
}

func ISOWeekEnd(t time.Time) time.Time {
	t = DayEnd(t)
	wd := t.Weekday()
	if wd == time.Sunday {
		return t
	}
	offset := int(time.Sunday - wd + 7)
	return t.AddDate(0, 0, offset)
}

// https://en.wikipedia.org/wiki/ISO_week_date
func ISOWeekStartByNumber(year int, week int, loc *time.Location) time.Time {
	jan4 := time.Date(year, 1, 4, 0, 0, 0, 0, loc)
	jan4wd := jan4.Weekday()
	if jan4wd == time.Sunday {
		jan4wd = 7
	}
	offset := (week-1)*7 + int(time.Monday-jan4wd)
	return jan4.AddDate(0, 0, offset)
}

func DaysDiff(year1 int, month1 time.Month, day1 int, year2 int, month2 time.Month, day2 int) int {
	d1 := time.Date(year1, month1, day1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(year2, month2, day2, 0, 0, 0, 0, time.UTC)
	return int(d2.Sub(d1) / (time.Hour * 24))
}

func DaysDiffOfTime(t1 time.Time, t2 time.Time) int {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return DaysDiff(y1, m1, d1, y2, m2, d2)
}

func DayAgo(t time.Time, days int) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d-days, 0, 0, 0, 0, t.Location())
}

func DaysSinceEpoch(t time.Time) int {
	y, m, d := t.Date()
	ref := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	epoch := time.Unix(0, 0).In(time.UTC)
	return int(ref.Sub(epoch) / (time.Hour * 24))
}

func DayStartByDaysSinceEpoch(days int, loc *time.Location) time.Time {
	return time.Date(1970, time.January, 1+days, 0, 0, 0, 0, loc)
}
