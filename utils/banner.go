package utils

import (
	"os"
)

type Args struct {
	BannerName string
	Text       string
	Output     string
}

func CheckArgs() (Args, string) {
	if len(os.Args) < 2 {
		return Args{}, ""
	}
	if len(os.Args) > 4 {
		return Args{}, ""
	}
	result := Args{BannerName: "standard", Text: os.Args[1]}

	if len(os.Args[1]) > 9 && os.Args[1][:9] == "--output=" {
		result.Output = os.Args[1][9:]
		if len(os.Args) >= 3 {
			result.Text = os.Args[2]
		}
		if len(os.Args) == 4 {
			result.BannerName = os.Args[3]
		}
	} else {
		result.BannerName = os.Args[2]
	}

	return result, "OK"
}
