-- my solution
select ds.date_id, ds.make_name, count(distinct ds.lead_id) "unique_leads", count(distinct ds.partner_id)"unique_partners"
from DailySales ds
group by 1,2