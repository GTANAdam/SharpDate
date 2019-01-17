<p align="center" >
    <img width="256" src ="https://i.imgur.com/FFTTmiD.png" />
</p>

# SharpDate
> DateTime a la .NET/C# Style in GoLang

https://i.imgur.com/FFTTmiD.png

## Setting up
```sh
go get github.com/GTANAdam/SharpDate
```

## Quickstart
The usage is quite easy and resembles .NET
#### Importing package
Before we continue, importing the package correctly is quite important as it involves a small trick in order to omit the use of package name.

Add just underneath the package name
```go
import . "github.com/GTANAdam/SharpDate"
```

## Usage Examples:
```go
package main

import "fmt"
import . "github.com/GTANAdam/SharpDate"

func main() {
	e := DateTime{}.Now().ToString("dd-MM-yyyy HH:mm:ss")
	fmt.Println(e) // Returns current time, i.e: // 17-05-2018 23:55:22
	
	e = DateTime{}.New(2018, 1, 17).ToString("dd/MMMM/yyyy")
	fmt.Println(e) // 17/January/2018
	
	e = NewDateTime(2018, 1, 17, 15, 5, 2).ToString("dd/MMMM/yyyy h:m:s TT")
	fmt.Println(e) // 17/January/2018 3:5:2 PM
	
	e = DateTime{Year: 2018, Month: 5, Day: 17}.ToString("dd-MM-yyyy")
	fmt.Println(e) // 17-05-2018
	
	e = DateTime{}
	e.Year = 2018
	e.Month = 5
	e.Day = 17
	fmt.Println(e.ToString("dd-MM-yyyy")) // 17-05-2018
	
	e = DateTime{}.New(2018, 1, 17).AddMonths(2).AddDays(5).AddMinutes(10)
	fmt.Println(e.ToString("dd-MM-yyyy HH:mm")) // Result: 22-03-2018 00:10
}

```
## Available functions:
```go
Now()

New(year, month, day)
NewDateTime(year, month, day)

New(year, month, day, hours)
NewDateTime(year, month, day, hours)

New(year, month, day, hours, minutes)
NewDateTime(year, month, day, hours, minutes)

New(year, month, day, hours, minutes, seconds)
NewDateTime(year, month, day, hours, minutes, seconds)

New(year, month, day, hours, minutes, seconds, nanoseconds)
NewDateTime(year, month, day, hours, minutes, seconds, nanoseconds)

Add(years, months, days)
Subtract(years, months, days)
AddDays(value)
AddMonths(value)
AddYears(value)

AddTime(hours, minutes, seconds)
SubtractTime(hours, minutes, seconds)
AddHours(value)
AddMinutes(value)
AddSeconds(value)
AddMilliseconds(value)
AddMicroseconds(value)
```


## Available formats
```sh
dddd: day string. example: Monday
ddd:  short day string, example: Mon
dd:   day number, from 01 to (end of month). example: 02
d:    day number, from 1 to (end of month). example: 2
MMMM: month string. example: January
MMM:  short month string. example: Jan
MM:   month number, from 01 to 12. example: 01
M:    month number, from 1 to 12. example: 1
yyyy: year. example: 2006
yy:   year, last 2 digits. example: 06
HH:   hours, from 00 to 23. example: 15
hh:   hours, from 01 to 12. example: 03
h:    hours, from 0 to 12. example: 3
mm:   minutes. example: 04
m:    minutes. example: 4
ss:   seconds. example: 05
s:    seconds. example: 5
f:    milliseconds. example: .000
ff:   microseconds. example: .999999
tt:   meridian designation. example: pm/am
TT:   meridian designation. example: PM/AM
Z:    Timezone. example: MST
zz:   Timezone offset. example: -07
zzz:  Timezone offset. example: -07:00
```

## License
MIT License © Guilherme Caruso
MIT License © GTANAdam
