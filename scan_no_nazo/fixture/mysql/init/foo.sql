CREATE TABLE `students` (
  id INT PRIMARY KEY,
  active TINYINT NULL,
  name VARCHAR(255) NULL,
  grade INT NULL,
  score DOUBLE NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `students` (id, active, name, grade, score) VALUES
  (1, 1, 'John Doe', 65535, 0.009876),
  (2, NULL, NULL, NULL, NULL),
  (3, NULL, 'Akira Toriyama', NULL, 100)
;
