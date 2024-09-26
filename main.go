package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
)

func bytesToStr(bs []byte) string {
	return *((*string)(unsafe.Pointer(&bs)))
}

func sub8(a, b uint16) uint16 {
	return a - b
}

func sSub8(a, b uint16) int16 {
	c := a - b
	return *((*int16)(unsafe.Pointer(&c)))
}

func main() {
	bindSimplePw := "tuantuan"
	h := sha256.New()
	h.Write([]byte(bindSimplePw))
	fmt.Printf("sha256 sumed: %x\n", h.Sum(nil))
	passbcrypt := "243261243130244B62463462656F7265504F762E794F324957746D656541326B4B46596275674A79336A476845764B616D65446169784E41384F4432"
	decoded, err := hex.DecodeString(passbcrypt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("decoded:", string(decoded))
	fmt.Println("compare result: ", bcrypt.CompareHashAndPassword(decoded, []byte(bindSimplePw)))
	cost, _ := bcrypt.Cost(decoded)
	fmt.Println("cost: ", cost)
	bHash, err := bcrypt.GenerateFromPassword([]byte(bindSimplePw), cost)
	if err != nil {
		fmt.Println("bcrypt error: ", err)
		return
	}
	fmt.Println("compare again: ", bcrypt.CompareHashAndPassword(bHash, []byte(bindSimplePw)), bytes.Equal(bHash, decoded))
	fmt.Println("decoded bcrypt result: ", string(bHash))
	fmt.Println("eq: ", string(bHash) == string(decoded))

	fmt.Println("bcrypt result: ", hex.EncodeToString(bHash))
}
