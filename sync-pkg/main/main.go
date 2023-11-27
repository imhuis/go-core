package main

import sync_pkg "go-core/sync-pkg"

func main() {
	sync_pkg.Once()

	//var wg sync.WaitGroup
	//wg.Add(2)
	//var ints = make([]int, 0, 1000)
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		ints = append(ints, i)
	//	}
	//	wg.Done()
	//}()
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		ints = append(ints, i)
	//	}
	//	wg.Done()
	//}()
	//wg.Wait()
	//fmt.Println(len(ints))
}
