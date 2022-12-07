package main

import (
	"strings"
	"testing"

	"github.com/josh-weston/advent-of-code-2022/day7/part1"
	"github.com/josh-weston/advent-of-code-2022/day7/part2"
)

func TestPart1(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt	
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	reader := strings.NewReader(input)
	result := part1.Run(reader)
	expected := 95437

	if result != expected {
		t.Fatalf("incorrect result. got=%d,want=%d", result, expected)
	}

}

func TestPart2(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt	
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	reader := strings.NewReader(input)
	result := part2.Run(reader)
	expected := 24933642

	if result != expected {
		t.Fatalf("incorrect result. got=%d,want=%d", result, expected)
	}

}
