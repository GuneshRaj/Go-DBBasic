# Go-DBBasic
Database Helper methods for GoLang
Based on DBBasic in Java, PHP

Methods for Simplifying development in Go Lang.

Update Statement

sql : Standard SQL statement thats meant to be Executed by the server - eg Insert, Update or Delete

```go
func updateStatement(db *sql.DB, sql string) error
```

```go

```



Select Statement

sql : Standard SQL statement thats meant to be Queried by the server - Usually Select

Returns an Array of Map of the result set. []map[string]string

```go
func selectStatement(db *sql.DB, sql string) ([]map[string]string, error)
```

```go

```

Select Single Statement

sql : Standard SQL statement for 1 record thats meant to be Queried by the server - Usually Select

Returns a Map of the result set. map[string]string

```go
func selectSingleStatement(db *sql.DB, sql string) (map[string]string, error)
```

```go

```

Select Data

sql : Standard SQL statement for 1 record and 1 data thats meant to be Queried by the server - Usually Select x from A limit 1

Returns a string of the data returned

```go
func selectData(db *sql.DB, sql string) (string, error)
```

```go

```



