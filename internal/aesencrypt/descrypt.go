package aesencrypt

import (
	"MySQLStatusWatcher/internal/logger"
	"encoding/hex"
	"fmt"
)

func Encrypt(code []byte) string {
	_encryped, err := AesEncrypt([]byte(code))
	if err != nil {
		logger.Errorf("encrypt error: %v", err)
		return ""
	}
	encrypted := fmt.Sprintf("%x", string(_encryped))

	return encrypted
}

func Decrypt(encrypted string) string {
	_encryped_byte, err := hex.DecodeString(encrypted)
	if err != nil {
		logger.Errorf("encrypt error: %v", err)
		return ""
	}
	_descryped, err := AesDecrypt(_encryped_byte)
	if err != nil {
		logger.Errorf("encrypt error: %v", err)
		return ""
	}
	return string(_descryped)
}
