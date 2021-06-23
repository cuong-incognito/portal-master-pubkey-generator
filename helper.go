package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/consensus_v2"
)

func readMiningKey(fileDir string) ([]string, error) {
	f, err := os.Open(fileDir)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	miningKeys := []string{}
	for scanner.Scan() {
		miningKeys = append(miningKeys, scanner.Text())
	}

	return miningKeys, nil
}

func generateSeedsFromMiningKey(miningKeyStrs []string) ([][]byte, error) {
	seeds := [][]byte{}
	for _, miningKeyStr := range miningKeyStrs {
		miningKeyObj, err := consensus_v2.GetMiningKeyFromPrivateSeed(miningKeyStr)
		if err != nil {
			return [][]byte{}, err
		}
		seeds = append(seeds, miningKeyObj.PriKey[common.BridgeConsensus])
	}
	return seeds, nil
}

func generateBTCMasterPubKeysFromSeeds(seeds [][]byte) [][]byte {
	masterPubKeys := [][]byte{}
	for _, s := range seeds {
		pubKey := generateBTCPubKeyFromSeed(s)
		masterPubKeys = append(masterPubKeys, pubKey)
	}
	return masterPubKeys
}

func generateBTCPubKeyFromSeed(seed []byte) []byte {
	// generate BTC master account
	BTCPrivateKeyMaster := chainhash.HashB(seed) // private mining key => private key btc
	return generateBTCPubKeyFromPrivKey(BTCPrivateKeyMaster)
}

func generateBTCPubKeyFromPrivKey(privateKey []byte) []byte {
	pkx, pky := btcec.S256().ScalarBaseMult(privateKey)
	pubKey := btcec.PublicKey{Curve: btcec.S256(), X: pkx, Y: pky}
	return pubKey.SerializeCompressed()
}

func writeFile(fileDir string, masterPubKeys [][]byte) error {
	f, err := os.Create(fileDir)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, pubKey := range masterPubKeys {
		str := fmt.Sprintf("%#v\n", pubKey)
		_, err := f.WriteString(str)
		if err != nil {
			return err
		}
	}

	return nil
}
