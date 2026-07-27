package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	data map[string]string
}

func New(envPath ...string) *Config {
	path := ".env"
	if len(envPath) > 0 {
		path = envPath[0]
	}

	c := &Config{data: make(map[string]string)}
	_ = godotenv.Load(path)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			c.data[pair[0]] = pair[1]
		}
	}

	return c
}

func (c *Config) Get(key string, defaultValue ...string) string {
	envKey := strings.ToUpper(key)
	envKey = strings.ReplaceAll(envKey, ".", "_")
	envKey = strings.ReplaceAll(envKey, "-", "_")

	if val, ok := c.data[envKey]; ok {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

func (c *Config) GetString(key string, defaultValue ...string) string {
	return c.Get(key, defaultValue...)
}

func (c *Config) GetInt(key string, defaultValue ...int) int {
	val := c.Get(key)
	if val == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return i
}

func (c *Config) GetBool(key string, defaultValue ...bool) bool {
	val := c.Get(key)
	if val == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	return b
}

func (c *Config) GetDuration(key string, defaultValue ...time.Duration) time.Duration {
	val := c.Get(key)
	if val == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return d
}

func (c *Config) GetStringSlice(key string, defaultValue ...[]string) []string {
	val := c.Get(key)
	if val == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []string{}
	}
	return strings.Split(val, ",")
}

func (c *Config) Set(key string, value string) {
	envKey := strings.ToUpper(key)
	envKey = strings.ReplaceAll(envKey, ".", "_")
	envKey = strings.ReplaceAll(envKey, "-", "_")
	c.data[envKey] = value
	os.Setenv(envKey, value)
}

func (c *Config) LoadConfigFile(path string) error {
	_ = godotenv.Load(path)
	return nil
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Driver:   c.Get("DB_CONNECTION", "mysql"),
		Host:     c.Get("DB_HOST", "127.0.0.1"),
		Port:     c.GetInt("DB_PORT", 3306),
		Database: c.Get("DB_DATABASE", "forge"),
		Username: c.Get("DB_USERNAME", "root"),
		Password: c.Get("DB_PASSWORD", ""),
	}
}

func (c *Config) DSN() string {
	db := c.GetDatabaseConfig()
	switch db.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.Username, db.Password, db.Host, db.Port, db.Database)
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			db.Host, db.Port, db.Username, db.Password, db.Database)
	case "sqlite":
		return db.Database + ".db"
	default:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.Username, db.Password, db.Host, db.Port, db.Database)
	}
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}
