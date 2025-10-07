package policy

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

func (c *Checker) yaml() error {
	if c.input != nil {
		return nil
	}

	content, err := os.ReadFile(c.fileName)
	if err != nil {
		return fmt.Errorf("failed to read input file '%s': %s", c.fileName, err)
	}

	if err := yaml.Unmarshal(content, &c.input); err != nil {
		return fmt.Errorf("failed to parse input: %s", err)
	}

	return nil
}
