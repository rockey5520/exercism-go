package clock

import "fmt"

// Clock struct stands for the given hour in int type
type Clock struct {
	minutes int
}

// New will instantiate new clock
func New(hour, minute int) Clock {
	minute += hour * 60
	minute %= 24 * 60
	if minute < 0 {
		minute += 24 * 60
	}
	return Clock{minutes: minute}
}

//Add adds minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

//Subtract reduces minutes to the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
