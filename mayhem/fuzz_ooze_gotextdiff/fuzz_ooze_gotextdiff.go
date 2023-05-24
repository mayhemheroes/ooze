package fuzz_ooze_gotextdiff

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/gtramontina/ooze/internal/gotextdiff"
)

func mayhemit(data []byte) int {

    fuzzConsumer := fuzz.NewConsumer(data)
    
    var diff gotextdiff.GoTextDiff
    fuzzConsumer.GenerateStruct(&diff)

    a, _ := fuzzConsumer.GetString()
    b, _ := fuzzConsumer.GetString()
    aData := []byte(a)
    bData := []byte(b)

    diff.Diff(a, b, aData, bData)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}