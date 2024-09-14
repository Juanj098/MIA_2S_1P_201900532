package comandos

import (
	"fmt"
	Structs "main/Structs"
	utils "main/Utils"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MKDISK struct {
	size int64
	unit string
	fit  string
	path string
}

func ParseMKDISK(tokens []string) {
	cmd := &MKDISK{}
	args := strings.Join(tokens, " ")
	re := regexp.MustCompile(`-size=\d+|-unit=[kKmM]|-fit=[bBfFwW]{2}|-path="[^"]+"|-path=[^\s]+`)
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
		case "-size":
			size, err := strconv.Atoi(value)
			if err != nil || size <= 0 {
				fmt.Printf("el tamano debde ser un numero entero positivo")
			}
			cmd.size = int64(size)
		case "-unit":
			if value != "M" && value != "K" {
				fmt.Printf("la unidad debe ser M o K")
			}
			cmd.unit = value
		case "-fit":
			value = strings.ToUpper(value)
			if value != "BF" && value != "FF" && value != "WF" {
				fmt.Print("el ajuste debe ser BF, FF, WF")
			}
			cmd.fit = value
		case "-path":
			if value == "" {
				fmt.Print("el path no puede estar vacío")
			}
			cmd.path = value
		default:
			fmt.Println("parametro desconocido: ", key)
		}
	}
	if cmd.size == 0 {
		fmt.Print("ingrese tamano de disco")
	}
	if cmd.fit == "" {
		cmd.fit = "FF"
	}
	if cmd.path == "" {
		fmt.Print("ingrese ubicacion del disco")
	}
	if cmd.unit == "" {
		cmd.unit = "M"
	}
	err := CommandDisk(cmd)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func CommandDisk(disk *MKDISK) error {
	sizeBytes, err := utils.ConvertTobytes(int(disk.size), disk.unit)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = createDisk(disk, sizeBytes)
	if err != nil {
		return err
	}
	err = createMBR(disk, sizeBytes)
	if err != nil {
		return err
	}
	return nil
}

func createDisk(mkdisk *MKDISK, sizeBytes int) error {
	err := os.MkdirAll(filepath.Dir(mkdisk.path), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(mkdisk.path)
	if err != nil {
		fmt.Println("error al crear archivo")
		return err
	}
	defer file.Close()
	buffer := make([]byte, 1024*1024)
	for sizeBytes > len(buffer) {
		writesize := len(buffer)
		if sizeBytes < writesize {
			writesize = sizeBytes
		}
		if _, err := file.Write(buffer[:writesize]); err != nil {
			return nil
		}
		sizeBytes -= writesize
	}
	return nil
}

func createMBR(disk *MKDISK, sizeB int) error {
	var fit byte
	switch disk.fit {
	case "FF":
		fit = 'F'
	case "BF":
		fit = 'B'
	case "WF":
		fit = 'W'
	default:
		fmt.Println("invalid fit type :)")
	}
	mbr := Structs.MBR{
		Mbr_size:           int32(sizeB),
		Mbr_fecha_creacion: float32(time.Now().Unix()),
		Mbr_dsk_signature:  rand.Int31(),
		Disk_fit:           [1]byte{fit},
		Mbr_partition: [4]Structs.Partition{
			{Partition_status: [1]byte{'N'}, Partition_type: [1]byte{'N'}, Partition_fit: [1]byte{'N'}, Partition_star: -1, Partition_name: [16]byte{'P'}, Partition_corr: 0, Partition_id: [4]byte{'N'}},
			{Partition_status: [1]byte{'N'}, Partition_type: [1]byte{'N'}, Partition_fit: [1]byte{'N'}, Partition_star: -1, Partition_name: [16]byte{'P'}, Partition_corr: 1, Partition_id: [4]byte{'N'}},
			{Partition_status: [1]byte{'N'}, Partition_type: [1]byte{'N'}, Partition_fit: [1]byte{'N'}, Partition_star: -1, Partition_name: [16]byte{'P'}, Partition_corr: 2, Partition_id: [4]byte{'N'}},
			{Partition_status: [1]byte{'N'}, Partition_type: [1]byte{'N'}, Partition_fit: [1]byte{'N'}, Partition_star: -1, Partition_name: [16]byte{'P'}, Partition_corr: 3, Partition_id: [4]byte{'N'}},
		},
	}
	err := mbr.SerializeMBR(disk.path)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	mbr.PrintMBR()
	return nil
}
