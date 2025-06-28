package auth

import (
	"mine-server-manager/internal/config"
	"os"
	"strings"
	"sync"
)

var (
	whitelist   []string
	whitelistMu sync.RWMutex
)

func LoadWhitelist(cfg *config.Config) {
	data, err := os.ReadFile(cfg.WhitelistPath)
	if err != nil {
		// should panic since it's a must
		panic(err)
	}

	whitelist = strings.Split(string(data), "\n")
	whitelistMu.Lock()
	defer whitelistMu.Unlock()

}
