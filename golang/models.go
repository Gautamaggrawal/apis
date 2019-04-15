package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Users struct {
	Emailid  string `json:"emailid,omitempty"`
	Username string `json:"username"`
	Phoneno  string `json:"phoneno"`
	Password string `json:"password"`
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

	rows, err := db.Query("SELECT `emailId`, `userName`, `phoneNo`, `password` FROM `userData`")
	if err != nil {
		log.Print(err, "cdsdc")
	}

	for rows.Next() {
		if err := rows.Scan(&users.Emailid, &users.Username, &users.Phoneno, &users.Password); err != nil {
			log.Fatal(err.Error())

		} else {

			arr_user = append(arr_user, users)

		}

	}

	response.Status = 1
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}

func userAddHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-type")
	phonere := regexp.MustCompile("^(6|[7-9])[0-9]{9}$")
	emailre := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	r.ParseForm()
	var username = r.FormValue("username")
	var emailid = r.FormValue("emailid")
	var password = r.FormValue("password")
	var phoneno = r.FormValue("phoneno")
	// w.Write([]byte("{error:" + username + "}"))
	if emailre.MatchString(emailid) == false {
		json.NewEncoder(w).Encode("Incorrect Email")
		return
	}

	if phonere.MatchString(phoneno) == false {

		json.NewEncoder(w).Encode("Incorrect Mobile")
		return
	}

	var db = connect()
	hash, _ := HashPassword(password)
	dt := time.Now()
	var datetime = dt.Format("2006-01-02 15:04:05")
	_, err := db.Exec("REPLACE INTO `userData` (`emailId`, `userName`, `phoneNo`, `password`, `dateTime`) VALUES (?,?,?,?,?)", emailid, username, phoneno, hash, datetime)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("success")
	return

	if contentType == "application/json" {

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

		// log.Print(phonere,emailre)

		if emailre.MatchString(k.Emailid) == false {
			w.Write([]byte("{error:" + "Incorrect emailid" + "}"))
			return
		}
		if phonere.MatchString(k.Phoneno) == false {
			w.Write([]byte("{error:" + "Incorrect phoneno" + "}"))
			return
		}

		log.Print(k.Username, "kya aya")

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

}

func insertInDatabase(data Users) error {
	var db = connect()
	//execute statement
	fmt.Println(data)
	log.Print(data)
	password := data.Password
	hash, _ := HashPassword(password)
	dt := time.Now()
	var datetime = dt.Format("2006-01-02 15:04:05")
	log.Print(datetime)

	_, err := db.Exec("REPLACE INTO `userData` (`emailId`, `userName`, `phoneNo`, `password`, `dateTime`) VALUES (?,?,?,?,?)", data.Emailid, data.Username, data.Phoneno, hash, datetime)
	return err

}

func emailsearchHandler(w http.ResponseWriter, r *http.Request) {
	var db = connect()
	var name string
	var em = r.FormValue("emailid")

	err := db.QueryRow("SELECT `emailId` FROM `userData` WHERE ` emailId`= ?", em).Scan(&name)
	if err != nil {
		json.NewEncoder(w).Encode("Not found")
		return
	}
	json.NewEncoder(w).Encode("Found")
	return
}

func emaildeleteHandler(w http.ResponseWriter, r *http.Request) {
	var db = connect()
	// var name string
	var em = r.FormValue("emailid")
	log.Print(em)
	var err error

	res, err := db.Exec("DELETE FROM `userData` WHERE `emailId`= ?", em)
	log.Print(res, err)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			log.Print(count)
			if count == 1 {
				json.NewEncoder(w).Encode("Deleted")
				return

			}
			json.NewEncoder(w).Encode("Not found")
			return
			/* check count and return true/false */
		}

	}
	log.Print("asf")
	return

}
