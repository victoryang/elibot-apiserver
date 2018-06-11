package pool

import (
	"errors"
	"fmt"
	"sync"
)

//PoolConfig
type PoolConfig struct {
	//initial connection number
	InitialCap int
	//max connection number
	MaxCap int
	//create a new conn
	Factory func() (interface{}, error)
	//close func
	Close func(interface{}) error
}

//ChannelPool
type ChannelPool struct {
	mu          sync.Mutex
	conns       chan *Conn
	factory     func() (interface{}, error)
	close       func(interface{}) error
}

type Conn struct {
	sync.RWMutex
	conn 		interface{}
}

//NewChannelPool
func NewChannelPool(poolConfig *PoolConfig) (Pool, error) {
	if poolConfig.InitialCap < 0 || poolConfig.MaxCap <= 0 || poolConfig.InitialCap > poolConfig.MaxCap {
		return nil, errors.New("invalid capacity settings")
	}

	c := &ChannelPool{
		conns:       make(chan *Conn, poolConfig.MaxCap),
		factory:     poolConfig.Factory,
		close:       poolConfig.Close,
	}

	for i := 0; i < poolConfig.InitialCap; i++ {
		conn, err := c.factory()
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &Conn{conn: conn}
	}

	return c, nil
}

//getConns
func (c *ChannelPool) getConns() chan *Conn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

//Get a conn from Pool
func (c *ChannelPool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, ErrClosed
	}

	select {
	case wrapConn := <-conns:
		if wrapConn == nil {
			return nil, ErrClosed
		}
		
		return wrapConn.conn, nil
	default:
		conn, err := c.factory()
		if err != nil {
			return nil, err
		}

		return conn, nil
	}
}

// Put conn back to pool
func (c *ChannelPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conns == nil {
		return c.Close(conn)
	}

	select {
	case c.conns <- &Conn{conn: conn}:
		return nil
	default:
		//pool is full, close it directly
		return c.Close(conn)
	}
}

//Close single one
func (c *ChannelPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.close(conn)
}

//Release free all connections
func (c *ChannelPool) Release() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.factory = nil
	closeFun := c.close
	c.close = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		closeFun(wrapConn.conn)
	}
}

//Len
func (c *ChannelPool) Len() int {
	return len(c.getConns())
}
