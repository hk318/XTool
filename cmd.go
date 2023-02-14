package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	opts := DefaultOptions()

	opts.Clear()
	opts.menu()

	var input int

Loop:
	for {
		fmt.Printf("%v%v: ", green("请输入数字"), blue(" (回车确认)"))
		if _, err := fmt.Scan(&input); err != nil {
			fmt.Printf("%v %v", red("[Warning]"), yellow("请输入一个正确的数字!\n\n"))
			continue
		}
		break
	}

	switch input {
	case 0:
		os.Exit(0)
	case 1:
		opts = opts.ConfirmModify()
		opts.CaddyInstallation()

	default:
		fmt.Printf("%v %v", red("[Warning]"), yellow("请输入一个正确的数字!\n\n"))
		goto Loop
	}
}

func (c *Config) Clear() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	if err := clear.Run(); err != nil {
		c.logger.Error(err.Error())
	}
}

func (c *Config) FileExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}

	return true
}
