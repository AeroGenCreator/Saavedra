package env

import (
	"log"
	"os"
	"strconv"
)

var IsProduction bool
var ApyKey string

func LoadGlobalEnvs() error {
	IsProduction, err := strconv.ParseBool(os.Getenv("IS_PRODUCTION"))
	if err != nil {
		log.Printf("Invalid value given to 'IS_PRODUCTION' (%v): default 'false'", err)
		IsProduction = false
	}
	log.Print("PROJECT STATUS SET TO:", IsProduction)

	ApyKey = (os.Getenv("SESSION_TOKEN"))

	return nil
}
