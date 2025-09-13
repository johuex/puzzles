-- first own solution
select
    'Low Salary' as "category",
    count(a1.account_id) as "accounts_count"
from Accounts a1
where income < 20000

union

select
    'Average Salary' as "category",
    count(a2.account_id) as "accounts_count"
from Accounts a2
where income >= 20000 and income <= 50000

union

select
    'High Salary' as "category",
    count(a3.account_id) as "accounts_count"
from Accounts a3
where income > 50000;

-- second solution (the best)
WITH Categorized AS (
  SELECT
    CASE
      WHEN income < 20000 THEN 'Low Salary'
      WHEN income > 50000 THEN 'High Salary'
      ELSE 'Average Salary'
    END AS category
  FROM Accounts
),
AllCategories AS (
  SELECT 'Low Salary' AS category
  UNION ALL
  SELECT 'Average Salary'
  UNION ALL
  SELECT 'High Salary'
)

SELECT 
  ac.category,
  COALESCE(COUNT(c.category), 0) AS accounts_count
FROM AllCategories ac
LEFT JOIN Categorized c ON ac.category = c.category
GROUP BY ac.category