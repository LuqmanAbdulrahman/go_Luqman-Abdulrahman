package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Cipher interface {
	Encode() string
	Decode() string
}

type SubstitutionCipher struct {
	key string
}

func (c *SubstitutionCipher) Encode(plainText string) string {
	var cipherText string
	for _, ch := range plainText {
		if ch >= 'a' && ch <= 'z' {
			idx := int(ch-'a') + len(c.key)
			idx %= 26
			cipherText += string('a' + idx)
		} else if ch >= 'A' && ch <= 'Z' {
			idx := int(ch-'A') + len(c.key)
			idx %= 26
			cipherText += string('A' + idx)
		} else {
			cipherText += string(ch)
		}
	}
	return cipherText
}

func (c *SubstitutionCipher) Decode(cipherText string) string {
	var plainText string
	for _, ch := range cipherText {
		if ch >= 'a' && ch <= 'z' {
			idx := int(ch-'a') - len(c.key)
			if idx < 0 {
				idx += 26
			}
			plainText += string('a' + idx)
		} else if ch >= 'A' && ch <= 'Z' {
			idx := int(ch-'A') - len(c.key)
			if idx < 0 {
				idx += 26
			}
			plainText += string('A' + idx)
		} else {
			plainText += string(ch)
		}
	}
	return plainText
}

func main() {
	var menu int
	var a student = student{}
	var c Cipher = &SubstitutionCipher{key: "hello"}

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		a.nameEncode = c.Encode(a.name)
		fmt.Print("\nEncode of student’s name " + a.name + "is : " + a.nameEncode)
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		a.name = c.Decode(a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + "is : " + a.name)
	}
}
