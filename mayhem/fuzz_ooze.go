package fuzz_ooze

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/gtramontina/ooze"
    "github.com/gtramontina/ooze/internal/gotextdiff"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            case 0:
                root, _ := fuzzConsumer.GetString()

                ooze.WithRepositoryRoot(root)
                return 0

            case 1:
                command, _ := fuzzConsumer.GetString()

                ooze.WithTestCommand(command)
                return 0

            case 2:
                temp, _ := fuzzConsumer.GetInt()
                threshold := float32(temp)

                ooze.WithMinimumThreshold(threshold)
                return 0

            case 3:
                pattern, _ := fuzzConsumer.GetString()

                ooze.IgnoreSourceFiles(pattern)
                return 0

            case 4:
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
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}