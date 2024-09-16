package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type MBR struct {
	Mbr_size           int32        //tamano de mbr en bytes
	Mbr_fecha_creacion float32      //fecha de creaion de mbr
	Mbr_dsk_signature  int32        //firma del disco
	Disk_fit           [1]byte      //tipo de ajuste
	Mbr_partition      [4]Partition //particiones de mbr
}

// serializacion MBR
func (mbr *MBR) SerializeMBR(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	err = binary.Write(file, binary.BigEndian, mbr)
	if err != nil {
		return err
	}
	return nil
}

// deserializar
func (mbr *MBR) DeserializeMBR(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	mbrsize := binary.Size(mbr)
	if mbrsize <= 0 {
		return fmt.Errorf("INvalid MBR size %d", mbrsize)
	}

	buffer := make([]byte, mbrsize)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(buffer)
	err = binary.Read(reader, binary.BigEndian, mbr)
	if err != nil {
		return err
	}
	return nil
}

func (mbr *MBR) PrintMBR() {
	creationTime := time.Unix(int64(mbr.Mbr_fecha_creacion), 0)
	diskFit := rune(mbr.Disk_fit[0])

	fmt.Println("<----------------->")
	fmt.Printf("MBR_size: %d\n", mbr.Mbr_size)
	fmt.Printf("Creation_Date: %s\n", creationTime)
	fmt.Printf("Disk_signature: %d\n", mbr.Mbr_dsk_signature)
	fmt.Printf("Disk_fit: %c\n", diskFit)
	fmt.Println("<----------------->")
}

func (mbr *MBR) GetAvailablePartition() (*Partition, int, int) {
	offset := binary.Size(mbr)
	for i := 0; i < len(mbr.Mbr_partition); i++ {
		if mbr.Mbr_partition[i].Partition_star == -1 {
			return &mbr.Mbr_partition[i], offset, i
		} else {
			offset += int(mbr.Mbr_partition[i].Partition_size)
		}
	}
	return nil, -1, -1
}

func (mbr *MBR) PrintPart() {
	for i, partition := range mbr.Mbr_partition {
		partStatus := rune(partition.Partition_status[0])
		partType := rune(partition.Partition_type[0])
		partFit := rune(partition.Partition_fit[0])

		partName := string(partition.Partition_name[:])
		partID := string(partition.Partition_id[:])

		fmt.Printf("Partition %d\n", i+1)
		fmt.Printf("status: %c\n", partStatus)
		fmt.Printf("Type: %c\n", partType)
		fmt.Printf("FIt: %c\n", partFit)
		fmt.Printf("Start: %d\n", partition.Partition_star)
		fmt.Printf("Size: %d\n", partition.Partition_size)
		fmt.Printf("Name: %s\n", partName)
		fmt.Printf("Correlativo: %d\n", partition.Partition_corr)
		fmt.Printf("ID: %s\n", partID)
	}
}

func (mbr *MBR) ContainExt() bool {
	for _, part := range mbr.Mbr_partition {
		if part.Partition_type == [1]byte{'E'} {
			return true
		}
	}
	return false
}
