package ristretto
import("testing";"time")
func TestHiddenWaitSurvivesConcurrentClear(t *testing.T){c,err:=NewCache(&Config{NumCounters:100,MaxCost:10,BufferItems:64});if err!=nil{t.Fatal(err)};defer c.Close();done:=make(chan struct{});go func(){for i:=0;i<10;i++{c.Wait()};close(done)}();for i:=0;i<10;i++{c.Clear()};select{case<-done:case<-time.After(time.Second):t.Fatal("Wait blocked after Clear discarded its marker")}}
