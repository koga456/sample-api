CREATE USER 'todo-app'@'%' IDENTIFIED BY 'todo-password';
GRANT SELECT,INSERT,UPDATE,DELETE ON todo.* TO 'todo-app'@'%';