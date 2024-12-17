package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Block struct {
	id    int
	start int
	end   int
}

func (block Block) Size() int {
	return block.end - block.start + 1
}

func main() {
	// data := getInput("smallInput.txt")
	data := getInput("input.txt")

	fileBlocks, freeBlocks := processData(data)

	fmt.Printf("%+v\n", fileBlocks[:5])
	fmt.Printf("%+v\n", freeBlocks[:5])
	// fmt.Printf("FileMap: \n%v\nFree Space Indexes: \n%v\n", fileMap, freeSpace)

	fillEmptySpaces(fileBlocks, freeBlocks)

	// fmt.Println(diskMap)

	fmt.Println(calcChecksum(fileBlocks))

}

func getInput(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()

	return s.Text()
}

func processData(data string) (fileBlocks []Block, freeBlocks []Block) {
	//data alternates between fileSize and length of freeSpace
	//files represented by their id (begins at 0)
	fileId := 0
	currIndex := 0
	for i := 0; i < len(data); i++ {
		block := Block{}
		//process fileMap
		blockLength, err := strconv.Atoi(string(data[i]))
		// fmt.Println(fileLength)
		if err != nil {
			log.Fatal(err)
		}
		block.start = currIndex
		block.end = currIndex + blockLength - 1
		currIndex += blockLength
		if i%2 == 0 {
			// fmt.Println(currIndex)
			block.id = fileId
			fileBlocks = append(fileBlocks, block)
			fileId++
		} else {
			//process freeSpace
			block.id = -1
			// fmt.Println(currIndex)
			freeBlocks = append(freeBlocks, block)
		}
	}
	return fileBlocks, freeBlocks
}

func fillEmptySpaces(fileBlocks []Block, freeBlocks []Block) {
	for i := len(fileBlocks) - 1; i >= 0; i-- {
		for j := 0; j < len(freeBlocks)-1; j++ {
			if freeBlocks[j].start > fileBlocks[i].start {
				continue
			}
			fileBlockSize := fileBlocks[i].Size()
			if freeBlocks[j].Size() >= fileBlockSize {
				fileBlocks[i].start = freeBlocks[j].start
				fileBlocks[i].end = fileBlocks[i].start + fileBlockSize - 1
				freeBlocks[j].start = fileBlocks[i].end + 1
				break
			}
		}
	}
}

func calcChecksum(fileBlocks []Block) int {
	total := 0
	for _, val := range fileBlocks {
		for i := val.start; i <= val.end; i++ {
			total += val.id * i
		}
	}
	return total
}
