package ristretto
import("testing";"time")
func TestHiddenStoreClearDropsExpiryIndex(t *testing.T){s:=newShardedMap[int]();it:=&Item[int]{Key:1,Conflict:2,Value:3,Expiration:time.Now().Add(time.Hour)};s.Set(it);if len(s.expiryMap.buckets)==0{t.Fatal("setup did not index expiration")};s.Clear(nil);if len(s.expiryMap.buckets)!=0{t.Fatalf("expiration index retained %d buckets",len(s.expiryMap.buckets))}}
