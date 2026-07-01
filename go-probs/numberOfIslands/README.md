# Number of Islands (Medium) — LeetCode #200

**Difficulty:** Medium  
**Pattern:** Graph traversal / matrix DFS or BFS

## Data Structure

- **2D grid** of `'1'` (land) and `'0'` (water).
- **Visited tracking** — mutate grid in-place (`'1'` → `'0'`) or use a `[][]bool` / set.

## Algorithm

1. Count = 0.
2. For each cell `(r, c)`:
   - If `grid[r][c] == '1'`, increment count and **flood-fill** the entire connected component (4-directional).
3. Flood-fill (DFS or BFS):
   - Mark current cell visited.
   - Recurse/queue all 4 neighbors that are in bounds and land.
4. Return count.

**Time:** O(m × n) — each cell visited once  
**Space:** O(m × n) worst case for recursion stack or BFS queue

## Edge Cases

| Case | Notes |
|------|-------|
| Empty grid | Return 0 |
| All water | Return 0 |
| All land | Return 1 |
| Single cell | `'1'` → 1, `'0'` → 0 |
| Diagonals don't connect | Only 4-directional adjacency |

## Go-Specific Notes

- **Grid type** — `[][]byte` or `[][]rune`; LeetCode often gives `[][]byte` with `'1'`/`'0'`.
- **DFS** — recursive; watch stack depth on huge grids (BFS with queue is iterative alternative).
- **Bounds check** — `r >= 0 && r < rows && c >= 0 && c < cols`.
- **Mutating input** — common interview pattern; mention if immutability is required.

## Other Notes

- Same template applies to: Max Area of Island, Surrounded Regions, Pacific Atlantic Water Flow.
- Union-Find also works but DFS/BFS is simpler here.
