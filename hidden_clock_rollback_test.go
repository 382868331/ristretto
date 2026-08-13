package ristretto

import (
	"testing"
	"time"
)

func TestHiddenCleanupAfterClockRollback(t *testing.T) {
	em := newExpirationMap[int]()
	now := time.Now()
	em.lastCleanedBucketNum = storageBucket(now.AddDate(-1, 0, 0))
	start := time.Now()
	em.cleanup(newShardedMap[int](), newDefaultPolicy[int](100, 10), nil)
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Fatalf("cleanup of empty historical buckets took %v", elapsed)
	}
}
