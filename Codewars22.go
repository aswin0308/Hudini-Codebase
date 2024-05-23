// import "fmt"

// func main(){
// 	num := FindOdd([]int{5,1,5,6,10,10,1})
// 	fmt.Println(num)
// }

// func FindOdd(seq []int) int {
// 	mapnum := make(map[int]int)
// 	for _,value := range seq{
// 	  mapnum[value]++
// 	}
// 	for value, count := range mapnum {
// 		  if(count%2!=0){
// 		return value
// 	  }
// 	  }
// 	return 0
// }

// package main

// import ("fmt")

// func main(){
// 	result:= GetSum(1,-3)
// 	fmt.Println(result)
// }
// func GetSum(a, b int) int {
// 	sum := 0
// 	if(a==b){
// 	  return a
// 	}else if b>a{
// 	for i:=a;i<=b;i++ {
// 	sum += i
// 	  }
// 	  return sum
//   }else{
// 	for i:=a;i>=b;i--{
// 		sum += i
// 		  }

//   }
// 	  return sum
// 	}

// package main

// import("strconv"
//        "fmt")

// func Grille(message string, code int) string {
//   newmsg:= ""
//   numstr := strconv.FormatInt(int64(code), 2)
//   for len(numstr)<len(message) {
//     numstr = "0"+numstr
//   }
//   for len(numstr)>len(message){
//     message = "#"+message
//   }

//   for i:= range numstr{
//     if(numstr[i]=='1' && message[i]!= '#'){
//       newmsg += string(message[i])
//     }

//   }
// 	return newmsg
// }

// func main(){
// 	result := Grille("",5)
// 	fmt.Println(result)
// }

// package main

// import "fmt"
// func SequenceSum(start, end, step int) int {
// 	sum:=0
// 	if(start>end){
// 	  return 0
// 	}else{
// 	for i:=start;i<=end;i+=step{
// 	  sum+=i
// 	}
// 	  return sum
// 	  }
//   }

//   func main(){
// 	result := SequenceSum(1,6,1)
// 	fmt.Println(result)
//   }

// Multiples of 3 or 5

// func Multiple3And5(number int) int {
// 	sum:=0
// 	for i := 1; i<number; i++{
// 	  if(i%3 == 0 && i%5==0 ){
// 		sum += i
// 	  }else if(i%3 == 0){
// 		sum += i
// 	  }else if(i%5 == 0){
// 		sum += i
// 	  }
// 	}
// 	return sum
//   }

// package main

// import "math"

// import ("fmt"
//         "math")

// func FindNextPower(val, pow int) int {
// 	num:= 0

// iterate from 1 to val
// calculate cube of i every turn
// compare if cube of i is greater than value
// if yes,
// return cube of i

// 	for i:=1;i<val;i++ {
// 		j :=1
// 		num:= 1
// 		for j<=pow{
// 				num *= i
// 				j++
// 					}
// 	  if(num > val){
// 		return num
// 	  }
// 	}
// 	return num
//   }

//	  func main(){
//		result := FindNextPower(12385, 3)
//		fmt.Println(result)
//		fmt.Println("Cube root of is", math.Pow(25, 1.0/3.0))
//	  }
package main

import "fmt"

// import("fmt"
//          "math")

//   func FindNextPower(val, pow int) int {
// 	num := 1.0
// 	root := math.Pow(float64(val), 1.0/float64(pow))
// 	upper := math.Ceil(root)
// 	j:=1
// 	for j<=pow{
// 				num *= upper
// 				j++
// 			  }
// 	return int(num)
//     }
//   func main(){
// 		result := FindNextPower(1185921, 4)
// 		fmt.Println(result)
// 	  }



func FindEvenIndex(arr []int) int {
  for i:=0; i<len(arr); i++{
    sum:=0
    for j := 0; j < i; j++ {
            sum += arr[j]
        }
  sum2:=0
  for j:=i+1; j<len(arr); j++{
    sum2+= arr[j]
  }
    if(sum == sum2){
      return i
    }
    }
  return -1
}

func main(){
	result:= FindEvenIndex([]int{8, 0})
	fmt.Println(result)
}
	  