package trainbit

import (
	"github.com/NodeSeekDev/nlist/v3/internal/driver"
	"github.com/NodeSeekDev/nlist/v3/internal/op"
)

type Addition struct {
	driver.RootID
	AUSHELLPORTAL string `json:"AUSHELLPORTAL" required:"true"`
	ApiKey        string `json:"apikey" required:"true"`
}

var config = driver.Config{
	Name:        "Trainbit",
	LocalSort:   false,
	OnlyLocal:   false,
	OnlyProxy:   false,
	NoCache:     false,
	NoUpload:    false,
	NeedMs:      false,
	DefaultRoot: "0_000",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Trainbit{}
	})
}
