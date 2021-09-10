package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/unidoc/unioffice/common/license"
)

// 515676efae26f59af2a0b72fdf7c8d546ac91d937c2027b19594769147558563

func init() {
	err := license.SetMeteredKey(`515676efae26f59af2a0b72fdf7c8d546ac91d937c2027b19594769147558563`)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var a string
	log.Println("Would you like to 'add' or 'terminate' this user?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	if a == "add" {
		//ext()
		log.Println("Initiating birthing protocol")
		//tcid()
		//everadd()
		//everperm()
		//adduser()
		log.Println("Would you like this user to have an individual email address? 'Yes' or 'No'")
		var m string
		_, err := fmt.Scanln(&m)
		if err != nil {
			log.Println(err)
		}
		sepem := strings.ToLower(m)
		if sepem == "yes" {
			//emailgroups()
			//gamadd()
			//gamgroup()
		} else if sepem == "no" {
			//alias()
			//gamalias()
		}
		doc()
		log.Println("For the time being please create employee in namify.axomo.com\ngive IsNamify permission, add to axomo role manager\nand add to phone directory")
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
