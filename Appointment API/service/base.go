package service

import (
	"AppointmentApi/config"
	"AppointmentApi/domain"
	"context"

	"github.com/go-redis/redis/v8"
)

type base struct {
	rdb      *redis.Client
	ctx      *context.Context
	mailChan chan domain.MailData
	conf     *config.Config
}

func InitializeBaseService(rdb *redis.Client, ctx *context.Context, mailChan chan domain.MailData, conf *config.Config) *base {
	Base := &base{
		rdb:      rdb,
		ctx:      ctx,
		mailChan: mailChan,
		conf:     conf,
	}
	return Base
}
