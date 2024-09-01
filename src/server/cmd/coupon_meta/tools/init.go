package tools

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/yitter/idgenerator-go/idgen"
	"log"
	"time"
)

var LocalCache *bigcache.BigCache

func Init() {
	idGen()
	bigCache()
}

func idGen() {
	var options = idgen.NewIdGeneratorOptions(1)
	options.SeqBitLength = 10
	idgen.SetIdGenerator(options)
}

func bigCache() {
	config := bigcache.Config{
		Shards:             2048,
		LifeWindow:         10 * time.Minute, // 对象生命时长
		CleanWindow:        5 * time.Minute,  // 间隔多久清理
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		OnRemove:           nil,
		OnRemoveWithReason: nil,
	}
	var err error
	LocalCache, err = bigcache.New(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
}
