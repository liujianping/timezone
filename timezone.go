package timezone

import (
	"fmt"
	"time"
)

const HOUR = 3600

var zones map[string]*time.Location

func AddZone(name string, hour int) {
	if _, err := time.LoadLocation(name); err != nil {
		zones[name] = time.FixedZone(name, hour*HOUR)
	}
}

func DelZone(name string) {
	delete(zones, name)
}

func GetZone(name string) (*time.Location, error) {
	if z, ok := zones[name]; ok {
		return z, nil
	}
	return time.LoadLocation(name)
}

func Zone(name string, t time.Time) (time.Time, error) {
	z, err := GetZone(name)
	if err != nil {
		return t, err
	}
	return t.In(z), nil
}

func ZoneHour(hour int, t time.Time) time.Time {
	hzone := time.FixedZone(fmt.Sprintf("hour-%d", hour), hour*HOUR)
	return t.In(hzone)
}

func ZoneDate(name string, year, month, day, hour, minute, second, nsec int) (time.Time, error) {
	z, err := GetZone(name)
	if err != nil {
		utc, _ := time.LoadLocation("UTC")
		return time.Date(year, time.Month(month), day, hour, minute, second, nsec, utc), err
	}
	return time.Date(year, time.Month(month), day, hour, minute, second, nsec, z), err
}

func init() {
	zones = make(map[string]*time.Location)
}
