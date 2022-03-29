/*
 * @FileName:
 * @Author:		xjj
 * @CreateTime:	下午6:23
 * @Description:
 */
package simpleCrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"sync"
)

var lock sync.Mutex

//Md5Enscrypto md5加密
func Md5Enscrypto(data string) string {
	m := md5.New()
	m.Write([]byte(data))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

//SimpleEncryptAesECB ECB加密，并对加密结果进行了base64编码
func SimpleEncryptAesECB(aesText []byte, aesKey []byte) string {
	lock.Lock()
	c, _ := aes.NewCipher(aesKey)
	encrypter := NewECBEncrypter(c)
	aesText = PKCS5Padding(aesText, c.BlockSize())
	d := make([]byte, len(aesText))
	encrypter.CryptBlocks(d, aesText)
	lock.Unlock()
	return base64.StdEncoding.EncodeToString(d)
}

//PKCS5Padding PKCS5明文填充方案，此方案使明文保持为块长度的倍数
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
