-- launch until receive 0 rows updated (manual or with some program language: python, go and etc.)
begin;

update table
set status = 'CANCELED'
where id in (
    select id
    from table
    where created_at between '2020-03-01' and '2020-04-01' and status = 'APPROVED'
    limit 10000 -- can be another
)
returning *;

commit;


-- another solution for if we are very worried via select and temp table
begin;

create temp table to_fix as ( -- create table will be deleted after session end
    select id
    from table
    where created_at between '2020-03-01' and '2020-04-01' and status = 'APPROVED'
    limit 10000
)

select * from to_fix; -- check result

update table
set status = 'CANCELED'
where id in (select id from to_fix);

commit;
