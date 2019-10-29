package read

import (
	"math/rand"
	"bufio"
	"io"
	"log"
	"os"
)

// const inputFile string = "./data/hg19_dbnsfp35a_1m.txt"

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// OpenFile open file return reader
func OpenFile(inputFile string) *os.File {
	log.Printf("open [%s]\n", inputFile)
	f, err := os.Open(inputFile)
	checkErr(err)

	return f
}

// OpenFileBuf open file return buf reader
func OpenFileBuf(inputFile string) *bufio.Reader {
	log.Printf("open [%s]\n", inputFile)
	f, err := os.Open(inputFile)
	checkErr(err)
	rd := bufio.NewReader(f)

	return rd
}

// PrintHeader print header of sfp file
func PrintHeader(inputFile string) {
	rd := OpenFileBuf(inputFile)

	line, err := rd.ReadString('\n')
	checkErr(err)

	log.Print(line)
}

// PassLines do nothing but run over every line in file
func PassLines(inputFile string) {
	log.Println("parse by line")
	rd := OpenFileBuf(inputFile)

	for {
		_, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}
		checkErr(err)
	}
	log.Println("File parsed")
}

// PassBytes pass through file and do nothing
func PassBytes(inputFile string, chunkSize int64) {
	log.Printf("parse bytes with chunkSize %d\n", chunkSize)
	rd := OpenFile(inputFile)
	for {
		b := make([]byte, chunkSize)
		_, err := rd.Read(b)
		if err == io.EOF {
			break
		}
		checkErr(err)
	}
}

func getFileSize(inputFile string) int64 {
	f, err := os.Open(inputFile)
	checkErr(err)
	fi, err := f.Stat()
	checkErr(err)
	return fi.Size()
}

// PassSeekInOrder pass file one time with certain time of seek 
func PassSeekInOrder(inputFile string, seekTime int, byteSize int64, printReadBytes bool) {
	log.Printf("parse in order seek [%d] times, and read [%d] each", seekTime, byteSize)
	rd := OpenFile(inputFile)
	fileSize := getFileSize(inputFile)
	seekBlockSize := fileSize / int64(seekTime)
	log.Printf("file size: [%d]", fileSize)
	log.Printf("seekBlockSize: [%d]", seekBlockSize)

	readTime := 0

	for seekPos:=int64(0); seekPos < fileSize; {
		// seek and read
		b := make([]byte, byteSize)
		rd.ReadAt(b, seekPos)
		if printReadBytes {
			log.Println("printing bytes:[" + string(b) + "]")
		}
		seekPos += seekBlockSize
		readTime++
	}
	log.Printf("read [%d] times\n", readTime)
}

// RandSeek rand seek in file
func RandSeek(inputFile string, seekTime int, byteSize int64) {
	log.Printf("rand seek in [%s] with seekTime [%d]", inputFile, seekTime)
	fileSize := getFileSize(inputFile)
	log.Printf("file size: [%d]", fileSize)

	rd := OpenFile(inputFile)

	for i:=0; i < seekTime; i++ {
		// get rand num
		seekPos := rand.Int63n(fileSize)
		b := make([]byte, byteSize)
		rd.ReadAt(b, seekPos)
	}
}