package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/common-nighthawk/go-figure"
	. "github.com/logrusorgru/aurora"
	"github.com/ns3777k/go-shodan/shodan"
)

func main() {
	fmt.Printf("\n")
	DNSolver := figure.NewFigure("DNSolver", "kban", true)
	DNSolver.Print()
	fmt.Printf("DNS|resolver%54.17s\n\n", "made by StanFaas")

	fmt.Println(`
DNSolver resolves all the IPs for the domains specified in a file.
Each domain should be on a seperate line.
	`)

	flag.Bool("h", false, "displays all options")
	flag.Bool("v", false, "displays the current version")
	flag.Bool("s", false, "Pulls data from Shodan, to see if the DNS is vulnerable")
	var domainList = flag.String("d", "", "file path to domain list [*Required]")
	var outputFile = flag.String("o", "", "file path to output ip list to")
	flag.Parse()

	var uniqueIPs []string

	flag.Visit(func(f *flag.Flag) {
		flagValue := f.Value.String()
		switch f.Name {
		case "s", "shodan":
			client := shodan.NewEnvClient(nil)
			dns, err := client.GetDNSResolve(context.Background(), []string{"google.com", "ya.ru"})

			if err != nil {
				log.Panic(err)
			} else {
				log.Println(dns["google.com"])
			}
		case "d", "domains":
			fmt.Println("Checking for file..")
			if fileExists(flagValue) {
				fmt.Println(Green("File FOUND"))
				uniqueIPsReturned := domainParser(flagValue)
				uniqueIPs = append(uniqueIPs, uniqueIPsReturned...)
				if *outputFile == "" {
					writeFile("target_ip_list.txt", uniqueIPsReturned)
				}
			} else {
				fmt.Println(Red("Domain file not FOUND"))
				fmt.Println(Red("Please check if the path to your file is correct"))
				os.Exit(3)
			}
		case "o", "output":
			if fileExists(flagValue) {
				fmt.Println(Red("\nFile already exists!"))
				fmt.Print(Red("Would you like to overwrite it?\n"))
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					text := scanner.Text()
					switch text {
					case "y", "Y":
						writeFile(flagValue, uniqueIPs)
					case "n", "N":
						fmt.Println(Red("Aborting.."))
						os.Exit(3)
					default:
						os.Exit(3)
					}
				}

				if scanner.Err() != nil {
					log.Fatal("Error")
				}
			} else {
				writeFile(flagValue, uniqueIPs)
			}
		case "h", "help":
			flag.PrintDefaults()
		case "v", "version":
			fmt.Println("The version is 0.1.1")
		default:
			if *domainList == "" {
				fmt.Println(Red("Oops, you forgot to point me towards your domain file, exiting"))
				flag.PrintDefaults()
				os.Exit(1)
			}
		}
	})
}

func domainParser(domainsFile string) []string {
	file, err := os.Open(domainsFile)
	if err != nil {
		fmt.Println(Red(err))
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(Red(err))
		}
	}()

	fileScanner := bufio.NewScanner(file)

	fmt.Println("Parsing domains..")

	IPs := make([]string, 0)
	index := 0
	for fileScanner.Scan() {
		domain := fileScanner.Text()
		ip, err := net.LookupHost(domain)

		if len(fileScanner.Text()) == 0 {
			continue
		}

		if err != nil {
			fmt.Println("Ip for", domain)
			fmt.Println(Red("No IP found.."))
			continue
		}

		fmt.Println("Ip for", domain)
		fmt.Println(Green(ip[0]))
		IPs = append(IPs, ip[0])
		index++
	}

	err = fileScanner.Err()
	if err != nil {
		log.Fatal(Red(err))
	}

	fmt.Println("\nFound", len(IPs), "IPs..")
	uniqueIPs := generateUniqueIPs(IPs)
	duplicates := len(IPs) - len(uniqueIPs)
	if duplicates > 0 {
		fmt.Println("but also found", Yellow(duplicates), "duplicates.. deduping now..")
		fmt.Println("Removed them.", Green(len(uniqueIPs)), Green("unique IPs remaining"), ".")
	}

	return uniqueIPs
}

func generateUniqueIPs(ipArray []string) []string {
	keys := make(map[string]bool)
	uniqueIPs := []string{}
	for _, entry := range ipArray {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueIPs = append(uniqueIPs, entry)
		}
	}
	return uniqueIPs
}

func writeFile(outputFile string, ipArray []string) {
	fmt.Println("\nWriting IPs to file", outputFile)

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	for _, value := range ipArray {
		fmt.Fprintln(file, value)
	}
	fmt.Println(Green("Jobs done!"))
	fmt.Println("Follow me on Twitter: @StanFaas")
	os.Exit(0)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
