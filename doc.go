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
	log.Println("Does this user need instructions for everest use?\nPlease enter 'Yes' or 'No'")
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
	new, err := os.Create(`C:\users\BenjaminMe\Desktop\Instructions.docx`)
	if err != nil {
		log.Println(err)
	}
	if strings.ToLower(a) == "yes" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\noemail.docx`)
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\nophone.docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "yes" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\noeverest.docx`)
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\noemailphone.docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\noeverestphone.docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\noeverestemail.docx`)
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\nophoneemaileverest.docx`)
	} else {
		og, _ = os.Open(`N:\IT\users\BenjaminMe\fulldoc.docx`)
	}
	io.Copy(new, og)
	login(a, b, c)
}

func login(a, b, c string) {
	var content []string
	file, err := os.Create(`C:\users\BenjaminMe\Desktop\Logins.txt`)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	windows := "Windows Username: " + ui.First + ui.Last[0:2] + "\nWindows Password: Namify4321\n"
	gmail := "Gmail: " + ui.First + "." + ui.Last + "\nGmail Password: " + ui.First + ui.Last + "4321\n" + ui.TWOST + "\n"
	tcid := "TCID: " + fmt.Sprint(ui.TCID) + "\n"
	everest := "Everest Username: " + ui.First + ui.Last[0:2] + "\nEverest Password: abcd4321!\n"
	phone := "Phone #: 801-704-3" + fmt.Sprint(ui.EXT) + "\nExtension: " + fmt.Sprint(ui.EXT) + "\n"
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
