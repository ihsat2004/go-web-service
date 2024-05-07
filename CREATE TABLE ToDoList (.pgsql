CREATE TABLE ToDoList (
    SINo int NOT NULL,
    Title varchar(255) NOT NULL,
    Description varchar(255) NOT NULL,
    Complete BOOLEAN NOT NULL,
    PRIMARY KEY (Title)
)