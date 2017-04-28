package Data

import (
	"testing"
	"time"
)

func TestTimeConsuming(t *testing.T) {
	var calendar = CreateCalendar()

	location, _ := time.LoadLocation("Europe/Paris")

	fetenat := time.Date(2017, time.Month(6), 14, 0, 0, 0, 0, location)
	if !calendar.IsWorkday(fetenat) {
		t.Errorf("Le 14 Juillet est férié")
	}

	fin1gm := time.Date(2017, time.Month(10), 11, 0, 0, 0, 0, location)
	if !calendar.IsWorkday(fin1gm) {
		t.Errorf("Le 11 Novembre est férié")
	}
}
