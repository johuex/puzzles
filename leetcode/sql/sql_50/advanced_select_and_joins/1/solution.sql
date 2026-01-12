select
    e1.reports_to as employee_id,
    e2.name,
    count(e1.reports_to) as reports_count,
    round(avg(e1.age)) as average_age
from Employees e1
right join Employees e2 on e1.reports_to = e2.employee_id
where e1.employee_id is not null
group by e1.reports_to, e2.name
order by e1.reports_to;