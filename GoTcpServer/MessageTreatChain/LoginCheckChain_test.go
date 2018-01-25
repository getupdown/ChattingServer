package MessageTreatChain

import (
	"testing"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func Test(t *testing.T) {
	db, err := sql.Open("mysql", "root:1shitouren@tcp(127.0.0.1:3306)/chatting")

	if err != nil {
		fmt.Println(err.Error())
	}

	if err = db.Ping() ; err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select * from user where username = 'admin'" )

	for rows.Next() {
		var username,password string
		rows.Scan(&username, &password)
	}


}
