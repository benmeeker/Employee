package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func fn() string {
	var a string
	fmt.Println("What is the new employee's FIRST name?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	return a
}

func ln() string {
	var a string
	fmt.Println("What is the new employee's LAST name?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	return a
}

func pou() string {
	var a string
	fmt.Println("What is the new employee's PRIMARY ou?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	return a
}

func sou() string {
	var a string
	fmt.Println("What is the new employee's SECONDARY ou?")
	fmt.Println("If there is no SECONDARY ou, please enter 'none'")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	return a
}

func tcid() int {
	var a int
	fmt.Println("What is the employee's TimeClock ID?")
	fmt.Println("If creating a new employee enter '0'")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	if a == 0 {
		input, err := ioutil.ReadFile(`N:\IT\users\BenjaminMe\employeeidcounter.txt`)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		byteNumber := input
		cnt, _ := strconv.Atoi(string(byteNumber))
		cnt++
		ui.TCID = cnt
		log.Println(ui.TCID)
		a = ui.TCID
		file := []byte(strconv.Itoa(ui.TCID))
		if err = ioutil.WriteFile(`N:\IT\users\BenjaminMe\employeeidcounter.txt`, file, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else if a != 0 {
			ui.TCID = a
		}
	}
	log.Println(a)
	return a
}

func fillps() {
	input, err := ioutil.ReadFile(`N:\IT\users\BenjaminMe\adduserscript.ps1`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input = bytes.Replace(input, []byte(`"dn"`), []byte(`"`+ui.First+ui.Last[0:2]+`"`), -1)
	input = bytes.Replace(input, []byte(`"n"`), []byte(`"`+ui.First+" "+ui.Last+`"`), -1)
	input = bytes.Replace(input, []byte(`"fn"`), []byte(`"`+ui.First+`"`), -1)
	input = bytes.Replace(input, []byte(`"ln"`), []byte(`"`+ui.Last+`"`), -1)
	input = bytes.Replace(input, []byte(`"an"`), []byte(`"`+ui.First+ui.Last[0:2]+`"`), -1)
	input = bytes.Replace(input, []byte(`"en`), []byte(`"`+ui.First+ui.Last[0:2]), -1)
	if ui.SOU != "none" {
		input = bytes.Replace(input, []byte(`pou`), []byte(ui.SOU+",OU="+ui.POU), -1)
	} else if ui.SOU == "none" {
		input = bytes.Replace(input, []byte(`pou`), []byte(ui.POU), -1)
	}
	if err = ioutil.WriteFile(ui.First+ui.Last+".ps1", input, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fillsql() {
	input, err := ioutil.ReadFile(`N:\IT\users\BenjaminMe\addeverest.sql`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input = bytes.Replace(input, []byte(`'n'`), []byte(`'`+ui.First+" "+ui.Last+`'`), -1)
	input = bytes.Replace(input, []byte(`'fn'`), []byte(`'`+ui.First+`'`), -1)
	input = bytes.Replace(input, []byte(`'ln'`), []byte(`'`+ui.Last+`'`), -1)
	input = bytes.Replace(input, []byte(`'an'`), []byte(`'`+ui.First+ui.Last[0:2]+`'`), -1)
	input = bytes.Replace(input, []byte(`'en`), []byte(`'`+ui.First+"."+ui.Last), -1)
	input = bytes.Replace(input, []byte(`'tcid'`), []byte(`'`+strconv.Itoa(ui.TCID)+`'`), -1)
	if err = ioutil.WriteFile(ui.First+ui.Last+".sql", input, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cleanup() {
	os.Remove(ui.First + ui.Last + ".ps1")
	os.Remove(ui.First + ui.Last + ".sql")
}