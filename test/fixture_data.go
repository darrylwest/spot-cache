package test

import (
    "fmt"
    "math/rand"
    "time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func CreateRandomId() string {

    return fmt.Sprintf("%x%x", random.Intn(9e7)+1e8, random.Intn(9e7)+1e8)
}

func CreateCommandId() []byte {
    id := []byte(CreateRandomId())

    return id
}


