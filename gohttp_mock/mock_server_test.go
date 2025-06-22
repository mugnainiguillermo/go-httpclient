package gohttp_mock

import (
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestMockServer_IsEnabledConcurrentAccess(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				_ = MockupServer.IsEnabled()
			}
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			MockupServer.Start()
		} else {
			MockupServer.Stop()
		}
	}
	wg.Wait()
}

func TestHttpClientMockConcurrentMapAccess(t *testing.T) {
	MockupServer.DeleteMocks()
	MockupServer.Start()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			for j := 0; j < 10; j++ {
				req, _ := http.NewRequest(http.MethodGet, "https://example.com", strings.NewReader(""))
				_, _ = MockupServer.GetMockedClient().Do(req)
			}
			wg.Done()
		}(i)
	}

	for i := 0; i < 50; i++ {
		MockupServer.AddMock(Mock{Method: http.MethodGet, Url: "https://example.com"})
		MockupServer.DeleteMocks()
	}

	wg.Wait()
	MockupServer.Stop()
}
