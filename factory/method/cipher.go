package method

type ICipher interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

type AESCipher struct {
}

func NewAESCipher() *AESCipher {
	return &AESCipher{}
}

func (AESCipher) Encrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (AESCipher) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

type DESCipher struct {
}

func NewDESCipher() *DESCipher {
	return &DESCipher{}
}

func (DESCipher) Encrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (DESCipher) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

//工厂类
type ICipherFactory interface {
	GetCipher() ICipher
}

type AESCipherFactory struct {

}

func (AESCipherFactory) GetCipher() ICipher {
	return NewAESCipher()
}

func NewAESCipherFactory() *AESCipherFactory {
	return &AESCipherFactory{}
}

type DESCipherFactory struct {

}

func (DESCipherFactory) GetCipher() ICipher {
	return NewDESCipher()
}

func NewDESCipherFactory() *DESCipherFactory {
	return &DESCipherFactory{}
}

func NewCipherFactory(cType string) ICipherFactory {
	switch cType {
	case "AES":
		return NewAESCipherFactory()
	case "DES":
		return NewDESCipherFactory()
	default:
		return nil
	}
}
