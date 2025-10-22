package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const (
	DFS = iota
	BFS
	GBFS
	ASTAR
	DIJKSTRA
)

func main() {
	var m Maze
	var maze, searchType string

	flag.StringVar(&maze, "file", "maze.txt", "maze file")
	flag.StringVar(&searchType, "search", "dfs", "searchType")
	flag.Parse()

	err := m.Load(maze)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	startTime := time.Now()

	switch searchType {
	case "dfs":
		m.SearchType = DFS
		SolveDFS(&m)
	default:
		fmt.Println("Invalid search Type")
		os.Exit(1)
	}

	if len(m.Solution.Actions) > 0 {
		fmt.Println("Solution:")
		m.printMaze()
		fmt.Println("Solution is ", len(m.Solution.Cells), "steps.")
		fmt.Println("Time to solve :", time.Since(startTime))
	} else {
		fmt.Println("no solution")
	}

	fmt.Println("Explored", len(m.Explored), "nodes.")
}

func (m *Maze) printMaze() {

	for r, row := range m.Walls {

		for c, col := range row {
			if col.wall {
				fmt.Print("\u2588")
			} else if m.Start.Row == col.State.Row && m.Start.Col == col.State.Col {
				fmt.Print("A")
			} else if m.Goal.Row == col.State.Row && m.Goal.Col == col.State.Col {
				fmt.Print("B")
			} else if m.inSolution(Point{Row: r, Col: c}) {
				fmt.Print("*")
			} else {

				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (m *Maze) inSolution(x Point) bool {

	for _, step := range m.Solution.Cells {
		if step.Row == x.Row && step.Col == x.Col {
			return true
		}
	}
	return false

}

func SolveDFS(m *Maze) {

	var s DepthFirstSearch
	s.Game = m

	fmt.Println("Goal is", s.Game.Goal)
	s.Solve()
}

func (g *Maze) Load(fileName string) error {

	f, error := os.Open("./mazes/" + fileName)
	if error != nil {
		fmt.Printf("error opening the file %s : %s\n", fileName, error)
	}

	defer f.Close()

	var fileContents []string

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New(fmt.Sprintf("error opening the file %s : %s\n", fileName, error))
		}

		fileContents = append(fileContents, line)
	}
	foundStart, foundEnd := false, false

	for _, line := range fileContents {
		if strings.Contains(line, "A") {
			foundStart = true
		}
		if strings.Contains(line, "B") {
			foundEnd = true
		}
	}
	if !foundStart {
		return errors.New("start location not found")
	}
	if !foundEnd {
		return errors.New("end location not found")
	}

	g.Height = len(fileContents)
	g.Width = len(fileContents[0])

	var rows [][]Wall

	for i, row := range fileContents {
		var Cols []Wall

		for j, col := range row {
			curLetter := fmt.Sprintf("%c", col)
			var wall Wall

			switch curLetter {
			case "A":
				g.Start = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false

			case "B":
				g.Goal = Point{Row: i, Col: j}
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false

			case " ":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = false

			case "#":
				wall.State.Row = i
				wall.State.Col = j
				wall.wall = true
			default:
				continue
			}
			Cols = append(Cols, wall)
		}
		rows = append(rows, Cols)
	}
	g.Walls = rows
	return nil
}
