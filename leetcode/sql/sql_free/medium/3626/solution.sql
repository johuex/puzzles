-- my solution
WITH ranked_products AS (
    -- range items, window == store
    SELECT 
        i.store_id,
        s.store_name,
        s.location,
        i.product_name,
        i.quantity,
        i.price,
        DENSE_RANK() OVER (PARTITION BY i.store_id ORDER BY i.price ASC) AS cheap_rank, -- rank in asc order
        DENSE_RANK() OVER (PARTITION BY i.store_id ORDER BY i.price DESC) AS expensive_rank, -- rank in decs order
        COUNT(*) OVER (PARTITION BY i.store_id) AS product_count -- uniq items' name in store
    FROM inventory i
    JOIN stores s ON i.store_id = s.store_id
),
store_products AS (
    -- select expensive and cheapest items info
    SELECT 
        store_id,
        store_name,
        location,
        product_count,
        MAX(CASE WHEN expensive_rank = 1 THEN product_name END) AS most_exp_product,
        MAX(CASE WHEN expensive_rank = 1 THEN quantity END) AS most_exp_quantity,
        MAX(CASE WHEN cheap_rank = 1 THEN product_name END) AS cheapest_product,
        MAX(CASE WHEN cheap_rank = 1 THEN quantity END) AS cheapest_quantity
    FROM ranked_products
    GROUP BY store_id, store_name, location, product_count
)

-- final conditions for store
SELECT 
    store_id,
    store_name,
    location,
    most_exp_product,
    cheapest_product,
    ROUND(cheapest_quantity::NUMERIC / most_exp_quantity, 2) AS imbalance_ratio
FROM store_products
WHERE 
    product_count >= 3
    AND most_exp_quantity < cheapest_quantity
ORDER BY 
    imbalance_ratio DESC,
    store_name ASC;


-- another faster solution
WITH MaxMinStore AS (
    SELECT store_id, MAX(price) AS max_price, MIN(price) AS min_price
    FROM Inventory
    GROUP BY store_id
    HAVING COUNT(1) > 2
)

SELECT s.store_id, s.store_name, s.location, i1.product_name AS most_exp_product,
    i2.product_name AS cheapest_product, ROUND(i2.quantity::DECIMAL / i1.quantity, 2) AS imbalance_ratio
FROM MaxMinStore mm JOIN Inventory i1 ON i1.store_id = mm.store_id AND i1.price = mm.max_price
JOIN Inventory i2 ON i2.store_id = mm.store_id AND i2.price = mm.min_price
JOIN Stores s ON s.store_id = mm.store_id
WHERE i1.quantity < i2.quantity
ORDER BY imbalance_ratio DESC, store_name ASC