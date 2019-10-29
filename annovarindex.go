package main

import (
	"fmt"
	"flag"
	"github.com/anarkia7115/annovar_index/read"
)

func main() {

	taskHelperText := "task choose in rb, rl, os, rs, ph"
	task := flag.String("task", "", taskHelperText)
	inputFile := flag.String("in", "./data/hg19_dbnsfp35a_1m.txt", "input file")
	chunkSize := flag.Int64("n", 1000, "chunk size for byte read")
	seekTime := flag.Int("sn", 1, "total time for seek")
	printReadBytes := flag.Bool("p", false, "print read bytes")
	
	flag.Parse()

	switch *task {
	case "rb":
		read.PassBytes(*inputFile, *chunkSize)
	case "rl":
		read.PassLines(*inputFile)
	case "ph":
		read.PrintHeader(*inputFile)
	case "os":
		read.PassSeekInOrder(*inputFile, *seekTime, *chunkSize, *printReadBytes)
	case "rs":
		read.RandSeek(*inputFile, *seekTime, *chunkSize)
	default:
		fmt.Printf("%s\nYou give [%s]\n", taskHelperText, *task)
	}
}
