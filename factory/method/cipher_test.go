package method

import (
	"reflect"
	"testing"
)

func TestCipherFactory(t *testing.T) {
	var f ICipherFactory = NewAESCipherFactory()
	if reflect.TypeOf(f.GetCipher()) != reflect.TypeOf(&AESCipher{}) {
		t.Fatalf("should be AESCipher")
	}

	f = NewDESCipherFactory()
	if reflect.TypeOf(f.GetCipher()) != reflect.TypeOf(&DESCipher{}) {
		t.Fatalf("should be DESCipher")
	}
}
