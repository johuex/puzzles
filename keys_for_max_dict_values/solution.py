from typing import Dict, List


def solution(d: Dict[str, int]) -> List[str]:
    key_sorted = sorted(d, key=d.get)  # sorted keys by values
    return key_sorted[-2::][::-1]
