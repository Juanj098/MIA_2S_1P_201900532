package structs

import "fmt"

type Partition struct {
	Partition_status [1]byte
	Partition_type   [1]byte
	Partition_fit    [1]byte
	Partition_star   int32
	Partition_size   int32
	Partition_name   [16]byte
	Partition_corr   int32
	Partition_id     [4]byte
}

/*
	Part status:
	-1 -> no creada
	0 -> inactiva
	1 -> montada
*/

func (part *Partition) CreatePartP(partStart, partSize int, partType, partFit, partName string) {
	part.Partition_status = [1]byte{'0'}
	part.Partition_star = int32(partStart)
	part.Partition_size = int32(partSize)
	if len(partType) > 0 {
		part.Partition_type[0] = partType[0]
	}
	if len(partFit) > 0 {
		part.Partition_fit[0] = partFit[0]
	}
	copy(part.Partition_name[:], []byte(partName))
}

func (part *Partition) CreatePartE(partStart, partSize int, partType, partFit, partName string) {
	part.Partition_status = [1]byte{'0'}
	part.Partition_star = int32(partStart)
	part.Partition_size = int32(partSize)
	if len(partType) > 0 {
		part.Partition_type[0] = partType[0]
	}
	if len(partFit) > 0 {
		part.Partition_fit[0] = partFit[0]
	}
	copy(part.Partition_name[:], []byte(partName))
}

func (part *Partition) Print() {
	fmt.Printf("Part_status: %c\n", part.Partition_status[0])
	fmt.Printf("Part_type: %c\n", part.Partition_type[0])
	fmt.Printf("Part_fit: %c\n", part.Partition_fit[0])
	fmt.Printf("Part_start: %d\n", part.Partition_star)
	fmt.Printf("Part_size: %d\n", part.Partition_size)
	fmt.Printf("Part_name: %s\n", part.Partition_name[:])
	fmt.Printf("Part_correlative: %d\n", part.Partition_corr)
	fmt.Printf("Part_id: %s\n", part.Partition_id)
}
