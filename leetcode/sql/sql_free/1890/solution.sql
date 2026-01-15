-- my solution
select l.user_id, max(l.time_stamp) as last_stamp
from Logins l
where extract('year' from l.time_stamp) = 2020
group by l.user_id;

-- another my solution (more faster on ??leetcode??)
select l.user_id, max(l.time_stamp) as last_stamp
from Logins l
where time_stamp >= '2020-01-01' and time_stamp <  '2021-01-01'
group by l.user_id;