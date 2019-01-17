package sharpdate

import (
	"testing"
	"time"
)

func TestDateTimeAddDays(t *testing.T) {
	expected := "31-12-2018 23:55:59"
	actual := NewDateTime(2018, 12, 28, 23, 55, 59).AddDays(3).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddMonths(t *testing.T) {
	expected := "01-05-2018 23:55:59"
	actual := NewDateTime(2018, 2, 1, 23, 55, 59).AddMonths(3).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddYears(t *testing.T) {
	expected := "01-05-2018 23:55:59"
	actual := NewDateTime(2015, 5, 1, 23, 55, 59).AddYears(3).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAdd(t *testing.T) {
	expected := "03-05-2018 23:55:59"
	actual := NewDateTime(2015, 4, 1, 23, 55, 59).Add(3, 1, 2).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeSubtract(t *testing.T) {
	expected := "01-04-2015 22:50:50"
	actual := NewDateTime(2018, 5, 3, 22, 50, 50).Subtract(3, 1, 2).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddTime(t *testing.T) {
	expected := "03-05-2018 23:55:59"
	actual := NewDateTime(2018, 5, 3, 22, 50, 50).AddTime(1, 5, 9).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeSubtractTime(t *testing.T) {
	expected := "03-05-2018 22:50:50"
	actual := NewDateTime(2018, 5, 3, 23, 55, 59).SubtractTime(1, 5, 9).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddMinutes(t *testing.T) {
	expected := "03-05-2018 23:55:59"
	actual := NewDateTime(2018, 5, 3, 23, 45, 59).AddMinutes(10).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddHours(t *testing.T) {
	expected := "03-05-2018 23:55:59"
	actual := NewDateTime(2018, 5, 3, 22, 55, 59).AddHours(1).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddSeconds(t *testing.T) {
	expected := "03-05-2018 23:55:59"
	actual := NewDateTime(2018, 5, 3, 23, 55, 49).AddSeconds(10).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddMilliseconds(t *testing.T) {
	expected := "03-05-2018 23:55:59.010"
	actual := NewDateTime(2018, 5, 3, 23, 55, 59, 0).AddMilliseconds(10).ToString("dd-MM-yyyy HH:mm:ss.f")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeAddMicroseconds(t *testing.T) {
	expected := "03-05-2018 23:55:59.000011"
	actual := NewDateTime(2018, 5, 3, 23, 55, 59, 0).AddMicroseconds(11).ToString("dd-MM-yyyy HH:mm:ss.ff")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestNewDateTimeToString(t *testing.T) {
	expected := "31-12-2018 23:55:59"
	actual := NewDateTime(2018, 12, 31, 23, 55, 59).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestNewDateTimeNoTimeToString(t *testing.T) {
	expected := "31-12-2018"
	actual := NewDateTime(2018, 12, 31).ToString("dd-MM-yyyy")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestNewDateTimeShortTimeToString(t *testing.T) {
	expected := "31-12-2018 23"
	actual := NewDateTime(2018, 12, 31, 23).ToString("dd-MM-yyyy HH")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestNewDateTimeShortTimeMinutesToString(t *testing.T) {
	expected := "31-12-2018 23:55"
	actual := NewDateTime(2018, 12, 31, 23, 55).ToString("dd-MM-yyyy HH:mm")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeStructToString(t *testing.T) {
	expected := "31-12-2018 23:55:59"
	actual := DateTime{Year: 2018, Month: 12, Day: 31, Hour: 23, Min: 55, Sec: 59}.ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeNewToString(t *testing.T) {
	expected := "31-12-2018 23:55:59"
	actual := DateTime{}.New(2018, 12, 31, 23, 55, 59).ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeNowToString(t *testing.T) {
	expected := time.Now().Format("02-01-2006 15:04:05")
	actual := DateTime{}.Now().ToString("dd-MM-yyyy HH:mm:ss")
	if expected != actual {
		t.Errorf("\nExpected '%s' but got '%s' instead.", expected, actual)
	}
}

func TestDateTimeUnix(t *testing.T) {
	expected := time.Now().Unix()
	actual := DateTime{}.Now().Time.Unix()
	if expected != actual {
		t.Errorf("\nExpected '%v' but got '%v' instead.", expected, actual)
	}
}

func TestDateTimeYearMonth(t *testing.T) {
	expected := "12-2018"
	actual := DateTime{Year: 2018, Month: 12}.ToString("MM-yyyy")
	if expected != actual {
		t.Errorf("\nExpected '%v' but got '%v' instead.", expected, actual)
	}
}

func TestDateTimeYear(t *testing.T) {
	expected := "2018"
	actual := DateTime{Year: 2018}.ToString("yyyy")
	if expected != actual {
		t.Errorf("\nExpected '%v' but got '%v' instead.", expected, actual)
	}
}

func TestNewDateTimeInvalidArguments(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("\nExpected panic but got no panic instead!")
		}
	}()
	NewDateTime("2018").ToString("yyyy")
}
