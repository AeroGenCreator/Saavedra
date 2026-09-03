package env

import (
	"log"
	"os"
	"strconv"
)

var IsProduction bool
var ApyKey string
var AdminId int
var RecordsPerSlice int

func LoadGlobalEnvs() error {
	IsProduction, err := strconv.ParseBool(os.Getenv("IS_PRODUCTION"))
	if err != nil {
		log.Printf("Invalid value given to 'IS_PRODUCTION' (%v): default 'false'", err)
		IsProduction = false
	}
	log.Print("PROJECT PRODUCTION STATUS SET TO: ", IsProduction)

	ApyKey = os.Getenv("SESSION_TOKEN")
	recordsPerSliceStr := os.Getenv("RECORDS_PER_SLICE")
	RecordsPerSlice, err = strconv.Atoi(recordsPerSliceStr)
	if err != nil {
		log.Fatal("Can't parse 'RECORDS_PER_SLICE'... MAKE SURE TO SPECIFY AN INTEGER NUMBER.")
	}

	return nil
}

func BufferAdminId(id int) {
	AdminId = id
}
