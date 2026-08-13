package ristretto
import "testing"
func TestHiddenPublicMethodsAfterClose(t *testing.T){c,err:=NewCache(&Config{NumCounters:100,MaxCost:10,BufferItems:64});if err!=nil{t.Fatal(err)};c.Close();defer func(){if r:=recover();r!=nil{t.Fatalf("public method panicked after Close: %v",r)}}();if c.Set("k","v",1){t.Fatal("Set accepted value after Close")};c.Get("k");c.Del("k");c.Clear();c.Wait();c.Close()}
