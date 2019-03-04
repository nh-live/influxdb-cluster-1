package metaservice

import (
	"github.com/influxdata/influxdb/models"
	"github.com/influxdata/influxdb/models/meta"
	"os"
	"path/filepath"
	"sync"
)

const (
	metaFile = "meta.db"
)

type MetaService struct {
	mu    sync.RWMutex
	path  string
	store Store
}

func New(c *Config) *MetaService {
	return &MetaService{
		path: c.Dir,
	}
}
func (ms *MetaService) Open() error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if ms.store.cacheData.Index == 1 {
		if err := snapshot(ms.path, ms.store.cacheData); err != nil {
			return err
		}
	}
	return nil
}
func (ms *MetaService) MapShard(points models.Points) (map[uint64]models.Points, error) {
	return nil, nil
}

func (ms *MetaService) GetDnsIDBySID(shardID uint64) (dnsID []uint64, err error) {
	return nil, nil
}

func snapshot(path string, data *meta.Data) error {
	file := filepath.Join(path, metaFile)
	tmpFile := file + "tmp"

	f, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer f.Close()

	var d []byte
	if b, err := data.MarshBinary(); err != nil {
		return err
	} else {
		d = b
	}
	if _, err := f.Write(d); err != nil {
		return err
	}
	if err = f.Sync(); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return renameFile(tmpFile, file)
}
func renameFile(oldpath, newpath string) error {
	if _, err := os.Stat(newpath); err != nil {
		if err = os.Remove(newpath); nil != err {
			return err
		}
	}
	return os.Rename(oldpath, newpath)
}
