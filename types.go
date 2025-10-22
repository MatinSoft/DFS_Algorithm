package main

type Point struct {
	Row int
	Col int
}

type Wall struct {
	State Point
	wall  bool
}

type Maze struct {
	Width       int
	Height      int
	Start       Point
	Goal        Point
	Debug       bool
	Walls       [][]Wall
	CurrentNode *Node
	Solution    Solution
	NumExplored int
	Explored    []Point
	SearchType  int
}

type Node struct {
	index       int
	State       Point
	Parent      *Node
	Action      string
	Steps       int
	NumExplored int
	Debug       bool
	SearchType  int
}

type Solution struct {
	Actions []string
	Cells   []Point
}

type DepthFirstSearch struct {
	Frontier []*Node
	Game     *Maze
}
