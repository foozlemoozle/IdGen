package IdGen

import (
	"time"
)

func Generate(prefix string, suffix string) string {
	return prefix + string(time.Now().Unix()) + suffix
}
