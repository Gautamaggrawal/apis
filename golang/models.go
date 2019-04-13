package main

import (
	
	"log"
	"encoding/json" 	 
	"net/http" 
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
