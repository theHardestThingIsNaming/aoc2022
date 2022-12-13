package main

import (
	"fmt"
	"os"
	"strings"
)

type dir struct {
	path       string
	subDirs    []string
	parentPath string
	totalSize  int
}

func addFileSizeToAllParentsTotalSize(fileSize int, currentDir dir) {
	if currentDir.parentPath != "" {
		parentDir := dirMap[currentDir.parentPath]
		parentDir.totalSize += fileSize
		dirMap[currentDir.parentPath] = parentDir
		addFileSizeToAllParentsTotalSize(fileSize, parentDir)
	}
}

func getDirPath(currentDir dir, dirName string) string {
	if currentDir.path == "/" {
		return "/" + dirName
	}
	return currentDir.path + "/" + dirName
}

func processList(line string) {
	currentDir := dirMap[cwd]
	if strings.HasPrefix(line, "dir") {
		var dirName string
		fmt.Sscanf(line, "dir %s", &dirName)
		dirPath := getDirPath(currentDir, dirName)
		if _, ok := dirMap[dirPath]; !ok {
			dirMap[dirPath] = dir{path: dirPath, parentPath: currentDir.path}
		}
		currentDir.subDirs = append(currentDir.subDirs, dirPath)
		dirMap[cwd] = currentDir
	} else {
		var fileSize int
		fmt.Sscanf(line, "%d %s", &fileSize)
		currentDir.totalSize += fileSize
		addFileSizeToAllParentsTotalSize(fileSize, currentDir)
		dirMap[cwd] = currentDir
	}
}

func processChangeDir(line string) {
	var dirName string
	fmt.Sscanf(line, "$ cd %s", &dirName)
	if dirName == "/" {
		return
	}
	currentDir := dirMap[cwd]
	if dirName == ".." {
		if currentDir.parentPath != "" {
			cwd = currentDir.parentPath
		}
		return
	}
	cwd = getDirPath(currentDir, dirName)
}

var cwd string = "/"
var dirMap = map[string]dir{"/": {path: "/"}}

func main() {
	input, _ := os.ReadFile("input.txt")
	termOuput := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, line := range termOuput {
		if strings.HasPrefix(line, "$ cd") {
			processChangeDir(line)
		} else if !strings.HasPrefix(line, "$") {
			processList(line)
		}
	}
	partOne, partTwo := 0, 70000000
	remainingSpace := partTwo - dirMap["/"].totalSize
	for _, dir := range dirMap {
		if dir.totalSize <= 100000 {
			partOne += dir.totalSize
		}
		if remainingSpace+dir.totalSize >= 30000000 && partTwo > dir.totalSize {
			partTwo = dir.totalSize
		}
	}
	fmt.Println("partOne: ", partOne)
	fmt.Println("partTwo: ", partTwo)
}
