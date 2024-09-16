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
		return err
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
