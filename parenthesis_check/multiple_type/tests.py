import pytest
from parenthesis_check.multiple_type.solution import solution, IncorrectRightParenthesisException

VALID_TEST_DATA = ["([{}])", "([({})])"]
INVALID_TEST_DATA = ["{[}]", "(({)"]


class TestMultipleParenthesisCheck:
    @pytest.mark.parametrize("in_data", VALID_TEST_DATA)
    def test_multiple_type_parenthesis_return_ok(self, in_data):
        assert solution(in_data)

    @pytest.mark.parametrize("in_data", INVALID_TEST_DATA)
    def test_multiple_type_parenthesis_return_error(self, in_data):
        with pytest.raises(IncorrectRightParenthesisException):
            solution(in_data)
