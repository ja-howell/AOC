package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {
	// data := getInput("smallInput.txt")
	data := getInput("input.txt")

	fileMap, freeSpace := processData(data)

	// fmt.Printf("FileMap: \n%v\nFree Space Indexes: \n%v\n", fileMap, freeSpace)

	diskMap := fillEmptySpaces(fileMap, freeSpace)

	fmt.Println(diskMap)

	fmt.Println(calcChecksum(diskMap))

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

func processData(data string) (fileMap []int, freeSpace []int) {
	//data alternates between fileSize and length of freeSpace
	//files represented by their id (begins at 0)
	fileId := 0
	currIndex := 0
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			//process fileMap
			fileLength, err := strconv.Atoi(string(data[i]))
			if err != nil {
				log.Fatal(err)
			}
			newFileMap := populate(fileLength, fileId, true)
			fileMap = append(fileMap, newFileMap...)
			currIndex += len(newFileMap)
			fileId++
		} else {
			//process freeSpace
			freeSpaceLength, err := strconv.Atoi(string(data[i]))
			if err != nil {
				log.Fatal(err)
			}
			newFreeSpace := populate(freeSpaceLength, currIndex, false)
			freeSpace = append(freeSpace, newFreeSpace...)
			currIndex += len(newFreeSpace)
		}
	}
	return fileMap, freeSpace
}

func populate(length int, id int, isFile bool) []int {
	newSlice := []int{}

	for i := 0; i < length; i++ {
		newSlice = append(newSlice, id)
		if !isFile {
			id++
		}
	}
	return newSlice
}

func fillEmptySpaces(fileMap []int, freeSpace []int) []int {
	fmt.Printf("File Map: \n%v\nFree Space: \n%v\n", fileMap, freeSpace)
	for i := 0; i < len(freeSpace); i++ {
		fileMapIndex := len(fileMap) - 1
		if freeSpace[i] >= fileMapIndex {
			break
		}
		// fmt.Printf("i: %d, fileMapIndex: %d, filemap[fileMapIndex]: %d, freespace[i]: %d\n", i, fileMapIndex, fileMap[fileMapIndex], freeSpace[i])
		fileMap = slices.Insert(fileMap, freeSpace[i], fileMap[fileMapIndex])
		fileMap = slices.Delete(fileMap, len(fileMap)-1, len(fileMap))
		// fmt.Println(fileMap)
	}
	return fileMap
}

func calcChecksum(fileMap []int) int {
	total := 0
	for i, val := range fileMap {
		total += val * i
	}
	return total
}
