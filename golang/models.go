package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"golang.org/x/crypto/bcrypt"
	"time"

)

type Users struct {
 Username string `json:"username"`
 Emailid   string `json:"emailid,omitempty"`
 Password  string `json:"password"`
 Phoneno string `json:"phoneno"`
}





type Response struct {
	Status  int    `json:" status "`
	Message string `json:" message "`
	Data    []Users
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	var db = connect()

	rows, err := db.Query("SELECT `userName`,` emailid`,` password`,` phoneNo` FROM `UserData`")
	if err != nil {
		log.Print(err, "cdsdc")
	}

	for rows.Next() {
		if err := rows.Scan(&users.Username, &users.Emailid, &users.Password, &users.Phoneno); err != nil {
			log.Fatal(err.Error(),)

		} else {

			arr_user = append(arr_user, users)
			log.Print(users)
		}

	}


	response.Status = 1
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}

func userAddHandler(w http.ResponseWriter, r *http.Request) {

	//make byte array
	out := make([]byte, 1024)

	//
	bodyLen, err := r.Body.Read(out)

	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	var k Users

	err = json.Unmarshal(out[:bodyLen], &k)
	
	phonere := regexp.MustCompile("^(6|[7-9])[0-9]{9}$")
	emailre := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	
	if emailre.MatchString(k.Emailid) ==false{
		w.Write([]byte("{error:" + "Incorrect emailid" + "}"))
		return 
	}
	if phonere.MatchString(k.Phoneno) ==false{
		w.Write([]byte("{error:" + "Incorrect phoneno" + "}"))
		return 
	}
	
	

	log.Print(k.Username,"kya aya")

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	err = insertInDatabase(k)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	w.Write([]byte(`{"msg":"success"}`))

}

func insertInDatabase(data Users) error {
	var db = connect()
	//execute statement
	fmt.Println(data)
	log.Print(data)
	password := data.Password
	hash, _ := HashPassword(password)
	dt := time.Now()
	var datetime=dt.Format("2006-01-02 15:04:05")
	log.Print(datetime)

	_, err := db.Exec("REPLACE INTO `UserData` (` emailid`, `userName`, ` phoneNo`, ` password`, ` datetime`) VALUES (?,?,?,?,?)", data.Emailid, data.Username, data.Phoneno, hash,datetime)
	return err

}
