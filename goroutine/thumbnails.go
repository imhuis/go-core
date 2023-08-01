package goroutine

import "fmt"

func makeThumbnail1(filenames []string) {
	for _, f := range filenames {
		if _, err := imageFile(f); err != nil {
			fmt.Println(err)
		}
	}
}

func makeThumbnail2(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func() {
			imageFile(f)
			ch <- struct{}{}
		}()
	}
	for range filenames {
		<-ch
	}
}

func imageFile(f string) (string, error) {
	return f, nil
}
