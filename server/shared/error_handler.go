package shared

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Default().Printf("error: %v", err)
		panic(err)
	}
}
