package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type EBR struct {
	Part_mount [1]byte
	Part_fit   [1]byte
	Part_start int32
	Part_s     int32
	Part_next  int32
	Part_name  [16]byte
}

func (ebr *EBR) SerializeEBR(path string, pos int64) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(pos, 0)
	if err != nil {
		return fmt.Errorf("aqui?: %v", err)
	}

	err = binary.Write(file, binary.BigEndian, ebr)
	if err != nil {
		return err
	}
	return nil
}

func (ebr *EBR) DeserializeEBR(path string, pos int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(int64(pos), 0)
	if err != nil {
		return err
	}

	ebrsize := binary.Size(ebr)
	if ebrsize <= 0 {
		return fmt.Errorf("invalid EBR size %d", ebrsize)
	}

	buffer := make([]byte, ebrsize)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(buffer)
	err = binary.Read(reader, binary.BigEndian, ebr)
	if err != nil {
		return err
	}
	return nil
}

func (ebr *EBR) PrintEBR() {
	fmt.Printf("p_mount: %c\n", ebr.Part_mount[0])
	fmt.Printf("p_fit: %c\n", ebr.Part_fit[0])
	fmt.Printf("p_start: %d\n", ebr.Part_start)
	fmt.Printf("p_s: %d\n", ebr.Part_s)
	fmt.Printf("p_next: %d\n", ebr.Part_next)
	fmt.Printf("p_name: %s\n", ebr.Part_name)
}

func (ebr *EBR) FIndLastEBR(starpos int, path string) (*EBR, int, error) {
	var currentEBR EBR
	pos := starpos
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()
	for {
		_, err = file.Seek(int64(pos), 0)
		if err != nil {
			return nil, 0, err
		}
		EBRsize := binary.Size(currentEBR)
		if EBRsize <= 0 {
			return nil, 0, fmt.Errorf("el tamano del ebr no es compatible: %d\n ", EBRsize)
		}
		buffer := make([]byte, EBRsize)
		_, err = file.Read(buffer)
		if err != nil {
			return nil, 0, err
		}
		reader := bytes.NewReader(buffer)
		err = binary.Read(reader, binary.BigEndian, &currentEBR)
		if err != nil {
			return nil, 0, err
		}
		if currentEBR.Part_next == -1 {
			return &currentEBR, pos, nil
		}
		pos = int(currentEBR.Part_next)
	}
}

func (ebr *EBR) Prints(posI int, path string) ([]EBR, error) {
	var currentEBR EBR
	pos := posI
	var ebrList []EBR
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	for {

		_, err = file.Seek(int64(pos), 0)
		if err != nil {
			return nil, err
		}
		EBRsize := binary.Size(currentEBR)
		if EBRsize <= 0 {
			return nil, fmt.Errorf("el tamano del ebr no es compatible: %d\n ", EBRsize)
		}
		buffer := make([]byte, EBRsize)
		_, err = file.Read(buffer)
		if err != nil {
			return nil, err
		}
		reader := bytes.NewReader(buffer)
		err = binary.Read(reader, binary.BigEndian, &currentEBR)
		if err != nil {
			return nil, err
		}
		ebrList = append(ebrList, currentEBR)
		if currentEBR.Part_next == -1 {
			break
		}
		pos = int(currentEBR.Part_next)
	}
	return ebrList, nil
}
