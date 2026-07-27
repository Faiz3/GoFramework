package commands

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type KeyGenerateCommand struct{}

func (c *KeyGenerateCommand) Signature() string {
	return "key:generate"
}

func (c *KeyGenerateCommand) Description() string {
	return "Generate a new application key"
}

func (c *KeyGenerateCommand) Handle(args []string) error {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	fmt.Printf("Application key generated: %s\n", hex.EncodeToString(key))
	return nil
}
