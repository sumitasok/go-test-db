package testdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"strings"
)

type TestDb struct {
	Mysql_username string
	Mysql_password string
	Dev_db_name    string
	Test_db_name   string
}

func (tdb TestDb) Prepare() (*sql.DB, error) {
	prepare_file, _ := os.Create("prepare.sh")
	file_content := fmt.Sprintf("mysqldump -u %s -p%s --no-data %s > schema.sql\n\nmysql -u %s -p%s %s < schema.sql\n", tdb.Mysql_username, tdb.Mysql_password, tdb.Dev_db_name, tdb.Mysql_username, tdb.Mysql_password, tdb.Test_db_name)
	prepare_file.WriteString(file_content)

	cmd := exec.Command("sh", "prepare.sh")
	cmd.Output()

	data_source_name := fmt.Sprintf("%s:%s@/%s?parseTime=true", tdb.Mysql_username, tdb.Mysql_password, tdb.Test_db_name)
	db, _ := sql.Open("mysql", data_source_name)

	err := db.Ping()

	if err != nil {
		if strings.Contains(err.Error(), "Error 1049: Unknown database") == true {

			fmt.Println("Caught ya! you need to create one test database for you!")
			fmt.Println(fmt.Sprintf("run: mysql -u %s -p%s -e 'CREATE DATABASE %s'", tdb.Mysql_username, tdb.Mysql_password, tdb.Test_db_name))
		}
		return nil, err
	}

	return db, nil

}

func (tdb TestDb) ClearTable(db *sql.DB, table_name string) error {
	del, err := db.Prepare(fmt.Sprintf("DELETE from %s", table_name))
	if err == nil {
		del.Exec()
	}
	return err
}
