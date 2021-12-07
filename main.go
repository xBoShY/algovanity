package main

import (
	"encoding/json"
	"fmt"
	"os"
  "strconv"
	"time"
  "sync/atomic"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func generateFast(channel chan crypto.Account, prefix string, counter *uint64) {
	var account crypto.Account
  var prefixlen = len(prefix)
	for {
    account = crypto.GenerateAccount()
    address := account.Address.String()
    atomic.AddUint64(counter, 1)
    if address[0:prefixlen] == prefix {
      channel <- account
    }
	}
}

func main() {
  prefix := os.Getenv("PREFIX")
	threads, err := strconv.ParseUint(os.Getenv("THREADS"), 10, 16)
  if err != nil {
		fmt.Println(err)
	}
  count, err := strconv.ParseUint(os.Getenv("COUNT"), 10, 16)
  if err != nil {
		fmt.Println(err)
	}

  fmt.Printf("PREFIX: %v\n", prefix)
  fmt.Printf("THREADS: %v\n", threads)
  fmt.Printf("COUNT: %v\n", count)

  var words string
  queue := make(chan crypto.Account, 0)
	counters := make([]uint64, threads)
  
  start := time.Now()

  var i uint64
  for i = 0; i < threads; i++ {
    counters[i] = 0
		go generateFast(queue, prefix, &(counters[i]))
	}

	for ; count > 0; count-- {
    account := <- queue
    duration := time.Since(start)

    var q uint64 = 0
    for i = 0; i < threads; i++ {
      q += atomic.LoadUint64(&(counters[i]))
    }
    fmt.Printf("====================\n")
    fmt.Printf("Generations: %v\n", q)
    fmt.Printf("Time: %v\n", duration)
    fmt.Printf("Address: %v\n", account.Address.String())
    words, _ = mnemonic.FromPrivateKey(account.PrivateKey)
    fmt.Printf("Mnemonic: %v\n", words)
    fmt.Printf("PublicKey: %v\n", account.PublicKey)
    fmt.Printf("PrivateKey: %v\n", account.PrivateKey)
  }

	os.Exit(0)
}
