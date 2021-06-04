package goindodax

import (
	"github.com/david-yappeter/go-indodax/api"
)

func NewClient(base string, privateBase string, option ...api.Option) *api.Client {
	return api.NewClient(base, privateBase, option...)
}

func NewDefaultClient(option ...api.Option) *api.Client {
	return api.NewDefaultClient(option...)
}
