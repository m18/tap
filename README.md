# tap — a water tap idiom for flow control

Tap allows you to start and stop code execution.

When a tap is closed, the stream doesn't flow, and code isn't executing. When a tap is open, the stream flows, and code is executing.

## Example
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/m18/tap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a (closed) tap
	t := tap.New()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.Stream():
				// execute code when the stream is flowing
				fmt.Println("💧")
				time.Sleep(time.Second)
			}
		}
	}()

	fmt.Println("tap is closed")
	time.Sleep(2 * time.Second)

	fmt.Println("opening tap...")
	// let the stream flow
	t.Open()
	time.Sleep(3 * time.Second)

	// stop the stream flow
	t.Close()
	fmt.Println("tap is closed")
	time.Sleep(2 * time.Second)

	fmt.Println("done")
}
```

Sample output:
```
tap is closed
opened tap
💧
💧
💧
closed tap
done
```
