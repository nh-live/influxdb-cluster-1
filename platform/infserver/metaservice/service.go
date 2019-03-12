package metaservice

import (
	"github.com/influxdata/influxdb"
	"github.com/influxdata/influxdb/models"
	"net"
	"sync"
)

type MetaService struct {
	mu     sync.RWMutex
	dir    string
	store  *Store
	closed chan struct{}
}

func New(c *Config, name string, nodes map[string]influxdb.Node, ln net.Listener) *MetaService {
	ms := &MetaService{
		dir: c.Dir,
	}
	ms.store = NewStore(name, nodes, ms.dir, ln)
	return ms
}
func (ms *MetaService) Open() error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if err := ms.store.Open(); err != nil {
		return err
	}
	return nil
}
func (ms *MetaService) MapShard(points models.Points) (map[uint64]models.Points, error) {
	return nil, nil
}

func (ms *MetaService) GetDnsIDBySID(shardID uint64) (dnsID []uint64, err error) {
	return nil, nil
}

//func snapshot(path string, data *meta.Data) error {
//	file := filepath.Join(path, metaFile)
//	tmpFile := file + "tmp"
//
//	f, err := os.Create(tmpFile)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	var d []byte
//	if b, err := data.MarshBinary(); err != nil {
//		return err
//	} else {
//		d = b
//	}
//	if _, err := f.Write(d); err != nil {
//		return err
//	}
//	if err = f.Sync(); err != nil {
//		return err
//	}
//	if err = f.Close(); err != nil {
//		return err
//	}
//	return renameFile(tmpFile, file)
//}
//func renameFile(oldpath, newpath string) error {
//	if _, err := os.Stat(newpath); err != nil {
//		if err = os.Remove(newpath); nil != err {
//			return err
//		}
//	}
//	return os.Rename(oldpath, newpath)
//}
