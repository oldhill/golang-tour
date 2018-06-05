// https://tour.golang.org/moretypes/23

package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    result := make(map[string]int)
    count := 0
    for _, word := range words {
        for _, candidate := range words {
            if word == candidate {
                count += 1
            }
        }
        result[word] = count
        count = 0
    }
    return result
}

func main() {
    wc.Test(WordCount)
}
