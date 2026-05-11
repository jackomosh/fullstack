package ascii

import (
	"os"
	"strings"
)

func GetBanner(filename string) ([]string, error) {

	data, err := os.ReadFile(filename) {
		if err != nil {
			return nil, err
		}
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n"))
	return strings.Split(content, "\n"), nil
}

func GenearteAscii(input string, banner string) (string, error) {

	bannerLines, err := GetBanner(banner, ".txt")
	if err != nil {
		return "", err
	}

	input := strings.ReplaceAll(input, "\\n", "\n")

	
}