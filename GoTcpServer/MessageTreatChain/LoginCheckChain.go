package MessageTreatChain

import (
	"GoTcpServer/Info"
	"GoTcpServer/Info/InfoParser"
	"errors"
	"database/sql"
	"fmt"
)

type LoginCheckChain struct {
	AbsChain
	Parser InfoParser.Parser
	DatabaseType string
}


func (lcc *LoginCheckChain) Treat(info *Info.MessageInfo, conn *Info.UserConnection) (err error) {

	inputUsername, inputPassword, err := lcc.findInfo(info.MessageContent)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	res, err := lcc.checkValidation_MySQL(inputUsername, inputPassword)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if res == true {
		return lcc.NextChain.Treat(info, conn)
	} else {
		
	}
}
//
//func (lcc *LoginCheckChain) check(content string) (err error) {
//
//	username , password , err := lcc.findInfo(content)
//
//	switch lcc.DatabaseType {
//		case "SQL": {
//			db, err := sql.Open("mysql", "192.168.0.117:3306/root:1shitouren/chatting")
//		}
//	}
//}


func (lcc *LoginCheckChain) checkValidation_MySQL(inputUsername, inputPassword string) (validate bool, err error) {
	db, err := sql.Open("mysql", "root:1shitouren@tcp(127.0.0.1:3306)/chatting")
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	if err = db.Ping() ; err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	rows, err := db.Query("select * from user where username = " + inputUsername)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	var username,password string
	for rows.Next() {
		rows.Scan(&username, &password)
	}

	if password != inputPassword {
		return false, nil
	} else {
		return true, nil
	}
}

func (lcc *LoginCheckChain) findInfo(content string) (username string, password string, err error) {
	var dic interface{}
	if err = lcc.Parser.Parse([]byte(content), &dic) ; err != nil {
		return "", "", err
	}

	if tmp, ok := dic.(map[string]string) ; ok {
		return tmp[username], tmp[password], nil
	} else {
		return "","",errors.New("type Fault in validation")
	}
}

