package comandos

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type RMDISK struct {
	path string
}

func ParseRMDISK(tokens []string) {
	cmd := &RMDISK{}
	args := strings.Join(tokens, " ")
	re := regexp.MustCompile(`-path="[^"]+"|-path=[^\s]+`)
	matches := re.FindAllString(args, -1)
	for _, match := range matches {
		kv := strings.SplitN(match, "=", 2)
		if len(kv) != 2 {
			fmt.Printf("formato de entrada invalido: %s\n", match)
		}
		key, value := strings.ToLower(kv[0]), kv[1]
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		}
		switch key {
		case "-path":
			if value == "" {
				fmt.Println("ingrese direccion de archivo :)")
			}
			cmd.path = value
		default:
			fmt.Println("parametro desconocido :)")
		}
	}
	if cmd.path == "" {
		fmt.Println("Falta parametro requerido : \"-path\"")
	}

	err := CommadRMDISK(cmd.path)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func CommadRMDISK(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("el disco no existe")
			return err
		}
	}
	err = os.Remove(path)
	if err != nil {
		fmt.Println("Error al eliminar el Disco :( ")
		return err
	}
	fmt.Println("Disco eliminado con exito UwU")
	return nil
}
