package error

import (
	"fmt"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute

	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, error := http.Get(url)
		if error == nil {
			return nil
		}
		fmt.Printf("server not responding, (%s), retring", error)
		return error
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
