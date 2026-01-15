-- my solution
select s.stock_name, sum(
    case
        when s.operation = 'Buy' then -s.price
        when s.operation = 'Sell' then s.price
    end
) "capital_gain_loss"
from Stocks s
group by s.stock_name;
