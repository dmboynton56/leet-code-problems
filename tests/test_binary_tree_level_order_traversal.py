from solutions.binary_tree_level_order_traversal import Solution, TreeNode

def test_example_1():
    solver = Solution()
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)
    assert solver.levelOrder(root) == [[3], [9, 20], [15, 7]]

def test_example_2():
    solver = Solution()
    root = TreeNode(1)
    assert solver.levelOrder(root) == [[1]]

def test_empty():
    solver = Solution()
    assert solver.levelOrder(None) == []
