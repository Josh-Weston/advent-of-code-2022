package part1

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// determine the total size of each directory
// total size = sum of all files the directory contains (including subdirectories)
// find all directories with a size <= 100_000 and calculate the sum of their total sizes

type Node struct {
	Parent    *Node
	Children  []*Node
	Filesizes int
}

var reIsCommand = regexp.MustCompile(`^\$`)
var reCommand = regexp.MustCompile(`^\$ (?P<Command>.*)`)
var reChangeDirectoryCommand = regexp.MustCompile(`^\$ cd (?P<Directory>.*)`)
var reIsFile = regexp.MustCompile(`^\d+`)
var reFileSize = regexp.MustCompile(`^(?P<FileSize>\d+)`)

func isCommand(s string) bool {
	return reIsCommand.MatchString(s)
}

func isListCommand(s string) bool {
	matches := reCommand.FindStringSubmatch(s)
	commandIndex := reCommand.SubexpIndex("Command")
	cmd := matches[commandIndex]
	return strings.HasPrefix("ls", cmd)
}

func parseChangeDirectoryCommand(s string) string {
	matches := reChangeDirectoryCommand.FindStringSubmatch(s)
	dirIndex := reChangeDirectoryCommand.SubexpIndex("Directory")
	return matches[dirIndex]
}

func isFile(s string) bool {
	return reIsFile.MatchString(s)
}

func parseFileSize(s string) int {
	matches := reFileSize.FindStringSubmatch(s)
	fileIndex := reFileSize.SubexpIndex("FileSize")
	size, err := strconv.ParseInt(matches[fileIndex], 10, 64)
	if err != nil {
		panic(err)
	}
	return int(size)
}

// Note: we can actually ignore the directories when using the list command because we only care about them when we traverse into them

func Run(input io.Reader) int {

	rootNode := &Node{
		Parent:    nil,
		Children:  []*Node{},
		Filesizes: 0,
	}

	// our starting node
	currNode := rootNode
	scanner := bufio.NewScanner(input)
	inListCommand := false
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		txt := scanner.Text()
		fmt.Println("Line:", txt) // Note: println adds the extra space

		if isCommand(txt) {
			if inListCommand {
				inListCommand = false
			}
			if isListCommand(txt) {
				inListCommand = true
			} else {
				dir := parseChangeDirectoryCommand(txt)
				if dir == ".." {
					currNode = currNode.Parent // move up a level of our tree
				} else {
					// Note: this assumes we only traverse the directory once
					childNode := &Node{
						Parent:    currNode,
						Children:  []*Node{},
						Filesizes: 0,
					}
					currNode.Children = append(currNode.Children, childNode)
					currNode = childNode
				}
			}
			continue
		}

		// Note: we ignore directories for now as we assume some may not be traversed
		if isFile(txt) {
			fileSize := parseFileSize(txt)
			currNode.Filesizes += fileSize
		}
	}

	// Now we traverse the tree totalling the directories
	var traverse func(n *Node) int

	totalSize := 0
	traverse = func(n *Node) int {
		if len(n.Children) == 0 {
			if n.Filesizes < 100_000 {
				totalSize += n.Filesizes
			}
			fmt.Println(n.Filesizes)
			return n.Filesizes
		}
		childrenSize := 0
		for _, c := range n.Children {
			childrenSize += traverse(c)
		}
		dirSize := childrenSize + n.Filesizes
		fmt.Println(dirSize)

		if dirSize < 100_000 {
			totalSize += dirSize
		}
		return dirSize
	}

	traverse(rootNode)
	return totalSize
}
