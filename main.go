package main

import (
	"flag"
	"fmt"
)

func main() {
	var inputDir, outputDir string
	flag.StringVar(&inputDir, "input", "private_mining_keys.txt", "private mining keys file directory")
	flag.StringVar(&outputDir, "output", "btc_public_keys.txt", "bitcoin public keys file directory")
	flag.Parse()

	// read private mining keys from file
	fmt.Printf("======== Reading mining keys from %v ========\n", inputDir)
	privMiningKeys, err := readMiningKey(inputDir)
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		panic(err)
	}

	// get seeds from private mining key
	seeds, err := generateSeedsFromMiningKey(privMiningKeys)
	if err != nil {
		fmt.Printf("Error while generating seeds: %v\n", err)
		panic(err)
	}

	fmt.Printf("======== Generating BTC Master PubKeys for %v Mining Key ========\n", len(seeds))
	masterPubKeys := generateBTCMasterPubKeysFromSeeds(seeds)

	// dump to output file
	fmt.Printf("======== Dumping to %v ========\n", outputDir)
	err = writeFile(outputDir, masterPubKeys)

	if err != nil {
		fmt.Printf("Error while writing file: %v\n", err)
		panic(err)
	}

	fmt.Printf("Generated BTC Master Public Keys successfully!\n")
}
