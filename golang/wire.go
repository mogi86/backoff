//+build wireinject

package main

import "github.com/google/wire"

//go:generate wire
func InitLoggerViaWire() Logger {
	wire.Build(InitMessage, InitLogger)
	return Logger{}
}
