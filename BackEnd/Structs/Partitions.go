package structs

type Partition struct {
	Partition_status [1]byte
	Partition_type   [1]byte
	Partition_fit    [1]byte
	Partition_star   int64
	Partition_size   int64
	Partition_name   [16]byte
	Partition_corr   int64
	Partition_id     [4]byte
}

/*
	Part status:
	0 -> inactiva
	1 -> montada
*/
