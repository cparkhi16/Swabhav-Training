ALTER USER 'root'@'%' IDENTIFIED WITH 'mysql_native_password' BY 'hello';
FLUSH PRIVILEGES;
CREATE TABLE todolist (
    id varchar(36),
    task varchar(255) NOT NULL,
    userid varchar(36),
    PRIMARY KEY (id)
);