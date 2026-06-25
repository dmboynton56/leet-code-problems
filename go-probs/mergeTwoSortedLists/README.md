# Merge Two Sorted Lists (Easy) — LeetCode #21

**Difficulty:** Easy  
**Pattern:** Linked list / two pointers / dummy head

## Data Structure

- **Singly linked list** — `ListNode` with `Val` and `Next`.
- **Dummy sentinel node** — simplifies head insertion logic.

## Algorithm

1. Create dummy node; `curr` points to tail of merged list.
2. While both lists non-empty, attach smaller head, advance that list and `curr`.
3. Attach remaining non-empty list (already sorted).
4. Return `dummy.Next`.

**Time:** O(n + m)  
**Space:** O(1) extra (relinks existing nodes)

## Edge Cases

| Case | Handling |
|------|----------|
| One list empty | Return the other |
| Both empty | Return nil |
| All of l1 < all of l2 | Drain l1 then append rest of l2 |
| Equal values | Either order OK; use `<=` for stability |

## Go-Specific Notes

- **Pointers everywhere** — `*ListNode`, `nil` for empty; no implicit references like Python.
- **`&ListNode{Val: v}`** — composite literal + address-of allocates node on heap.
- **Dummy head pattern** — `return dummy.Next` skips sentinel; very common in Go LL problems.
- **Nil checks** — `for list1 != nil && list2 != nil`; never dereference nil.
- LeetCode provides `ListNode` in stub — don't redefine on submit if already there.
- Helper `buildList` / `listToSlice` are for local testing only.

## Other Notes

- Recursive merge is O(n+m) stack space — iterative preferred in interviews.
- Contrast with merge step in merge sort on arrays (different memory profile).
