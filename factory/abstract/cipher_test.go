package abstract

import (
	"reflect"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var f = NewCipherFactory("AES")
	if reflect.TypeOf(f.GetCipher64()) != reflect.TypeOf(&AESCipher64{}) {
		t.Fatalf("should be AESCipher64")
	}

	if reflect.TypeOf(f.GetCipher128()) != reflect.TypeOf(&AESCipher128{}) {
		t.Fatalf("should be AESCipher128")
	}

	f = NewCipherFactory("DES")
	if reflect.TypeOf(f.GetCipher64()) != reflect.TypeOf(&DESCipher64{}) {
		t.Fatalf("should be DESCipher64")
	}

	if reflect.TypeOf(f.GetCipher128()) != reflect.TypeOf(&DESCipher128{}) {
		t.Fatalf("should be DESCipher128")
	}
}
