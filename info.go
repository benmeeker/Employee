package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func mantcid() int {
	var a int
	log.Println("What is the TimeClock ID?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	return a
}

//func tcid() int {
//	var a int
//	fmt.Println("What is the employee's TimeClock ID?")
//	fmt.Println("If creating a new employee enter '0'")
//	_, err := fmt.Scanln(&a)
//	if err != nil {
//		log.Println(err)
//	}
//	if a == 0 {
//		input, err := ioutil.ReadFile(`idcounter.txt`)
//		if err != nil {
//			log.Println(err)
//			os.Exit(1)
//		}
//		byteNumber := input
//		cnt, _ := strconv.Atoi(string(byteNumber))
//		cnt++
//		ui.TCID = cnt
//		a = ui.TCID
//		file := []byte(strconv.Itoa(ui.TCID))
//		if err = ioutil.WriteFile(`idcounter.txt`, file, 0666); err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		} else if a != 0 {
//			ui.TCID = a
//		}
//	}
//	log.Println(a)
//	return a
//}

func ext() int {
	var a string
	log.Println("What is the employee's extension?")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	b, err := strconv.Atoi(a)
	if err != nil {
		log.Println(err)
	}
	return b
}

func date() string {
	now := time.Now()
	y, m, d := now.Date()
	date := (fmt.Sprint(d) + "/" + fmt.Sprint(int(m)) + "/" + fmt.Sprint(y))
	return date
}

func alias() {
	var a string
	fmt.Println("Please enter the email you would like this user to be an alias of")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	ui.ALIAS = a
}

func groups() []string {
	var a string
	var c []string
	fmt.Println("Would you like to add this user to any groups other than 'Everyone'?")
	fmt.Println("Please enter 'Yes' or 'No'")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	if strings.ToLower(a) == "yes" {
		var b string
		fmt.Println("List Groups Here if Applicable")
		fmt.Println("Please enter the groups you want to add\nline by line in a case sensitive format\ntype 'done' when finished")
		for {
			_, err := fmt.Scanln(&b)
			if err != nil {
				log.Println(err)
			}
			if strings.ToLower(b) == "done" {
				c = append(c, "EVERY")
				return c
			} else {
				c = append(c, b)
				continue
			}
		}
	}
	c = append(c, "EVERY")
	return c
}

func emailgroups() {
	var a string
	var c []string
	fmt.Println("Would you like to add this user to any email groups other than 'Everyone'?\nIf you plan to terminate this employee please enter 'No'")
	fmt.Println("Please enter 'Yes' or 'No'")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	if strings.ToLower(a) == "yes" {
		var b string
		fmt.Println("List Email Groups Here")
		fmt.Println("Please enter the email groups you want to add type 'done' when finished")
		for {
			_, err := fmt.Scanln(&b)
			if err != nil {
				log.Println(err)
				break
			}
			if strings.ToLower(b) == "done" {
				c = append(c, "everyone")
				ui.EMAILGROUPS = c
				break
			} else {
				c = append(c, b)
				continue
			}
		}
	}
}

func fillps() {
	input, err := ioutil.ReadFile(`adduserscript.ps1`)
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
	input, err := ioutil.ReadFile(`addeverest.sql`)
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
	//input = bytes.Replace(input, []byte(`'start'`), []byte(`'`+ui.DATE+`'`), -1)
	if err = ioutil.WriteFile(ui.First+ui.Last+".sql", input, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addperm(inf string) string {
	input, err := ioutil.ReadFile(`addeverestperms.sql`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input = bytes.Replace(input, []byte(`'mo'`), []byte(`'`+fmt.Sprint(inf)+`'`), -1)
	input = bytes.Replace(input, []byte(`'an'`), []byte(`'`+ui.First+ui.Last[0:2]+`'`), -1)
	if err = ioutil.WriteFile(ui.First+ui.Last+"perms.sql", input, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ui.First + ui.Last + "perms.sql"
}

func cleanup() {
	os.Remove(ui.First + ui.Last + ".ps1")
	os.Remove(ui.First + ui.Last + ".sql")
	os.Remove(ui.First + ui.Last + "perms.sql")
}
