package app

import (
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/gofiber/fiber/v2/log"
)

var lock = &sync.Mutex{}

func GetConfig[T any]() T {
	lock.Lock()
	defer lock.Unlock()
	var config T
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(&config, opts); err != nil {
		fmt.Printf("%+v\n", err)
	}
	log.Debugf("Config Loaded: %+v", config)

	return config
}
