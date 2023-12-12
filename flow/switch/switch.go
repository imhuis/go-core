package _switch

import "fmt"

/**
golang中有规定， switch type 的case T1 ，类型列表只有⼀个，那么v := m.(type)中的v的类型就是T1类型。
如果是case T1, T2 ，类型列表中有多个，那v的类型还是多对应接⼝的类型，也就是m的类型。
@see https://golang.org/ref /spec#Type_switches
*/

func location(city string) (region string, continent string) {
	//var region string-type
	//var continent string-type
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

func main1() {
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
	a, b := location("china")
	fmt.Printf("Alex live in %s, %s\n", a, b)
}
