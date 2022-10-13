## Project Sample for Golang Unit Testing

### Architecture

1. PostgreSQL
2. Gin Framework

### Database

```sql
CREATE DATABASE db_sample_go;

CREATE TABLE customer (
    id varchar(60) primary key,
    name varchar(50) not null,
    address text
);
```

### Make sure this project able to run

#### Windows

```
SET DB_HOST=localhost
SET DB_PORT=5432
SET DB_NAME=db_sample_go
SET DB_USER=yourdbuser
SET DB_PASSWORD=yourdbpassword
SET API_PORT=8888
SET API_HOST=localhost
go run .
```

#### Linux / Mac

```
DB_HOST=localhost DB_PORT=5432 DB_NAME=db_sample_go DB_USER=yourdbuser DB_PASSWORD=yourdbpassword API_PORT=8888 API_HOST=localhost go run .
```

### What will be tested

1. Repository
2. Usecase
3. Controller
