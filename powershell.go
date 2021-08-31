package main

import (
	"fmt"
	"log"
	"os/exec"
)

func adduser() {
	if au, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "boba", "-Credential", "BenjaminMe", "-FilePath", ui.First+ui.Last+".ps1").CombinedOutput(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", au)
	}
	enableuser()
	changepass()
}

func enableuser() {
	if eu, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "boba", "-ScriptBlock", "{Enable-ADAccount -Identity "+ui.First+ui.Last[0:2]+"}").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", eu)
	}
}

func changepass() {
	if cp, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "boba", "-ScriptBlock", "{Set-ADUser -Identity "+ui.First+ui.Last[0:2]+" -ChangePasswordAtLogon $true}").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", cp)
	}
}

func gamadd() {
	if ga, err := exec.Command("cmd", "/c", "gam", "create", "user", ui.First+"."+ui.Last+"@namify.com", "firstname", ui.First, "lastname", ui.Last, "password", ui.First+ui.Last+"4321", "changepassword", "on").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", ga)
	}
}

func gamgroup() {
	if gg, err := exec.Command("cmd", "/c", "gam", "update", "group", "everyone@namify.com", "add", "member", ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", gg)
	}
}

func termuser() {
	if tu, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "boba", "-Credential", "BenjaminMe", "-ScriptBlock", "{Remove-ADUser -Identity "+ui.First+ui.Last[0:2]+" -Confirm:$false}").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", tu)
	}
	gamterm()
}

func gamterm() {
	if gt, err := exec.Command("cmd", "/c", "gam", "update", "user", ui.First+"."+ui.Last+"@namify.com", "email", "x"+ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", gt)
	}
	if gt, err := exec.Command("cmd", "/c", "gam", "delete", "alias", ui.First+"."+ui.Last+"@namify.com", "user", "x"+ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", gt)
	}
}

func everadd() {
	if ea, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `BB8\SQL`, "-InputFile", ui.First+ui.Last+`.sql`).CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", ea)
	}
}
