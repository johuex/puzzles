-- my solution
select a.player_id, min(a.event_date) "first_login"
from Activity a
group by a.player_id;