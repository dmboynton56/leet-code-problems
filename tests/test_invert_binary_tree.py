from solutions.invert_binary_tree import Solution, TreeNode

def test_example_2():
    # Input: [2, 1, 3] -> Output: [2, 3, 1]
    solver = Solution()
    root = TreeNode(2)
    root.left = TreeNode(1)
    root.right = TreeNode(3)
    
    inverted = solver.invertTree(root)
    assert inverted.val == 2
    assert inverted.left.val == 3
    assert inverted.right.val == 1

def test_empty():
    solver = Solution()
    assert solver.invertTree(None) is None

def test_single_node():
    solver = Solution()
    root = TreeNode(1)
    inverted = solver.invertTree(root)
    assert inverted.val == 1
    assert inverted.left is None
    assert inverted.right is None
