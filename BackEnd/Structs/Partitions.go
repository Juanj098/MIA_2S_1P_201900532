package structs

type Partition struct {
	Partition_status byte
	Partition_type   byte
	partition_fit    byte
	partition_star   int64
	partition_size   int64
	Partition_name   [16]byte
	Partition_corr   int64
	Partition_id     [4]byte
}
