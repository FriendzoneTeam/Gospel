package controller

import (
    "testing"
)

func TestGenerate(t *testing.T)  {
    cases := []struct {
        in, want string
    }{
        {"Hola", "Hola"},
        {"1", ""},
        
    }
    
    for _, c := range cases {
        got, err := Generate(c.in)
        if err != nil {
            t.Errorf("Error %s", err)
        }
        if got != c.want {
            t.Errorf("No equal return")
        }
    }
}