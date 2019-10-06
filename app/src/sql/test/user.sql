# // User is an account that owns Tasks.
# type User struct {
#     ID       int32  `json:"id"`
#     Email    string `json:"email"`
#     Password string `json:"password"`
# }

CREATE TABLE account(
   ID serial PRIMARY KEY,
   Email VARCHAR (355) UNIQUE NOT NULL,
   Password VARCHAR(1048)
);


INSERT INTO account(Email, Password)
VALUES
('test@cool.com', 'test123') RETURNING ID;

INSERT INTO account(Email, Password)
VALUES
('blue@fire.com', 'cool345') RETURNING ID;

