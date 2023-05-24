package fuzz_ooze_options

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/gtramontina/ooze"
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
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}