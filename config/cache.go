package config

import "go-framework/framework/config"

func Cache(cfg *config.Config) map[string]interface{} {
	return map[string]interface{}{
		"default": cfg.Get("CACHE_DRIVER", "file"),
		"prefix":  cfg.Get("CACHE_PREFIX", "go_cache"),
	}
}
