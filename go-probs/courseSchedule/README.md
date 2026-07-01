# Course Schedule (Medium) — LeetCode #207

**Difficulty:** Medium  
**Pattern:** Graph / topological sort (cycle detection)

## Data Structure

- **Adjacency list** — `map[int][]int` or `[][]int` for prerequisites graph.
- **Indegree array** — count incoming edges per node (Kahn's BFS variant).
- Or **3-state DFS visit** — unvisited / visiting / visited for cycle detection.

## Algorithm (Kahn's BFS)

1. Build graph: edge `prereq → course`.
2. Compute indegree for each of `numCourses` nodes.
3. Enqueue all nodes with indegree 0.
4. While queue non-empty:
   - Pop node, increment processed count.
   - For each neighbor, decrement indegree; enqueue if it hits 0.
5. Return `processed == numCourses` (no cycle iff all nodes processed).

**Alternative (DFS):** Mark nodes visiting/visited; back edge to visiting node ⇒ cycle.

**Time:** O(V + E)  
**Space:** O(V + E)

## Edge Cases

| Case | Notes |
|------|-------|
| No prerequisites | All courses indegree 0 → true |
| Single course | true unless self-loop |
| Cycle in prereqs | false |
| Disconnected components | Still process all nodes |

## Go-Specific Notes

- **`[][]int` prerequisites** — each pair is `[a, b]` meaning `b → a`.
- **Queue** — slice as queue with index head, or `container/list`.
- **Slice indegree** — `make([]int, numCourses)`.

## Other Notes

- Course Schedule II asks for an actual ordering (topological order).
- Same pattern as Alien Dictionary, Build Order in CTCI.
