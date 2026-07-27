package config

import "go-framework/framework/config"

func Mail(cfg *config.Config) map[string]interface{} {
	return map[string]interface{}{
		"driver":       cfg.Get("MAIL_DRIVER", "smtp"),
		"host":         cfg.Get("MAIL_HOST", "localhost"),
		"port":         cfg.Get("MAIL_PORT", "587"),
		"username":     cfg.Get("MAIL_USERNAME", ""),
		"password":     cfg.Get("MAIL_PASSWORD", ""),
		"encryption":   cfg.Get("MAIL_ENCRYPTION", "tls"),
		"from_address": cfg.Get("MAIL_FROM_ADDRESS", "app@localhost"),
		"from_name":    cfg.Get("MAIL_FROM_NAME", "GoFramework"),
	}
}
