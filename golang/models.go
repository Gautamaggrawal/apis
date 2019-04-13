package main

import (
	"fmt"
	"log"
	"encoding/json" 	 
	"net/http" 
	"io"
) 

type Users struct {  
	username string `form:" username "json:" username "` 
	emailid string `form:" emailid "json:" emailid "` 
	password string `form:" password "json:" password "`
	phoneno string `form:" phoneno "json:" phoneno "`

} 

type Response struct { 
	Status int `json:" status "` 
	Message string `json:" message "` 
	Data [] Users 
}

func returnAllUsers (w http.ResponseWriter, r * http.Request) { 
	var users Users 
	var arr_user [] Users 
	var response Response 

	var db=connect()

	rows, err:= db.Query("SELECT `userName`,` emailid`,` password`,` phoneNo` FROM `UserData`")
	if err != nil { 
		log.Print(err,"cdsdc") 
	} 

	for rows.Next() { 
		if err:= rows.Scan(&users.username , &users.emailid, &users.password, &users.phoneno  ); err != nil { 
			log.Fatal(err.Error(),"kya") 

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
       out := make([]byte,1024)

       //
       bodyLen, err := r.Body.Read(out)

       if err != io.EOF {
              fmt.Println(err.Error(),"kya hua")
              w.Write([]byte("{error:" + err.Error() + "}"))
              return
       }

       var k Users

       err = json.Unmarshal(out[:bodyLen],&k)
       log.Print(&k)


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
	var db=connect()
      //execute statement
      fmt.Println(data)
      log.Print(data)

       _, err := db.Exec("INSERT INTO `UserData` (` emailid`, `userName`, ` phoneNo`, ` password`) VALUES (?,?, ?,?)",data.emailid , data.username, data.phoneno,data.password)
       return err

}
