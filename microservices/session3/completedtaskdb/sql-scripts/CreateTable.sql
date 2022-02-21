ALTER USER 'root'@'%' IDENTIFIED WITH 'mysql_native_password' BY 'hello';
FLUSH PRIVILEGES;
CREATE TABLE completedtasks (
    id varchar(36),
    task varchar(255) NOT NULL,
    PRIMARY KEY (id)
);