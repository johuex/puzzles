-- my solution
SELECT p1.product_id, 'store1' "store", p1.store1 "price"
FROM Products p1
WHERE p1.store1 is not null

UNION ALL
SELECT p2.product_id, 'store2' "store",p2.store2 "price"
FROM Products p2
WHERE p2.store2 is not null

UNION ALL
SELECT p3.product_id, 'store3' "store",p3.store3 "price"
FROM Products p3
WHERE p3.store3 is not null;



-- another solution (same planning)
SELECT product_id, store, price
FROM (
    SELECT 
        product_id,
        'store1' AS store, store1 AS price
    FROM Products
    UNION ALL
    SELECT 
        product_id,
        'store2' AS store, store2 AS price
    FROM Products
    UNION ALL
    SELECT 
        product_id,
        'store3' AS store, store3 AS price
    FROM Products
) AS t
WHERE price IS NOT NULL;