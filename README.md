# BTC Master PubKeys Generator

- ***Input:*** Incognito Private Mining Keys (Mining Key could be also called Private Seed or Validator Key)
- ***Output:*** BTC Master PubKeys used for Portal V4

## USAGE
1. Prepare the mining keys file. See an example at ``` template_private_mining_keys.txt```
2. Run script
```bash
INPUT={{mining_keys_file_directory}} OUTPUT={{output_file_directory}} ./run.sh
```

Example:
```bash
INPUT=private_mining_keys.txt OUTPUT=master_keys.txt ./run.sh
```
