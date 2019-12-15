package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIP(input string) string {
	grammer := `[0-9]{1,3}\.[0-9]{1,3}[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`
	matchMe := regexp.MustCompile(grammer)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 1 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		r := bufio.NewScanner(f)
		for r.Scan() {
			line := r.Text()
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			} else {
				fmt.Println(ip)
			}
		}
	}
}
