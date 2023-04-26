from typing import Dict, List
import pytest
from keys_for_max_dict_values.solution import solution


params = [({
    'a': 100,
    'b': 45,
    'c': 39,
    "d": 2000,
    },
    ['d', 'a'],)
]


@pytest.mark.parametrize("in_data, answer", params)
def test_max_values_dict_keys_return_ok(in_data: Dict[str, int], answer: List[str]) -> None:
    res = solution(in_data) 
    assert res[0] == answer[0] and res[1] == answer[1]
    