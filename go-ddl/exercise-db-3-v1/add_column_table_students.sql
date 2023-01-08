-- TODO: answer here
ALTER TABLE students ADD COLUMN street VARCHAR(255);
ALTER TABLE students ADD COLUMN city varchar(100);
ALTER TABLE students ADD COLUMN province varchar(100);
ALTER TABLE students ADD COLUMN country varchar(100);
ALTER TABLE students ADD COLUMN postal_code VARCHAR(50);

ALTER TABLE students ADD COLUMN date_of_birth DATE NOT NULL;