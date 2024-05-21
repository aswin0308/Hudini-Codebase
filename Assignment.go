package kata

func monkeyCount(n int) []int {
  count := make([]int,n)
     j:=1
   for i:=0;i<n;i++ {
     count[i] = j
     j++
   }
  return count
   
}

package kata

func MakeNegative(x int) int {
  if(x<0 || x==0){
    return x
  }else{
    return -x
  }
  return 0
}

package kata

func FindMultiples(integer, limit int) []int {
  var multiples []int
    for i:=integer;i<=limit;i+=integer{
      if(i%integer==0){
        multiples = append(multiples,i)
      }
    }   
   return multiples
  }

  package kata


func CountBy(x, n int) []int {
  var count []int
  var a int
  for i:=1;i<=n;i++ {
    a = x* i
    count = append(count,a)
  }
  return count
}

package kata
import "math"

func PowersOfTwo(n int) []uint64 {
  var powers []uint64
  for i:=0;i<=n;i++{
    a := math.Pow(2,float64(i))
    powers = append(powers,uint64(a))
    
  }
  return powers
}

package kata

func Points(games []string) int {
  var score = 0
  var tmp string
  for i:=0;i<len(games);i++{
    tmp = games[i]
    if(tmp[0]>tmp[2]){
      score +=3
    }else if(games[i][0]==games[i][2]){
      score+=1
    }else{
      continue
    }
  }
  return score
}


package kata


func GetSum(a, b int) int {
  sum:=0
  if(a == b){
    return a
  }else if(a>b){
    for i:=a;i>=b;i-- {
    sum = sum + i
  }
    }else{
    for i:=a;i<=b;i++ {
    sum = sum + i
      }
}
  return sum
  }

  package kata
import("strings")

func FindShort(s string) int {
  sen2 := strings.Split(s," ")
  short := len(sen2[0])
  for i:=1;i<len(sen2);i++ {
    if len(sen2[i])<short {
      short = len(sen2[i])
    }
  }
  return short
}
// problem 11
package kata

func SquareSum(numbers []int) int {
  sum:=0
    for i:=0;i<len(numbers);i++ {
      sum+= numbers[i]* numbers[i]
    }
  return sum
}