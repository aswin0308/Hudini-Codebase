package main
import "strconv"
import "unicode"
import "fmt"

func main(){
	fmt.Printf("The sum of integers is %d,\n",sumofintegersinstring("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))

}

func sumofintegersinstring(str string) int{

 str2 := []rune(str)
 numberinstring := ""
 convertednum := 0
 sum :=0

 for i:=0;i<len(str2);i++ {
	if(unicode.IsDigit(str2[i])){
		numberinstring = string(str2[i])
	}else{
		if(string(str2[i])==""){
			convertednum,_ = strconv.Atoi(numberinstring)
			sum = sum + convertednum 
		}
	}
	convertednum,_ = strconv.Atoi(numberinstring)
	sum = sum + convertednum 
 }
return sum
}