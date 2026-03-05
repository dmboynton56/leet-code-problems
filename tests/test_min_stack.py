from solutions.min_stack import MinStack

def test_min_stack_flow():
    min_stack = MinStack()
    min_stack.push(-2)
    min_stack.push(0)
    min_stack.push(-3)
    assert min_stack.getMin() == -3
    min_stack.pop()
    assert min_stack.top() == 0
    assert min_stack.getMin() == -2

def test_min_stack_all_same():
    min_stack = MinStack()
    min_stack.push(1)
    min_stack.push(1)
    min_stack.push(1)
    assert min_stack.getMin() == 1
    min_stack.pop()
    assert min_stack.getMin() == 1

def test_min_stack_increasing():
    min_stack = MinStack()
    min_stack.push(1)
    min_stack.push(2)
    min_stack.push(3)
    assert min_stack.getMin() == 1
    min_stack.pop()
    assert min_stack.getMin() == 1
