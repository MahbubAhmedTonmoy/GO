package handlers

import "AppointmentApi/config"

var Base *base

type base struct {
	conf *config.Config
}

func (b *base) Initialize(conf *config.Config) {
	Base = &base{
		conf: conf,
	}
}
