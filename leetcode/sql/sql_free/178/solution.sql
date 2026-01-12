-- my own solution
select s.score, dense_rank() OVER (ORDER BY s.score desc) as rank
from Scores s
order by s.score desc

-- solution from LeetCode 
SELECT
    score,
    DENSE_RANK() OVER(ORDER BY score DESC) AS rank
FROM Scores;
