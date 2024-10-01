package utils

import (
	"os"
)

func GetBannerName() (string, string) {
	bannerName := "standard"

	if len(os.Args) < 2 {
		return "", ""
	}
	if len(os.Args) > 3 {
		return "", ""
	}
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	return bannerName, "OK"
}
