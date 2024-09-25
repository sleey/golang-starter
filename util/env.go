package util

import "os"

var Environment = os.Getenv("ENVIRONMENT")

func IsLocalDev() bool {
	return Environment == "dev" || Environment == ""
}
