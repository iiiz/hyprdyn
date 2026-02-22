package hyprdyn

import (
	"github.com/charmbracelet/log"
)

// generic
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
