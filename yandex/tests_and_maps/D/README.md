# Second solution (on python)

```
n, k = map(int, input().split())
problems = list(map(int, input.split()))
uniq_problems = set(problems)
tour = list(uniq_problems)
if len(tour) < k:
    for problem in problems:
        if problem not in uniq_problems:
            tour.append(problem)
        else:
            uniq_problems.descard(problem)
print(*tour[:k]) # отбрасываем лишнее
```