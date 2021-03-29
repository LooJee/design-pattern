package abstract

type ICipher interface {
	Encrypt(data, key[]byte) ([]byte, error)
	Decrypt(data, key[]byte) ([]byte, error)
}

type AESCipher64 struct {
}

func NewAESCipher64() *AESCipher64 {
	return &AESCipher64{}
}

func (AESCipher64) Encrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

func (AESCipher64) Decrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

type AESCipher128 struct {
}

func NewAESCipher128() *AESCipher128 {
	return &AESCipher128{}
}

func (AESCipher128) Encrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

func (AESCipher128) Decrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

type DESCipher64 struct {
}

func NewDESCipher64() *DESCipher64 {
	return &DESCipher64{}
}

func (DESCipher64) Encrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

func (DESCipher64) Decrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

type DESCipher128 struct {
}

func NewDESCipher128() *DESCipher128 {
	return &DESCipher128{}
}

func (DESCipher128) Encrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

func (DESCipher128) Decrypt(data, key []byte) ([]byte, error) {
	return nil, nil
}

//工厂类
type ICipherFactory interface {
	GetCipher64() ICipher
	GetCipher128() ICipher
}

type AESCipherFactory struct {

}

func (AESCipherFactory) GetCipher64() ICipher {
	return NewAESCipher64()
}

func (AESCipherFactory) GetCipher128() ICipher {
	return NewAESCipher128()
}

func NewAESCipherFactory() *AESCipherFactory {
	return &AESCipherFactory{}
}

type DESCipherFactory struct {

}

func (DESCipherFactory) GetCipher64() ICipher {
	return NewDESCipher64()
}

func (DESCipherFactory) GetCipher128() ICipher {
	return NewDESCipher128()
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
