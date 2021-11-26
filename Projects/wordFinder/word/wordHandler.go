package word

import (
    "fmt"
    "strings"
)

const body = "" + "If you go down in the woods you'd better go in deguise"

func TeddyBearPicnic() {
    words := strings.Fields(body)


    sentence:
    for _, w := range words {
        fmt.Println(w)
        break sentence
            }
        }

