package main

import (
	"fmt"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func splitArray(array []string, num int64) [][]string {
	if int64(num) > int64(len(array)) {
		return nil
	}
	result := make([][]string, 0)
	// ---->>>>> 5 1
	// i!=num [sha256 v1 createtime expiretime true]
	// END 4
	// ---->>>>> 10 2
	// i!=num [sha256 v2 createtime expiretime false]
	// END 8
	// ---->>>>> 15 3
	// i=num [sha256 v2 createtime expiretime false]
	// END 12
	// Average number of elements
	averNumElements := int64(len(array)) / num
	steps := int64(0)
	for i := int64(1); i <= num; i++ {
		maxIterEm := i * averNumElements
		if i != num {
			result = append(result, array[i-1+steps:maxIterEm])
		} else {
			result = append(result, array[i-1+steps:])
		}
		steps = maxIterEm - i
	}
	return result
}

func main() {
	data := map[string][]string{"T21": []string{"T21", "sha256", "v1", "createtime", "expiretime", "false", "T21", "sha256", "v2", "createtime", "expiretime", "false", "T21", "sha256", "v2", "createtime", "expiretime", "true"},
		"T22": []string{"T22", "sha256", "v2", "createtime", "expiretime", "true"},
		"T23": []string{"T23", "sha256", "v3", "createtime", "expiretime", "false"},
		"T24": []string{"T24", "sha256", "v4", "createtime", "expiretime", "true"},
	}
	isExpired := "true"
	eNums := 6
	tableHeader := []string{}
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	// fmt.Printf("%#v\n", data)
	for _, v := range data {
		if cap(v) == eNums {
			// if v[5] == "true" {
			// 	table.Append(v)
			// }
			if v[len(v)-1] == isExpired {
				table.Append(v)
			}
		}
		if cap(v) > eNums {
			count := cap(v) / eNums
			result := splitArray(v, int64(count))
			for _, v2 := range result {
				// if v2[5] == "true" {
				// 	table.Append(v2)
				// }
				if v2[len(v2)-1] == isExpired {
					table.Append(v2)
				}
			}
		}

	}
	// fmt.Printf("%#v\n", res)
	// []string{"sha256", "v1", "createtime", "expiretime", "true"}
	// []string{"sha256", "v2", "createtime", "expiretime", "false"}
	// []string{"sha256", "v2", "createtime", "expiretime", "false"}
	table.SetRowLine(true)
	table.SetHeader(tableHeader)
	table.SetCaption(true, "Movie ratings.")
	table.Render() // Send output

	fmt.Println(tableString.String())
}
