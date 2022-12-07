package part2

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
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

	usedSize := 0
	// first traversal is to get the total size
	traverse = func(n *Node) int {

		if len(n.Children) == 0 {
			return n.Filesizes
		}
		childrenSize := 0
		for _, c := range n.Children {
			childrenSize += traverse(c)
		}
		dirSize := childrenSize + n.Filesizes

		// we do this to avoid double dipping on the parent node
		if n.Parent == nil {
			usedSize += dirSize
		}
		return dirSize
	}

	traverse(rootNode)

	totalDiskSize := 70_000_000
	requiredSpace := 30_000_000
	freeSpace := totalDiskSize - usedSize
	neededSpace := requiredSpace - freeSpace

	fmt.Printf("Total used size: %d\n", usedSize)
	fmt.Printf("Total required size: %d\n", requiredSpace)
	fmt.Printf("Total freeSpace: %d\n", freeSpace)
	fmt.Printf("Additional space required: %d\n", neededSpace)

	// Now we traverse the tree looking for the smallest directory that will work
	var find func(n *Node) int
	candidateDirectories := []int{}
	find = func(n *Node) int {

		if len(n.Children) == 0 {
			if n.Filesizes >= neededSpace {
				candidateDirectories = append(candidateDirectories, n.Filesizes)
			}
			return n.Filesizes
		}
		childrenSize := 0
		for _, c := range n.Children {
			childrenSize += find(c)
		}
		dirSize := childrenSize + n.Filesizes

		if n.Parent != nil {
			if dirSize >= neededSpace {
				candidateDirectories = append(candidateDirectories, dirSize)
			}
		}
		return dirSize
	}

	find(rootNode)

	sort.Ints(candidateDirectories)
	fmt.Println(candidateDirectories)
	return candidateDirectories[0]
}

// total disk size available = 70_000_000
// free space required = 30_000_00
// find the smallest directory, that if deleted, would free up enough space on the filesystem to run the update. What is the total size of that directory?
