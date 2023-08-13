package main

import (
	"sync"

	"github.com/m/v2/config"
	"github.com/m/v2/db"
	supplybot "github.com/m/v2/supply-bot"
	"github.com/m/v2/watchers"
)

func main() {
	config.IntializeConfig()
	db.InitializeDB()
	var wg sync.WaitGroup
	wg.Add(2)
	go watchers.InitializeWatcher(&wg)
	go supplybot.InitializeSupplyBot(&wg)
	wg.Wait()
}
