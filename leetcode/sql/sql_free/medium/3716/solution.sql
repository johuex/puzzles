-- my soultion
with ranked_events_user as (
    -- rank event grouped by user_id (last event have 1 rank)
    select *, dense_rank() over (partition by se.user_id order by se.event_date DESC) "rank"
    from subscription_events se
), risked_user_id as (
    select
        reu.user_id,
        min(reu.plan_name) filter (where reu.rank = 1) "current_plan",
        min(reu.monthly_amount) filter (where reu.rank = 1) "current_monthly_amount",
        max(reu.monthly_amount) "max_historical_amount",
        max(reu.event_date)::date - min(reu.event_date)::date "days_as_subscriber"
    from ranked_events_user reu
    -- where reu.rank = 1 and reu.event_type <> 'cancel'
    group by reu.user_id
    having
        count(reu.event_id) filter (where reu.event_type = 'downgrade') > 0
        and max(reu.monthly_amount) > 2 * (min(reu.monthly_amount) filter (where reu.rank = 1))
        and max(reu.event_date)::date - min(reu.event_date)::date >= 60
        and (min(reu.monthly_amount) filter (where reu.rank = 1)) <> 0.00 -- actual status is not cancelled
)

select *
from risked_user_id
order by 5 DESC, 1 ASC;


-- more clear by syntax solution by deepseek
WITH user_metrics AS (
    SELECT 
        user_id,
        FIRST_VALUE(event_type) OVER w AS last_event_type,
        FIRST_VALUE(monthly_amount) OVER w AS current_amount,
        FIRST_VALUE(plan_name) OVER w AS current_plan,
        LAST_VALUE(event_date) OVER w AS first_date,
        MAX(monthly_amount) OVER w AS max_amount,
        COUNT(*) FILTER (WHERE event_type = 'downgrade') OVER w AS downgrade_count
    FROM subscription_events
    WINDOW w AS (PARTITION BY user_id ORDER BY event_date 
                 ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING)
)
SELECT DISTINCT
    user_id,
    current_plan,
    current_amount,
    max_amount,
    last_date - first_date AS days_as_subscriber
FROM user_metrics
WHERE 
    last_event_type != 'cancel'
    AND downgrade_count > 0
    AND current_amount * 2 < max_amount
    AND last_date - first_date >= 60
ORDER BY days_as_subscriber DESC, user_id ASC;