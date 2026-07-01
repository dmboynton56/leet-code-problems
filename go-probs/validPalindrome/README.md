# Valid Palindrome (Easy) — LeetCode #125

**Difficulty:** Easy  
**Pattern:** Two pointers / string filtering

## Data Structure

- **Two indices** `left` and `right` on the string.
- Alphanumeric check via helper or `unicode.IsLetter` / `unicode.IsDigit`.

## Algorithm

1. Set `left = 0`, `right = len(s) - 1`.
2. While `left < right`:
   - Skip non-alphanumeric from left and right.
   - Compare lowercased chars at both pointers.
   - If mismatch → false.
   - Else increment `left`, decrement `right`.
3. Return true.

**Time:** O(n)  
**Space:** O(1)

## Edge Cases

| Case | Notes |
|------|-------|
| Empty string | true |
| Only punctuation/spaces | true (no pairs to fail) |
| Mixed case | Normalize to lowercase |
| Single alphanumeric | true |
| Numbers in string | Treat as valid chars |

## Go-Specific Notes

- **`unicode` package** — `unicode.IsLetter(r)`, `unicode.IsDigit(r)` on runes.
- **String indexing** — `s[i]` is byte; for ASCII alnum this is fine; full Unicode needs `[]rune(s)`.
- **`strings.ToLower`** on single byte works for ASCII letters.

## Other Notes

- Valid Palindrome II (#680) allows deleting one char — greedy or DP extension.
- Is Palindrome for linked list uses reverse half + compare.
