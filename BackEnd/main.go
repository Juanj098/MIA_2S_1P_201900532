package main

import (
	"bufio"
	"fmt"
	COMMANDS "main/Comandos"

	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		Analizer(input)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error al leer")
	}

}

/*
  ->  Comandos
  - Administracion de discos
  ** MKDISK -
  ** RMDISK -
  ** FDISK -
  ** MOUNT -
  - Admin. sistema de archivos
  ** MKFS -
  ** CAT
  - Admin. Grupos y usuarios
  ** LOGIN
  ** LOGOUT
  ** MKGRP
  ** MKUSR
  ** RMUSR
  ** CHGRP
  - Admin. de carpetas archivos y permisos
  ** MKFILE -
  ** MKDIR -
*/

func Analizer(cadena string) {
	if cadena != "" {
		lines := strings.Split(cadena, "\\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			f := strings.Fields(line)
			com := strings.ToUpper(f[0])
			switch com {
			case "MKDISK":
				COMMANDS.ParseMKDISK(f[1:])
			case "RMDISK":
				COMMANDS.ParseRMDISK(f[1:])
			case "FDISK":
				fmt.Println("3.FDISK")
			case "MOUNT":
				fmt.Println("4.MOUNT")
			case "MKFS":
				fmt.Println("5.MKFS")
			case "CAT":
				fmt.Println("6.CAT")
			case "LOGIN":
				fmt.Println("7.LOGIN")
			case "LOGOUT":
				fmt.Println("8.LOGOUT")
			case "MKGRP":
				fmt.Println("9.MKGRP")
			case "MKUSR":
				fmt.Println("10.MKUSR")
			case "RMUSR":
				fmt.Println("11.RMUSR")
			case "CHGRP":
				fmt.Println("12.CHGRP")
			case "MKFILE":
				fmt.Println("13.MKFILE")
			case "MKDIR":
				fmt.Println("14.MKDIR")
			default:
				fmt.Println("comando desconocido ", f[0])
			}
		}
	} else {
		println("no se leyeron comandos")
	}
}
