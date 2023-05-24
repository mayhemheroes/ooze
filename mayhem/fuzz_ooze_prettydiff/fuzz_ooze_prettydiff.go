package fuzz_ooze_prettydiff

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/gtramontina/ooze/internal/prettydiff"
)

func mayhemit(data []byte) int {

    fuzzConsumer := fuzz.NewConsumer(data)

    var diff prettydiff.PrettyDiff
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