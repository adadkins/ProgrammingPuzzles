package programmingpuzzles

func MazeSolutions(maze [][]int, startX, startY int) int {
	if startX >= len(maze[0]) {
		return 0
	}

	if startY >= len(maze) {
		return 0
	}
	if maze[startY][startX] == 1 {
		return 0
	}

	if startX == len(maze[0])-1 && startY == len(maze)-1 {
		return 1
	}

	return MazeSolutions(maze, startX+1, startY) + MazeSolutions(maze, startX, startY+1)
}
