package internal

import (
	"log"
	"os"
)

func EnvShow(key string) {
	if val, ok := os.LookupEnv(key); !ok {
		log.Printf("%s is not set.\n", key)
	} else {
		log.Printf("%s=%s\n", key, val)
	}
}
