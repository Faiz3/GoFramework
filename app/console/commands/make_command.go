package commands

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type MakeCommand struct{}

func (c *MakeCommand) Signature() string {
	return "make"
}

func (c *MakeCommand) Description() string {
	return "Generate a new class (controller, model, migration, seeder)"
}

var makeStubs = map[string]string{
	"controller": `package controllers

import (
	"go-framework/framework/http"
	"github.com/gofiber/fiber/v2"
)

type {{Name}}Controller struct {
	http.BaseController
}

func (c *{{Name}}Controller) Handle(ctx *fiber.Ctx) error {
	return c.Success(ctx, fiber.Map{"message": "{{Name}}Controller"})
}
`,
	"model": `package models

import "time"

type {{Name}} struct {
	ID        uint      ` + "`" + `gorm:"primaryKey" json:"id"` + "`" + `
	CreatedAt time.Time ` + "`" + `json:"created_at"` + "`" + `
	UpdatedAt time.Time ` + "`" + `json:"updated_at"` + "`" + `
}

func (m *{{Name}}) TableName() string {
	return "{{Snake}}"
}
`,
	"migration": `package migrations

import (
	"go-framework/framework/database"
	"go-framework/framework/logging"
)

type {{Name}} struct{}

func (m *{{Name}}) Signature() string {
	return "{{Timestamp}}"
}

func (m *{{Name}}) Description() string {
	return "{{Description}}"
}

func (m *{{Name}}) Up() error {
	logging.Info("Running migration: %s", m.Description())
	return nil
}

func (m *{{Name}}) Down() error {
	return nil
}
`,
	"seeder": `package seeds

import (
	"go-framework/framework/database"
	"go-framework/framework/logging"
)

type {{Name}}Seeder struct{}

func (s *{{Name}}Seeder) Run() error {
	logging.Info("Seeding {{Table}} table")
	return nil
}
`,
}

func (c *MakeCommand) Handle(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: make <type> <name>\n  types: controller, model, migration, seeder")
	}

	genType := args[0]
	name := args[1]

	stub, ok := makeStubs[genType]
	if !ok {
		return fmt.Errorf("unknown type: %s. Available types: controller, model, migration, seeder", genType)
	}

	replacements := map[string]string{
		"{{Name}}":        name,
		"{{Snake}}":       toSnakeCase(name),
		"{{Timestamp}}":   time.Now().Format("2006_01_02_150405"),
		"{{Description}}": fmt.Sprintf("Create %s table", toSnakeCase(name)),
		"{{Table}}":       toSnakeCase(name),
	}

	for k, v := range replacements {
		stub = strings.ReplaceAll(stub, k, v)
	}

	var filePath string
	switch genType {
	case "controller":
		filePath = fmt.Sprintf("app/http/controllers/%s_controller.go", toSnakeCase(name))
	case "model":
		filePath = fmt.Sprintf("app/models/%s.go", toSnakeCase(name))
	case "migration":
		filePath = fmt.Sprintf("database/migrations/%s_%s.go", time.Now().Format("2006_01_02_150405"), toSnakeCase(name))
	case "seeder":
		filePath = fmt.Sprintf("database/seeds/%s_seeder.go", toSnakeCase(name))
	}

	if err := os.WriteFile(filePath, []byte(stub), 0644); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	fmt.Printf("Created: %s\n", filePath)
	return nil
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(r + 32)
		} else {
			result.WriteRune(r)
		}
	}
	return strings.ToLower(result.String())
}
