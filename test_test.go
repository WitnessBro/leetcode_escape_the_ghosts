package partitionlabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var ghost = [][]int{{2, 0}} // Input array
	var target = []int{1, 0}
	var expectedResult = true          // The expected answer with correct length.
	res := escapeGhosts(ghost, target) // Calls your implementation

	assert.Equal(t, expectedResult, res)
}
func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func escapeGhosts(ghosts [][]int, target []int) bool {
	enemyPaths := make([]int, len(ghosts))
	myPath := manhattanDistance(0, 0, target[0], target[1])
	for i := range ghosts {
		enemyPaths[i] = manhattanDistance(target[0], target[1], ghosts[i][0], ghosts[i][1])
	}
	for _, val := range enemyPaths {
		if val <= myPath {
			return false
		}
	}
	return true
}
