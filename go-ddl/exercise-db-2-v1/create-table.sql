CREATE TABLE IF NOT EXISTS persons (
      id INT primary key,
      NIK VARCHAR(255) NOT NULL unique,
      fullname VARCHAR(255) NOT NULL,
      gender VARCHAR(50) NOT NULL,
      birth_date DATE NOT NULL,
      is_married BOOLEAN,
      height FLOAT,
      weight FLOAT,
      address TEXT
);