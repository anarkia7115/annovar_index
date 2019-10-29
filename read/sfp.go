package read

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFile string = "./data/hg19_dbnsfp35a_1m.txt"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// OpenFile open file return reader
func OpenFile() *bufio.Reader {
	f, err := os.Open(inputFile)
	checkErr(err)

	rd := bufio.NewReader(f)

	return rd
}

// PrintHeader print header of sfp file
func PrintHeader() {
	rd := OpenFile()

	line, err := rd.ReadString('\n')
	checkErr(err)

	fmt.Print(line)
}

// PassLines do nothing but run over every line in file
func PassLines() {
	rd := OpenFile()

	for {
		_, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		checkErr(err)
	}
	fmt.Println("File parsed")
}

// PassBytes pass through file and do nothing
func PassBytes() {
	const chunkSize int = 1000
	rd := OpenFile()
	for {
		b := make([]byte, chunkSize)
		_, err := rd.Read(b)
		if err == io.EOF {
			break
		}
		checkErr(err)
	}
}