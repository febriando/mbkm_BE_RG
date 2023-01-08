select 
	id,
	concat(first_name || ' ' || last_name) as fullname, 
	SPLIT_PART(exam_id, '-', 1) as class, 
	(select coalesce(bahasa_indonesia+bahasa_inggris+matematika+ipa)/4) as avg_score
from final_scores
where exam_status = 'pass' and (fee_status = 'full' or fee_status = 'installment')
order by avg_score desc
limit 5;