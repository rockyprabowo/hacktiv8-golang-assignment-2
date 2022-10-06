package database

func Config() DbConfigMap {
	return DbConfigMap{
		"host":         "localhost",
		"user":         "postgres",
		"password":     "postgres",
		"databaseName": "orders_by",
		"sslMode":      "disable",
	}
}
