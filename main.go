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
		everadd()
		everperm()
		adduser()
		log.Println("Would you like this user to have an individual email address? 'Yes' or 'No'")
		var m string
		_, err := fmt.Scanln(&m)
		if err != nil {
			log.Println(err)
		}
		sepem := strings.ToLower(m)
		if sepem == "yes" {
			emailgroups()
			gamadd()
			gamgroup()
			gamtwost()
		} else if sepem == "no" {
			alias()
			gamalias()
		}
		instructions()
		log.Println("\n\nFor the time being please create employee in namify.axomo.com\ngive IsNamify permission, add to axomo role manager\nadd to Jive (if needed), and add to the phone directory\nDon't forget to print out the forms!\n\n")
		cleanup()
	}
	if a == "terminate" {
		log.Println("Commencing self destruct")
		termuser()
		cleanup()
	} else {
		log.Println("Sorry I don't understand people that suck at typing")
	}
}
