-- my solution
select e.employee_id, 
    case
        when e.employee_id % 2 <> 0 and left(e.name, 1) <> 'M' then e.salary
        else 0
    end as bonus
from Employees e
order by 1;

-- another solution, it's faster
select e.employee_id, 
    case
        -- because direct comparasion and like is faster that left function
        when e.employee_id % 2 = 1 and name not like 'M%' then e.salary
        else 0
    end as bonus
from Employees e
order by 1;