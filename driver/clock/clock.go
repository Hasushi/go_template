package clock

import (
	"time"

	"go_template/usecase/output_port"
)

type Clock struct{}

func New() output_port.Clock {
	return Clock{}
}

func (c Clock) Now() time.Time {
	return time.Now()
}
