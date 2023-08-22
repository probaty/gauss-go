package service

import (
	"flag"
	"fmt"
	"os"
)

type AppOption struct {
	fileName string
	raduis   int
}

func GetArgs() (AppOption, error) {
	fileName := flag.String("f", "", "The file name to read")
	raduis := flag.Int("r", 1, "Raduis of blur")
	flag.Parse()
	var errorArgs error
	if *fileName == "" {
		errorArgs = fmt.Errorf("The file name is not specified.")
	}

	_, err := os.Open(*fileName)
	if err != nil {
		errorArgs = err
	}

	return AppOption{
		fileName: *fileName,
		raduis:   *raduis,
	}, errorArgs
}
