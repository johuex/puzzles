select a.activity_date as day, count(distinct a.user_id) as active_users
from activity a
WHERE activity_date BETWEEN DATE '2019-06-28' AND DATE '2019-07-27'
group by a.activity_date;