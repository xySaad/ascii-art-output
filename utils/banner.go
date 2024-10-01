package utils

import (
	"os"
)

func GetBannerName() (string, error) {
	bannerName := "standard"

	if len(os.Args) < 2 {
		return "", Error("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	if len(os.Args) > 3 {
		return "", Error("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	return bannerName, nil
}
