from solutions.maximum_depth_of_binary_tree import Solution, TreeNode

def test_example_1():
    solver = Solution()
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)
    assert solver.maxDepth(root) == 3

def test_example_2():
    solver = Solution()
    root = TreeNode(1)
    root.right = TreeNode(2)
    assert solver.maxDepth(root) == 2

def test_empty():
    solver = Solution()
    assert solver.maxDepth(None) == 0

def test_single_node():
    solver = Solution()
    root = TreeNode(1)
    assert solver.maxDepth(root) == 1
