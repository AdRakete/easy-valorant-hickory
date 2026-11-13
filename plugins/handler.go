package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.4" )

func ayogWuncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NThTTM6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9ojDZcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7li38HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eiSvGdbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBD8FEWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1HnmDX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jmit0amYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPF5vGZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dh3uPL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXcmGeCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtG5AFTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxXlKC7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdTNg8WKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oG4LLFp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PVowQggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8H7aoAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3j7J5MvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPByVluzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzPnJkHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Khz09AxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 159HfBswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5y9miduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmIxRvRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hz5JTQtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q49DrVoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YbheUK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9W3hhUOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSubYwo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCQ04mbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGG4S26oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDFDlyLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m26Vi6J6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oF1r0ARpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0RSM0XlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRV2BwZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ5lyBM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6S7nLXGeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fVNz1geWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBKWHHaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVoic1pdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXzymgTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AYOUe3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s015Hq2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chR4WKAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORFjg2kCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMqHHQYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNsIxze4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjpL4eW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOLLhQFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcUl2vy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQRnX33HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8wQ85EUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjUIlW6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7C9A7fqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dV3QyiwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jYmG5kAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfmWpVy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fa7vk9FKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsRFYN3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TXkKwNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAfqlkjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAKVi10gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5aUZrY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCJYpAi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Q8ihmv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9O9LrUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1CjsXapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1obDNOk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w341mh3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOXr3SarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXvE1f4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InAPkkPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IADyDyvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRNTkWUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oe7HHzBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Le1XWTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGkBesruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUy6v7mIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJh7DP3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7tAGcQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ynu6ImdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONVDy3bwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObiMAZ9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baGHh6B3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uR3MZN2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JteKrodaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJNUFhI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJ4hhbi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4UzDPibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2CZt6ezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdYKB8TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PB3i6UjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FohmmBc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Un6EFM1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llL6WdyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRBhaTkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZYZ17RDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G19FUGfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z44vTsgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDi62Q6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raI0g9DaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eb5lFTi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iln5lmOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCqa1SC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZoZn0JDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJoEf4DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Od9NeBW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Rh3LBs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhmIFW0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8U3SytqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c04WVKdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAgFMG3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SvxFXBeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvVtYYLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUq6ePA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVyckzC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiA4ZM3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WVn7IVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNHA4ZCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwVYGhSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dpp76zoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w51sej5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNiV6ZSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7R09c1xVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjUW54oQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5mK6mzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9TjCUvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLJexgSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S7ex6pHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rC1vTLw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGnzQG6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOsnewKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y94SmUdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koE6nf0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgmUvefSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbeIAF6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMWvUpwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OorwHGaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zC59IQc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxOvOpBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AIImR0p2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxONSERwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNusU0JaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoXtPNyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITga5bZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YFZ6kAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiADWLbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kllI4vvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kKSNce2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXVaNw5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBT28kumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjgFp4tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8tbM13WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TloU8eclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uUaEVbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHxYrcBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzl3nNF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ71nw0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1PahTKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBjm44RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vqK9kZOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYOK0HGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iy8Wtn6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzQkytdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K099Wcg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKFh8J9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfDxZKMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDAKLu55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5fFwwe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTvulqMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWAs0e4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTUF7hdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQCoOrS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awOJDpBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkH8QHkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJFYHm7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmgSyvELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEwBLXo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nemYMuBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8m54iI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEcXMWYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juIqtGQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNIbJDehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvD11R9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHRKa6eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bscHcbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYXuywMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z54zVmYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZMIfK6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sh07Faa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7KSASMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q66hLTYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jbghYG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjacVakvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cN3CH3A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JWDlHznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIhA0hAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgXg6LdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXxqPDC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1e1F5QWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFX4S9rWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhOIyl6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1srbkGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSxW1Z96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6VCGjtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGBzG1VhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqwWHb08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30JYcmmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsE78vkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLHrZTXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7YKXExdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZvcfJRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qHEIhFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVzaCNgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvPJMbkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nkl1chgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtClTvbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNN0oluMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BonAUWp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MeoAiKApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPk28nNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKcqGfhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoDIgElcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylRyDN7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYFj2ZsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BaR9ROvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQq4Kez2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uDgb2IZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypXmlxw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JNeef3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dem4cfFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8e6cBOkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhKoPIVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGiMOAGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mI2vQ2kLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvoqF4adWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PukwkH78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pKH7dWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IWxOKE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TgnuiK4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8Efzo2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ySU9J6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xyHVbxfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYjXtslzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHlKvBvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AkOtsB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4kRdpHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PG3VeNewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7VnWVQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDg8Okf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGLSDG4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xk5HRV6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWcvn67TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUYzGaPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5bf4y7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8n5Zau9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfGTa9bvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXS9sAs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iebla8XRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2lvrASiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3S8kFJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWxwUI7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWlX037SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pX79ccZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPt3GOtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmwIzJwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrqOVgWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DfxHgTi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJ6Y6lP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xC945eHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3S7HwTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXsYOvzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsxXGDQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ug78qzPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wl3Mrz06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AK9CP9BhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTSrlCFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsyWyoj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVDYIAbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCnt07iBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wcCBUavYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6kav5MbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhrsXHu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVxwYNGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUQGmTkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXfmXdr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGt2ps8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sojnx3uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUIlE1ZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLQOMldsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qtKVY1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UXGLDrVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsMw4CeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzBU6AXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mymOXOT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uC2reCy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm9JrohGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSVk7e1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2Bv6LX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cn5l1Gu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pj6fMNQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FySF5KIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U6yn3bS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Qj4LpHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoG7Hv3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CSZGSx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JloF9XuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r95QYRosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlGCh1PaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLYMOJEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6HUzpTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjhAI5i0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLHQSoBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8TcxUZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3BJ8h3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kX45Wwl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtbVGYhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5w2iHXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOsBJyEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngQJFqeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esG1gbPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izLcihDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1u0T9L3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJSu6VWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0QXfcX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0a5boMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCx58sxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2GAvuJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQmBXYccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRfzNgSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRyBfRpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E22NPm6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsEjW0K2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOfpjq2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KoaOMyWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POlcDSzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lc2j6hXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDuoVqGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9Fcf8qrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AApccdlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hxx896B0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYtFtImAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDBQJ0rPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdLc8Xh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpGDX7pWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lV0c28D6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hwkmzrj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlSd2Fp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trqY7XfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LM4NHgmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQWdOpqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnQgKel8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hpxn5j7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aO28bqGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6Z5oWolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzHprS4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUBAot63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DLJF5atWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBCz9a0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yrkv2F3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msM96ARrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxBr0HgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUj1K72AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYEv5JVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TqXPb0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYeDjmfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfOmsTCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsMguJniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WErhD4RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czjVhE66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urAvAAdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tPwKFRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMBEvovCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOHlGDfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHIMQc01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8m8wxGj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05RwtQHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzxNLpFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uggn0rGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLulQrGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RR4QTmlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaqI6bf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2Ro29g4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlVyReyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJtKXQgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhck9znGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6b5E3g9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rh5fYmA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyjP6SwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sp2o5vqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbCW6bSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wywLIWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wh1Dku7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVtcw6taWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvwplBMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlIvJ0lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDoTPPl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDaGBtrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQjPu9BbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBVIpvFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nESKPBC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyvYOkiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dltKrjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hs7MBsWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwAk23hPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UkBCzhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KKEPzJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KOgCR1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmeuisRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oiReGcXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5skIYzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwF4kvOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZkBdAnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NedTz28IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KOarqzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjOvPHtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqo7LriPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MekjTQmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kissxM9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SIf5qkCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fbf7gj3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0j2J3IPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmPtiSUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func We4aZFqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1XMbPeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIFK1rpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0Yea8f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sB3tv9yhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhMSTiIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1LGQfthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDX4lA7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVoMEAGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXbmEYjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SIJcs4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUf1WoXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcL4aph7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUJAEliVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94sblE0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMhyplwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hNEPexuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceOLSVvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1wTGD0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myIcDeYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93fRGImDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6pOngujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtY5cpVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1Cq83WtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSDav0sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NuriOVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGZcfwYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5SsLrKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2Di6eE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrMYlUsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeersrhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p15ll3MnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6Go8f8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f78t2nIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hDirmGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jI36EzY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozgslHI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwDFXRCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loCHomM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hx6Gso9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oatfG9MQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vnNi8R6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkergk05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hk9ijEbYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaHLN49ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wo7ELVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R30QMUNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q57qiB7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7MlQaVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dxqAtopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjyAsVO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XSodDzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QstiBqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOQnGBEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2s43Oe9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3Bndx5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tb3e8g3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcPB1cKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7Wcbx7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gsl3ncX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHxBJTggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHH45KbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofN5IbLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPK5Nk4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1KdGIPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNz71yqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALg8wSXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDmlooesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaHqoMUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOeI3DD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34d90Nw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fFUbdSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0wrDAjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FdZeQO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pkl4RyKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PesQAoeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMx93BRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3U2huMavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrPXi6jGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jIuwlIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHWPPvEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Su7EM6eZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCHcTJr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jp8gzD39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAmLrqeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAFutTLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSEgaE0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Vu7L6JMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PupbnqGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnyMMySfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0O17EvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpBEEcx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CO7vzq8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44nnQXs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHBCZaYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDoEBzv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OD4zNldCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lha0acUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yuxb5ZEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJkmxdnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5avmIqc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFadu7d9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhs7Cy1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 295R57CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQWxTPOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WUyswSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1axn1j3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHoYh1pJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZP1m0APUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkKSwEgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVEJYXNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGBDtDNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOvPzM2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sVSDU3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHgcuCbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0TKSv5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jh6rtbkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNE2dYgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XU6fZtjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yyvZjsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6FlEYJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xUw3Jz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMleqaEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67WJf9AMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHaJ81AoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqkYrNRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rc94xYSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iF4A3ubvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZAjovseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PC4e7g4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idamxKzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9np0sboHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4EVDHBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8qTEPdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4oekUi0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlkKZnrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yft2GqIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hbgzgmu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWCNtdApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6Z4MsL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xhU3JrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXDQ3NqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujzyiLWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OD169bt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBLPRAkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzxB9j0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiING7ZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pv6POtgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmuxdOtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3G7R5NLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWmZv1IWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYt3BfjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBsyltZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1eo9tEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htOYQb0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9soPaF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35SumAOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isp0vVk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIbw4CtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s39WTAbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glxbEwYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ng845rv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rrxip66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYvUHFDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdOObkvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEuDznRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bZOBEAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkVRlm0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNVe7CTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWM6oZGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jj3i6plxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qG86cZGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMYDE27hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDdFItG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaWxUGaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HE9f7CMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KN9q7koFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6O4CYeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epQK2S9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzFr3kqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1lHikvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awmyhlMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCrlrPRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJ0yII7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MM0vIBbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func feZdZaZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjRgfddZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VO3RU1B3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxVMcvz1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XaLalY3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihPZMAY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkRxRcttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRutIIgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6Eg5USkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYLW4EATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSDMf4zHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yia9bzFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfUeYgqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qltKDTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sgDLdzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCi2OJDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BySOsWlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1YMYSLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K7GRnpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WN9xpmtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMZl9xEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w53VGnBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxgQvPghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bloh1KkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1fp4YXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLOHLtOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpFKDr0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyNehBvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xUkbd3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afqoOKOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtQZ6FTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uHzXFkfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93thnegdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STkTKASwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16AJTUedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5UG9Tp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsMOdOOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0Dqi5CmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9AyOBLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnrePVNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73UdSD8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07GQ6pIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KifTEj52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6s4v1UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUrEeh2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5T3yAhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsiIMfr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UykppUe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjdzL69qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wf2kHxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVqVC7SgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMJdDAdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2zHzexvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHXSmqmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoNHRIaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYZr0j1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4T9XsOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KkkilWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW1Ph8tSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQhW7WH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uySsFOgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMHBb9QHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fggj0bHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrBQoNR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AN4laO6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func celHqkiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEZDoEi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIwnOdzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZvuMwdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uksLFqgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYaXJ27bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7WJvuVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lmq09Ko5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nG21FRUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaudXTrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SBUjesAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVkwC1xFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAVbJa3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ny0XOEeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zItytDc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgqxcgVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbCoOC3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaO6EWyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJwsIlRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Z9wneQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJ2ZqQSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAd2ftsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEfPsKbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woG5O6PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDluKbGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udQZWmOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RP91Aq0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OAJKmV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kc9uD1NlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHbYtb4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3i4Ag9m7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2pOy9BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnQIQ5utWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0PWvbeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVJrJj4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAIf9vusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sh1iSHemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynr3h9d0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UA8GRgOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVjmgFD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqDN67PEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DFtbAYX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6tOgjTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKMPSGWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXCeJPK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhBBUyToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfYpxCqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivocJW76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ve99ImwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08jaqVvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aHFlDhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcyV36KSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8uuG939Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfYmS0WZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJS3iTilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcI7TlhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LH5x6wCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EydUbSyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzFXgqLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WMQuyJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqRy7bJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCDwxtBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W267GQA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWQ1Wuo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRE1weffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlJcClqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0fvHPLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nxNIirGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwSZllNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjaeMftWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPoUktTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxNdqrDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rVo01AoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA6dx9bfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qV5kcgLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCVFrjIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZux9HHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMWHm4poWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtUKD7VsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NDopYb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpBweIWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTiWhhqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J64UBBtZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3rYGeKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KE4oHfWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHnj1Md3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7m72J8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzLEtDoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmuVSfr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gATl5prfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWjPTyf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2lC7PZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dUw9tOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qKNBKSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2odmMl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqSD4NBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vONht5LTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDkeAcu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gfeFe4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 37QV0j4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcuMuIXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJ4lJtaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdDfxDNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEWqwvEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cYOa9shWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HYMWHUNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtxA268JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBrFfRiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpAtgwoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3pbORLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63ztXeftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQCLdL3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2irMA4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6tgP1hHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q11nTI4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ILoibIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meUmFLaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VTsCUXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHkp3P72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPqLOIdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfjc0x8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hgQOtXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeS835hYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUpWWu6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAS3PU7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZ6WWv4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qe0LxctDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdttrvFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pk2k6UoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNyx4WOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLsbnG9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5xwOXpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEHoX5P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHuBdhATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCd3oyXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gc4GY9iAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKVnJb7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONSaJIq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BVQk4BEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOQ2ArIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pG2AS8KBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3xlG1sEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvjfr9dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmLOOSxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Wa0ZZLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6GMof8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NC0vZN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDKvdPxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRkjO4ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Py2Y1bohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82ZgLcRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bvYmaczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcH2uOO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPlm6zUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxZl8YeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDtkhOwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnwB3xiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WT2nBN5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmazfEXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3ctYSRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCzaULvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UntM3sckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKXf4O0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76iWmgrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pru77FOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPQ1pc8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnvX4UiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFui1n16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDx3OL0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGA1NFBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVlcGTuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9Mpf2RQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiJrAbG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyJtusVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ez7rpz8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ym4PbFFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 079uqzmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3AZwEDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6sIlgaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejbBdxujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5nzNX9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGZFaXl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JeIr5muDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJNAZPgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izOjzOCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJyTrwzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJ2CsOr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLx3H031Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKVSnxqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HH7IT4N5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qVDvxzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

