package structs

//Bloque de carpetas
type B_carpetas struct {
	B_content [4]B_content
}

//Bloque de archivos
type B_Arch struct {
	B_content [64]byte
}

//Bloque de apuntadores
type B_pointers struct {
	B_pinters [16]int64
}

type B_content struct {
	B_name  [12]byte
	B_inode int64
}
