package context_pkg

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
)

var client = http.Client{}

func CallBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	fmt.Println("done with both")

}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request failed", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "request failed", err)
		return err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read body failed", err)
		return err
	}
	result := string(bytes)
	fmt.Println(label, "result", result)
	if result == "error" {
		return fmt.Errorf(label + " server returned an error")
	}
	return nil
}
