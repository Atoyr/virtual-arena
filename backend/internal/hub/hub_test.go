// internal/hub/hub_test.go
package hub

import (
	"testing"
	"time"
)

// モッククライアント：send チャネルに届いたメッセージを記録
type mockClient struct {
	send chan []byte
}

func newMockClient() *mockClient {
	return &mockClient{send: make(chan []byte, 1)}
}

func (m *mockClient) Send(msg []byte) {
	m.send <- msg
}

func (m *mockClient) Receive(timeout time.Duration) ([]byte, bool) {
	select {
	case b := <-m.send:
		return b, true
	case <-time.After(timeout):
		return nil, false
	}
}

func TestHub_Broadcast(t *testing.T) {
	h := NewHub()
	go h.Run()

	// モッククライアントを2つ登録
	c1 := newMockClient()
	c2 := newMockClient()
	h.register <- c1
	h.register <- c2

	// 少し待って登録完了を確実に
	time.Sleep(10 * time.Millisecond)

	// ブロードキャスト
	payload := []byte("hello world")
	h.broadcast <- payload

	// 各クライアントに届くことをチェック
	if msg, ok := c1.Receive(100 * time.Millisecond); !ok || string(msg) != "hello world" {
		t.Errorf("client1 did not receive broadcast, got=%v ok=%v", msg, ok)
	}
	if msg, ok := c2.Receive(100 * time.Millisecond); !ok || string(msg) != "hello world" {
		t.Errorf("client2 did not receive broadcast, got=%v ok=%v", msg, ok)
	}

	// クライアント解除
	h.unregister <- c1
	time.Sleep(10 * time.Millisecond)

	// 再度ブロードキャストして c1 に届かないことを確認
	h.broadcast <- []byte("second")
	if _, ok := c1.Receive(50 * time.Millisecond); ok {
		t.Error("client1 should not receive after unregister")
	}
	if msg, ok := c2.Receive(100 * time.Millisecond); !ok || string(msg) != "second" {
		t.Errorf("client2 did not receive second broadcast, got=%v ok=%v", msg, ok)
	}
}
