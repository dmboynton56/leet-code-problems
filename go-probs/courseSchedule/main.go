package main

import "fmt"

// solution returns true if all courses can be finished (no cycle in prereq graph).
func solution(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		course, prereq := p[0], p[1]
		graph[prereq] = append(graph[prereq], course)
		indegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	processed := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		processed++

		for _, next := range graph[node] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return processed == numCourses
}

func main() {
	fmt.Println(solution(2, [][]int{{1, 0}}))       // true
	fmt.Println(solution(2, [][]int{{1, 0}, {0, 1}})) // false
}
