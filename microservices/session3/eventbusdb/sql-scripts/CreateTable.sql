ALTER USER 'root'@'%' IDENTIFIED WITH 'mysql_native_password' BY 'hello';
FLUSH PRIVILEGES;
CREATE TABLE events (
    id int NOT NULL AUTO_INCREMENT,
    type varchar(255) NOT NULL,
    data varchar(255),
    PRIMARY KEY (id)
);