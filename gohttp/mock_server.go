package gohttp

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

func FlushMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.GetMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) GetMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))

	return hex.EncodeToString(hasher.Sum(nil))
}

func (m *mockServer) GetMock(method, url, body string) *Mock {
	if !m.enabled {
		return nil
	}

	if mock := m.mocks[m.GetMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: errors.New(fmt.Sprintf("no mock matching %s from '%s' with given body", method, url)),
	}
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)

	if body == "" {
		return ""
	}

	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")

	return body
}
