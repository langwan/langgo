---
title: "rsa"
---

# rsa

## 使用

## 函数列表

* `CreateKeyX509PKCS1(bits int) (pub string, pri string)` 创建公钥私钥
* `Decrypt(key *rsa.PrivateKey, src []byte) (data []byte, err error)` 解密数据
* `Encrypt(key *rsa.PublicKey, src []byte) (data []byte, err error)` 加密数据
* `PrivateKeyFromX509PKCS1(pri string) (*rsa.PrivateKey, error)` 从X509PKCS1解析私钥
* `PublicKeyFromX509PKCS1(pub string) (*rsa.PublicKey, error)` 从X509PKCS1解析公钥
* `Sign(key *rsa.PrivateKey, src []byte) (sign []byte, err error)` 使用私钥签名
* `Verify(key *rsa.PublicKey, sign, src []byte) (err error)` 使用公钥校验
* `PrivateKeyToPKCS1(pri *rsa.PrivateKey) string` 私钥转成字符串
* `PublicKeyToPKCS1(pub *rsa.PublicKey) string` 公钥转成字符串
