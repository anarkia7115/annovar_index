package main

import (
	"fmt"
	"flag"
	"github.com/anarkia7115/annovar_index/read"
)

func main() {

	taskHelperText := "task choose in rb, rl, os, rs, ph, mos, mrs"
	task := flag.String("task", "", taskHelperText)
	inputFile := flag.String("in", "./data/hg19_dbnsfp35a_1m.txt", "input file")
	chunkSize := flag.Int64("n", 1000, "chunk size for byte read")
	seekTime := flag.Int("sn", 1, "total time for seek")
	printReadBytes := flag.Bool("p", false, "print read bytes")
	loop := flag.Int("l", 1, "loop seek util")
	
	flag.Parse()

	switch *task {
	case "rb":
		read.PassBytes(*inputFile, *chunkSize)
	case "rl":
		read.PassLines(*inputFile)
	case "ph":
		read.PrintHeader(*inputFile)
	case "os":
		read.PassSeekInOrder(*inputFile, *seekTime, *chunkSize, *printReadBytes, *loop)
	case "rs":
		read.RandSeek(*inputFile, *seekTime, *chunkSize, *loop)
	case "mos":
		read.PassMmapSeekInOrder(*inputFile, *seekTime, *chunkSize, *printReadBytes, *loop)
	case "mrs":
		read.RandMmapSeek(*inputFile, *seekTime, *chunkSize, *loop)
	default:
		fmt.Printf("%s\nYou give [%s]\n", taskHelperText, *task)
	}
}
