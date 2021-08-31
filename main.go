package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var a string
	log.Println("Would you like to 'add' or 'terminate' this user?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	if a == "add" {
		log.Println("Initiating birthing protocol")
		fillps()
		tcid()
		fillsql()
		everadd()
		adduser()
		log.Println("Would you like this user to have an individual email address? 'Yes' or 'No'")
		var m string
		_, err := fmt.Scanln(&m)
		if err != nil {
			log.Println(err)
		}
		sepem := strings.ToLower(m)
		if sepem == "yes" {
			gamadd()
			gamgroup()
		} else if sepem == "no" {
			log.Println("Code for adding alias here")
		}
		cleanup()
	}
	if a == "terminate" {
		log.Println("Commencing self destruct")
		termuser()
	} else {
		log.Println("Sorry I don't understand people that suck at typing")
	}
}
