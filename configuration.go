package astibranch

import (
	"flag"

	astihttp "github.com/asticode/go-astitools/http"
)

// Flags
var (
	Key  = flag.String("branch-key", "", "the Branch key")
)

// Configuration represents the lib's configuration
type Configuration struct {
	Key    string `toml:"key"`
	Sender astihttp.SenderOptions
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		Key:  *Key,
	}
}
