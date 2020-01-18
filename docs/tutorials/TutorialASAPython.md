Tutorial – Algorand Standard Assets (ASA)
=========================================

The Algorand platform is a general-purpose economic exchange system which
represents an extremely broad market. A given platform’s attractiveness and
effectiveness as a means of economic exchange can be defined by the combination
of what you can own and how you can transact.

Algorand Standard Assets represents Algorand’s ability to digitize any asset and
have both it and its ownership represented on chain. These assets could be
fungible (for example: currencies, stable coins, loyalty points, system credits,
in-game points, etc) or non-fungible (for example: real estate, collectables,
supply chain, in-game items, tickets, etc). In addition, our functionality
allows restrictions to be placed on the assets where needed (for example:
securities, certifications, compliance, etc).

Algorand has implemented named assets as a truly layer 1 asset. This allows any
asset created on Algorand to enjoy:

**Increased security** - New assets will enjoy the same security and safety as
Algos, the native currency on Algorand

**Inherent compatibility** - Apps that support any Algorand asset will support
all Algorand assets

**High ease of use** - Create your asset with a single transaction to the
network


Step 1A - Create 3 Accounts and add Algos to the Accounts
--------------------------------------------------

Assets are created at the account level. Each account can have up to 1000 assets for creator accounts as well as for consumer accounts. So, before starting the ASA tutorial, 3 new accounts will be created for this step for ASA transactions. Once created, copy off the account mnumonic and account address values.  

**Task 1A-1** Create an empty code file for desired language of choice (CreateNewAccounts.js, CreateNewAccounts.py, CreateNewAccounts.java, or CreateNewAccounts.go). Then simply copy the and insert the snippet. Update the node token and url, and run. Then copy off the account addresses and mnumonics. In Step 1B, we will paste those into the TutorialASA code and recover the account using the mnemonics created in this step. Mnumonics are for demonstration purposes. **NEVER** reveal secret mnemonics in practice.

In order to run ASA transactions, or any transactions for that matter, the accounts need to have TestNet Algo funds. Load all 3 accounts from the Algorand TestNet Dispenser which is located here: <https://bank.testnet.algorand.network/>

**Note:**
    See the appropriate GitHub repository for installing the supported SDKs:

[JavaScript](https://github.com/algorand/js-algorand-sdk)

[Python](https://github.com/algorand/py-algorand-sdk)

[Java](https://github.com/algorand/java-algorand-sdk)

[Go Lang](https://github.com/algorand/go-algorand-sdk)

Community SDKs 

[.NET / C#](https://github.com/RileyGe/dotnet-algorand-sdk) All Algorand 2.0 functions, including ASA, are available and demonstrated in the sample code github repository (just not in this tutorial yet)

[Rust (no Algorand 2.0 functions at this time)](http://mraof.com/temp/algosdk-doc/algosdk/)



**Note:**
    If you do not have a node setup, see these [instructions](https://developer.algorand.org/docs/introduction-installing-node). Also, another alternative is to use [PureStake](https://www.purestake.com/algorand-api).



```python tab="Python"
#python
import json
from algosdk import account, mnemonic

acct = account.generate_account()
address1 = acct[1]
print("Account 1")
print(address1)
mnemonic1 = mnemonic.from_private_key(acct[0])
print("Account 2")
acct = account.generate_account()
address2 = acct[1]
print(address2)
mnemonic2 = mnemonic.from_private_key(acct[0])
print("Account 3")
acct = account.generate_account()
address3 = acct[1]
print(address3)
mnemonic3 = mnemonic.from_private_key(acct[0])
print ("")
print("Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/")
print("copy off the following mnemonic code for use in Step 1B")
print("")
print("mnemonic1 = \"{}\"".format(mnemonic1))
print("mnemonic2 = \"{}\"".format(mnemonic2))
print("mnemonic3 = \"{}\"".format(mnemonic3))

# terminal output should look similar to this
# Account 1
# KECA5Z2ZILJOH2ZG7OPKJ5KMFXP5XBAOC7H36TLPJOQI3KB5UIYUD5XTZU
# Account 2
# DWQ4IA7EK5BKHSPNCJBA5SOVU66TKGUCCGO2SHQCI5UB2JAO3G2GWXMGPA
# Account 3
# TABDMZ2EUTNOR3S74SJWW37DLHE7BDGS5XB5JPLFQ2VQVJOE2DXKX722VU

# Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https: // bank.testnet.algorand.network/
# copy off the following mnemonic code for use in Step 1B

# mnemonic1 = "consider round clerk soldier hurt dynamic floor video output spoon deliver virtual zoo inspire rubber doll nose warfare improve abstract recall choice size above actor"
# mnemonic2 = "boil explain enlist adapt science hub universe knife ghost scheme lazy payment must gas coconut forget goddess author filter civil tumble antique delay absorb lend"
# mnemonic3 = "place elbow thumb bid taste strong sting tube swarm comic wave dinosaur congress sword zebra need proud primary brief rotate story pilot garbage abstract black"
```


Step 1B Setup Accounts, Utility Functions and Tools
--------------------------------------------------

This tutorial will use three TestNet accounts that have been pre-created in Step 1A. Be sure to dispense Algos to these accounts before continuing, using the TestNet Dispenser.

The TestNet dispenser is located here:
<https://bank.testnet.algorand.network/>

The accounts used in this tutorial are: (yours will be different)

Account 1
`THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM`

Account 2
`AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU`

Account 3
`3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU`

**Info:**
    You may want to verify account information periodically as well as transactions
    with asset information during the course of this tutorial. You can use either
    the [Algo TestNet Explorer](https://testnet.algoexplorer.io/) or use the Purestake's [Goalseeker](https://goalseeker.purestake.io/algorand/testnet), which also
    facilitates search by asset ID.

![Figure Step 1A-1 Use Purestake’s [Goalseeker](https://goalseeker.purestake.io/algorand/testnet) to search on Address, Transaction, Block or AssetID](../imgs/TutorialASA-01.png)
**Figure Step 1B-1** Purestake’s Goalseeker used to search Address,
Transaction, Block or AssetID.
<!-- <center>![Goalseeker](../imgs/TutorialASA-01.png)</center>
<center>**Figure Step 1A-1** Use Purestake’s Goalseeker to search Address,
Transaction, Block or AssetID.</center> -->

The tutorial code below is separated into snippets categorized by ASA core functions and is laid out in order. The solution should be coded as a single script for each respective language. 

**Task 1B-1** Create an empty code file for desired language of choice (TutorialASA.js, TutorialASA.py, TutorialASA.java, or TutorialASA.go). Then simply copy the code below and paste into the empty file. Then append the each snippet after the last line of code in the prior step as you read through this tutorial.


```python tab="Python"
# Python
import json
from algosdk import account, algod, mnemonic, transaction
# Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
# Change these values with the mnemonics from Step 1A
mnemonic1 = "PASTE your phrase for account 1 from Step 1A"
mnemonic2 = "PASTE your phrase for account 2 from Step 1A"
mnemonic3 = "PASTE your phrase for account 3 from Step 1A"

# For ease of reference, add account public and private keys to
# an accounts dict.

accounts = {}
counter = 1
for m in [mnemonic1, mnemonic2, mnemonic3]:
accounts[counter] = {}
accounts[counter]['pk'] = mnemonic.to_public_key(m)
accounts[counter]['sk'] = mnemonic.to_private_key(m)

counter += 1
# Specify your node address and token. This must be updated.
# algod_address = "" \# ADD ADDRESS
# algod_token = "" \# ADD TOKEN

algod_address = "http://hackathon.algodev.network:9100"
algod_token = "ef920e2e7e002953f4b29a8af720efe8e4ecc75ff102b165e0472834b25832c1"

# Initialize an algod client
algod_client = algod.AlgodClient(algod_token, algod_address)
# Get network params for transactions.
params = algod_client.suggested_params()
first = params.get("lastRound")
last = first + 1000
gen = params.get("genesisID")
gh = params.get("genesishashb64")
min_fee = params.get("minFee")
# Utility function to wait for a transaction to be confirmed by network
def wait_for_tx_confirmation(txid):
# Wait until the transaction's round info is confirmed, i.e. no longer 0
pendinginfo = algod_client.pending_transaction_info(txid)
while pendinginfo['round'] == 0:
pendinginfo = algod_client.pending_transaction_info(txid)
print("Account 1 address: {}".format(accounts[1]['pk']))
print("Account 2 address: {}".format(accounts[2]['pk']))
print("Account 3 address: {}".format(accounts[3]['pk']))


# your terminal output should look similar to the following

# Account 1 account: THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
# Account 2 account: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
# Account 3 account: 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
```

Step 2 Create a New Asset
---------------------------

The ability to create asserts directly on the blockcahin is an exciting capability of the Algorand Blockchain. Possible uses include currency, game leader boards, points in a loyalty system, shares of an asset, and securities such as stocks, bonds, and derivaties.  


**Info:**
    The decimals value determines the placement of the decimal. For example, when `decimals = 2`, and the `amount = 1000`, the acatual amount is 10.00. So, when a  transfer of amount of 10 is made, the actual transfer is .10




**Task 2-1** Account 1 creates an asset called latinum and sets Account 2 as the manager, reserve, freeze, and clawback address.


```python tab="Python"
# python
# insert after step 1 code
# Account 1 creates an asset called latinum and sets Account 2 as the manager, reserve, freeze, and clawback address.

data = {
    "sender": accounts[1]['pk'],
    "fee": min_fee,
    "first": first,
    "last": last,
    "gh": gh,
    "total": 1000,
    "default_frozen": False,
    "unit_name": "LATINUM",
    "asset_name": "latinum",
    "manager": accounts[2]['pk'],
    "reserve": accounts[2]['pk'],
    "freeze": accounts[2]['pk'],
    "clawback": accounts[2]['pk'],
    "url": "https://path/to/my/asset/details",
    "flat_fee": True,
    "decimals": 0
}

# Construct Asset Creation transaction
txn = transaction.AssetConfigTxn(**data)

# Sign with secret key of creator
stxn = txn.sign(accounts[1]['sk'])

# Send the transaction to the network and retrieve the txid.
txid = algod_client.send_transaction(stxn)
print(txid)

# Retrieve the asset ID of the newly created asset by first
# ensuring that the creation transaction was confirmed,
# then pulling account info of the creator and grabbing the
# asset with the max asset ID.

# Wait for the transaction to be confirmed
wait_for_tx_confirmation(txid)

try:
    # Pull account info for the creator
    account_info = algod_client.account_info(accounts[1]['pk'])
    # Get max asset ID
    asset_id = max(
        map(lambda x: int(x), account_info.get('thisassettotal').keys()))
    print("Asset ID: {}".format(asset_id))
    print(json.dumps(account_info['thisassettotal'][str(asset_id)], indent=4))
except Exception as e:
    print(e)

# terminal output should be similar to below
# BO6D5D2FKNNEOJLDCJBFEXQCQ5VOAFFZTKWTQJR7US2JJAM6N7RA
# Asset ID: 150003
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "total": 1000,
#     "decimals": 0,
#     "defaultfrozen": false,
#     "unitname": "LATINUM",
#     "assetname": "latinum",
#     "url": "https://path/to/my/asset/details",
#     "managerkey": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
#     "reserveaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
#     "freezeaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
#     "clawbackaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU"
# }
```


Step 3 Configure Asset Manager
------------------------------
Assets can be managed as to which accounts have roles for overall manager, reserve, freeze, and clawback functions. By default all of these roles are set to the creator account. 

Asset reconfiguration allows the address specified as manager to change any of the special addresses for the asset, such as the reserve address. To keep an address the same, it must be re-specified in each new configuration transaction. Supplying an empty address is the same as turning the associated feature off for this asset. Once a special address is set to the empty address, it can never change again. For example, if an asset configuration transaction specifying clawback="" were issued, the associated asset could never be revoked from asset holders, and clawback="" would be true for all time. The  strictEmptyAddressChecking argument can help with this behavior: when set to its default true, makeAssetConfigTxn will throw an error if any undefined management addresses are passed.

**Task 3-1** Here, the current manager (Account 2) issues an asset configuration transaction that assigns Account 1 as the new manager.


```python tab="Python"
# Python
# insert after Step 2 code
# Update manager address.
# The current manager(Account 2) issues an asset configuration transaction 
# that assigns Account 1 as the new manager.
# Keep reserve, freeze, and clawback address same as before, i.e. account 2
data = {
    "sender": accounts[2]['pk'],
    "fee": min_fee,
    "first": first,
    "last": last,
    "gh": gh,
    "index": asset_id,
    "manager": accounts[1]['pk'],
    "reserve": accounts[2]['pk'],
    "freeze": accounts[2]['pk'],
    "clawback": accounts[2]['pk'],
    "flat_fee": True
}
txn = transaction.AssetConfigTxn(**data)
# sign by the current manager
stxn = txn.sign(accounts[2]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)

# Wait for the transaction to be confirmed
wait_for_tx_confirmation(txid)

# Check asset info to view change in management.
asset_info = algod_client.asset_info(asset_id)
print(json.dumps(asset_info, indent=4))

# terminal output should be similar to...
# CPZZG7ZTMYKXUYDFHLPGKL4E6FF5BONF5ZZCD5U6PAJPNUMSSGTQ
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "total": 1000,
#     "decimals": 0,
#     "defaultfrozen": false,
#     "unitname": "LATINUM",
#     "assetname": "latinum",
#     "url": "https://path/to/my/asset/details",
#     "managerkey": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "reserveaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
#     "freezeaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU",
#     "clawbackaddr": "AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU"
# }
```


Step 4 Opt-in to Receive Asset
------------------------------

Once the asset has been created, the next thing to do is send assets to other accounts. 

Before a user can begin transacting with an asset, the user must first issue an asset acceptance transaction. This is a special case of the asset transfer transaction, where the user sends 0 assets to themself. After issuing this transaction, the user can begin transacting with the asset. Each new accepted asset increases the user's minimum balance.

**Task 4-1** Account 3 opts-in to receive the new asset by sending a 0 amount transfer of the asset to itself.


```python tab="Python"
# Python
# insert after Step 3 code

# Check if asset_id is in account 2's asset holdings prior to opt-in
account_info = algod_client.account_info(accounts[3]['pk'])
holding = None
if 'assets' in account_info:
    holding = account_info['assets'].get(str(asset_id))

if not holding:
    # Get latest network parameters
    data = {
        "sender": accounts[3]['pk'],
        "fee": min_fee,
        "first": first,
        "last": last,
        "gh": gh,
        "receiver": accounts[3]["pk"],
        "amt": 0,
        "index": asset_id,
        "flat_fee": True
    }

    # Use the AssetTransferTxn class to transfer assets
    txn = transaction.AssetTransferTxn(**data)
    stxn = txn.sign(accounts[3]['sk'])
    txid = algod_client.send_transaction(stxn)
    print(txid)
    # Wait for the transaction to be confirmed
    wait_for_tx_confirmation(txid)
    # Now check the asset holding for that account.
    # This should now show a holding with a balance of 0.
    account_info = algod_client.account_info(accounts[3]['pk'])
    print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

# terminal output should look similar to this...

# 4ZSLST43HZZQNTF7Y2ZM5ENWLHOTS3XCZX574YEPDUI3PFUSODKQ
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "amount": 0,
#     "frozen": false
# }
```


Step 5 Transfer an Asset
------------------------

Transfer an asset allows users to transact with assets, after they have issued asset acceptance transactions. The optional closeRemainderTo argument can be used to stop transacting with a particular asset. Now that the opt-in has been done on a potential receiving account, assets can be transferred.

**Note:**
    A frozen account can always close out to the asset creator.

**Task 5-1** This code has Account 1 sending 10 latinum to Account 3.


```python tab="Python"
# Python
# insert after Step 4 code

# transfer asset from account 1 to account 3
data = {
    "sender": accounts[1]['pk'],
    "fee": min_fee,
    "first": first,
    "last": last,
    "gh": gh,
    "receiver": accounts[3]["pk"],
    "amt": 10,
    "index": asset_id,
    "flat_fee": True
}

txn = transaction.AssetTransferTxn(**data)
stxn = txn.sign(accounts[1]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_tx_confirmation(txid)
# The balance should now be 10.
account_info = algod_client.account_info(accounts[3]['pk'])
print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

# terminal output should look similar to this...

# BX72JE7T2NTOO33VZFK75MT4JNQXFVEUOMFKV2LRPZFH5VDL27AQ
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "amount": 10,
#     "frozen": false
# }
```


Step 6 Freeze an Asset
-----------------------

To freeze or unfreeze an asset, this transaction must be sent from the account specified as the freeze manager for the asset.

**Task 6-1** The freeze address (Account 2) freezes Account 3's latinum holdings.


```python tab="Python"
# Python
# insert code after Step 5
# The freeze address (Account 2) freezes Account 3's latinum holdings.
data = {
    "sender": accounts[2]['pk'],
    "fee": min_fee,
    "first": first,
    "last": last,
    "gh": gh,
    "index": asset_id,
    "target": accounts[3]["pk"],
    "new_freeze_state": True
}

txn = transaction.AssetFreezeTxn(**data)
stxn = txn.sign(accounts[2]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_tx_confirmation(txid)
# The balance should now be 10.
account_info = algod_client.account_info(accounts[3]['pk'])
print(json.dumps(account_info['assets'][str(asset_id)], indent=4))


# terminal output should look similar to this...
# ZRSYHNYRMF3A2HCHWN4RKDFLMCOF6TASFGSKQSJZP4XZN3KZGOJA
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "amount": 10,
#     "frozen": true
# }
```


Step 7 Revoke an Asset
----------------------

Revoking an asset allows an asset's revocation manager to transfer assets on behalf of another user. It will only work when issued by the asset's revocation manager.

**Task 7-1** The clawback address (Account 2) revokes 10 latinum from Account 3 and places it back with Account 1.


```python tab="Python"
# Python
# insert after Step 6 code
# The clawback address (Account 2) revokes 10 latinum from Account 3 and places it back with Account 1.
data = {
    "sender": accounts[2]['pk'],
    "fee": min_fee,
    "first": first,
    "last": last,
    "gh": gh,
    "receiver": accounts[1]["pk"],
    "amt": 10,
    "index": asset_id,
    "revocation_target": accounts[3]['pk'],
    "flat_fee": True
}
# Must be signed by the account that is the clawback address
txn = transaction.AssetTransferTxn(**data)
stxn = txn.sign(accounts[2]['sk'])
txid = algod_client.send_transaction(stxn)
print(txid)
# Wait for the transaction to be confirmed
wait_for_tx_confirmation(txid)
# The balance of account 3 should now be 0.
account_info = algod_client.account_info(accounts[3]['pk'])
print(json.dumps(account_info['assets'][str(asset_id)], indent=4))
# The balance of account 1 should increase by 10 to 1000.
account_info = algod_client.account_info(accounts[1]['pk'])
print(json.dumps(account_info['assets'][str(asset_id)], indent=4))

# terminal output should be similar to..

# 3BAJ6EO6V5EVQRUO6BDGBUXKQVWTQ7AFXRRSDQHWE7Q6QANDFVOQ
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "amount": 0,
#     "frozen": true
# }
# {
#     "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM",
#     "amount": 1000,
#     "frozen": false
# }
```


Step 8 Destroy an Asset
-----------------------

Asset destruction allows the creator to remove the asset from the ledger, if all outstanding assets are held by the creator. 

**Task 8-1** With all assets back in the creator's account, the manager (Account 1) destroys the asset.


```python tab="Python"
# Python
# insert code after Step 7's code]
# With all assets back in the creator's account,
# the manaager (Account 1) destroys the asset.

data = {
    "sender": accounts[1]['pk'],
    "fee": min_fee,
    "first": first,
    "last": last,
    "gh": gh,
    "index": asset_id,
    "strict_empty_address_check": False,
    "flat_fee": True
}

# Construct Asset Creation transaction
txn = transaction.AssetConfigTxn(**data)

# Sign with secret key of creator
stxn = txn.sign(accounts[1]['sk'])

# Send the transaction to the network and retrieve the txid.
txid = algod_client.send_transaction(stxn)
print(txid)

# Wait for the transaction to be confirmed
wait_for_tx_confirmation(txid)



# This should raise an exception since the asset was deleted.
try:
    asset_info = algod_client.asset_info(asset_id)
except Exception as e:
    print(e)
# terminal output should look similar to this
# HSH2T3E6Z4BMV5DZE47IWDNVMUF7DSHPLQPCQTC6X4NHTQCTWRHQ
# failed to retrieve asset creator from the ledger
```


Conclusion
----------

ASA is a very powerful layer 1 feature of the Algorand Blockchain. We created an asset in this tutotial and showed how to do the following functions:

* Create
* Opt-In
* Manage
* Transfer
* Freeze
* Revoke
* Destroy

Source code for the completed solution can be found here <https://github.com/rfustino/TutorialASA>.
