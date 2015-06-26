For testing database, we need to have a database with exactly same schema of that of development database.

TO-DO: This package gives functionality that will duplicate the database without data in it. Which we can use to populate our test data, and every time we call the database, it is recreated.

```
	ddb := testdb.TestDb{"mysql_username", "mysql_password", "development_db_name", "test_db_name"}
	db, _ := ddb.Prepare()
```

This is based on `go-sql-driver`, and `var db` is returned with the db instance.

> We have to make sure, the code is written in a manner where we can inject this db instance instead of the development instance.
