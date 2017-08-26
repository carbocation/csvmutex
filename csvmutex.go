package csvmutex

import (
	"encoding/csv"
	"io"
	"sync"

	"github.com/carbocation/pfx"
)

type CSVMutex struct {
	mutex  *sync.Mutex
	Writer *csv.Writer
}

func NewCSVMutex(writer io.Writer) *CSVMutex {
	return &CSVMutex{
		Writer: csv.NewWriter(writer),
		mutex:  &sync.Mutex{},
	}
}

func (c *CSVMutex) Write(row []string) error {
	c.mutex.Lock()
	if err := c.Writer.Write(row); err != nil {
		c.mutex.Unlock()
		return pfx.Err(err)
	}
	c.mutex.Unlock()

	return nil
}

func (c *CSVMutex) Flush() {
	c.mutex.Lock()
	c.Writer.Flush()
	c.mutex.Unlock()
}
