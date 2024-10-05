package utils

import (
	"os"
	"strings"
)

type Output struct {
	Path string
	Ok   bool
}

type Args struct {
	BannerName string
	Text       string
	Output     Output
}

func GetArgs() (Args, string) {
	if len(os.Args) < 2 {
		return Args{}, ""
	}
	if len(os.Args) > 4 {
		return Args{}, ""
	}
	args := Args{BannerName: "standard", Text: os.Args[1]}

	if len(os.Args[1]) >= 9 && os.Args[1][:9] == "--output=" {

		if len(os.Args) <= 2 {
			return Args{}, ""
		}
		if len(os.Args[1]) <= 9 {
			return Args{}, ""
		}
		args.Output.Ok = true
		outputArg := strings.Split(os.Args[1][9:], "/")
		outputFile := outputArg[len(outputArg)-1]
		if len(outputFile) <= 4 || outputFile[len(outputFile)-4:] != ".txt" {
			return Args{}, ""
		}
		args.Output.Path = outputFile
		if len(os.Args) >= 3 {
			args.Text = os.Args[2]
		}
		if len(os.Args) == 4 {
			args.BannerName = os.Args[3]
		}
	} else {
		if len(os.Args) > 3 {
			return Args{}, ""
		}
		if len(os.Args) > 2 {
			args.BannerName = os.Args[2]
		}
	}

	return args, "OK"
}
