package main
import (
	"fmt"
	"Configure/ini"
)

func main(){
	fmt.Println(ini.GetPrivateProfileString("C", "D", "Test", "D:\\33e9_work\\Learn\\go\\SectionRead\\test.txt"))
}