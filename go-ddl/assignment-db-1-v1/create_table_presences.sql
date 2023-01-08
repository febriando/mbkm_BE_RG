-- TODO: answer here
CREATE TABLE IF NOT EXISTS presences (
      id INT NOT NULL,
      user_id INT NOT NULL,
      presence_date DATE NOT NULL,
      status VARCHAR(50) NOT NULL,
      location VARCHAR(50) NULL,
      description VARCHAR(255) NULL,
      image_presence VARCHAR(255) NULL,
      image_location VARCHAR(255) NULL
);