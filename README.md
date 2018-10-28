# Payment channel CLI tool

Little CLI tool written a golang to show how payment channels work

## Instructions of use

1) Run `ganache` on port `7545`
2) Copy and paste the mnemonic and run `payment-channel init "<ADD MNEMONIC HERE>"`

## Further information

This tool was created to show a payment channel between two fictional characters, Bob and Alice.

To close this payment channel, both users need to sign a valid hashed message that includes

1) The contract address
2) The value
3) A nonce

Simply put, we do `keccack256(contractAddress, value, nonce)` and that will be the message that we sign with a private key.

To challenge this channel, only 1 signature from the opponent is needed, with a higher nonce than the one used when closing the channel.

Basic flow of the application:

`init` -> `deploy` -> `open` -> `sign` -> `close` -> `challenge` (maybe) -> `timewarp` -> `finalize`
