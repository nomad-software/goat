package variable

import (
	"fmt"

	"github.com/nomad-software/goat/internal/ui/element/hash"
)

// GenerateName generates a custom variable name.
func GenerateName(args ...string) string {
	args = append(args, "variable")
	hash := hash.Generate(args...)

	return fmt.Sprintf("variable-%s", hash)
}
