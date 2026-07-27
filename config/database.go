package config

import "go-framework/framework/config"

func Database(cfg *config.Config) map[string]interface{} {
	return map[string]interface{}{
		"connection": cfg.Get("DB_CONNECTION", "mysql"),
		"host":       cfg.Get("DB_HOST", "127.0.0.1"),
		"port":       cfg.Get("DB_PORT", "3306"),
		"database":   cfg.Get("DB_DATABASE", "forge"),
		"username":   cfg.Get("DB_USERNAME", "root"),
		"password":   cfg.Get("DB_PASSWORD", ""),
	}
}
