package csvmutex

import (
	"encoding/csv"
	"io"
	"sync"

	"github.com/carbocation/pfx"
)

type MutexCSV struct {
	mutex  *sync.Mutex
	Writer *csv.Writer
}

func NewMutexCSV(writer io.Writer) *MutexCSV {
	return &MutexCSV{
		Writer: csv.NewWriter(writer),
		mutex:  &sync.Mutex{},
	}
}

func (c *MutexCSV) Write(row []string) error {
	c.mutex.Lock()
	if err := c.Writer.Write(row); err != nil {
		c.mutex.Unlock()
		return pfx.Err(err)
	}
	c.mutex.Unlock()

	return nil
}

func (c *MutexCSV) Flush() {
	c.mutex.Lock()
	c.Writer.Flush()
	c.mutex.Unlock()
}
