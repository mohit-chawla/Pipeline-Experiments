package main

import "testing"
import "os"
import "strings"
// import "fmt"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10
    found := false

    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        if pair[0] == "MY_ENV_VAR" {
        	t.Logf("yay! found %v", pair)
        	found = true
        }
        
    }
    if !found{
    	t.Errorf("Did not find MY_ENV_VAR")
    }

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }

}