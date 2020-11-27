package dsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

const (
	Path       = "data/"
	SignFolder = "signs/"
	KeyFolder  = "keys/"
	Ext        = ".txt"
	Bits       = 2048
)

func SignFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, Bits)
	if err != nil {
		return err
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(data)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, newhash, hashed, &opts)
	if err != nil {
		return err
	}

	path := strings.Split(fileName, "/")
	fileName = path[len(path)-1]

	signFile := Path + SignFolder + fileName + Ext
	err = ioutil.WriteFile(signFile, signature, 0666)
	if err != nil {
		return err
	}

	err = writePublicKey(fileName, &privateKey.PublicKey)
	if err != nil {
		return err
	}

	return nil
}

func VerifyFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(data)
	hashed := pssh.Sum(nil)

	path := strings.Split(fileName, "/")
	fileName = path[len(path)-1]

	signFile := Path + SignFolder + fileName + Ext
	signature, err := ioutil.ReadFile(signFile)
	if err != nil {
		return err
	}

	keyFile := Path + KeyFolder + fileName + Ext
	publicKey, err := readPublicKey(keyFile)
	if err != nil {
		return err
	}

	err = rsa.VerifyPSS(publicKey, newhash, hashed, signature, &opts)
	return err
}

func writePublicKey(fileName string, publicKey *rsa.PublicKey) error {
	keyFile := Path + KeyFolder + fileName + Ext
	keyData := fmt.Sprintf("%v %v", publicKey.N, publicKey.E)
	err := ioutil.WriteFile(keyFile, []byte(keyData), 0666)
	return err
}

func readPublicKey(keyFile string) (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	keys := strings.Split(string(keyData), " ")
	if len(keys) != 2 {
		return nil, errors.New("Wrong data in key file!")
	}

	N, ok := new(big.Int).SetString(keys[0], 10)
	if !ok {
		return nil, errors.New("Wrong data in key file!")
	}

	E, err := strconv.Atoi(keys[1])
	if err != nil {
		return nil, err
	}

	publicKey := &rsa.PublicKey{N: N, E: E}
	return publicKey, nil
}
