package goindodax

import (
	"go-indodax/api/indodax"
)

func NewClient(base string, privateBase string, option ...indodax.Option) *indodax.Client {
	return indodax.NewClient(base, privateBase, option...)
}

func NewDefaultClient(option ...indodax.Option) *indodax.Client {
	return indodax.NewDefaultClient(option...)
}
