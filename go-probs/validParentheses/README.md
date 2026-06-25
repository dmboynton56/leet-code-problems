# Valid Parentheses (Easy) — LeetCode #20

**Difficulty:** Easy  
**Pattern:** Stack

## Data Structure

- **Stack** (slice in Go, list in Python): holds unmatched opening brackets in order.

## Algorithm

1. Scan left to right.
2. On `(`, `[`, `{` — push onto stack.
3. On `)`, `]`, `}` — if stack empty or top doesn't match, return false; else pop.
4. After scan, valid iff stack is empty.

**Time:** O(n)  
**Space:** O(n) worst case (all opens)

## Edge Cases

| Case | Expected |
|------|----------|
| Empty string | Valid (stack empty) |
| Only opens `"((("` | Invalid |
| Only closes `"))"` | Invalid (stack empty on first close) |
| Interleaved wrong `"([)]"` | Invalid |
| Nested `"{[]}"` | Valid |
| Single pair | Valid |

## Go-Specific Notes

- **`[]rune` vs `[]byte`** — `range` on `string` yields **runes** (Unicode code points). For ASCII brackets, either works; runes are safer for general strings.
- **No negative indexing** — use `stack[len(stack)-1]`; Python allows `stack[-1]`.
- **Pop via reslice** — `stack = stack[:len(stack)-1]`; no built-in pop. Python: `stack.pop()`.
- **Empty check** — `len(stack) == 0`; don't rely on nil/falsey like Python's `if not stack` (though `len(nil slice) == 0` in Go).
- **Char comparisons** — rune literals like `'('` are typed; comparing rune to rune is clean.
- Optional: map closing→opening char like Python for less branching.

## Other Notes

- Cannot solve with a simple open/close counter — nesting order matters (stack required).
- Matching must be immediate LIFO, not global counting.
