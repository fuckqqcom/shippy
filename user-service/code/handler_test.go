package code

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestEqual(t *testing.T) {
	passwd := []byte("123456890445dfsf")
	hash, _ := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)
	fmt.Println(string(hash))
	if err := bcrypt.CompareHashAndPassword(hash, passwd); err != nil {
		t.Errorf("bcrypt.CompareHashAndPassword error, got %#v", err)
	}
	fmt.Println(string(passwd))
}
