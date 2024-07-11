package graph

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var whitespaces = regexp.MustCompile(`\s+`)

func TestGraph(t *testing.T) {
	tmp := initGraph(t, NewGraph, "tiny.txt")
	g := tmp.(*Graph)
	fmt.Printf("graph: %+v\n", *g)
	dp := NewDfsPaths(g, 0)

	dfsPath := dp.PathTo(4)
	fmt.Printf("dfs path: %+v\n", dfsPath)

	bp := NewBfsPaths(g, 0)

	bfsPath := bp.PathTo(4)
	fmt.Printf("bfs path: %+v\n", bfsPath)
}

func initGraph(t *testing.T, newFunc func(v int) Grapher, input string) Grapher {
	t.Helper()

	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	v, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	// skip edges num
	scanner.Scan()

	g := newFunc(v)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " \n")
		// replace multiple spaces with one
		line = whitespaces.ReplaceAllString(line, " ")
		splitted := strings.Split(line, " ")

		val := strings.Trim(splitted[0], " \n")
		v, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		val = strings.Trim(splitted[1], " \n")
		w, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		g.AddEdge(v, w)
	}

	return g
}
