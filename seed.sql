DROP TABLE IF EXISTS temperatures;
DROP TABLE IF EXISTS thresholds;
DROP TABLE IF EXISTS threshold_alerts;
DROP TABLE IF EXISTS sensors;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id SERIAL NOT NULL,
  user_name VARCHAR(255) NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE sensors (
  id SERIAL NOT NULL,
  user_id INT references users(id),
  sensor_name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE(user_id, sensor_name)
);

CREATE TABLE temperatures (
  id SERIAL NOT NULL,
  temperature NUMERIC(8, 4) NOT NULL,
  sensor_id INT references sensors(id),
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);

CREATE INDEX sensor_time ON temperatures(sensor_id, created_at);

CREATE TABLE thresholds (
  temperature NUMERIC(8, 4) NOT NULL,
  sensor_id INT UNIQUE references sensors(id),
  PRIMARY KEY (sensor_id)
);

CREATE TABLE threshold_alerts (
  id SERIAL NOT NULL,
  temperature NUMERIC(8, 4) NOT NULL,
  sensor_id INT references sensors(id),
  threshold NUMERIC(8, 4) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO users (user_name, hashed_password) VALUES
('rlopez', '$2a$10$arSaVKo66cOIyRTNolSP6ubLnRkKAZaiQJpttCXUzOLpIWoNAGGLG'),
('bbryski', '$2a$10$arSaVKo66cOIyRTNolSP6ubLnRkKAZaiQJpttCXUzOLpIWoNAGGLG'),
('dburns', '$2a$10$arSaVKo66cOIyRTNolSP6ubLnRkKAZaiQJpttCXUzOLpIWoNAGGLG');

INSERT INTO sensors (sensor_name, user_id) VALUES
('tennis court', 1),
('office', 1),
('throne room', 2),
('garden', 2),
('dungeon', 3);

-- Creating seed data
CREATE OR REPLACE FUNCTION generate_temp_seed_data(sensor_id INT, start_temp NUMERIC(8, 4), num_days TEXT, time_interval TEXT) RETURNS void
AS
$$
DECLARE
  curr_time TIMESTAMP;
  end_time TIMESTAMP DEFAULT NOW()::timestamp;
  temp_data NUMERIC(8, 4)[] DEFAULT ARRAY[]::NUMERIC(8,4)[];
  time_data TIMESTAMP[] DEFAULT ARRAY[]::TIMESTAMP[];
  sensor_ids INT[] DEFAULT ARRAY[]::INT[];
  decrease INT[] := ARRAY [65, 65, 65, 55, 45, 35, 25, 15, 10, 05, 05, 05, 05, 05, 05, 15, 25, 35, 45, 55, 60, 65, 65, 65];
  stay INT[] := ARRAY [95, 95, 95, 85, 75, 65, 55, 45, 40, 35, 35, 35, 35, 35, 35, 45, 55, 65, 75, 85, 90, 95, 95, 95];
  hour INT;
  temp ALIAS FOR start_temp;
  rand_perc INT;
  rand_delta NUMERIC(8, 4);
  loops INT DEFAULT 0;
BEGIN
  curr_time = end_time - num_days::INTERVAL;
  temp_data = array_append(temp_data, temp);
  time_data = array_append(time_data, curr_time);
  sensor_ids = array_append(sensor_ids, sensor_id);
  curr_time = curr_time + time_interval::INTERVAL;
  WHILE curr_time < end_time LOOP
    time_data = array_append(time_data, curr_time);
    sensor_ids = array_append(sensor_ids, sensor_id);

    hour = EXTRACT(HOUR FROM curr_time);
    rand_perc = floor(random() * 100 + 1)::INT;
    rand_delta = (random() / 10)::NUMERIC(8, 4);

    IF decrease[hour] >= rand_perc THEN
      temp = temp - rand_delta;
    ELSIF stay[hour] < rand_perc THEN
      temp = temp + rand_delta;
    END IF;

    temp_data = array_append(temp_data, temp);

    curr_time = curr_time + time_interval::INTERVAL;
    loops = loops + 1;
    IF loops % 5000 = 0 THEN
      PERFORM insert_temp_seed_data(temp_data, time_data, sensor_ids);
      temp_data = ARRAY[]::NUMERIC(8,4)[];
      time_data = ARRAY[]::TIMESTAMP[];
      sensor_ids = ARRAY[]::INT[];
    END IF;
  END LOOP;
  PERFORM insert_temp_seed_data(temp_data, time_data, sensor_ids);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION insert_temp_seed_data(temps NUMERIC(8, 4)[], t_times TIMESTAMP[], ids INT[]) RETURNS void
AS
$$
  INSERT INTO temperatures (temperature, created_at, sensor_id)
  SELECT * FROM unnest(temps, t_times, ids);
$$
LANGUAGE sql STRICT;

SELECT generate_temp_seed_data(1, 50.0, '90 days', '1 minute');
SELECT generate_temp_seed_data(2, 70.0, '90 days', '1 minute');
SELECT generate_temp_seed_data(3, 90.0, '90 days', '1 minute');
SELECT generate_temp_seed_data(4, 60.0, '90 days', '1 minute');
SELECT generate_temp_seed_data(5, 40.0, '90 days', '1 minute');