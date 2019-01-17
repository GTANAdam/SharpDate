package sharpdate

import (
	"regexp"
	"strings"
	"time"
)

// **DateTime** Represents an instant in time, typically expressed as a date and time of day.
//
// **Arguments:**
//
// year: int
//
// month: int
//
// day: int
//
// hour: int
//
// min: int
//
// sec: int
//
// nsec: int
//
// location: *time.Location
//
// **Properties:**
//
// Time: time.Time
type DateTime struct {
	Year     int
	Month    int
	Day      int
	Hour     int
	Min      int
	Sec      int
	Nsec     int
	Location *time.Location
	Time     time.Time
}

var (
	months = map[int]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}

	predefs = map[string]string{
		"dddd": "Monday",
		"ddd":  "Mon",
		"dd":   "02",
		"d":    "2",
		"MMMM": "January",
		"MMM":  "Jan",
		"MM":   "01",
		"M":    "1",
		"yyyy": "2006",
		"yy":   "06",
		"HH":   "15",
		"hh":   "03",
		"h":    "3",
		"mm":   "04",
		"m":    "4",
		"ss":   "05",
		"s":    "5",
		"f":    "000",
		"ff":   "999999",
		"TT":   "PM",
		"tt":   "pm",
		"Z":    "MST",
		"zz":   "-07",
		"zzz":  "-07:00",
	}
)

// **ToString** converts the value of the current DateTime instance to its equivalent string representation using the specified format.
//
// **Avaialble formats:**
//
// dddd: day string. example: Monday
//
// ddd:  short day string, example: Mon
//
// dd:   day number, from 01 to (end of month). example: 02
//
// d:    day number, from 1 to (end of month). example: 2
//
// MMMM: month string. example: January
//
// MMM:  short month string. example: Jan
//
// MM:   month number, from 01 to 12. example: 01
//
// M:    month number, from 1 to 12. example: 1
//
// yyyy: year. example: 2006
//
// yy:   year, last 2 digits. example: 06
//
// HH:   hours, from 00 to 23. example: 15
//
// hh:   hours, from 01 to 12. example: 03
//
// h:    hours, from 0 to 12. example: 3
//
// mm:   minutes. example: 04
//
// m:    minutes. example: 4
//
// ss:   seconds. example: 05
//
// s:    seconds. example: 5
//
// f:	 milliseconds. example: .000
//
// ff:	 microseconds. example: .999999
//
// tt:	 meridian designation. example: pm/am
//
// TT:	 meridian designation. example: PM/AM
//
// Z: 	 Timezone. example: MST
//
// zz:	 Timezone offset. example: -07
//
// zzz:	 Timezone offset. example: -07:00
func (d DateTime) ToString(resultFormat string) string {
	regex := regexp.MustCompile(`(?m)(M{4})|(M{3})|(M{2})|(M{1})|(y{4})|(y{2})|(D{2})|(D{1})|(d{2})|(d{1})|(H{2})|(h{2})|(h{1})|(m{2})|(m{1})|(s{2})|(s{1}|(f{2})|(f{1})|(TT{1})|(tt{1})|(Z{1})|(zzz{1})|(zz{1}))`)
	for _, match := range regex.FindAllString(resultFormat, -1) {
		for key, val := range predefs {
			if match == key {
				resultFormat = strings.Replace(resultFormat, match, val, -1)
			}
		}
	}
	if d.Year != 0 {
		if d.Day == 0 {
			d.Day = 1
		}
		if d.Month == 0 {
			d.Month = 1
		}
		return d.New(d.Year, d.Month, d.Day, d.Hour, d.Min, d.Sec, d.Nsec).Time.Format(resultFormat)
	}
	return d.Time.Format(resultFormat)
}

// **New function** returns a DateTime instance from specified date and/or time values
//
// **Overloads:**
//
// New(year, month, day)
//
// New(year, month, day, hours)
//
// New(year, month, day, hours, minutes)
//
// New(year, month, day, hours, minutes, seconds)
//
// New(year, month, day, hours, minutes, seconds, nanoseconds)
func (DateTime) New(args ...interface{}) *DateTime {
	var d DateTime
	switch len(args) {
	case 3: // New(year, month, day)
		d.Time = time.Date(args[0].(int), months[args[1].(int)], args[2].(int), 0, 0, 0, 0, time.UTC)
	case 4: // New(year, month, day, hours)
		d.Time = time.Date(args[0].(int), months[args[1].(int)], args[2].(int), args[3].(int), 0, 0, 0, time.UTC)
	case 5: // New(year, month, day, hours, minutes)
		d.Time = time.Date(args[0].(int), months[args[1].(int)], args[2].(int), args[3].(int), args[4].(int), 0, 0, time.UTC)
	case 6: // New(year, month, day, hours, minutes, seconds)
		d.Time = time.Date(args[0].(int), months[args[1].(int)], args[2].(int), args[3].(int), args[4].(int), args[5].(int), 0, time.UTC)
	case 7: // New(year, month, day, hours, minutes, seconds, nanoseconds)
		d.Time = time.Date(args[0].(int), months[args[1].(int)], args[2].(int), args[3].(int), args[4].(int), args[5].(int), args[6].(int), time.UTC)
	default:
		panic("invalid function arguments")
	}
	return &d
}

// **NewDateTime** function returns a DateTime instance from specified date and/or time values
//
// **Overloads:**
//
// NewDateTime(year, month, day)
//
// NewDateTime(year, month, day, hours)
//
// NewDateTime(year, month, day, hours, minutes)
//
// NewDateTime(year, month, day, hours, minutes, seconds)
//
// NewDateTime(year, month, day, hours, minutes, seconds, nanoseconds)
func NewDateTime(args ...interface{}) *DateTime {
	return DateTime{}.New(args...)
}

// **Now** function returns DateTime instance from specified date and/or time values
func (DateTime) Now() *DateTime {
	return &DateTime{Time: time.Now()}
}

// **Add** Returns a new DateTime that adds the value of the specified date values to the value of this instance.
func (d DateTime) Add(years, months, days int) *DateTime {
	return &DateTime{Time: d.Time.AddDate(years, months, days)}
}

// **AddDays** Returns a new DateTime that adds the specified number of days to the value of this instance.
func (d DateTime) AddDays(value int) *DateTime {
	return &DateTime{Time: d.Time.AddDate(0, 0, value)}
}

// **AddMonths** Returns a new DateTime that adds the specified number of months to the value of this instance.
func (d DateTime) AddMonths(value int) *DateTime {
	return &DateTime{Time: d.Time.AddDate(0, value, 0)}
}

// **AddYears** Returns a new DateTime that adds the specified number of years to the value of this instance.
func (d DateTime) AddYears(value int) *DateTime {
	return &DateTime{Time: d.Time.AddDate(value, 0, 0)}
}

// **AddTime** Returns a new DateTime that adds the specified time values to the value of this instance.
func (d DateTime) AddTime(hours, minutes, seconds int) *DateTime {
	return &DateTime{Time: d.Time.Add(time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds))}
}

// **AddHours** Returns a new DateTime that adds the specified number of hours to the value of this instance.
func (d DateTime) AddHours(value int) *DateTime {
	return &DateTime{Time: d.Time.Add(time.Hour * time.Duration(value))}
}

// **AddMinutes** Returns a new DateTime that adds the specified number of minutes to the value of this instance.
func (d DateTime) AddMinutes(value int) *DateTime {
	return &DateTime{Time: d.Time.Add(time.Minute * time.Duration(value))}
}

// **AddSeconds** Returns a new DateTime that adds the specified number of seconds to the value of this instance.
func (d DateTime) AddSeconds(value int) *DateTime {
	return &DateTime{Time: d.Time.Add(time.Second * time.Duration(value))}
}

// **AddMilliseconds** Returns a new DateTime that adds the specified number of milliseconds to the value of this instance.
func (d DateTime) AddMilliseconds(value int) *DateTime {
	return &DateTime{Time: d.Time.Add(time.Millisecond * time.Duration(value))}
}

// **AddMicroseconds** Returns a new DateTime that adds the specified number of microseconds to the value of this instance.
func (d DateTime) AddMicroseconds(value int) *DateTime {
	return &DateTime{Time: d.Time.Add(time.Microsecond * time.Duration(value))}
}

// **Subtract** Returns a new DateTime that subtracts the specified date values from the value of this instance.
func (d DateTime) Subtract(years, months, days int) *DateTime {
	return &DateTime{Time: d.Time.AddDate(-years, -months, -days)}
}

// **SubtractTime** Returns a new DateTime that subtracts the specified time values from the value of this instance.
func (d DateTime) SubtractTime(hours, minutes, seconds int) *DateTime {
	return &DateTime{Time: d.Time.Add(-(time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)))}
}
