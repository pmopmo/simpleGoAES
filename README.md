# Simple Go AES encrypt/decrypt

forked from https://github.com/sezzle/simpleGoAES

A simple golang library to encrypt and decrypt strings or byte arrays using AES.
My goal with this fork was to make it as simple as possible to use.

All you need now is a password string __(pwd)__, and a string you want to encrypt __(encryptThis)__


Example usage:

    package main

    import (
        "fmt"
        "github.com/pmopmo/simpleGoAES"
    )

    func main() {

	    pwd := "My secret password"
	    encryptThis := "Simple Go AES encrypt/decrypt"

	    cipher, err := simpleGoAES.Encrypt(pwd, encryptThis)

	    fmt.Println("res: ", cipher, " error:", err)
	    cleartext, err := simpleGoAES.Decrypt("My secret password", cipher)
	    fmt.Println("res: ", cleartext, " error:", err)
    }

<!-- for some reason this pulls an old release [Try it](https://play.golang.org/p/HpaS1-Mpq7G) -->

----

Code licensed under an [MIT-style License](./LICENSE).\
Documentation licensed under [CC BY 4.0](http://creativecommons.org/licenses/by/4.0/).

Library home: https://github.com/pmopmo/simpleGoAES