package SetFlags

import (
	"flag"
	"fmt"
)

func SetFlags() {
	// showHelp := flag.String("h", "help", "displays all options")
	// showVersion := flag.String("v", "version", "displays the current version")
	domainsFile := initFlag("d", "domains", "domains")
	// outputFile := flag.String("o", "output", "file path to output ip list to")
	flag.Parse()

	fmt.Println(domainsFile)
}

func initFlag(name string, shortName string, defaultGopher string) string {
	var flagName string
	const (
		usage = "file path to read from"
	)
	flag.StringVar(&flagName, shortName, defaultGopher, usage)
	flag.StringVar(&flagName, name, defaultGopher, usage+" (shorthand)")

	fmt.Println("FLAGNAME" + flagName)
	return flagName
}
