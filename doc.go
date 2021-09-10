package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/unidoc/unioffice/document"
)

func doc() {
	var a, b, c, d string
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
	if strings.ToLower(a) == "yes" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "yes" {
		d = `N:\IT\users\BenjaminMe\fulldoc.docx`
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		d = `N:\IT\users\BenjaminMe\noemail.docx`
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		d = `N:\IT\users\BenjaminMe\nophone.docx`
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "yes" {
		d = `N:\IT\users\BenjaminMe\noeverest.docx`
	} else if strings.ToLower(a) == "yes" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		d = `N:\IT\users\BenjaminMe\noemailphone.docx`
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "yes" {
		d = `N:\IT\users\BenjaminMe\noeverestphone.docx`
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "yes" && strings.ToLower(c) == "no" {
		d = `N:\IT\users\BenjaminMe\noeverestemail.docx`
	} else if strings.ToLower(a) == "no" && strings.ToLower(b) == "no" && strings.ToLower(c) == "no" {
		d = `N:\IT\users\BenjaminMe\nophoneemaileverest.docx`
	} else {
		d = `N:\IT\users\BenjaminMe\fulldoc.docx`
	}
	doc, err := document.Open(d)
	if err != nil {
		log.Println(err)
	}
	defer doc.Close()
	paragraphs := []document.Paragraph{}
	for _, p := range doc.Paragraphs() {
		paragraphs = append(paragraphs, p)
	}

	for _, sdt := range doc.StructuredDocumentTags() {
		for _, p := range sdt.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}
	}

	for _, p := range paragraphs {
		for _, r := range p.Runs() {
			switch r.Text() {
			case "'an'":
				r.ClearContent()
				r.AddText(ui.First + ui.Last[0:2])
			case "'en'":
				r.ClearContent()
				r.AddText(ui.First + "." + ui.Last)
			case "'n'":
				r.ClearContent()
				r.AddText(ui.First + ui.Last)
			case "'twostep'":
				r.ClearContent()
				r.AddText(fmt.Sprint(ui.TWOST))
			case "'tcid'":
				r.ClearContent()
				r.AddText(fmt.Sprint(ui.TCID))
			case "'ext'":
				r.ClearContent()
				r.AddText(fmt.Sprint(ui.EXT))
			}
		}
	}
	doc.SaveToFile(ui.First + ui.Last + ".docx")
}
