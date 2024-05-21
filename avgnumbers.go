package main

import (
	"fmt"
)

func main() {
	fmt.Println(avgofnum([]int{10, 20, 30, 40, 50}))
	fmt.Println(checkage(-6))
	fmt.Printf("The largest number is : %v\n", largestofnum([]int{10, 20, 30, 40, 50}))
	var counter = counter{count:0}
	counter = counter.add()
	counter = counter.add()
	counter = counter.add()
    counter.display()
    counter = counter.sub()
    counter.display()
    counter = counter.reset()
    counter.display()
	

	fmt.Println(reversestr("Hell2o"))
}


func avgofnum(numbers []int) float64 {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += float64(numbers[i])
	}
	return sum / float64(len(numbers))
}

func checkage(age int) string {
	if age >= 0 && age < 150 {
		if age < 18 {
			return "Minor"
		} else if age >= 18 && age <= 22 {
			return "Young Adult"
		} else {
			return "Adult"
		}
	} else {
		return "Ivalid age"
	}
}

// func reversestr(word string) string{
// 	str := ""
// 	for i:=len(word)-1;i>=0;i--{
// 		str += string(word[i])
// 	}
// 	return str

// }
func reversestr(word string) string{
	revers := []rune(word)
	fmt.Println(revers)
	for i,j:=0,len(word)-1;i<j;i,j=i+1,j-1 {
			revers[i],revers[j] = revers[j], revers[i]
	}
	return string(revers)
}




func largestofnum(numbers []int) int {
	largest := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}

	}
	return largest

}

type counter struct{
	count int
}
func(c counter) add()counter{
	c.count = c.count + 1
	return c
}
func(c counter) sub()counter{
	c.count = c.count - 1
	return c
}
func(c counter) display()counter{
	fmt.Println(c.count)
	return c
}
func(c counter) reset()counter{
	c.count = 0
	return c
}
