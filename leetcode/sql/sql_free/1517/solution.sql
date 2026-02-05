SELECT
    u.user_id,
    u.name,
    u.mail 
FROM users u
WHERE u.mail ~ '^[A-Za-z][A-Za-z0-9_.-]*@leetcode\.com$' 