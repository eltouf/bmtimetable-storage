package main

import (
	"os"

	"google.golang.org/api/option"
)

func main() {
	home, err := os.UserHomeDir()
	opt := option.WithCredentialsFile(home + string(os.PathSeparator) + ".config" + string(os.PathSeparator) + "bmtimetable-381d0e19e193.json")

}
