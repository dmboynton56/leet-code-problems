# Add Two Numbers (Medium) — LeetCode #2

**Difficulty:** Medium  
**Pattern:** Linked list / simulation / carry

## Data Structure

- **Two singly linked lists** — digits in **reverse** order (ones place at head).
- **Result list** built with dummy head.

## Algorithm

1. Loop while `l1`, `l2`, or `carry` remain.
2. Sum = carry + optional l1 digit + optional l2 digit.
3. Append `sum % 10`, carry = `sum / 10` (integer division).
4. Advance non-nil pointers.

**Time:** O(max(n, m))  
**Space:** O(max(n, m)) for result list

## Edge Cases

| Case | Example |
|------|---------|
| Different lengths | `[9,9] + [1]` → `[0,0,1]` |
| Final carry | `[5] + [5]` → `[0,1]` |
| One list nil mid-loop | Still process carry |
| Zeros in list | `[0] + [0]` → `[0]` |
| Empty lists | Not in typical constraints |

## Go-Specific Notes

- **Loop condition** — `l1 != nil || l2 != nil || carry != 0`; easy to forget carry-only iteration.
- **Integer division** — `sum / 10` for positive ints equals floor division here.
- **Nil before dereference** — always check `l1 != nil` before `l1.Val`.
- **Allocating nodes** — `&ListNode{Val: sum % 10}` each digit; no slice backing.
- Digits are 0–9; sum ≤ 27 per step — no overflow concerns.

## Other Notes

- Digits stored **reversed** — no manual reverse needed; if stored forward, reverse or use stack.
- Follow-up: Add without reversing (LC 445) uses two stacks or reverse lists first.
