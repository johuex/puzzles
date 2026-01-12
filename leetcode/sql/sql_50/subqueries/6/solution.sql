-- my solution
select round(sum(tiv_2016)::numeric,2) as tiv_2016
from Insurance i
where i.pid in (
    select i1.pid
    from Insurance i1
    cross join Insurance i2
    where i1.tiv_2015 = i2.tiv_2015 and i1.lat <> i2.lat and i1.lon <> i2.lon and i1.pid <> i2.pid
    group by i1.pid
    except
    select i11.pid
    from Insurance i11
    cross join Insurance i21
    where i11.lat = i21.lat and i11.lon = i21.lon and i11.pid <> i21.pid
    group by i11.pid    
    );

-- another my solution (after shock by previous solution plan)
select round(sum(tiv_2016)::numeric,2) as tiv_2016
from Insurance i
where (i.lat, i.lon) not in (
    select lat,lon from Insurance group by lat, lon having count(*) > 1  
) and i.pid in (
    select i1.pid
    from Insurance i1
    cross join Insurance i2
    where i1.tiv_2015 = i2.tiv_2015 and i1.lat <> i2.lat and i1.lon <> i2.lon and i1.pid <> i2.pid
    group by i1.pid
);

-- another solution
select round(sum(tiv_2016)::numeric,2) as tiv_2016
from insurance 
where tiv_2015 in(
    select tiv_2015
    from insurance 
    group by 1
    having count(1)>1
) and (lat, lon) in (
    select lat, lon from insurance
    group by lat, lon
    having count(1)=1
);
