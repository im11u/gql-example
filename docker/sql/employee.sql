CREATE TABLE employee (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    department_id INT NOT NULL
);

INSERT INTO employee VALUES(null, '従業員01', 1);
INSERT INTO employee VALUES(null, '従業員02', 2);
INSERT INTO employee VALUES(null, '従業員03', 1);
