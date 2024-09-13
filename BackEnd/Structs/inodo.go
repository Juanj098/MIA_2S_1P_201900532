package structs

import "time"

type Inodo struct {
	i_uid   int64
	i_gid   int64
	i_s     int64
	i_atime time.Time
	i_ctime time.Time
	i_mtime time.Time
	i_block int64
	i_type  byte
	i_perm  [3]byte
}
