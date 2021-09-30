package main

import (
	
	"fmt" //for Output

	"reflect" //implements run-time reflection, allowing a program to manipulate objects with arbitrary types

	"github.com/spf13/viper" //for converting YAML to GO STRUCT

	"github.com/boltdb/bolt" //for BOLT DB

	"log" //It defines a type, Logger, with methods for formatting output

	"time" // for measuring and displaying time
)

type EMP struct {
	Jobs struct {
		Role1 string `yaml:"Role1"`
		Role2 string `yaml:"Role2"`
	} `yaml:"Jobs"`
	Skills struct {
		FullStack   string `yaml:"FullStack"`				//STRUCT FOR CONTAINING YAML
		AiDeveloper string `yaml:"AiDeveloper"`
	} `yaml:"Skills"`
	Candidates struct {
		Ace    string `yaml:"Ace"`
		Bose   string `yaml:"Bose"`
		Crank  string `yaml:"Crank"`
		Dole   string `yaml:"Dole"`
		Easter string `yaml:"Easter"`
		Folk   string `yaml:"Folk"`
	} `yaml:"Candidates"`
}

func main() {
	vi := viper.New()		//Viper Object
	vi.SetConfigFile("test.yaml")						
	vi.ReadInConfig()	//Reading yaml file
	var emp EMP		
	vi.Unmarshal(&emp)		//converting yaml to struct
		
	db,err:= bolt.Open("EMP.db", 0600, &bolt.Options{Timeout: 1 * time.Second}) //Creating DataBase in Bolt Db
	// bolt options for Timeouting DB lock if DB is in deadlock after 1s
	if err != nil{
		log.Fatalf("Error creating DB: %v",err)
	}
	defer db.Close()
	
	//getting values from Struct using reflect
	v := reflect.ValueOf(emp)
    typeOfS := v.Type()
		
		//Creating Buckets Jobs,Skills and Candidates
    	for i:=0;i<v.NumField();i++{
			db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucket([]byte(typeOfS.Field(i).Name)) //typeOfS.Field(i).Name gives the 
			if err != nil {											 //Field type name of EMP struct which is the bukcet name
				return fmt.Errorf("create bucket: %s", err)		//To print error and exit
			}
			return nil
		})
		
		}
	//inserting into Jobs bucket
	key_value:=reflect.ValueOf(emp.Jobs)
		typeX:=key_value.Type()
		for i:=0;i<key_value.NumField();i++{
			valuename:=fmt.Sprint(key_value.Field(i).Interface())
			db.Update(func(tx *bolt.Tx)error{
				b:=tx.Bucket([]byte("Jobs"))
				err := b.Put([]byte(typeX.Field(i).Name), []byte(valuename)) 
				return err
			})
	}
	//insering in SKills Bucket
	key_value=reflect.ValueOf(emp.Skills)
	typeX=key_value.Type()
	for i:=0;i<key_value.NumField();i++{
		valuename:=fmt.Sprint(key_value.Field(i).Interface())
		db.Update(func(tx *bolt.Tx)error{
			b:=tx.Bucket([]byte("Skills"))
			err := b.Put([]byte(typeX.Field(i).Name), []byte(valuename))
			return err
		})
	}
	//inserting in candidates Bucket
	key_value=reflect.ValueOf(emp.Candidates)
	typeX=key_value.Type()
	for i:=0;i<key_value.NumField();i++{
		valuename:=fmt.Sprint(key_value.Field(i).Interface())
		db.Update(func(tx *bolt.Tx)error{
			b:=tx.Bucket([]byte("Candidates"))
			err := b.Put([]byte(typeX.Field(i).Name), []byte(valuename))
			return err
		})
	}
	
	db.Close()	//Closing the Database connection
	phello("EMP.db") //Calling function an passing DB name
}



	