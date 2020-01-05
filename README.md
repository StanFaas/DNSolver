# DNSolver v0.1.1

Used for reconnaissance, DNSolver parses multiline domain files, resolves their IPs, removing duplicates and writes the unique IPs to a new file.

## Table of contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Example domain file](#example-domain-file)
- [Screenshot](#screenshot)
- [Generate domain file](#generate-domain-file)
- [Todo](#todo)

## <a name="prerequisites"></a>Prerequisites

DNSolver is built with Golang. This means that you need Go installed before you can use this tool.

Here you find Go's [installation instructions](https://golang.org/doc/install).

## <a name="installation"></a>Installation

Installation is easy, just copy and paste the following line into your terminal:

`go get -u github.com/StanFaas/DNSolver`

Go will then download, build and install the program for you.

## <a name="usage"></a>Usage

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

## <a name="example-domain-file"></a>Example domain file

```
google.com
github.com
admin.github.com
stanfaas.com
facebook.com
thisisnotasubdomain.facebook.com
```

## <a name="screenshot"></a>Screenshot

![DNS resolver](/screenshot.png?raw=true 'DNS resolver')

## <a name="generate-domain-file"></a>Generate domain file

To generate a file with domains I would recommend [AssetFinder](https://github.com/tomnomnom/assetfinder) by [TomNomNom](https://github.com/tomnomnom)

## <a name="todo"></a>Todo

- Shodan API to include open ports to the report
