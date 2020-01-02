module mytest

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1-0.20191212001955-b66d043e6c89 // indirect
	mycore v0.0.0
)

replace mycore v0.0.0 => ../mycore
