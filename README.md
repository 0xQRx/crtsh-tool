
# crtsh-tool

A Go-based command-line tool for querying `crt.sh` to retrieve domain certificates. This tool fetches certificates for subdomains of a given domain and outputs the results to the console or an optional file.

## Features

- Fetches certificate data from `crt.sh` for a given domain.
- Outputs subdomains in a clean, deduplicated format.
- Allows saving the results to a file.

## Installation

To install the `crtsh-tool` CLI tool, use the following `go install` command:

```bash
GOPRIVATE=github.com/0xQRx/crtsh-tool go install github.com/0xQRx/crtsh-tool/cmd/crtsh-tool@latest
```

Ensure your `GOPATH/bin` directory is in your system's `PATH` so the installed binary can be accessed directly.

## Usage

### Basic Usage
Run the tool to query `crt.sh` for subdomains of a domain:
```bash
crtsh-tool --domain example.com
```

### Save Results to a File
Use the `--outputfile` option to save the output to a file:
```bash
crtsh-tool --domain example.com --outputfile results.txt
```

## Development

### Clone the Repository
If you want to modify or build the tool locally, clone the repository:
```bash
git clone https://github.com/0xQRx/crtsh-tool.git
cd crtsh-tool
```

### Build Locally
You can build the project using the `go build` command:
```bash
go build -o crtsh-tool ./cmd/crtsh-tool
```

Run the compiled binary:
```bash
./crtsh-tool --domain example.com
```

### Testing
Test the tool with different domains and output options to verify its functionality.


## Acknowledgments
- Special thanks to [crt.sh](https://crt.sh/) for providing free and open access to certificate data.

---

Enjoy using the tool!
