-- first own solution
select e.employee_id
from Employees e
where salary < 30000 and manager_id not in (
    select e2.employee_id
    from Employees e2
)
order by employee_id;

-- second own solution
select e1.employee_id
from Employees e1
left join Employees e2 on e1.manager_id = e2.employee_id
where e1.salary < 30000 and e1.manager_id is not null and e2.employee_id is null
order by 1;