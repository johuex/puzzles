-- my own solution
select d.name as Department, ae.name as Employee, ae.salary as Salary
from (
    select *, dense_rank() over (partition by e.departmentId order by e.salary DESC)
    from Employee e
) as ae
join Department d on d.id = ae.departmentId
where dense_rank <= 3;