package persistence

import "sync"

type orderRepository struct {
	mu *sync.Mutex
}
