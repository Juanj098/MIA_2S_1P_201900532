rmdisk -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disco1.mia
mkdisk -size=10 -unit=M -fit=WF -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disk1.mia
fdisk -size=2 -type=P -unit=M -fit=BF -name=part1 -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disk1.mia
fdisk -size=2 -type=P -unit=M -fit=BF -name=Part2 -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disk1.mia
fdisk -size=3 -type=E -unit=M -fit=BF -name=EXT1 -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disk1.mia  
fdisk -size=2 -type=P -unit=M -fit=WF -name=part3 -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disk1.mia
fdisk -size=1 -type=L -unit=M -fit=BF -name=LOG1 -path=/home/jgeraardi/Escritorio/MIA_2S_1P_201900532/BackEnd/Discos/Disk1.mia
