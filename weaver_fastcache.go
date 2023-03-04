package weaverutils

import (
	"context"
	"sync"

	"github.com/ServiceWeaver/weaver"
	"github.com/VictoriaMetrics/fastcache"
	"github.com/cespare/xxhash/v2"
)

const WeaverFastCacheDefaultSize = 32 * 1 << 20 // 32MB

// WeaverFastCache is a FastCache component for Service Weaver.
//
// https://github.com/VictoriaMetrics/fastcache
//
// Configuration Options:
//
// ["gopkg.eu.org/weaverutils/WeaverFastCache"]
//
// Size = 33554432
type WeaverFastCache interface {
	Get(_ context.Context, key []byte) ([]byte, error)
	Set(_ context.Context, key []byte, value []byte) error
	Del(_ context.Context, key []byte) error
	Has(_ context.Context, key []byte) (bool, error)
	HasGet(_ context.Context, key []byte) ([]byte, bool, error)
}

type weaverFastCacheOptions struct {
	Size int
}

var _ WeaverFastCache = (*weaverFastCache)(nil)

type weaverFastCache struct {
	weaver.Implements[WeaverFastCache]
	weaver.WithConfig[weaverFastCacheOptions]
	weaver.WithRouter[weaverFastCacheRouter]
	o sync.Once
	c *fastcache.Cache
}

func (w *weaverFastCache) doInit() {
	size := w.Config().Size
	if size == 0 {
		size = WeaverFastCacheDefaultSize
	}
	w.c = fastcache.New(size)
}

func (w *weaverFastCache) Get(_ context.Context, key []byte) ([]byte, error) {
	w.o.Do(w.doInit)
	return w.c.Get(nil, key), nil
}

func (w *weaverFastCache) Set(_ context.Context, key []byte, value []byte) error {
	w.o.Do(w.doInit)
	w.c.Set(key, value)
	return nil
}

func (w *weaverFastCache) Del(_ context.Context, key []byte) error {
	w.o.Do(w.doInit)
	w.c.Del(key)
	return nil
}

func (w *weaverFastCache) Has(_ context.Context, key []byte) (bool, error) {
	w.o.Do(w.doInit)
	return w.c.Has(key), nil
}

func (w *weaverFastCache) HasGet(_ context.Context, key []byte) ([]byte, bool, error) {
	w.o.Do(w.doInit)
	v, ok := w.c.HasGet(nil, key)
	return v, ok, nil
}

type weaverFastCacheRouter struct{}

func (weaverFastCacheRouter) Get(_ context.Context, key []byte) uint64 { return xxhash.Sum64(key) }
func (weaverFastCacheRouter) Set(_ context.Context, key []byte, _ []byte) uint64 {
	return xxhash.Sum64(key)
}
func (weaverFastCacheRouter) Del(_ context.Context, key []byte) uint64    { return xxhash.Sum64(key) }
func (weaverFastCacheRouter) Has(_ context.Context, key []byte) uint64    { return xxhash.Sum64(key) }
func (weaverFastCacheRouter) HasGet(_ context.Context, key []byte) uint64 { return xxhash.Sum64(key) }
