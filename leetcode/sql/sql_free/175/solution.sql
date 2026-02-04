select p.firstName "firstName", p.lastName "lastName", a.city, a.state
from Person p
left join Address a on p.personId = a.personId;