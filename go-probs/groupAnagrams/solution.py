from collections import defaultdict


# solution groups strings that are anagrams of each other.
# Go: map[string][]string with sorted runes as key; Python uses sorted tuple or char count key.
def solution(strs: list[str]) -> list[list[str]]:
    # defaultdict auto-inits empty lists; Go map requires append on missing key
    # (append to nil slice works in Go).
    groups: dict[tuple[str, ...], list[str]] = defaultdict(list)

    for s in strs:
        key = tuple(sorted(s))  # Go: anagramKey(s) → sort runes, return string(runes)
        groups[key].append(s)  # append creates slice if key new

    return list(groups.values())  # Go: explicit loop building result slice from map values


if __name__ == "__main__":
    print(solution(["eat", "tea", "tan", "ate", "nat", "bat"]))
