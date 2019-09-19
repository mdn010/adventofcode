package main

import (
	"fmt"
	"github.com/mdn010/adventofcode/common"
	"bufio"
	"sort"
	"strings"
)

type graph struct {
	// [A] [C,D,F]
	// [B] [C]
	mapOfVertices map[string][]string
	parents []string 
}

func newGraph() *graph {
	g := new(graph)
	g.mapOfVertices = make(map[string][]string)
	g.parents = strings.Split("ZYXWVUTSRQPONMLKJIHGFEDCBA", "")
	return g
}

func (g *graph) addEdge(source string, target string) {
	array := (*g).getMapOfVertices(source)
	(*g).mapOfVertices[source] = append(array, target)
	sort.Strings((*g).mapOfVertices[source])
	(*g).markAsNotParent(target)
}

func (g *graph) markAsNotParent(target string) {
	filteredLetters := (*g).parents[:0]
    for _, parent := range (*g).parents {
        if target != parent {
            filteredLetters = append(filteredLetters, parent)
        }
    }
    (*g).parents = filteredLetters
}

func (g *graph) getMapOfVertices(source string) []string {
	if (*g).mapOfVertices[source] == nil {
		return []string{}
	} else {
		return (*g).mapOfVertices[source]
	}
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func getNext(key string, visited map[string]bool, temp *[]string, mapOfVertices map[string][]string, stack *[]string) {
	if visited[key] {
		return
	} else {
		visited[key] = true
	}

	*temp = append(*temp, key)
	for i := 1; i <= len(mapOfVertices[key]); i++ {
		pop:=mapOfVertices[key][len(mapOfVertices[key])-i]
		getNext(pop, visited, temp, mapOfVertices, stack)
	}
	pop:=(*temp)[len(*temp)-1]
    *temp=(*temp)[:len(*temp)-1]
	*stack = append(*stack, pop)}

func main() {
	elves := newGraph()

	file := common.NewFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text();
		var source, target string
		_, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &source, &target)
		common.Checkerror(err)
		elves.addEdge(source, target)
	}
	
	// fmt.Println(elves.mapOfVertices)

	stack := make([]string, 0, 5)
	temp := make([]string, 0, 5)
	visited := make(map[string]bool)

	for i := 26; i >= 1; i-- {
		key := toCharStr(i)
		getNext(key, visited, &temp, elves.mapOfVertices, &stack)
	}

	for index := len(stack)-1; index >= 0; index-- {
		fmt.Print(stack[index])
	}
}