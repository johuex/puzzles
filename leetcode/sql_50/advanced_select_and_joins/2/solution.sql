-- my own solution (#2 speed place)
-- employee with multiple deparments
select e1.employee_id, e1.department_id
from Employee e1
where primary_flag = 'Y'
union
-- employee with one department
select e3.employee_id, e3.department_id
from Employee e3
where employee_id in (
    select e2.employee_id
    from Employee e2
    group by e2.employee_id
    having count(e2.employee_id) = 1
);

-- another solution from top (#3 speed place)
with employee_total as (
    select
        employee_id,
        department_id,
        primary_flag,
        count(*) over (partition by employee_id) as cnt
    from employee
)

select
    employee_id,
    department_id
from employee_total
where cnt = 1 or primary_flag ='Y'

-- another solution from top (#1 speed place)
select employee_id,
       department_id
from Employee 
where primary_flag = 'Y'
   or employee_id in (select employee_id from Employee group by employee_id having count(employee_id)=1) 