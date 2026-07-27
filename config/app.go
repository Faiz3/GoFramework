package config

import "go-framework/framework/config"

func App(cfg *config.Config) map[string]interface{} {
	return map[string]interface{}{
		"name":         cfg.Get("APP_NAME", "GoFramework"),
		"env":          cfg.Get("APP_ENV", "production"),
		"debug":        cfg.GetBool("APP_DEBUG", false),
		"url":          cfg.Get("APP_URL", "http://localhost:3000"),
		"port":         cfg.Get("APP_PORT", "3000"),
		"read_timeout": cfg.Get("APP_READ_TIMEOUT", "10s"),
		"write_timeout": cfg.Get("APP_WRITE_TIMEOUT", "10s"),
		"body_limit":   cfg.Get("APP_BODY_LIMIT", "10"),
	}
}
