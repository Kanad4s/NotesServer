package db

func selectAllTables() string {
	return "SELECT TABLE_SCHEMA FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA='public'"
}

func selectAllFromTable(table string) string {
	return "SELECT * FROM " + table
}
