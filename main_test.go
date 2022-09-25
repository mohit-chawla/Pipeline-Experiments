package main

import "testing"
import "os"
import "strings"
import "fmt"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 11

    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        if pair[0] == "MY_ENV_VAR" {
        	fmt.Println(pair)	
        }
        
    }

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }

}