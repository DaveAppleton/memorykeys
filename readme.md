# Memory Keys

This library allows us to generate and use keys on an ad-hoc basis without storing them for re-use on a subsequent run. 

We reference these keys by name and obtain

- the private key (for signing)
- a transaction object encapsulating that key for future use with ABIGEN
- the address

## Restrictions

It is not expected to be necessary for concurrent creation of keys so keys are stored in a map which is not safe for concurrent writes.

## Functions

`GetPrivateKey(keyname)` create a keypair associated with the name on first call, subsequent calls return the previous value.

``` go

   launcher,err := memorykeys.GetPrivateKey("launcher")

```

`GetAddress(keyName string)` gets the address associated with a key. Creates the key if it does not exist

``` go

   recipient,err := GetAddress("recipient")

```

`GetTransactor(keyName string)` gets a transaction object for use with ABIGEN objects. Creates the key if necessary.

``` go

   deployer, err := GetTransactor("deployer")

```

`ImportPrivateKey(keyName, hexKey)` imports a hex encoded private key for use.

**BE CAREFUL NOT TO USE PRODUCTION KEYS**

``` go

    privateKey := "d31a46c5322e8e8a7e11f51cf9c4073fea42d33b431b5e7e76a82518fc178ea8"
    key, err := ImportPrivateKey("imported", privateKey)

```
