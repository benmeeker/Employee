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
		log.Println(err)
	} else {
		fmt.Printf("%s\n", eu)
	}
}

func changepass() {
	if cp, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "boba", "-ScriptBlock", "{Set-ADUser -Identity "+ui.First+ui.Last[0:2]+" -ChangePasswordAtLogon $true}").CombinedOutput(); err != nil {
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
		if gg, err := exec.Command("cmd", "/c", "gam", "update", "group", fmt.Sprint(inf)+"@namify.com", "add", "member", ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gg)
		}
	}
}

func gamalias() {
	if gg, err := exec.Command("cmd", "/c", "gam", "create", "alias", ui.First+"."+ui.Last+"@namify.com", "user", ui.ALIAS).CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", gg)
	}
}

func gamtwost() {
	if gts, err := exec.Command("cmd", "/c", "gam", "user", ui.First+"."+ui.Last+"@namify.com", "update", "backupcodes").CombinedOutput(); err != nil {
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
		if gt, err := exec.Command("cmd", "/c", "gam", "update", "user", ui.First+"."+ui.Last+"@namify.com", "email", "x"+ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gt)
		}
		if gt, err := exec.Command("cmd", "/c", "gam", "delete", "alias", ui.First+"."+ui.Last+"@namify.com", "user", "x"+ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gt)
		}
	} else if strings.ToLower(a) == "no" {
		if gt, err := exec.Command("cmd", "/c", "gam", "delete", "alias", ui.First+"."+ui.Last+"@namify.com").CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", gt)
		}
	}
}

func termuser() {
	if tu, err := exec.Command("powershell", "Invoke-Command", "-ComputerName", "boba", "-Credential", ui.WHO, "-ScriptBlock", "{Remove-ADUser -Identity "+ui.First+ui.Last[0:2]+" -Confirm:$false}").CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", tu)
	}
	gamterm()
	everterm()
}

func everadd() {
	fillsql()
	if ea, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `BB8\SQL`, "-InputFile", ui.First+ui.Last+`.sql`).CombinedOutput(); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%s\n", ea)
	}
}

func everperm() {
	for _, inf := range ui.GROUPS {
		if ep, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `BB8\SQL`, "-InputFile", addperm(fmt.Sprint(inf))).CombinedOutput(); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("%s\n", ep)
		}
		os.Remove(ui.First + ui.Last + "perms.sql")
	}
}

func everterm() {
	fillsql()
	if et, err := exec.Command("powershell", "Invoke-SqlCmd", "-ServerInstance", `BB8\SQL`, "-Database", "EVEREST_TNS", "-Query", `"update dbo.PERSONAL set ACTIVE='F' where IDNO='`+ui.First+ui.Last[0:2]+`'"`).CombinedOutput(); err != nil {
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
