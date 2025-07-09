package user

import "sync"

type repository struct {
	data map[string]*repoModel.User
	m    sync.RWMutex
}
