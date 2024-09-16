package comandos

import (
	"fmt"
	structs "main/Structs"
	utils "main/Utils"
	"regexp"
	"strconv"
	"strings"
)

type FDISK struct {
	Size int
	Unit string
	Path string
	Type string
	Fit  string
	Name string
}

func ParseFDISK(tokens []string) {
	cmd := &FDISK{}
	args := strings.Join(tokens, " ")
	re := regexp.MustCompile(`-size=\d+|-unit=[bBkKmM]|-path="[^\"]"|-path=[^\s]+|-fit=[bBfFwW]{2}|-type=[pPeElL]|-name=[^\s]+|-name="[^\"]"`)
	matches := re.FindAllString(args, -1)
	for _, match := range matches {
		kv := strings.SplitN(match, "=", 2)
		if len(kv) != 2 {
			fmt.Println("formato de parametro invalido: ", kv)
		}
		key, value := strings.ToLower(kv[0]), kv[1]
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
			value = strings.Trim(value, "\"")
		}
		switch key {
		case "-size":
			size, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("el tamano de la particion debe ser un numero entero positivo")
			}
			cmd.Size = size
		case "-path":
			if value == "" {
				fmt.Println("ingrese direccion de disco")
			}
			cmd.Path = value
		case "-unit":
			if value != "K" && value != "B" && value != "M" {
				fmt.Println("la unidad debe ser B, M o K")
			}
			cmd.Unit = value
		case "-type": //P-E-L
			value = strings.ToUpper(value)
			if value != "P" && value != "E" && value != "L" {
				fmt.Println("El tipo de particion debe ser P,L o E")
			}
			cmd.Type = value
		case "-fit": // BF - FF - WF
			if value != "WF" && value != "BF" && value != "FF" {
				fmt.Println("El ajuste debe ser BF, FF o WF")
			}
			cmd.Fit = value
		case "-name":
			if value == "" {
				fmt.Println("ingrese nombre de particion :)")
			}
			cmd.Name = value
		default:
			fmt.Println("parametro desconocido: ", key)
		}
	}
	if cmd.Size == 0 {
		fmt.Println("falta parametro requerido: size")
	}
	if cmd.Name == "" {
		fmt.Println("falta parametro requerido: name")
	}
	if cmd.Path == "" {
		fmt.Println("falta parametro requerido: path")
	}
	if cmd.Fit == "" {
		cmd.Fit = "WF"
	}
	if cmd.Type == "" {
		cmd.Type = "P"
	}
	if cmd.Unit == "" {
		cmd.Unit = "K"
	}
	err := CommandFDISK(cmd)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func CommandFDISK(disk *FDISK) error {
	sizeB, err := utils.ConvertTobytes(disk.Size, disk.Unit)
	if err != nil {
		fmt.Println("error al convertir: ", err)
		return err
	}
	if disk.Type == "P" {
		err = P_Primaria(disk, sizeB)
		if err != nil {
			fmt.Println("err: ", err)
		}
	} else if disk.Type == "E" {
		err = P_Extendida(disk, sizeB)
		if err != nil {
			fmt.Println("err: ", err)
		}

	} else if disk.Type == "L" {
		err = P_Logica(disk, sizeB)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}
	return nil
}

// particion Primaria
func P_Primaria(disk *FDISK, sizeB int) error {
	var mbr structs.MBR
	err := mbr.DeserializeMBR(disk.Path)
	if err != nil {
		return err
	}

	availablePart, startPart, index := mbr.GetAvailablePartition()
	if availablePart == nil {
		fmt.Println("no hay particion disponible :(")
	}

	fmt.Println("<Particion Disponible>")
	availablePart.Print()
	fmt.Print("\n")

	availablePart.CreatePartP(startPart, sizeB, disk.Type, disk.Fit, disk.Name)
	if availablePart != nil {
		mbr.Mbr_partition[index] = *availablePart
	}

	fmt.Println("<Particiones>")
	mbr.PrintPart()
	fmt.Print("\n")

	err = mbr.SerializeMBR(disk.Path)
	if err != nil {
		return err
	}
	return nil
}

// particion Extendida
func P_Extendida(disk *FDISK, sizeB int) error {
	var mbr structs.MBR
	err := mbr.DeserializeMBR(disk.Path)
	if err != nil {
		return err
	}

	availablePart, startPart, index := mbr.GetAvailablePartition()
	if availablePart == nil {
		fmt.Println("no hay particion disponible :(")
	}

	fmt.Println("<Particion Disponible>")
	availablePart.Print()
	fmt.Print("\n")

	availablePart.CreatePartE(startPart, sizeB, disk.Type, disk.Fit, disk.Name)
	if availablePart != nil {
		mbr.Mbr_partition[index] = *availablePart
	}

	err = CreateEBR(disk, availablePart)
	if err != nil {
		return err
	}

	err = mbr.SerializeMBR(disk.Path)
	if err != nil {
		return err
	}
	return nil
}

// particion logica
func P_Logica(disk *FDISK, sizeB int) error {
	var mbr structs.MBR
	err := mbr.DeserializeMBR(disk.Path)
	if err != nil {
		return err
	}

	resp := mbr.ContainExt()
	if resp {
		fmt.Println("contiene particion extendida")
	} else {
		fmt.Println("no contiene particion extendida")
	}

	err = mbr.SerializeMBR(disk.Path)
	if err != nil {
		return err
	}
	return nil
}

func CreateEBR(f *FDISK, p *structs.Partition) error {
	ebr := structs.EBR{
		Part_mount: [1]byte{'N'},
		Part_fit:   [1]byte{'N'},
		Part_start: p.Partition_star,
		Part_s:     0,
		Part_next:  -1,
		Part_name:  [16]byte{'N'},
	}
	err := ebr.SerializeEBR(f.Path, int64(p.Partition_star))
	if err != nil {
		return err
	}
	fmt.Println("<EBR>")
	ebr.PrintEBR()
	return nil
}
