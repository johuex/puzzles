select sp1.name
from SalesPerson sp1
where sp1.sales_id not in (
    -- select all salesperson that have orders with 'RED'
    select sp.sales_id
    from SalesPerson sp
    join Orders o on o.sales_id = sp.sales_id
    join Company c on o.com_id = c.com_id
    where c.name = 'RED'
);

-- other solution
select s.name 
from salesperson s
where sales_id not in (
    select sales_id
    from orders
    where com_id in (
        select com_id from company
        where name =  'RED'
    )
);