# Payment channel CLI tool

Little CLI tool written a golang to show how payment channels work

## TODO

- [X] Init (creates the file and 2 new keys for bob and alice)
- [X] Version
- [X] Deploy -- deploys our contract with the needed initial value (for now only alice can deploy the contract)
- [X] Open a channel (opens a channel with the needed initial value -- only alice can open the channel)
- [X] Sign -- create a signature with the needed arguments so both bob AND alice sign it
- [X] Verify -- verifies the signature given and returns a boolean.
- [X] Close a channel (closes the channel with a specific set of signatures -- only alice can od it)
- [X] Challenge a channel
- [X] Check balances
- [X] Output current info 
    - [X] on chain balance
    - [X] payment channel balance
    - [X] latest transaction message