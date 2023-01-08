-- TODO: answer here
select 
    reports.id as id,
    students.fullname as fullname,
    students.class as class,
    students.status as status,
    reports.study as study,
    reports.score as score
from reports
right join students
on reports.student_id = students.id
where score < 70 and status = 'active'
order by score asc;