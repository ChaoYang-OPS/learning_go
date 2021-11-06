package main

import "fmt"

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
	data := map[string][]string{"projectId1": []string{"sha256", "v1", "createtime", "expiretime", "true", "sha256", "v2", "createtime", "expiretime", "false", "sha256", "v2", "createtime", "expiretime", "false"},
		"projectId2": []string{"sha256", "v2", "createtime", "expiretime", "true"},
		"projectId3": []string{"sha256", "v3", "createtime", "expiretime", "false"},
		"projectId4": []string{"sha256", "v4", "createtime", "expiretime", "true"},
	}
	res := make([]string, 0)
	// fmt.Printf("%#v\n", data)
	for _, v := range data {
		// fmt.Printf("%#v----%#v\n", v, cap(v))
		// for _, k := range v {
		// 	// fmt.Printf("%#v---%#v\n", s, k)
		// 	tmp := strings.Split(k, " ")
		// 	fmt.Printf("%#v\n", tmp[:])
		// }
		// fmt.Printf("%#v\n", v[:5])
		if cap(v) == 5 {
			// fmt.Printf("%#v\n", v[4])
			if v[4] == "true" {
				// fmt.Printf("%#v\n", v)
				res = append(res, v...)
			}
		}
		if cap(v) > 5 {
			count := cap(v) / 5
			res := splitArray(v, int64(count))
			// fmt.Printf("%#v\n", res)
			for _, v1 := range res {
				fmt.Printf("%#v\n", v1)
			}
		}

	}
	// fmt.Printf("%#v\n", res)
	// []string{"sha256", "v1", "createtime", "expiretime", "true"}
	// []string{"sha256", "v2", "createtime", "expiretime", "false"}
	// []string{"sha256", "v2", "createtime", "expiretime", "false"}
}
