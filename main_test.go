package main

import "testing"
import "os"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 11

    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
    
    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }

}