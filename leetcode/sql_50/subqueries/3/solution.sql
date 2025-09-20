-- my own solution
(select u.name as results
from Users u
inner join (
    select mv.user_id, count(mv.user_id) "review_count"
    from MovieRating mv
    group by mv.user_id
) mvc on u.user_id = mvc.user_id
order by mvc."review_count" DESC, u.name ASC
limit 1)

union all

(
    select m.title as results
from Movies m
inner join (
     select mv2.movie_id, avg(mv2.rating) as "movie_rating"
     from MovieRating mv2
     where mv2.created_at BETWEEN '2020-02-01' AND '2020-02-29'
     group by mv2.movie_id
) as mvs on m.movie_id = mvs.movie_id
order by mvs."movie_rating" DESC, m.title ASC
limit 1
);

-- another solution: w/o subqueries
(
  SELECT u.name AS results
  FROM movierating AS mr
  JOIN users AS u ON mr.user_id = u.user_id
  GROUP BY u.name
  ORDER BY COUNT(*) DESC, u.name
  LIMIT 1
)
UNION ALL
(
  SELECT m.title AS results
  FROM movierating AS mr
  JOIN movies AS m ON mr.movie_id = m.movie_id
  WHERE mr.created_at BETWEEN '2020-02-01' AND '2020-02-29'
  GROUP BY m.title
  ORDER BY AVG(mr.rating) DESC, m.title
  LIMIT 1
);