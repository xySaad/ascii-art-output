package utils

import (
	"os"
)

func GetBannerName() (string, error) {
	args := os.Args
	bannerName := "standard"

	if len(args) < 2 {
		return "", Error("missing arguments")
	}
	if len(args) > 3 {
		return "", Error("too many arguments")
	}
	if len(args) == 3 {
		bannerName = args[2]
	}

	return bannerName, nil
}
