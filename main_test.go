package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	privMiningKeys := []string{
		"12NjwfVK7YoBENaZwmHHPUaspK3uUtv3M22761TFk3H3FEuT83U",
		"12pLJqKTpcyQxdkUBthhydx2PKSt7G5YhBDRGjVWquGZLB8Mrvk",
		"1mGVTi2HmraVLx5vpLgpYGT9ZT76pNbb5Hh222hurZJf1257uE",
		"1vd6YYHi4MdPib8YymiFkdEHpZ3qNUKmqQVDGpwYYGgtDN3dhy",
	}
	seeds, err := generateSeedsFromMiningKey(privMiningKeys)
	if err != nil {
		fmt.Printf("Error while generating seeds: %v\n", err)
		panic(err)
	}

	masterPubKeys := generateBTCMasterPubKeysFromSeeds(seeds)
	for _, pubKey := range masterPubKeys {
		fmt.Printf("%#v\n", pubKey)
	}
}
