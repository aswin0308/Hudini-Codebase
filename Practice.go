package main

import (
	"fmt"
	"maps"
	"math"
)

func main(){
	var numItems int
    fmt.Print("Enter the number of items in the shopping cart: ")
    fmt.Scanln(&numItems)

    // Initialize an empty slice to store items
    cart := make([]shoppingcart, numItems)
	for i := 0; i < numItems; i++ {
        fmt.Printf("Enter details for item %d:\n", i+1)
        fmt.Print("Name: ")
        fmt.Scanln(&cart[i].name)
        fmt.Print("Price: ")
        fmt.Scanln(&cart[i].price)
        fmt.Print("Quantity: ")
        fmt.Scanln(&cart[i].quantity)
    }

    // Calculate the total cost of the shopping cart
    totalCost := totalcost(cart)

    // Print the total cost
    fmt.Printf("Total cost of the shopping cart: $%.2f\n", totalCost)
}
	


	// var things = []int{1,2,3,4}
	// fmt.Printf("The memory location of things is %p\n",&things)
	// v := square(&things)
	// fmt.Println(v)


	//merging maps
	// mymap1 := map[int]int{1:1,2:2,3:3}
	// mymap2 := map[int]int{4:4,5:5,6:6}
	// merge(mymap1,mymap2)
	// fmt.Printf("The new map after merging map1 and map2 is %v",mymap1)

	// neworder := shoppingcart{"rice",200,5,0}
	// neworder.total = totalcost(neworder)
	// fmt.Println(neworder)
	// }


    //finding max min
	// maxval,minval :=findmaxminel([]int{1,2,3,4,5,15,10,9})
	// fmt.Printf("THe max and min value of array are %v , %v",maxval,minval)

	//sorting using pointers
	// newsorted := sortasc([]int{1,2,3,4,5,15,10,9})
	// fmt.Printf("THe soreted array is %v",newsorted)


	//shopping cart
	

	type shoppingcart struct{
		name string
		price int
		quantity float64
		total float64
	}

	func totalcost(items []shoppingcart) float64{
		totalcost := 0
		for _,value := range items{
			totalcost += value.price * int(value.quantity)
		}
		return float64(totalcost)
	}


	// var a = []int{1,2,3,4}
	// copy := a
	// copy[3] = 5
	// fmt.Println(copy,a)
	// var mystring = "123456"
	// var str1 = []string{"s","u","b","s"}
	// str2 := ""
	// for i:= range str1{
	// 	str2 += str1[i]
	// }
	// fmt.Printf("%v\n",string(str2))
	// fmt.Printf("The index id %v \nand value %T",mystring[3],mystring[3])
	// for i,v := range mystring{
	// 	fmt.Println(i,v)
	// }

// }

func square(things *[]int) []int{
	for i:= range *things{
		(*things)[i] = (*things)[i]*(*things)[i]
	}
	fmt.Printf("The memory location of things is %p\n",things)
	return *things
}

func merge(mymap map[int]int, mymap2 map[int]int){
	maps.Copy(mymap,mymap2)
}
func findmaxminel(arr []int) (int,int){
	maxnum := math.MinInt
	minnum := math.MaxInt
	for i:= range arr{
		if arr[i]<minnum {
			minnum = arr[i]
		}else if arr[i]>maxnum {
			maxnum = arr[i]
		}
	}
	return maxnum,minnum

	}

	func sortasc(arr []int)([]int){
     for i:= range arr{
		for j:= range arr{
			if(i == j){
				continue
			}else{
             if (arr)[j]>(arr)[i]{
				 arr[i],arr[j] = arr[j],arr[i]
			 }
			}
		}
	 }
	 return arr
	}