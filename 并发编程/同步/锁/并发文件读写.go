package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
)

type Data []byte
type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

func main() {
	dataLen := uint32(3)
	path := filepath.Join(os.TempDir(), "data_file_test.txt")
	dataFile, _ := NewDataFile(path, dataLen)
	var wg sync.WaitGroup
	wg.Add(5)
	max := 10000
	// 写入
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			var prevWSN int64 = -1
			for j := 0; j < max; j++ {
				data := Data{
					byte(rand.Int31n(256)),
					byte(rand.Int31n(256)),
					byte(rand.Int31n(256)),
				}
				wsn, err := dataFile.Write(data)
				if err != nil {
					fmt.Println("错误1")
				}
				if prevWSN >= 0 && wsn <= prevWSN {
					fmt.Println("错误2")
				}
				prevWSN = wsn
			}
		}()
	}
	// 读取
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			var prevRSN int64 = -1
			for i := 0; i < max; i++ {
				rsn, data, err := dataFile.Read()
				if err != nil {
					fmt.Println("错误3")
				}
				if data == nil {
					fmt.Println("错误4")
				}
				if prevRSN >= 0 && rsn <= prevRSN {
					fmt.Println("错误5")
				}
				prevRSN = rsn
			}
		}()
	}
	wg.Wait()
	defer dataFile.Close()
	//fmt.Println(err)
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	woffset int64
	roffset int64
	wmutex  sync.RWMutex
	rmutex  sync.RWMutex
	dataLen uint32
}

func (m *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量。
	var offset int64
	m.rmutex.Lock()
	offset = m.roffset
	m.roffset += int64(m.dataLen)
	m.rmutex.Unlock()

	//读取一个数据块。
	rsn = offset / int64(m.dataLen)
	bytes := make([]byte, m.dataLen)
	for {
		m.fmutex.RLock()
		_, err = m.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				m.fmutex.RUnlock()
				continue
			}
			m.fmutex.RUnlock()
			return
		}
		d = bytes
		m.fmutex.RUnlock()
		return
	}
}

func (m *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量。
	var offset int64
	m.wmutex.Lock()
	offset = m.woffset
	m.woffset += int64(m.dataLen)
	m.wmutex.Unlock()

	//写入一个数据块。
	wsn = offset / int64(m.dataLen)
	var bytes []byte
	if len(d) > int(m.dataLen) {
		bytes = d[0:m.dataLen]
	} else {
		bytes = d
	}
	m.fmutex.Lock()
	defer m.fmutex.Unlock()
	_, err = m.f.Write(bytes)
	return
}

func (m *myDataFile) RSN() int64 {
	m.rmutex.Lock()
	defer m.rmutex.Unlock()
	return m.roffset / int64(m.dataLen)
}

// 获取最后写入的数据块的序列号
func (m *myDataFile) WSN() int64 {
	m.wmutex.Lock()
	defer m.wmutex.Unlock()
	return m.woffset / int64(m.dataLen)
}

func (m *myDataFile) DataLen() uint32 {
	return m.dataLen
}

func (m *myDataFile) Close() error {
	if m.f == nil {
		return nil
	}
	return m.f.Close()
}

func NewDataFile(path string, datalen uint32) (DataFile, error) {
	d, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if datalen == 0 {
		return nil, errors.New("长度为 0")
	}
	m := &myDataFile{f: d, dataLen: datalen}
	return m, nil
}
