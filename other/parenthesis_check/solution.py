
class Stack:
    def __init__(self):
        self.content = []

    def insert(self, value: str):
        self.content.append(value)

    def pop(self) -> str | bool:
        if self.check_len():
            return self.content.pop(-1)
        else:
            return False

    def check_len(self) -> int:
        return len(self.content)


def solution(in_data: str, parenthesis_types: dict) -> bool:
    """
    Main func
    """
    stack = Stack()
    for now_elem in in_data:
        if now_elem in parenthesis_types.values():
            stack.insert(now_elem)
        elif now_elem in parenthesis_types.keys():
            poped_elem = stack.pop()
            if not poped_elem or poped_elem != parenthesis_types[now_elem]:
                return False

    if stack.check_len() == 0:
        return True
