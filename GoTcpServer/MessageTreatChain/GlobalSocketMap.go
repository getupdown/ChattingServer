package MessageTreatChain

import (
	"net"
	"sync"
	"errors"
	"fmt"
)

/*
When the ip address of the online users extremely share the prefix.
Trie-tree may work better than Map.
 */
type GlobalSocketMap struct {
	Socketmap map[string]net.Conn
	Mylock sync.Mutex
}


func (mp *GlobalSocketMap) Get(key string) (conn net.Conn, err error) {
	mp.Mylock.Lock()
	fmt.Println(key)
	conn, ok := mp.Socketmap[key]
	if !ok {
		err = errors.New("Ip Not Found!")
	}
	mp.Mylock.Unlock()
	return conn, err
}


func (mp *GlobalSocketMap) Set(key string, value net.Conn) {
	mp.Mylock.Lock()
	mp.Socketmap[key] = value
	mp.Mylock.Unlock()
}

