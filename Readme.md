## Golang BAQ Heroku-Circle

This is a simple app to show TDD workflow with Heroku and CircleCI.

### Setting up the Database

Run the following to setup the database:

```sql
CREATE DATABASE gophers;
CREATE DATABASE gophers_test;

GRANT ALL PRIVILEGES ON DATABASE gophers TO gophers;
GRANT ALL PRIVILEGES ON DATABASE gophers_test TO gophers;
```
