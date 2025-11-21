package hyprdyn

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

// generic
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrettyPrint(d any) {
	jd, _ := json.MarshalIndent(d, "", "\t")

	fmt.Println(string(jd))
}
