package passwords

import (
	"log"
	"testing"

	"github.com/AndreyLM/go-vue-server/pkg/system/passwords"
)

func init() {
	log.Println("Testing password")
}

// TestBasicLog - testing password
func TestBasicLog(t *testing.T) {
	pass := "TEST"
	hash, err := passwords.Encrypt(pass)
	if err != nil {
		t.Error(err.Error())
	}

	t.Log("hash:", hash)
	ok := passwords.IsValid(hash, pass)
	if !ok {
		t.Error("Password not matching... hash not working")
	}
	t.Log("Success")
}
