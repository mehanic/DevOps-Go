package main

import (
    "fmt"
)
//We preallocate the slice result with the capacity equal to the length of the dictionary. This helps
// in avoiding multiple allocations when appending keys.
// keys extracts the keys from the map.
func keys(dictionary map[string]interface{}) []string {
    result := make([]string, 0, len(dictionary))

    // Extract keys
    for k := range dictionary {
        result = append(result, k)
    }

    return result
}

func main() {
    myDict := map[string]interface{}{
        "name": "Sebastian",
        "age":  21,
    }

    keysResult := keys(myDict)

    fmt.Println(keysResult)
}

//[name age]