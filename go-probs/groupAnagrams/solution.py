from collections import defaultdict


def solution(strs: list[str]) -> list[list[str]]:
    """
    defaultdict(list) auto-inits empty lists; Go map requires append on missing key
    (append to nil slice works in Go).
    """
    groups: dict[tuple[str, ...], list[str]] = defaultdict(list)

    for s in strs:
        key = tuple(sorted(s))
        groups[key].append(s)

    return list(groups.values())


if __name__ == "__main__":
    print(solution(["eat", "tea", "tan", "ate", "nat", "bat"]))
