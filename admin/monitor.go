package main 

import (
	"context"
	"net/http"
	"sync"
)

const (
	Unknown = iota
	Online
	Offline
)

type RegisterReq struct {
	name		string
	config 		*BackendConfig
}

type Monitor struct {
	backends		map[string]*Backend
	ctx 			context.Context
	cancel 			context.CancelFunc
	registerCh		chan RegisterReq
	unregisterCh 	chan string
	backendsCh		chan bool
}

func (m *Monitor) RegisterBackend(name string, b *Backend) {
	m.backends[name] = b
}

func (m *Monitor) UnRegisterBackend(name string) {
	delete(m.backends, name)
}

func (m *Monitor) GetStats(name string) [][]byte {
	if b, ok := m.backends[name]; ok {
		return b.metrics
	}
	return nil
}

func (m *Monitor) kickoff(ctx context.Context) {
	var wg sync.WaitGroup
	for _, backend := range m.backends {
		wg.Add(1)
		go func() {
			defer wg.Done()

			backend.startcheckHealth(ctx)
		}()
	}
	wg.Wait()
}

func (m *Monitor) work() {
	ctx, cancel := context.WithCancel(m.ctx)
	defer cancel()

	go m.kickoff(ctx)
	for {
		select {
			case <-m.backendsCh:
				cancel()
				go m.kickoff()
			case stat := <- statCh:

			case <-m.ctx.Done():
				return
		}	
	}
}

func (m *Monitor) Register(){
	for {
		select {
		case req := <-m.registerCh:
			backend := NewBackend(req.config)
			m.RegisterBackend(req.name, backend)
			m.backendsCh <- true
		case name :=<-m.unregisterCh:
			m.UnRegisterBackend(name)
		case <-m.ctx.Done():
			return
		}
	}
}

func (m *Monitor) StopMonitor() {
	close(m.registerCh)
	close(m.unregisterCh)
	close(m.backendsCh)
	m.cancel()
}

func NewMonitor() *Monitor {
	m := new(Monitor)

	m.backend = make(map[string]*Backend)
	m.registerCh = make(chan RegisterReq, 1)

	DefaultBackendConfig := NewBackendConfig(Options{Interval: DefaultInterval}, "127.0.0.1:9000/health", DefaultRequestTimeout)
	DefaultBackend := NewBackend(DefaultBackendConfig)
	m.RegisterBackend("default", DefaultBackend)

	m.ctx, m.cancel = context.WithCancel(context.Background())

	go m.work()
	go m.Register()
	return m
}