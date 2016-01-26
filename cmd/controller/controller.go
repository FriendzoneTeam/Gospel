package controller

import (
    "errors"
)
func Generate(name string)(string, error){
    if name != "" {
        return name, nil
    }
    return "", errors.New("Nombre vacio")
}