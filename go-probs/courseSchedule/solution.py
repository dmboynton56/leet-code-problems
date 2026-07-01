from collections import deque


def solution(num_courses: int, prerequisites: list[list[int]]) -> bool:
    graph: list[list[int]] = [[] for _ in range(num_courses)]
    indegree = [0] * num_courses

    for course, prereq in prerequisites:
        graph[prereq].append(course)
        indegree[course] += 1

    queue = deque(i for i in range(num_courses) if indegree[i] == 0)
    processed = 0

    while queue:
        node = queue.popleft()
        processed += 1
        for nxt in graph[node]:
            indegree[nxt] -= 1
            if indegree[nxt] == 0:
                queue.append(nxt)

    return processed == num_courses


if __name__ == "__main__":
    print(solution(2, [[1, 0]]))          # True
    print(solution(2, [[1, 0], [0, 1]]))  # False
