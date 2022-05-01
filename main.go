package main

import(
    "fmt"
)

func main(){
    var array []int32 = []int32{1,2,3,4,5}
    var result = doThing(array)
    fmt.Println(result)


}

func doThing(array []int32) int32{
    var operation int32 = 1
    var score int32 = 0
    var currentSum int32 = 0
    for _, value := range array {
        // Sum up array
        currentSum += value
    }

    var leftIndex int = 0
    var rightIndex int = len(array) - 1

    for leftIndex <= rightIndex {
        if operation % 2 == 0 {
            // Current operation is even so we subtract
            score -= currentSum
            if array[leftIndex] <= array[rightIndex] {
                // If the next operation is going to be odd AND the left value is lower than Right we can get rid of the left value
                //Get rid of left indexed value
                currentSum -= array[leftIndex] 
                leftIndex++ 
            } else {
                // Else if the next operation is odd and AND the leftValue is higher than the right we get rid of the right
                currentSum -= array[rightIndex]
                rightIndex--
            } 
        } else {
            // current operation is odd so we add
            score += currentSum
            if array[leftIndex] <= array[rightIndex] {
                // If the operation is odd the next is even. So we want to get rid of the larger number
                currentSum -= array[rightIndex]
                rightIndex--
            } else {
                currentSum -= array[leftIndex]
                leftIndex++
            }
        }
        operation++
    }
    
    return score
}
