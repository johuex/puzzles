-- my solution
-- Write your PostgreSQL query statement below
select 
    d.id,
    sum (CASE WHEN month = 'Jan' THEN revenue else NULL END) "Jan_Revenue",
    sum (CASE WHEN month = 'Feb' THEN revenue else NULL END) "Feb_Revenue",
    sum (CASE WHEN month = 'Mar' THEN revenue else NULL END) "Mar_Revenue",
    sum (CASE WHEN month = 'Apr' THEN revenue else NULL END) "Apr_Revenue",
    sum (CASE WHEN month = 'May' THEN revenue else NULL END) "May_Revenue",
    sum (CASE WHEN month = 'Jun' THEN revenue else NULL END) "Jun_Revenue",
    sum (CASE WHEN month = 'Jul' THEN revenue else NULL END) "Jul_Revenue",
    sum (CASE WHEN month = 'Aug' THEN revenue else NULL END) "Aug_Revenue",
    sum (CASE WHEN month = 'Sep' THEN revenue else NULL END) "Sep_Revenue",
    sum (CASE WHEN month = 'Oct' THEN revenue else NULL END) "Oct_Revenue",
    sum (CASE WHEN month = 'Nov' THEN revenue else NULL END) "Nov_Revenue",
    sum (CASE WHEN month = 'Dec' THEN revenue else NULL END) "Dec_Revenue"
from Department d
group by d.id;