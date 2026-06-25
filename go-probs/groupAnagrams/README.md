# Group Anagrams (Medium) — LeetCode #49

**Difficulty:** Medium  
**Pattern:** Hash map / grouping key

## Data Structure

- **Hash map** — canonical anagram key → list of original strings.
- **Key options:**
  - Sorted characters (e.g. `"aet"` for `"eat"`)
  - Char frequency count (26-letter array or map) — O(k) key for length-k strings

## Algorithm

1. For each string, compute canonical key.
2. Append string to `groups[key]`.
3. Return all map values as `[][]string`.

**Time:** O(n · k log k) with sort keys; O(n · k) with frequency vector keys  
**Space:** O(n · k) for stored strings

## Edge Cases

| Case | Notes |
|------|-------|
| Single string | One group |
| Empty string | Valid; key is empty |
| All same anagram | One group |
| No anagram pairs | Each string own group |
| Unicode | Sort on runes, not bytes |

## Go-Specific Notes

- **`append` on nil slice** — `groups[key] = append(groups[key], s)` works when key missing (nil slice).
- **No defaultdict** — manual map + append is idiomatic; or check `if _, ok := groups[key]`.
- **Key as string** — sort `[]rune(s)` then `string(runes)`; sorting bytes breaks multi-byte UTF-8.
- **Map iteration order** — random in Go; LeetCode doesn't care about group order, only contents.
- **Frequency key** — `[26]int` array converted to string with fmt or custom encoding — faster for long strings.

## Other Notes

- Interview tradeoff: sorted key is simpler; count key avoids O(k log k) per string.
- Related: Valid Anagram is the pairwise check.
