# Day 28

Inserting data to a relational database

```bash
export DBUSER=myusername
export DBPASS=mypassword
```

```bash
go run main.go
```

## Setting up the database

```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);
```

**Credit:** [Tutorial: Accessing a relational
database](https://go.dev/doc/tutorial/database-access)
