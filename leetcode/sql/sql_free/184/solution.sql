with cte as (
    select *, rank() over (partition by e.departmentId order by salary desc)
    from Employee e
)


select d.name "Department", c.name "Employee", c.salary "Salary"
from cte c
join Department d on c.departmentId = d.id
where c.rank = 1;

