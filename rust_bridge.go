// rust_bridge.go

package main

// #cgo LDFLAGS: -L${SRCDIR}/target/release -lcrypto_rust
// #include <stdlib.h>
// #include "crypto.h"
import "C"
import (
  "fmt"
  "unsafe"
)

func GenerateKeyPair() (string, string) {
  sk, pk := C.generate_key_pair()

  // Convert C strings to Go strings
  gosk := C.GoString(sk)
  gopk := C.GoString(pk)

  return gosk, gopk
}

func main() {
  sk, pk := GenerateKeyPair()
  fmt.Println("Private Key:", sk)
  fmt.Println("Public Key:", pk)
}
