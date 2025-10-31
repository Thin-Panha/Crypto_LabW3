package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func proofMe(txt1, txt2 string) {
	// 1. MD5
	fmt.Println("Hash (MD5):")
	md5_1 := func_md5(txt1)
	md5_2 := func_md5(txt2)
	fmt.Println(matching(md5_1, md5_2))

	// 2. SHA1
	fmt.Println("Hash (SHA1):")
	sha1_1 := func_sha1(txt1)
	sha1_2 := func_sha1(txt2)
	fmt.Println(matching(sha1_1, sha1_2))
	// 3. SHA256
	fmt.Println("Hash (SHA256):")
	sha256_1 := func_sha256(txt1)
	sha256_2 := func_sha256(txt2)
	fmt.Println(matching(sha256_1, sha256_2))
	// 4. SHA512
	fmt.Println("Hash (SHA512):")
	sha512_1 := func_sha512(txt1)
	sha512_2 := func_sha512(txt2)
	fmt.Println(matching(sha512_1, sha512_2))
	// 5. SHA3
	fmt.Println("Hash (SHA3):")
	sha3_1 := func_sha3(txt1)
	sha3_2 := func_sha3(txt2)
	fmt.Println(matching(sha3_1, sha3_2))
}

func func_md5(str1 string) string {
	hash := md5.Sum([]byte(str1))
	return hex.EncodeToString(hash[:])
}

func func_sha1(str1 string) string {
	hash := sha1.Sum([]byte(str1))
	return hex.EncodeToString(hash[:])
}
func func_sha256(str1 string) string {
	hash := sha256.Sum256([]byte(str1))
	return hex.EncodeToString(hash[:])
}

func func_sha512(str1 string) string {
	hash := sha512.Sum512([]byte(str1))
	return hex.EncodeToString(hash[:])
}

func func_sha3(str1 string) string {
	hash := sha3.Sum256([]byte(str1))
	return hex.EncodeToString(hash[:])
}

func matching(hash1, hash2 string) string {
	fmt.Println("Value 1: ", hash1)
	fmt.Println("Value 2: ", hash2)
	if hash1 == hash2 {
		return "=> Match!"
	} else {
		return "=> Not Match!"
	}
}

func main() {
	var str1, str2 string
	fmt.Println("======== Name + Hashing Program ========")
	fmt.Print("Input valuse 1: ")
	fmt.Scanln(&str1)
	fmt.Print("Input valuse 2: ")
	fmt.Scanln(&str2)
	proofMe(str1, str2)
}
