import pytest
from parenthesis_check.solution import solution

MULTIPLE_VALID_DATA = ["([{}])", "([({})])"]
MULTIPLE_INVALID_DATA = ["{[}]", "(({)"]
MULTIPLE_PARENTHESIS_TYPE = {
    ")": "(",
    "}": "{",
    "]": "["
}
MULTIPLE_OK_TEST_PARAMS = [(i, MULTIPLE_PARENTHESIS_TYPE) for i in MULTIPLE_VALID_DATA]
MULTIPLE_ERROR_TEST_PARAMS = [(i, MULTIPLE_PARENTHESIS_TYPE) for i in MULTIPLE_INVALID_DATA]

SINGLE_VALID_DATA = ["(()())", "((()))"]
SINGLE_INVALID_DATA = ["((())", "())", ")("]
SINGLE_PARENTHESIS_TYPE = {
    ")": "(",
    "}": "{",
    "]": "["
}
SINGLE_OK_TEST_PARAMS = [(i, SINGLE_PARENTHESIS_TYPE) for i in SINGLE_VALID_DATA]
SINGLE_ERROR_TEST_PARAMS = [(i, SINGLE_PARENTHESIS_TYPE) for i in SINGLE_INVALID_DATA]


class TestParenthesisCheck:
    @pytest.mark.parametrize("in_data, parenthesis_type", MULTIPLE_OK_TEST_PARAMS)
    def test_multiple_type_parenthesis_return_ok(self, in_data, parenthesis_type):
        assert solution(in_data, parenthesis_type)

    @pytest.mark.parametrize("in_data, parenthesis_type", MULTIPLE_ERROR_TEST_PARAMS)
    def test_multiple_type_parenthesis_return_error(self, in_data, parenthesis_type):
        assert not solution(in_data, parenthesis_type)

    @pytest.mark.parametrize("in_data, parenthesis_type", SINGLE_OK_TEST_PARAMS)
    def test_single_type_parenthesis_return_ok(self, in_data, parenthesis_type):
        assert solution(in_data, parenthesis_type)

    @pytest.mark.parametrize("in_data, parenthesis_type", SINGLE_ERROR_TEST_PARAMS)
    def test_single_type_parenthesis_return_error(self, in_data, parenthesis_type):
        assert not solution(in_data, parenthesis_type)
