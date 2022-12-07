package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Instruction struct {
	Cmd     string
	Files   []FileDesc
	Folders []string
}

type FileDesc struct {
	Name string
	Size int
}

func getPartOne(rows []Instruction) int {
	sizePerFolder := make(map[string]int, 0)
	folderMap := make(map[string][]string, 0)
	currentFolderStack := make([]string, 0)
	for _, row := range rows {
		size := 0
		if strings.HasPrefix(row.Cmd, "cd ") {
			if row.Cmd == "cd .." {
				currentFolderStack = currentFolderStack[:len(currentFolderStack)-1]
				continue
			} else {
				currentFolderStack = append(currentFolderStack, strings.Replace(row.Cmd, "cd ", "", 1))
			}
		}
		currentFolder := ""
		for _, f := range currentFolderStack {
			if f == "/" {
				currentFolder += "/"
			} else {
				currentFolder += f + "/"
			}
		}
		for _, file := range row.Files {
			size += file.Size
		}
		for _, folder := range row.Folders {
			folderMap[currentFolder] = append(folderMap[currentFolder], currentFolder+folder+"/")
		}
		sizePerFolder[currentFolder] = size
	}

	sum := 0
	for folder := range sizePerFolder {
		s := findFolderSize(folder, sizePerFolder, folderMap)
		if s <= 100000 {
			sum += s
		}
	}

	return sum
}

func findFolderSize(name string, sizePerFolder map[string]int, folderMap map[string][]string) int {
	size := 0
	if val, ok := sizePerFolder[name]; ok {
		size += val
	}
	if _, ok := folderMap[name]; !ok {
		return size
	}
	for _, sub := range folderMap[name] {
		size += findFolderSize(sub, sizePerFolder, folderMap)
	}
	return size
}

func getPartTwo(rows []Instruction) int {
	sizePerFolder := make(map[string]int, 0)
	folderMap := make(map[string][]string, 0)
	currentFolderStack := make([]string, 0)
	for _, row := range rows {
		size := 0
		if strings.HasPrefix(row.Cmd, "cd ") {
			if row.Cmd == "cd .." {
				currentFolderStack = currentFolderStack[:len(currentFolderStack)-1]
				continue
			} else {
				currentFolderStack = append(currentFolderStack, strings.Replace(row.Cmd, "cd ", "", 1))
			}
		}
		currentFolder := ""
		for _, f := range currentFolderStack {
			if f == "/" {
				currentFolder += "/"
			} else {
				currentFolder += f + "/"
			}
		}
		for _, file := range row.Files {
			size += file.Size
		}
		for _, folder := range row.Folders {
			folderMap[currentFolder] = append(folderMap[currentFolder], currentFolder+folder+"/")
		}
		sizePerFolder[currentFolder] = size
	}

	total := 70000000
	wanted := 30000000
	actual := findFolderSize("/", sizePerFolder, folderMap)
	res := make([]int, 0)
	for folder := range sizePerFolder {
		s := findFolderSize(folder, sizePerFolder, folderMap)
		if total-actual+s > wanted {
			res = append(res, s)
		}
	}

	sort.Ints(res)

	return res[0]
}

func getRows(filename string) []Instruction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := make([]Instruction, 0)
	i := -1
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "$") {
			var inst Instruction
			inst.Cmd = strings.Replace(row, "$ ", "", 1)
			instructions = append(instructions, inst)
			i++
		} else {
			if strings.Contains(row, "dir") {
				var folder string
				fmt.Sscanf(row, "dir %s",
					&folder)
				instructions[i].Folders = append(instructions[i].Folders,
					folder)
			} else {
				var fileDesc FileDesc
				fmt.Sscanf(row, "%d %s",
					&fileDesc.Size, &fileDesc.Name)
				instructions[i].Files = append(instructions[i].Files,
					fileDesc)
			}
		}
	}
	return instructions
}

func main() {
	//printIf(4)
	//printIf(14)
	fmt.Println("Part one:", getPartOne(getRows("../input.txt")))
	fmt.Println("Part two:", getPartTwo(getRows("../input.txt")))
}
