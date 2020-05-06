package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// type P struct {
// 	X, Y, Z int
// 	Name    string
// }

// type Q struct {
// 	X, Y *int32
// 	Name string
// }

// func main1() {
// 	var network bytes.Buffer
// 	enc := gob.NewEncoder(&network)
// 	dec := gob.NewDecoder(&network)
// 	// Encode (send) the value.
// 	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
// 	if err != nil {
// 		log.Fatal("encode error:", err)
// 	}
// 	// Decode (receive) the value.
// 	var q Q
// 	err = dec.Decode(&q)
// 	if err != nil {
// 		log.Fatal("decode error:", err)
// 	}
// 	fmt.Println(q)
// 	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)

// }

type P struct {
	X, Y, Z int
	Name    interface{}
}

type Q struct {
	X, Y *int32
	Name interface{}
}

type Inner struct {
	Test int
}

func main() {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.

	gob.Register(Inner{})

	// Encode (send) the value.
	inner := Inner{1}
	err := enc.Encode(P{1, 2, 3, inner})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(q)
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
