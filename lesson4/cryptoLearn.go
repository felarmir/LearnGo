package lesson4

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func LearnCRC32() {
	h := crc32.NewIEEE()
	h.Write([]byte("Denis Abdreev"))
	v := h.Sum32()
	fmt.Println(v)
}

//===========================================
// hash file

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)

	return h.Sum32(), nil
}

func HashFileLearn() {
	h1, err := getHash("test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}

// ============================================
// sha1
func SHA1Learn() {
	h := sha1.New()
	h.Write([]byte("Denis Andreev"))
	bs := h.Sum([]byte{})
	fmt.Println(string(bs))
}
