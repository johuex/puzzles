-- my solution
select p.user_id, count(p.user_id) "prompt_count", round(avg(p.tokens), 2) "avg_tokens"
from prompts p
group by p.user_id
having count(p.user_id) >= 3 and max(p.tokens) > avg(p.tokens)
order by 3 DESC, 1 ASC;