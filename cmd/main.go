package main

import (
	"fmt"
    "NotesServer/internal/bd"
)

func main() {
	a := [...]int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(a))
}

func removeDuplicates(nums [6]int) int {
    bd.Connect()
    k := 1
    reps := 0
    fmt.Println(nums)
    for i := 1; i < len(nums); i++ {
        fmt.Printf("i: %d, k: %d, reps: %d\n", i, k, reps)
        if nums[k - 1] == nums[i] {
            reps += 1
            if reps < 2 {
                k++
            }
        } else {
            if reps >= 1 {
                tmp := nums[k]
                nums[k] = nums[i]
                nums[i] = tmp
                k++
            }
            reps = 0
        }
        fmt.Printf("i: %d, k: %d, reps: %d\n", i, k, reps)
        fmt.Println(nums)
    }
    // fmt.Println(nums)
    return k
}