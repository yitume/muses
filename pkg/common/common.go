package common

import "time"

type Caller interface {
	// Init cfg returns parse cfg error.
	InitCfg(cfg []byte) error
	// Init Caller returns init caller error
	InitCaller() error
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type CallerFunc func() Caller
