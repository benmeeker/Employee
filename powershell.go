package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func adduser() {
	fillps()
	if au, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "Computer Name Here", "-Credential", "Credentials Here", "-FilePath", ".ps1").CombinedOutput(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", au)
	}
	enableuser()
	changepass()
}

func enableuser() {
	if eu, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "Computer Name Here", "-ScriptBlock", "{Enable-ADAccount -Identity AD Username Here}").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", eu)
	}
}

func changepass() {
	if cp, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "Computer Name Here", "-ScriptBlock", "{Set-ADUser -Identity AD Username Here" -ChangePasswordAtLogon $true}").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", cp)
	}
}

func gamadd() {
	if ga, err := exec.Command("cmd", "/c", "gam", "create", "user", ui.First+"."+ui.Last+"@namify.com", "firstname", ui.First, "lastname", ui.Last, "password", ui.First+ui.Last+"4321", "changepassword", "on").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", ga)
	}
}

func gamgroup() {
	for _, inf := range ui.EMAILGROUPS {
		if gg, err := exec.Command("cmd", "/c", "gam", "update", "group", "email groups here", "add", "member", "User email here").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gg)
		}
	}
}

func gamalias() {
	if gg, err := exec.Command("cmd", "/c", "gam", "create", "alias", "User email here", "user", "Alias here").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", gg)
	}
}

func gamtwost() {
	if gts, err := exec.Command("cmd", "/c", "gam", "user", "User email here", "update", "backupcodes").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		ui.TWOST = fmt.Sprintf("%s\n", gts)
	}
}

func gamterm() {
	var a string
	log.Println("Does this user have an individual email?\nPlease enter 'Yes' or 'No'")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Println(err)
	}
	if strings.ToLower(a) == "yes" {
		if gt, err := exec.Command("cmd", "/c", "gam", "update", "user", "User email here", "email", "Terminated user email here").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gt)
		}
		if gt, err := exec.Command("cmd", "/c", "gam", "delete", "alias", "User email here", "user", "Terminated user email here").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gt)
		}
	} else if strings.ToLower(a) == "no" {
		if gt, err := exec.Command("cmd", "/c", "gam", "delete", "alias", "User email here").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gt)
		}
	}
}

func termuser() {
	if tu, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "Computer Name Here", "-Credential", ui.WHO, "-ScriptBlock", "{Remove-ADUser -Identity AD Username Here -Confirm:$false}").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", tu)
	}
	gamterm()
	everterm()
}

func everadd() {
	fillsql()
	if ea, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `Server Here`, "-InputFile", `.sql`).CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", ea)
	}
}

func everperm() {
	for _, inf := range ui.GROUPS {
		if ep, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `Server Here`, "-InputFile", addperm(fmt.Sprint(inf))).CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", ep)
		}
		os.Remove("perms.sql")
	}
}

func everterm() {
	fillsql()
	if et, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `Server Here`, "-Database", "EDatabase Here", "-Query", `"update dbo."database here" set ACTIVE='F' where IDNO='ID Number Here'"`).CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", et)
	}
}

func whoami() string {
	var who string
	if wai, err := exec.Command("cmd", "/c", "gam", "whoami").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		output := fmt.Sprintf("%s\n", wai)
		who = strings.Trim(output, `namify\`)
	}
	return who
}
