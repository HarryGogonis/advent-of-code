package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func groupByLine(text []byte) [][]string {
	groupedLines := [][]string{}
	group := []string{}
	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			groupedLines = append(groupedLines, group)
			group = []string{}
		} else {
			group = append(group, lines[i])
		}
	}
	return groupedLines
}

func groupSum(group []string) (int, error) {
	sum := 0
	for _, x := range group {
		n, err := strconv.Atoi(x)
		if err != nil {
			return 0, err
		}
		sum += n
	}
	return sum, nil
}

func main() {
	// read input
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	groups := groupByLine(content)
	sums := []int{}
	for _, group := range groups {
		sum, err := groupSum(group)
		if err != nil {
			log.Fatal(err)
		}
		sums = append(sums, sum)
	}
	sort.Ints(sums)
	log.Printf("sums: %v", sums)
	log.Printf("%+v", sums[len(sums)-1])

}
