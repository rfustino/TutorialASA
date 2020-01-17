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

**Info:**
    See the appropriate GitHub repository for installing the supported SDKs:

[JavaScript](https://github.com/algorand/js-algorand-sdk)

[Python](https://github.com/algorand/py-algorand-sdk)

[Java](https://github.com/algorand/java-algorand-sdk)

[Go Lang](https://github.com/algorand/go-algorand-sdk)

Community SDKs 

[.NET / C#](https://github.com/RileyGe/dotnet-algorand-sdk) All Algorand 2.0 functions, including ASA, are available and demonstrated in the sample code github repository (just not in this tutorial yet)

[Rust (no Algorand 2.0 functions at this time)](http://mraof.com/temp/algosdk-doc/algosdk/)



**Info:**
    If you do not have a node setup, see these [instructions](https://developer.algorand.org/docs/introduction-installing-node). Also, another alternative is to use [PureStake](https://www.purestake.com/algorand-api).


```javascript tab="JavaScript"
//JavaScript
const algosdk = require('algosdk');

// This code generates 3 accounts
// Once created, copy off the values which we will past into the TutorialASA code in the next Step
// Then, Add funds to all three
// The Algorand TestNet Dispenser is located here: 
// https://bank.testnet.algorand.network/

// Retrieve the token, server and port values for your installation in the algod.net
// and algod.token files within the data directory
// or use PureStake API service https://www.purestake.com/

// UPDATE THESE VALUES
const token = "TOKEN";
const server = "SERVER";
const port = PORT;

var acct = null;

acct = algosdk.generateAccount();

account1 = acct.addr;
console.log("Account 1 = " + account1);
var account1_mnemonic = algosdk.secretKeyToMnemonic(acct.sk);
console.log("Account Mnemonic 1 = "+ account1_mnemonic);
var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
var isValid = algosdk.isValidAddress(recoveredAccount1.addr);
console.log("Is this a valid address: " + isValid);
console.log("Account created. Save off Mnemonic and address");

acct = algosdk.generateAccount();

account2 = acct.addr;
console.log("Account 2 = " + account2);
var account2_mnemonic = algosdk.secretKeyToMnemonic(acct.sk);
console.log("Account Mnemonic 2 = " +account2_mnemonic);
var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
var isValid = algosdk.isValidAddress(recoveredAccount2.addr);
console.log("Is this a valid address: " + isValid);
console.log("Account created. Save off Mnemonic and address");

acct = algosdk.generateAccount();

account3 = acct.addr;
console.log("Account 3 = " + account3);
var account3_mnemonic = algosdk.secretKeyToMnemonic(acct.sk);
console.log("Account Mnemonic 3 = " +account3_mnemonic);
var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
var isValid = algosdk.isValidAddress(recoveredAccount3.addr);
console.log("Is this a valid address: " + isValid);
console.log("Account created. Save off Mnemonic and address");
console.log("");
console.log("Add funds to all of these accounts using the TestNet Dispenser at https://bank.testnet.algorand.network/ ");
console.log("");
console.log("Copy off these 3 lines of code and they will be pasted in the ASA Tutorial code");
console.log("");
console.log("var account1_mnemonic = \"" + account1_mnemonic + "\"");
console.log("var account2_mnemonic = \"" + account2_mnemonic + "\"");
console.log("var account3_mnemonic = \"" + account3_mnemonic + "\"");
//
// your terminal output should look similar to this:

// Account 1 = GNZLAV7QIYGISQ4EYY4Q2LUKHBXKXKPVMQJDWLZXDY57Z46HATOFMBWP7A
// Account Mnemonic 1 = hidden glove heart now claw away business behind echo best select merry secret upset shed margin outer reflect holiday rib among script scorpion above oval
// Is this a valid address: true
// Account created.Save off Mnemonic and address

// Account 2 = 4N6OZZU73MUIS7IWW2OVSDXRIK6JSI5QAOWHU2CZGIVF4WNQUYFOMY7UEA
// Account Mnemonic 2 = plug cabin shuffle hope similar exact enable speed amateur shift ramp machine scrap comfort shove carry test marble hip protect melody win wink ability hybrid
// Is this a valid address: true
// Account created.Save off Mnemonic and address

// Account 3 = LCNNXW27VDCFYZI3AERPBHJ3DWIC3PJ4DQYDQBFHWR4UNDYFZVYT2VJ5LU
// Account Mnemonic 3 = divert feel ball make absorb deny student turn fuel north slice wolf small tribe pull work scout empty nose riot zoo prepare loan absent kingdom
// Is this a valid address: true
// Account created.Save off Mnemonic and address

// Add funds to all of these accounts using the TestNet Dispenser at https://bank.testnet.algorand.network/ 

// Copy off these 3 lines of code and they will be pasted in the ASA Tutorial code

// var account1_mnemonic = "hidden glove heart now claw away business behind echo best select merry secret upset shed margin outer reflect holiday rib among script scorpion above oval"
// var account2_mnemonic = "plug cabin shuffle hope similar exact enable speed amateur shift ramp machine scrap comfort shove carry test marble hip protect melody win wink ability hybrid"
// var account3_mnemonic = "divert feel ball make absorb deny student turn fuel north slice wolf small tribe pull work scout empty nose riot zoo prepare loan absent kingdom"
```

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

```java tab="Java"
//Java
package com.algorand.Tutorials;
import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.crypto.Address;
public class CreateNewAccounts {
public static void main(final String args[]) throws Exception {
// Create a random new account
Account act1 = new Account();
// Get the new account address
Address addr1 = act1.getAddress();
// Get the backup phrase
String backup1 = act1.toMnemonic();
System.out.println("Account Address 1: " + addr1.toString());
// Create a random new account
Account act2 = new Account();
// Get the new account address
Address addr2 = act2.getAddress();
// Get the backup phrase
String backup2 = act2.toMnemonic();
System.out.println("Account Address 2: " + addr2.toString());
// Create a random new account
Account act3 = new Account();
// Get the new account address
Address addr3 = act3.getAddress();
// Get the backup phrase
String backup3 = act3.toMnemonic();
System.out.println("Account Address 3: " + addr3.toString());
System.out.println("");
System.out.println(
        "Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/");
System.out.println("");
System.out.println("copy off the following code for use in Step 1B");
System.out.println("");
System.out.println("final String account1_mnemonic = \"" + backup1 + "\"");
System.out.println("final String account2_mnemonic = \"" + backup2 + "\"");
System.out.println("final String account3_mnemonic = \"" + backup3 + "\"");
// Terminal output should look similar to this...

// Account Address 1: PSF2S7OTAIP6YEQBOF4R7JOYDBZMTMAG7KCAUCQUM6UPCELHNRS5LQPS3M
// Account Address 2: FIBG65O25CMDTEG5ARZBDWABUDO3Z6ZRVNTAJN4DCLYCI6Z4YZFSEWF55M
// Account Address 3: AWICWFUDL5CO3XYSWGRMXLS7ML2RGOSCWQNXNSFVR2VOP73P6IMEZ7FUFY

// Copy off accounts above and add TestNet Algo funds using the TestNet
// Dispenser at https://bank.testnet.algorand.network/

// copy off the following code for use in Step 1B

// final String account1_mnemonic = "tongue mention field lottery diagram yellow
// walk erupt supply link bike setup learn excite recycle lunar flag dash lucky
// dad evoke refuse ticket able please"
// final String account2_mnemonic = "over fantasy smile pupil visual film club
// balcony jacket tilt amateur whip crisp allow clay govern canoe fat group
// develop sorry test unknown absent dinner"
// final String account3_mnemonic = "category sick rough intact design board
// episode hedgehog build bundle siege voyage rent season fossil weird cannon
// pipe grow winner wall urge report about weapon"
}
}
```

```go tab="Go"
//Go
package main

import (
	json "encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	account1 := crypto.GenerateAccount()
	account2 := crypto.GenerateAccount()
	account3 := crypto.GenerateAccount()
	address1 := account1.Address.String()
	address2 := account2.Address.String()
	address3 := account3.Address.String()

	mnemonic1, err := mnemonic.FromPrivateKey(account1.PrivateKey)
	if err != nil {
		return
	}
	mnemonic2, err := mnemonic.FromPrivateKey(account2.PrivateKey)
	if err != nil {
		return
	}
	mnemonic3, err := mnemonic.FromPrivateKey(account3.PrivateKey)
	if err != nil {
		return
	}
	fmt.Printf("Account 1 = %s\n", address1)
	fmt.Printf("Account 2 = %s\n", address2)
	fmt.Printf("Account 3 = %s\n", address3)
	fmt.Printf("")
	fmt.Printf("Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/\n")	
	fmt.Printf("Copy off the following mnemonic code for use in Step 1B\n")
	fmt.Printf("\n")	
	fmt.Printf("mnemonic1 := \"%s\"\n", mnemonic1)
	fmt.Printf("mnemonic2 := \"%s\"\n", mnemonic2)
	fmt.Printf("mnemonic3 := \"%s\"\n", mnemonic3)
}

// Terminal output should look similar to this...

// Account 1 = BQBRKJ7KSQO6WWX3QRSITPCZOEDOT6DX6MJVQQFQI72PDPTKWU5WKOOJ2I
// Account 2 = LBMZ2DNGB5N7AJNZQZPMTPDJRW3S6OB6YR4EHE5RYODRS67TILA5H23K4A
// Account 3 = 6AVEBUP35AHIKDA6UGSFM6LP433AUHVRZNW4KHAXOYCOYDYOLFJWZ7J3I4
// Copy off accounts above and add TestNet Algo funds using the TestNet Dispenser at https://bank.testnet.algorand.network/
// Copy off the following mnemonic code for use in Step 1B

// mnemonic1 := "fatigue laundry small early radar code supreme diary skin record slice distance bike skirt guard surround miss turtle horror frame train taxi column ability forest"
// mnemonic2 := "ill live coconut risk east flower snack develop boring enroll nest rice mistake pioneer rival account coffee depend bachelor orient rebel inform throw absent option"
// mnemonic3 := "broom bid found recall stick gas sample copy network mistake mind relief rely file disorder east asthma program filter hedgehog legal walnut wait about slogan"
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


```javascript tab="JavaScript"
//JavaScript
const algosdk = require('algosdk');
//Retrieve the token, server and port values for your installation in the algod.net
//and algod.token files within the data directory
// UPDATE THESE VALUES
const token = "TOKEN";
const server = "SERVER";
const port = PORT;
// Structure for changing blockchain params
var cp = {
fee: 0,
firstRound: 0,
lastRound: 0,
genID: "",
genHash: ""
}

// Utility function to update params from blockchain
var getChangingParms = async function( algodclient ) {
let params = await algodclient.getTransactionParams();
cp.firstRound = params.lastRound;
cp.lastRound = cp.firstRound + parseInt(1000);
let sfee = await algodclient.suggestedFee();
cp.fee = sfee.fee;
cp.genID = params.genesisID;
cp.genHash = params.genesishashb64;
}

// Function used to wait for a tx confirmation
var waitForConfirmation = async function(algodclient, txId) {
while (true) {
b3 = await algodclient.pendingTransactionInformation(txId);
if (b3.round != null && b3.round \> 0) {
//Got the completed Transaction
console.log("Transaction " + b3.tx + " confirmed in round " + b3.round);
break;
}
}
};

// Recover accounts created in Step 1A
// paste in mnemonic phrases here for each account

var account1_mnemonic = "PASTE your phrase for account 1 from Step 1A";
var account2_mnemonic = "PASTE your phrase for account 2 from Step 1A";
var account3_mnemonic = "PASTE your phrase for account 3 from Step 1A"


var recoveredAccount1 = algosdk.mnemonicToSecretKey(account1_mnemonic);
var recoveredAccount2 = algosdk.mnemonicToSecretKey(account2_mnemonic);
var recoveredAccount3 = algosdk.mnemonicToSecretKey(account3_mnemonic);
console.log(recoveredAccount1.addr);
console.log(recoveredAccount2.addr);
console.log(recoveredAccount3.addr);
// Instantiate the algod wrapper
let algodclient = new algosdk.Algod(token, server, port);
//console/terminal output should look similar to this
//THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
//AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
//3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
```


Step 2 Create a New Asset
---------------------------

The ability to create asserts directly on the blockcahin is an exciting capability of the Algorand Blockchain. Possible uses include currency, game leader boards, points in a loyalty system, shares of an asset, and securities such as stocks, bonds, and derivaties.  


**Info:**
    The decimals value determines the placement of the decimal. For example, when `decimals = 2`, and the `amount = 1000`, the acatual amount is 10.00. So, when a  transfer of amount of 10 is made, the actual transfer is .10




**Task 2-1** Account 1 creates an asset called latinum and sets Account 2 as the manager, reserve, freeze, and clawback address.

```javascript tab="JavaScript"
// JavaScript
(async () => {
    // Asset Creation:
    // The first transaciton is to create a new asset
    // Get last round and suggested tx fee
    // We use these to get the latest round and tx fees
    // These parameters will be required before every 
    // Transaction
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);
    let note = undefined; // arbitrary data to be stored in the transaction; here, none is stored
    //Asset creation specific parameters
    // The following parameters are asset specific
    // Throughout the example these will be re-used. 
    // We will also change the manager later in the example
    let addr = recoveredAccount1.addr;
    // Whether user accounts will need to be unfrozen before transacting    
    let defaultFrozen = false;
    // integer number of decimals for asset unit calculation
    let decimals = 0;
    // total number of this asset available for circulation   
    let totalIssuance = 1000;
    // Used to display asset units to user    
    let unitName = "LATINUM";
    // Friendly name of the asset    
    let assetName = "latinum";
    // Optional string pointing to a URL relating to the asset
    let assetURL = "http://someurl";
    // Optional hash commitment of some sort relating to the asset. 32 character length.
    let assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
    // The following parameters are the only ones
    // that can be changed, and they have to be changed
    // by the current manager
    // Specified address can change reserve, freeze, clawback, and manager
    let manager = recoveredAccount2.addr;
    // Specified address is considered the asset reserve
    // (it has no special privileges, this is only informational)
    let reserve = recoveredAccount2.addr;
    // Specified address can freeze or unfreeze user asset holdings 
    let freeze = recoveredAccount2.addr;
    // Specified address can revoke user asset holdings and send 
    // them to other addresses    
    let clawback = recoveredAccount2.addr;

    // signing and sending "txn" allows "addr" to create an asset
    let txn = algosdk.makeAssetCreateTxn(addr, cp.fee, cp.firstRound, cp.lastRound, note,
        cp.genHash, cp.genID, totalIssuance, decimals, defaultFrozen, manager, reserve, freeze,
        clawback, unitName, assetName, assetURL, assetMetadataHash);

    let rawSignedTxn = txn.signTxn(recoveredAccount1.sk)
    let tx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + tx.txId);
    let assetID = null;
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, tx.txId);
    // Get the new asset's information from the creator account
    let ptx = await algodclient.pendingTransactionInformation(tx.txId);
    assetID = ptx.txresults.createdasset;
    console.log("AssetID = " + assetID);

    // your terminal/console output should be similar to this
    // Transaction: RXSAJUYVPDWUF4XNGA2VYQX3NUVT5YJEZZ5SJXIIASZK5M55LVVQ
    // Transaction RXSAJUYVPDWUF4XNGA2VYQX3NUVT5YJEZZ5SJXIIASZK5M55LVVQ confirmed in round 4272786
    // AssetID = 149657

})().catch(e => {
    console.log(e);
    console.trace();
});

```


Step 3 Configure Asset Manager
------------------------------
Assets can be managed as to which accounts have roles for overall manager, reserve, freeze, and clawback functions. By default all of these roles are set to the creator account. 

Asset reconfiguration allows the address specified as manager to change any of the special addresses for the asset, such as the reserve address. To keep an address the same, it must be re-specified in each new configuration transaction. Supplying an empty address is the same as turning the associated feature off for this asset. Once a special address is set to the empty address, it can never change again. For example, if an asset configuration transaction specifying clawback="" were issued, the associated asset could never be revoked from asset holders, and clawback="" would be true for all time. The  strictEmptyAddressChecking argument can help with this behavior: when set to its default true, makeAssetConfigTxn will throw an error if any undefined management addresses are passed.

**Task 3-1** Here, the current manager (Account 2) issues an asset configuration transaction that assigns Account 1 as the new manager.

**Note:**
    With the JavaScript code solution, you will need to paste snippets above this catch statement.

```javascript
    })().catch(e => {
    console.log(e);
    console.trace();
});
```

```javascript tab="JavaScript"
    // JavaScript
    // add this code after console.log("AssetID = " + assetID); in Step 2 above
    
    // Change Asset Configuration:
    // Change the manager using an asset configuration transaction

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);

    // Asset configuration specific parameters
    // all other values are the same so we leave 
    // Them set.
    // specified address can change reserve, freeze, clawback, and manager
    manager = recoveredAccount1.addr;

    // Note that the change has to come from the existing manager
    let ctxn = algosdk.makeAssetConfigTxn(recoveredAccount2.addr, cp.fee,
        cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID,
        assetID, manager, reserve, freeze, clawback);

    // This transaction must be signed by the current manager
    rawSignedTxn = ctxn.signTxn(recoveredAccount2.sk)
    let ctx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + ctx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ctx.txId);

    //Get the asset information for the newly changed asset
    let assetInfo = await algodclient.assetInformation(assetID);
    //The manager should now be the same as the creator

    console.log(assetInfo);
// Asset info should look similar to this in the Terminal Output:
    // Transaction: QGXMMYIKRV3TE3NWCL6FVNJQ3DCODML3GBC3BR4TSB3D6S3YQVCQ
    // Transaction QGXMMYIKRV3TE3NWCL6FVNJQ3DCODML3GBC3BR4TSB3D6S3YQVCQ confirmed in round 4273376
    // {
    //     creator: 'THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM',
    //     total: 1000,
    //     decimals: 0,
    //     defaultfrozen: false,
    //     unitname: 'LATINUM',
    //     assetname: 'latinum',
    //     url: 'http://someurl',
    //     metadatahash: 'MTZlZmFhMzkyNGE2ZmQ5ZDNhNDgyNDc5OWE0YWM2NWQ=',
    //     managerkey: 'THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM',
    //     reserveaddr: 'AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU',
    //     freezeaddr: 'AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU',
    //     clawbackaddr: 'AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU'
    // }
```


Step 4 Opt-in to Receive Asset
------------------------------

Once the asset has been created, the next thing to do is send assets to other accounts. 

Before a user can begin transacting with an asset, the user must first issue an asset acceptance transaction. This is a special case of the asset transfer transaction, where the user sends 0 assets to themself. After issuing this transaction, the user can begin transacting with the asset. Each new accepted asset increases the user's minimum balance.

**Task 4-1** Account 3 opts-in to receive the new asset by sending a 0 amount transfer of the asset to itself.

```javascript tab="JavaScript"
    // JavaScript
    // insert this code after console.log(assetInfo); in Step 3

    // Opting in to an Asset:
    // Opting in to transact with the new asset
    // Allow accounts that want recieve the new asset
    // Have to opt in. To do this they send an asset transfer
    // of the new asset to themseleves 
    // In this example we are setting up the 3rd recovered account to 
    // receive the new asset
    let sender = recoveredAccount3.addr;
    let recipient = sender;
    // We set revocationTarget to undefined as 
    // This is not a clawback operation
    let revocationTarget = undefined;
    // CloseReaminerTo is set to undefined as
    // we are not closing out an asset
    let closeRemainderTo = undefined;
    // We are sending 0 assets
    amount = 0;

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);

    // signing and sending "txn" allows sender to begin accepting asset specified by creator and index
    let opttxn = algosdk.makeAssetTransferTxn(sender, recipient, closeRemainderTo, revocationTarget,
        cp.fee, amount, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);

    // Must be signed by the account wishing to opt in to the asset    
    rawSignedTxn = opttxn.signTxn(recoveredAccount3.sk);
    let opttx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + opttx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, opttx.txId);

    //You should now see the new asset listed in the account information
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets[assetID]));

// your console/terminal out put should look as follows
//   Transaction YT2U2WWBUWB4P54M24OUCMIIN3VDZ64LIDCBMED6AGQIQZKT6PTQ confirmed in round 4273745
//   Account Information for: { "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM", "amount": 0, "frozen": false }

```


Step 5 Transfer an Asset
------------------------

Transfer an asset allows users to transact with assets, after they have issued asset acceptance transactions. The optional closeRemainderTo argument can be used to stop transacting with a particular asset. Now that the opt-in has been done on a potential receiving accout, assets can be transferred.

**Note:**
    A frozen account can always close out to the asset creator.

**Task 5-1** This code has Account 1 sending 10 latinum to Account 3.

```javascript tab="JavaScript"
// JavaScript
// add this code after Step 4's code - console.log("Account Information for: " + JSON.stringify(act.assets[assetID]));
  
    // Transfer New Asset:
    // Now that account3 can recieve the new tokens 
    // we can tranfer tokens in from the creator
    // to account3
    sender = recoveredAccount1.addr;
    recipient = recoveredAccount3.addr;
    revocationTarget = undefined;
    closeRemainderTo = undefined;
    //Amount of the asset to transfer
    amount = 10;

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);

    // signing and sending "txn" will send "amount" assets from "sender" to "recipient"
    let xtxn = algosdk.makeAssetTransferTxn(sender, recipient, closeRemainderTo, revocationTarget,
        cp.fee, amount, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
    // Must be signed by the account sending the asset  
    rawSignedTxn = xtxn.signTxn(recoveredAccount1.sk)
    let xtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + xtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, xtx.txId);

    // You should now see the 10 assets listed in the account information
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets[assetID]));
// your console/terminal out put should look similar to this:
//    Transaction LM2QOJP6FKLX2XIKCDLTM37IEE2IGQ5IMVOEYOCFHKOAUMU5G6ZQ confirmed in round 4273946
//    Account Information for: { "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM", "amount": 10, "frozen": false }
```


Step 6 Freeze an Asset
-----------------------

To freeze or unfreeze an asset, this transaction must be sent from the account specified as the freeze manager for the asset.

**Task 6-1** The freeze address (Account 2) freezes Account 3's latinum holdings.


```javascript tab="JavaScript"
// JavaScript
// insert this code after Step 5's code

// The asset was created and configured to allow freezing an account
    // If the freeze address is set "", it will no longer be possible to do this.
    // In this example we will now freeze account3 from transacting with the 
    // The newly created asset. 
    // Thre freeze transaction is sent from the freeze acount
    // Which in this example is account2 
    from = recoveredAccount2.addr;
    freezeTarget = recoveredAccount3.addr;
    freezeState = true;

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);


    // The freeze transaction needs to be signed by the freeze account
    let ftxn = algosdk.makeAssetFreezeTxn(from, cp.fee, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID,
        assetID, freezeTarget, freezeState)

    // Must be signed by the freeze account   
    rawSignedTxn = ftxn.signTxn(recoveredAccount2.sk)
    let ftx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + ftx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, ftx.txId);

    // You should now see the asset is frozen listed in the account information
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets[assetID]));

//you should see console/terminal output similar to this witht he frozen vales set to true
//    Transaction: CZ7VYA7UFGRVJ4EHQOLKCPRESDJC4ULHUTWM3QCAPZR3K4KVT2IQ
//    Transaction CZ7VYA7UFGRVJ4EHQOLKCPRESDJC4ULHUTWM3QCAPZR3K4KVT2IQ confirmed in round 4274065
//    Account Information for: { "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM", "amount": 10, "frozen": true }
```


Step 7 Revoke an Asset
----------------------

Revoking an asset allows an asset's revocation manager to transfer assets on behalf of another user. It will only work when issued by the asset's revocation manager.

**Task 7-1** The clawback address (Account 2) revokes 10 latinum from Account 3 and places it back with Account 1.

```javascript tab="JavaScript"
//JavaScript
//insert this code affter Step 6 code
    // Revoke an Asset:
    // The asset was also created with the ability for it to be revoked by 
    // the clawbackaddress. If the asset was created or configured by the manager
    // to not allow this by setting the clawbackaddress to "" then this would 
    // not be possible.
    // We will now clawback the 10 assets in account3. account2
    // is the clawbackaccount and must sign the transaction
    // The sender will be be the clawback adress.
    // the recipient will also be be the creator in this case
    // that is account3
    sender = recoveredAccount2.addr;
    recipient = recoveredAccount1.addr;
    revocationTarget = recoveredAccount3.addr;
    closeRemainderTo = undefined;
    amount = 10;

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);

    // signing and sending "txn" will send "amount" assets from "revocationTarget" to "recipient",
    // if and only if sender == clawback manager for this asset
    let rtxn = algosdk.makeAssetTransferTxn(sender, recipient, closeRemainderTo, revocationTarget,
        cp.fee, amount, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
    // Must be signed by the account that is the clawback address    
    rawSignedTxn = rtxn.signTxn(recoveredAccount2.sk)
    let rtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + rtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, rtx.txId);

    // You should now see 0 assets listed in the account information
    // for the third account
    console.log("Asset ID: " + assetID);
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets[assetID]));

    //you should see console/terminal output similar to below for account 3
    //Transaction MW2ZKQ2GXMVVJSF23AXFUSKQTF43EANJHXJJAXRQ7BSBKDUWZMTA confirmed in round 4274200
    //Asset ID: 149774
    //Account Information for: { "creator": "THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM", "amount": 0, "frozen": true }

```


Step 8 Destroy an Asset
-----------------------

Asset destruction allows the creator to remove the asset from the ledger, if all outstanding assets are held by the creator. 

**Task 8-1** With all assets back in the creator's account, the manager (Account 1) destroys the asset.

```javascript tab="JavaScript"
    // JavaScript
    // insert this code after Step 7's code

    // Destroy and Asset:
    // All of the created assets should now be back in the creators
    // Account so we can delete the asset.
    // If this is not the case the asset deletion will fail

    // First update changing transaction parameters
    // We will account for changing transaction parameters
    // before every transaction in this example
    await getChangingParms(algodclient);


    // The address for the from field must be the manager account
    // Which is currently the creator addr1
    addr = recoveredAccount1.addr;

    // if all assets are held by the asset creator,
    // the asset creator can sign and issue "txn" to remove the asset from the ledger. 
    let dtxn = algosdk.makeAssetDestroyTxn(addr, cp.fee, cp.firstRound, cp.lastRound, note, cp.genHash, cp.genID, assetID);
    // The transaction must be signed by the manager which 
    // is currently set to account1
    rawSignedTxn = dtxn.signTxn(recoveredAccount1.sk)
    let dtx = (await algodclient.sendRawTransaction(rawSignedTxn));
    console.log("Transaction : " + dtx.txId);
    // wait for transaction to be confirmed
    await waitForConfirmation(algodclient, dtx.txId);

    // The account3 and account1 should no longer contain the asset as it has been destroyed
    console.log("Asset ID: " + assetID);
    act = await algodclient.accountInformation(recoveredAccount3.addr);
    console.log("Account Information for: " + JSON.stringify(act.assets[assetID]));

    // your console/terminal output should look similar to this:   
    // Transaction: FGMK2Q5Y2AXSOVSFO3XB6GBWQBUODF7A6QOU5BPHSIJIRA3UJHHQ
    // Transaction FGMK2Q5Y2AXSOVSFO3XB6GBWQBUODF7A6QOU5BPHSIJIRA3UJHHQ confirmed in round 4274286
    // Asset ID: 149783
    // Account Information for: { "creator": "", "amount": 0, "frozen": true }
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
