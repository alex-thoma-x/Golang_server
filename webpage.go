package main

import (
	
	"net/http"//For Creating GO Server
	
	"html/template"//For parsing Html Template
	
	"github.com/boltdb/bolt" //for BOLT DB

	"log" //It defines a type, Logger, with methods for formatting output
	)
type web struct {  //Struct for getting values from DB
    Name    string
    Role    string
    Skills	string
}
type dat struct{
	
	myslice []web  //slice of type Web
}



func phello(n string){
	db,err:= bolt.Open(n, 0600, nil) //Creating DB in Bolt Db
	//bolt options for Timeouting DB lock if DB is in deadlock after 1s
	if err != nil{
		log.Fatalf("Error  %v",err)
	}
	defer db.Close()
	
	tmpl := template.Must(template.ParseFiles("emp_detail.html"))//Parsing Html template 

	data:=dat{
	myslice:make([]web,6)}
	point:=0
	candidate:=5 //candidate count
	
	for candidate>0{
		db.View(func(tx *bolt.Tx)error{
			cand:=tx.Bucket([]byte("Candidates"))//object creation for Bucket:candidates
			c:=cand.Cursor()
	
			for k, v := c.First(); k != nil; k, v = c.Next() {//retriving values as key value pair from bucket

				job:=tx.Bucket([]byte("Jobs"))//object creation for Bucket :Jobs
				role:=job.Get([]byte(v))     //getting role of candidate 

				skill:=tx.Bucket([]byte("Skills"))//object creation for Bukcet :Skills
				skill_set:=skill.Get([]byte(role))
				
				//Storing the values into myclice []web
				data.myslice[point]=web{Name:string(k),Role: string(role),Skills:string(skill_set) }
				point++						
				candidate--
			
	
			}
			return nil
		})
		
	}
	
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
			tmpl.Execute(w, data.myslice)//executing Html Template and passing value to the template
		})
		http.ListenAndServe(":80", nil)//server creation at port :80
	}


	
