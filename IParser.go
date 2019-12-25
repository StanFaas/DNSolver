package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	iParser := figure.NewFigure("iparser", "poison", true)
	iParser.Print()
	fmt.Printf("\niParser%56.17s\n\n", "made by StanFaas")

	fmt.Println(`
iParser resolves all the IPs for the domains specified in a file.
Each domain should be on a seperate line.
	`)

	// flag.Visit(func(f *flag.Flag) {
	// 	fmt.Println(f.Name)
	// })

	// f, err := os.Open(*domainsFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	if err = f.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// fileScanner := bufio.NewScanner(f)
	// for fileScanner.Scan() {
	// 	fmt.Println(fileScanner.Text())
	// }

	// err = fileScanner.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func flagParser(flag string) {
	switch flag {
	case "d", "domains":
		fmt.Println("domains.")
	case "o", "output":
		fmt.Println("output file.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("No flags found")
	}

}

func questions() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		} else {
			fmt.Println("Thats not a valid option")
		}

	}
}
