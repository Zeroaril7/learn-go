package pkg

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	println("In Initialize File")
	println(connection)
	println()
	return connection
}
