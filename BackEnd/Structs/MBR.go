package structs

type MBR struct {
	Mbr_tamano         int64        //tamano de mbr en bytes
	Mbr_fecha_creacion float32      //fecha de creaion de mbr
	Mbr_dsk_signature  int64        //firma del disco
	Disk_fit           byte         //tipo de ajuste
	Mbr_partition      [4]Partition //particiones de mbr
}
