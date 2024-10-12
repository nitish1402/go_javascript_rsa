package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {
	// Generate RSA key pair (2048-bit)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	// Data to encrypt
	plaintext := []byte("Hello, RSA Encryption!")

	// Encrypt the data using the public key
	label := []byte("") // Optional label
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, plaintext, label)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Encode the encrypted data in base64 for easy transmission
	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("Encrypted message:", encodedCiphertext)

	// Base64 encoded ciphertext from JavaScript

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)

	// Optionally, save the private key for decryption in JavaScript

	// converting in pem format
	pubkeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	prikeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	fmt.Println("Public Key (PEM):", string(pubkeyPem))
	fmt.Println("Private Key (PEM):", string(prikeyPem))

	decryptMessage()
}

func decryptMessage() {
	privateKeyNew, err := readPrivateKey()
	if err != nil {
		fmt.Println("Error reading private key:", err)
		return
	}
	ciphertextBase64 := "u/Pk0I36vl0hJmcy/mSHvki7hzKrCEW5rejbhNY+4T2zfprb4Pwgtz1Te00GxuvKAlLSTC9NfbHU91fxpS52A66RZVtPX3nwQm5LfQ26rTwd2bDSlqCkdDrO8zFvzYdbH3EyRSDoBDAOB7VQ6o763k+ytAb6VaMzD8i+q4BogT96aNp5L05tRkDod3ZC8SDqKaipLe9Bj2uaGLnMktXM0raTqMyEqDhAouqipQVckO8GoEhb+jGkYZ2nVyQvUKonpFZc5+hcN6jOwAJVsUryVUwDIZmoBHZd5K4ITNfo5shG9W7UK3b4PvzLVaswP1sDZuZpR6b75PrTlaSmX/p+6A==" // Replace this with the JS ciphertext
	ciphertextDec, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		fmt.Println("Error decoding ciphertext:", err)
		return
	}
	// Decrypt the data
	label := []byte("") // Optional label
	hash := sha256.New()

	decryptedText, err := rsa.DecryptOAEP(hash, rand.Reader, privateKeyNew, ciphertextDec, label)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	// Output the decrypted message
	fmt.Println("Decrypted message:", string(decryptedText))
}

func readPrivateKey() (*rsa.PrivateKey, error) {
	pKey := `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDCRvPjdIkfNOaZ
E2gBq68S+Cuw8CblgJaxfelB4CouKO07VbHYJ9ZMi2co95rZ/ilnRPgm0NibTRFb
6G0Qr45JTuqCdE6Mbz0L1fM0iqMRiq/+cnf63EOKif5+s3UfE1YZNRcZiI86K+v/
BbviOM4mVdHg+KK1yK5RjpztwhJWZidZzZY20nfDbF1LeN/aFIUF4h0Pz+Rjt0cO
Y6T9cUBJKCNolEt6zs6D8sABRh6h3y5tCydZAZ50lnFRMPjZR9IuH7jElggub68L
dA8QTXc/odHpinS6qdDjYb5L7+yo6DQ8kW6tf6h9pWHU62bi9Vm1CyajOmuIEh5o
dNQ2zw+3AgMBAAECggEABXyxiLwS32HEHZCxwFJUCIof2ch+oL7IAs1WmDi3mEQp
pyJdeohtgg3x5PNRWVTXYIZPJ/GAHKrJkbn30p/SuflToEmeqlK9+6aYTuSXhHhR
TjN3dgtgkPoiyPtSlIUcSmv4cg0sVm1FJhfIXbRTBjwoSF1dYxr+6WjIv+JaWDxP
xt/cpHcLX2kEdnxWMgboqDZH7mkKCIVE/r2Dm8kngm9Qw+S9oVmgBDjvz8CkfkCV
qlZfmsmKjbogxh8VAPVrNI762zN1Z1Q/fkJcfNxgnyYM4PNOOKBgRWzCp8M6o+8h
/gqyuNcQz2/3Hk+anbdlHYzujDaUXcpO+QJUMg8S2QKBgQDntA266RS0eI1h+Y/u
mxYSq5zhSEsd0xEUUS9f/tgv35aelDfG2/uAQUsOU4wF0EWDRto5N9idnE8NC5fy
dFh/KpZFUZVKHWHYbkfvjwtHvMEDtb6nObFoC0zxvmnXv9GpoFY0Vc9CgBm8G5iP
spqby4qjoMFKWX2vihz7eAzaGwKBgQDWpjbdVUDCQu8K4PGPsRR7DI1ofn+Y2PCs
KxXLYlj9G1AS1qSAdW/z+B2pj8ph/vP+APtqKmgFyWLYRiu/cFsKRO3NUlHX0gkp
HG6LNnR2HJQudigHDCjUeMkiyC40AZljCAM/YqT3sV8JZxJDTEpOldJfuODG9PPB
zojqiF46lQKBgQCEcgHfM9joCHkY5iUGSZRme76jcEWv+LSsnnOsNeqyAucAIs13
WMv81lXnDI7fy9vQXLHlPy0Newoc9OGYcDUeC+P2H9pskTsbEn8ASw4xpY3XZw5i
XLIyPTNwhF3QPA/HHKXqKJre8obDZirhCUEjiUfonL2gmGMl8pb6j/cYvQKBgQCc
v+bi/SX8dwq/xANDrspJDaKag35ErXAcGp6g1uIre+2exDZ/RMOLw05ODB58L8YY
YLk5D1zFlQpk6+F4vOEO/8U1In93/v8Hkaa8bPjhY/9maozSkLOkbcxcZRkwi1Zr
NmfAuxXDeDjx0d5JXQYKm2h01Lr5L9puVpgvxYQifQKBgCmTOsl0i/A+KF2bDCtE
NUh5kQh21KBSktos8fKS9bOjva/BFQuDr/2CeDw6txo91p9n7Cevka/TQDo+0gpC
ivVbAodzc7OT4bwB+Ti7wAx/kc7hyLvds3dKB3IvPRxr/KNxER6uUbKlzqtMXnKZ
1qCijkA2OH+K0ut2Sr54avqL
-----END PRIVATE KEY-----`

	privateKeyBytes := []byte(pKey)

	block, _ := pem.Decode(privateKeyBytes)

	fmt.Println("Block is ", block.Type)

	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	fmt.Println("Block is ", block.Type)

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey.(*rsa.PrivateKey), nil
}

/**
Encrypted message: IXLQgqoto7Qw2G0YJAC7FwMw+LhconMKNsME/rb94IAT2XAL6YHCgrVQJqBhXWKiGOipP5vgpMwYo/s8rw0plJpL8WXtwPYQDU1+9T9w9PBWR0SRbZ/KLgFHZVk1RsgjS1mOKowH76Eba7v4M+7IpeMP4BOhB8DL7rB6IFuRSr/rXEFROyuGyvI1SQWB2J+YEOTlptAZQnLW0yaJBUDGu906LihGntTWdddgb7gkLYbWgzRHWY4SQzamOhi7rN8+qAaxA/qoSfCHAj26G3vebGk2nfDLPpsZeyMGJxi/UOSYtfvWMcKgFRuqRr4T9/BftyP3t1QOask5wupki0UPLQ==
Public Key (PEM): -----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwkbz43SJHzTmmRNoAauv
EvgrsPAm5YCWsX3pQeAqLijtO1Wx2CfWTItnKPea2f4pZ0T4JtDYm00RW+htEK+O
SU7qgnROjG89C9XzNIqjEYqv/nJ3+txDion+frN1HxNWGTUXGYiPOivr/wW74jjO
JlXR4PiitciuUY6c7cISVmYnWc2WNtJ3w2xdS3jf2hSFBeIdD8/kY7dHDmOk/XFA
SSgjaJRLes7Og/LAAUYeod8ubQsnWQGedJZxUTD42UfSLh+4xJYILm+vC3QPEE13
P6HR6Yp0uqnQ42G+S+/sqOg0PJFurX+ofaVh1Otm4vVZtQsmozpriBIeaHTUNs8P
twIDAQAB
-----END PUBLIC KEY-----

Private Key (PEM): -----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDCRvPjdIkfNOaZ
E2gBq68S+Cuw8CblgJaxfelB4CouKO07VbHYJ9ZMi2co95rZ/ilnRPgm0NibTRFb
6G0Qr45JTuqCdE6Mbz0L1fM0iqMRiq/+cnf63EOKif5+s3UfE1YZNRcZiI86K+v/
BbviOM4mVdHg+KK1yK5RjpztwhJWZidZzZY20nfDbF1LeN/aFIUF4h0Pz+Rjt0cO
Y6T9cUBJKCNolEt6zs6D8sABRh6h3y5tCydZAZ50lnFRMPjZR9IuH7jElggub68L
dA8QTXc/odHpinS6qdDjYb5L7+yo6DQ8kW6tf6h9pWHU62bi9Vm1CyajOmuIEh5o
dNQ2zw+3AgMBAAECggEABXyxiLwS32HEHZCxwFJUCIof2ch+oL7IAs1WmDi3mEQp
pyJdeohtgg3x5PNRWVTXYIZPJ/GAHKrJkbn30p/SuflToEmeqlK9+6aYTuSXhHhR
TjN3dgtgkPoiyPtSlIUcSmv4cg0sVm1FJhfIXbRTBjwoSF1dYxr+6WjIv+JaWDxP
xt/cpHcLX2kEdnxWMgboqDZH7mkKCIVE/r2Dm8kngm9Qw+S9oVmgBDjvz8CkfkCV
qlZfmsmKjbogxh8VAPVrNI762zN1Z1Q/fkJcfNxgnyYM4PNOOKBgRWzCp8M6o+8h
/gqyuNcQz2/3Hk+anbdlHYzujDaUXcpO+QJUMg8S2QKBgQDntA266RS0eI1h+Y/u
mxYSq5zhSEsd0xEUUS9f/tgv35aelDfG2/uAQUsOU4wF0EWDRto5N9idnE8NC5fy
dFh/KpZFUZVKHWHYbkfvjwtHvMEDtb6nObFoC0zxvmnXv9GpoFY0Vc9CgBm8G5iP
spqby4qjoMFKWX2vihz7eAzaGwKBgQDWpjbdVUDCQu8K4PGPsRR7DI1ofn+Y2PCs
KxXLYlj9G1AS1qSAdW/z+B2pj8ph/vP+APtqKmgFyWLYRiu/cFsKRO3NUlHX0gkp
HG6LNnR2HJQudigHDCjUeMkiyC40AZljCAM/YqT3sV8JZxJDTEpOldJfuODG9PPB
zojqiF46lQKBgQCEcgHfM9joCHkY5iUGSZRme76jcEWv+LSsnnOsNeqyAucAIs13
WMv81lXnDI7fy9vQXLHlPy0Newoc9OGYcDUeC+P2H9pskTsbEn8ASw4xpY3XZw5i
XLIyPTNwhF3QPA/HHKXqKJre8obDZirhCUEjiUfonL2gmGMl8pb6j/cYvQKBgQCc
v+bi/SX8dwq/xANDrspJDaKag35ErXAcGp6g1uIre+2exDZ/RMOLw05ODB58L8YY
YLk5D1zFlQpk6+F4vOEO/8U1In93/v8Hkaa8bPjhY/9maozSkLOkbcxcZRkwi1Zr
NmfAuxXDeDjx0d5JXQYKm2h01Lr5L9puVpgvxYQifQKBgCmTOsl0i/A+KF2bDCtE
NUh5kQh21KBSktos8fKS9bOjva/BFQuDr/2CeDw6txo91p9n7Cevka/TQDo+0gpC
ivVbAodzc7OT4bwB+Ti7wAx/kc7hyLvds3dKB3IvPRxr/KNxER6uUbKlzqtMXnKZ
1qCijkA2OH+K0ut2Sr54avqL
-----END PRIVATE KEY-----
**/

/**
u/Pk0I36vl0hJmcy/mSHvki7hzKrCEW5rejbhNY+4T2zfprb4Pwgtz1Te00GxuvKAlLSTC9NfbHU91fxpS52A66RZVtPX3nwQm5LfQ26rTwd2bDSlqCkdDrO8zFvzYdbH3EyRSDoBDAOB7VQ6o763k+ytAb6VaMzD8i+q4BogT96aNp5L05tRkDod3ZC8SDqKaipLe9Bj2uaGLnMktXM0raTqMyEqDhAouqipQVckO8GoEhb+jGkYZ2nVyQvUKonpFZc5+hcN6jOwAJVsUryVUwDIZmoBHZd5K4ITNfo5shG9W7UK3b4PvzLVaswP1sDZuZpR6b75PrTlaSmX/p+6A==
**/
