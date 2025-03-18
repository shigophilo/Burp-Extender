package s

import (
	"sync"
)

var (
	HostPath    string
	ExcludeHost []string
	lock        sync.RWMutex
	Wg          sync.WaitGroup
)
