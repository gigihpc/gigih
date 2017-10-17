package main

import(
  "fmt"
)

func main()  {
    var input =[]int{1,4,5,6,8,2}
    verticalChart(input)
}
func verticalChart(input []int)  {
  var maxval int
  maxval = getMaxval(input)
  drawChart(input,maxval)
}

func getMaxval(input []int) int  {
  var index,maxval, temp int
  maxval = input[0]

  for index = 0;  index< len(input); index++ {
    temp = input[index]
    if maxval < temp {
      maxval = temp
    }
  }
  return maxval
}

func drawChart(input []int,maxVal int) {
  var result [8][8]string
  var in,i int

  for in = 0;  in< len(input); in++ {
    var val int
    val = input[in]
    for i = maxVal;  i>0 ; i-- {
      if val > 0 {
        result[i-1][in] = "|"
        val--
      }else{
        result[i-1][in] = " "
      }
    }
  }

  for i = 0;  i<maxVal ; i++ {
    for in := 0;  in< len(input); in++ {
      fmt.Print(result[i][in])
    }
    fmt.Print("\n")
  }
  for in := 0;  in< len(input); in++ {
    fmt.Print(input[in])
  }
  fmt.Print("\n")
}
