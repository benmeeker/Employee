package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func instructions() {
	var a, b, c string
	log.Println("Does this user need instructions for 'program1' use?\nPlease enter 'Yes' or 'No'")
	_, e := fmt.Scanln(&a)
	if e != nil {
		log.Println(e)
	}
	log.Println("Does this user need instructions for phone use?\nPlease enter 'Yes' or 'No'")
	_, er := fmt.Scanln(&b)
	if er != nil {
		log.Println(er)
	}
	log.Println("Does this user need instructions for email use?\nPlease enter 'Yes' or 'No'")
	_, err := fmt.Scanln(&c)
	if err != nil {
		log.Println(err)
	}
	var og *os.File
	new, err := os.Create(`Instructions.docx`)
	if err != nil {
		log.Println(err)
	}
	if strings.ToLower(a) == "yes" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`reference docx`)
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		og, _ = os.Open(`reference docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "yes" {
		og, _ = os.Open(`reference docx`)
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`reference docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		og, _ = os.Open(`reference docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`reference docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`reference docx`)
	} else {
		og, _ = os.Open(`reference docx`)
	}
	io.Copy(new, og)
	login(a, b, c)
}

func login(a, b, c string) {
	var content []string
	file, err := os.Create(`Logins.txt`)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	windows := "Windows Username: Windows Username Here\nWindows Password: Windows Password Here"
	gmail := "Gmail: Gmail Username Here\nGmail Password: Gmail Password Here"
	tcid := "TCID: " + fmt.Sprint(ui.TCID) + "\n"
	everest := "Username: Username Here" "\nPassword: Password Here\n"
	phone := "Phone #: (Number Base Here)" + fmt.Sprint(ui.EXT) + "\nExtension: " + fmt.Sprint(ui.EXT) + "\n"
	if strings.ToLower(a) == "yes" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		content = []string{windows, tcid, everest, phone}
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		content = []string{windows, gmail, tcid, everest}
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "yes" {
		content = []string{windows, gmail, tcid, phone}
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		content = []string{windows, tcid, everest}
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		content = []string{windows, gmail, tcid}
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		content = []string{windows, tcid, phone}
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		content = []string{windows, tcid}
	} else {
		content = []string{windows, gmail, tcid, everest, phone}
	}
	for _, content := range content {
		_, err := file.WriteString(content + "\n")
		if err != nil {
			log.Println(err)
		}
	}
}
