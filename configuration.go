package astibranch

import (
	"flag"

	"github.com/asticode/go-astikit"
)

// Flags
var (
	Key = flag.String("branch-key", "", "the Branch key")
)

// Configuration represents the lib's configuration
type Configuration struct {
	Key    string `toml:"key"`
	Sender astikit.HTTPSenderOptions
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		Key: *Key,
	}
}
