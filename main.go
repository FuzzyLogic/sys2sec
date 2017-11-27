package main

import(
    "fmt"
    "bufio"
    "os"
    "io"
    "errors"
    "regexp"
    "strings"
)

func main() {
    // Read command line arguments
    if len(os.Args) != 2 {
        panic(errors.New("Incorrect number of command line arguments"))
    }
    trigger := os.Args[1]

    // Compile regex to extract a syscall
    syscallExp, err := regexp.Compile(`(?P<syscall>[a-zA-Z0-9_]+)\(`)
    if err != nil {
        panic(err)
    }

    // Read input from STDIN, line by line
    input := bufio.NewReader(os.Stdin)
    triggerHit := false
    syscallSlice := []string{}
    for {
        line, err := input.ReadString('\n')
        if err == io.EOF {
            break
        }

        // Wait for trigger and capture subsequent syscalls
        if triggerHit == false {
            if strings.Contains(line, trigger) == true {
                triggerHit = true
                fmt.Println("Triggered on: " + line)
            }
        } else {
            matchSyscall := syscallExp.FindStringSubmatch(line)
            isNew := true
            if len(matchSyscall) > 1 {
                // Add syscall if not already in syscallSlice
                for _, sc := range syscallSlice {
                    if sc ==  matchSyscall[1] {
                        isNew = false
                        break
                    }
                }

                if isNew == true {
                    syscallSlice = append(syscallSlice, matchSyscall[1])
                }
            }
        }
    }

    // Output syscall list
    for _, sc := range syscallSlice {
        fmt.Println(sc)
    }
}
