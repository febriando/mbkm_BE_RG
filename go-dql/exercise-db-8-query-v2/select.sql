-- TODO: answer here
select id, NIK, concat (first_name, ' ', last_name) as fullname, date_of_birth, weight, address 
from people 
where gender = 'laki-laki' order by weight DESC
FETCH FIRST 5 ROW ONLY;