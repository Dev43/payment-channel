# Payment channel CLI tool

## THIS IS STILL A WORK IN PROGRESS

## TODO

- [X] Init (creates the file and 2 new keys for bob and alice)
- [X] Version
- [ ] Deploy -- deploys our contract with the needed initial value (for now only alice can deploy the contract)
- [ ] Open a channel (opens a channel with the needed initial value -- only alice can open the channel)
- [ ] Sign -- create a signature with the needed arguments so both bob AND alice sign it
- [ ] Verify -- verifies the signature given and returns a boolean.
- [ ] Close a channel (closes the channel with a specific set of signatures -- only alice can od it)
- [ ] Challenge a channel
- [ ] Check balances
- [ ] Output current info

    - [ ] on chain balance
    - [ ] payment channel balance
    - [ ] latest transaction message
    - [ ] Alice's and Bob's public and private key