-- my solution
select coalesce(a.employee_id,b.employee_id) employee_id
from Employees a
full outer join Salaries b
on a.employee_id = b.employee_id
where b.salary is null or a.name is null