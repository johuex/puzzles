
class Stack:
    def __init__(self):
        self.content = []

    def insert(self, value: str):
        self.content.append(value)

    def pop(self) -> str:
        res = self.content[-1]
        self.content.pop(-1)
        return res

    def check_len(self) -> int:
        return len(self.content)


class IncorrectRightParenthesisException(Exception):
    def __init__(self, element: str, message: str = "Incorrect Right Parenthesis: {}"):
        self.message = message.format(element)
        super().__init__(self.message)


LEFT_PARENTHESIS_TYPE = ["(", "[", "{"]
RIGHT_PARENTHESIS_TYPE = [")", "]", "}"]
RIGHT_TO_LEFT_TYPE = {
    ")": "(",
    "}": "{",
    "]": "["
}


def solution(in_data: str) -> bool:
    """
    Main func
    """
    stack = Stack()
    for now_elem in in_data:
        if now_elem in LEFT_PARENTHESIS_TYPE:
            stack.insert(now_elem)
        elif now_elem in RIGHT_PARENTHESIS_TYPE:
            poped_elem = stack.pop()
            if poped_elem != RIGHT_TO_LEFT_TYPE[now_elem]:
                raise IncorrectRightParenthesisException(poped_elem)

    if stack.check_len() == 0:
        return True
