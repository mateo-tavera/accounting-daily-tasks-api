CREATE TABLE IF NOT EXISTS tasks (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user VARCHAR(255),
    summary VARCHAR(2500),
    date DATETIME,
    status VARCHAR(255)
);