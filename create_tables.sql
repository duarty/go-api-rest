CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE,
  username VARCHAR(80) NOT NULL UNIQUE,
  phone VARCHAR(20) NOT NULL UNIQUE,
  age INTEGER NOT NULL
);

CREATE TABLE gyms (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  address VARCHAR(255),
  placeID VARCHAR(50) UNIQUE,
  longitude REAL,
  latitude REAL
);

CREATE TABLE user_gyms (
  user_id INTEGER NOT NULL,
  gym_id INTEGER NOT NULL,
  PRIMARY KEY (user_id, gym_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (gym_id) REFERENCES gyms(id)
);
