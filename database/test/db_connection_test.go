package test

import (
	"blog/database"
	"sync"
	"testing"
)

func TestGetBlogDBConnection(t *testing.T) {
	const c = 100
	wg := sync.WaitGroup{}
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()
			database.GetBlogDBConnection()
		}()
	}
	wg.Wait()
}
// go test -v ./database/test -run=^TestGetBlogDBConnection$ -count=1