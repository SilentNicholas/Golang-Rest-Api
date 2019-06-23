################ Project Review ###############
In this project you can see 4 directories ad 1 main file, which include and use all directories in itself.

#####
1. Controllers. This folder contains files that are responsible for for project logic.
#####
2. Driver. This folder contains files that are responsible for connection with database.
#####
3. Models. This folder contains files that are responsible for code executable functions.
#####
4. Repository. This folder contains files that are responsible for the queries to the database.
#####
5. .env file contain your URL of postgres remote host on https://api.elephantsql.com .
#####
6. main.go is an executable file of this REST Api application.

############### How to use ################

If you want to use this application you can know next:
#####
1. When you run this code your host will be :8000 (for example http://localhost:8000).
#####
2. You have to create account on https://api.elephantsql.com and write your URL in .env file.
#####
3. If you want to see all items of database table you can write next url path /books (for example http://localhost:8000/books)
for GET method.
#####
4. If you want to see one item of database table you can write next url path /books/{id} (for example http://localhost:8000/books/3)
for GET method.
#####
5. If you want to add new item for database table you can write next url path /books (for example http://localhost:8000/books)
for POST method with json row.
#####
6. If you want to update item of table you can write next url path /books (for example http://localhost:8000/books)
for PUT method with json row.
#####
7. If you want to remove item of table you can write next url path /books/{id} (for example http://localhost:8000/books/3)
for DELETE method.

############### Good Luck With Golang ################