package config

import "go-framework/framework/config"

func Session(cfg *config.Config) map[string]interface{} {
	return map[string]interface{}{
		"driver": cfg.Get("SESSION_DRIVER", "cookie"),
		"lifetime": cfg.Get("SESSION_LIFETIME", "120"),
		"expire_on_close": cfg.GetBool("SESSION_EXPIRE_ON_CLOSE", false),
		"cookie":  cfg.Get("SESSION_COOKIE", "go_session"),
	}
}
