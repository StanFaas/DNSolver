# DNSolver v0.0.1

Used for reconnaissance, DNSolver parses multiline files, gets their IPs, removing duplicates and writes the unique IPs to a new file.

## Prerequisites

DNSolver is built with Golang. This means that you need Go installed before you can use it.

Here you find Go's [installation instructions](https://golang.org/doc/install).

## Installation

Installation is easy, just copy and paste the following line into your terminal:

`go get -u github.com/StanFaas/DNSolver`

Go will then download, build and install the program for you.

## Usage

- See all options:  
  `dnsolver -h`

- Parse domain file:  
  `dnsolver -d domainlist.txt`  
  Be sure your domain file has 1 domain per line, no comma's etc.

- Custom output file:  
  `dnsolver -d domainlist.txt -o target_ip_list.txt`
  This writes the unique IPs to a file called **target_ip_list.txt**.

- Parse IP's with Shodan and output open ports:  
  `dnsolver -d domainslist.txt -o target_ip_list.txt -s`

## Generate domainlist

To gerenate a domainlist I would recommend [AssetFinder](https://github.com/tomnomnom/assetfinder) by [TomNomNom](https://github.com/tomnomnom)
