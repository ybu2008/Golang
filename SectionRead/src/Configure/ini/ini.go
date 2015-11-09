/*************************************************************************
  > File Name: ini.go
  > Author: Zhong Baihong
  > Mail: zbaihong@live.com
  > Created Time: 2015-11-09 11:13
 ************************************************************************/

package ini
import (
	"strings"
	"io"
	"bufio"
	"os"
)

func isTitle(line, title, delim string)bool{
	bRet := false
	cmpTitle := "[" + title + "]" + delim

	if(strings.EqualFold(line, cmpTitle)){
		bRet = true
	}
	return bRet
}

func decideDelim(line string)string{
	delim := ""
	length := len(line)
	

	if 2 <= length && '\r' == line[length - 2]{
		delim = "\r"
	}
	
	if 1 <= length && '\n' == line[length - 1]{
		delim += "\n"
	}
	
	return delim
}

func getValueByKey(line, key string)(string, bool){
	bRet := false
	strRet := ""
	item := strings.SplitN(line, "=", 2)

	if 2 == len(item) && strings.EqualFold(item[0], key){	
		strRet = item[1]
		bRet = true
	}
	return strRet,bRet
}

/*************************************************************************
  > Function: GetPrivateProfileString
  > Description: Retrieves a string from the specified section in an initialization file.
  > Param 1: IN, string. The name of the section.
  > Param 2: IN, string. The name of the key.
  > Param 3: IN, A default string.
  > Param 4: The name of the initialization file
  > Return : The retrieved string.
 ************************************************************************/
func GetPrivateProfileString(appName, keyName, defValue, filePath string)string{
	bFlag :=  false
	retValue := defValue
	src, err := os.Open(filePath)
	
	if nil == err{
		defer src.Close()
		
		delim := ""
		r := bufio.NewReader(src)
		for{
			buf, err := r.ReadString('\n')
			if io.EOF == err{
				break
			}
			
			if 0 == len(delim){
				delim = decideDelim(buf)
			}
			
			if 0 == len(delim){
				continue
			}
			
			if !bFlag && isTitle(buf, appName, delim){
				bFlag = true;
			}else if bFlag{
				value, retFlag := getValueByKey(buf, keyName)
				if(retFlag){
					retValue = value
					break;
				}
			}
		}
	}
	return retValue
}


