# CSV Generator

CSV Generator is a command-line tool for generating CSV files with random data based on specified parameters. This tool creates CSV files according to the specified file size, number of columns, number of files, and output directory.

## Features

- Generate CSV files with random data
- Specify file size, number of columns, and number of files
- Automatically create the output directory if it doesn't exist

## Installation

This project is written in Go. Follow these steps to install the necessary dependencies:

1. Install Go (https://golang.org/doc/install)
2. Install the Cobra library:

```bash
go get -u github.com/spf13/cobra@latest
```

## Usage

Run the following command from the command line to generate CSV files:

```bash
go run main.go --dir /path/to/output --file-size 1048576 --num-columns 100 --num-files 1
```

### Arguments

- --dir or -d: Output directory (default is "outdir")
- --prefix or -p: File name prefix (default is "out")
- --file-size or -s: File size in bytes  (default is 1024 bytes)
- --num-columns or -c: Number of columns (default is 5)
- --num-files or -f: Number of files (default is 1)

## Example

The following command creates a CSV file with a size of 1MB, 100 columns, and outputs it to the specified directory:

```bash
go run main.go --dir ./out -p filename --file-size 1048576 --num-columns 100 --num-files 1
```

### License

This project is licensed under the MIT License. See the LICENSE file for details.