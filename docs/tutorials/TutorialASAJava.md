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

```java tab="Java"
//Java
package com.algorand.Tutorials;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.Map.Entry;
import com.algorand.algosdk.account.Account;
import com.algorand.algosdk.algod.client.AlgodClient;
import com.algorand.algosdk.algod.client.ApiException;
import com.algorand.algosdk.algod.client.api.AlgodApi;
import com.algorand.algosdk.algod.client.auth.ApiKeyAuth;
import com.algorand.algosdk.algod.client.model.AssetParams;
import com.algorand.algosdk.algod.client.model.TransactionID;
import com.algorand.algosdk.algod.client.model.TransactionParams;
import com.algorand.algosdk.crypto.Address;
import com.algorand.algosdk.crypto.Digest;
import com.algorand.algosdk.transaction.SignedTransaction;
import com.algorand.algosdk.transaction.Transaction;
import com.algorand.algosdk.util.Encoder;

// To be completed... Creating, modifying, sending and listing assets

public class TutorialASAJava {
// Inline class to handle changing block parameters
// Throughout the example
static class ChangingBlockParms {
public BigInteger fee;
public BigInteger firstRound;
public BigInteger lastRound;
public String genID;
public Digest genHash;
public ChangingBlockParms() {
this.fee = BigInteger.valueOf(0);
this.firstRound = BigInteger.valueOf(0);
this.lastRound = BigInteger.valueOf(0);
this.genID = "";
this.genHash = null;
}
};

// Utility function to wait on a transaction to be confirmed

private static void waitForTransactionToComplete(AlgodApi algodApiInstance,
String txID) throws Exception {
while (true) {
try {
// Check the pending tranactions
com.algorand.algosdk.algod.client.model.Transaction b3 = algodApiInstance.pendingTransactionInformation(txID);
if (b3.getRound() != null && b3.getRound().longValue() \> 0) {
// Got the completed Transaction
System.out.println("Transaction " + b3.getTx() + " confirmed in round " +b3.getRound().longValue());
break;
}
} catch (Exception e) {
throw (e);
}
}
}

// Utility function to update changing block parameters

public static ChangingBlockParms getChangingParms(AlgodApi algodApiInstance)
throws Exception {
ChangingBlockParms cp = new TutorialASAJava.ChangingBlockParms();
try {
TransactionParams params = algodApiInstance.transactionParams();
cp.fee = params.getFee();
cp.firstRound = params.getLastRound();
cp.lastRound = cp.firstRound.add(BigInteger.valueOf(1000))
cp.genID = params.getGenesisID();
cp.genHash = new Digest(params.getGenesishashb64());
} catch (ApiException e) {
throw (e);
}
return (cp);
}

// Utility function for sending a raw signed transaction to the network
public static TransactionID submitTransaction(AlgodApi algodApiInstance,
SignedTransaction signedTx)
throws Exception {
try {
// Msgpack encode the signed transaction
byte[] encodedTxBytes = Encoder.encodeToMsgPack(signedTx)
TransactionID id = algodApiInstance.rawTransaction(encodedTxBytes);
return (id);
} catch (ApiException e) {
throw (e);
}
}

public static void main(String args[]) throws Exception {
final String ALGOD_API_ADDR = "NODEADDRESS";
final String ALGOD_API_TOKEN = "NODETOKEN";
AlgodClient client = (AlgodClient) new
AlgodClient().setBasePath(ALGOD_API_ADDR);
ApiKeyAuth api_key = (ApiKeyAuth) client.getAuthentication("api_key");
api_key.setApiKey(ALGOD_API_TOKEN);
AlgodApi algodApiInstance = new AlgodApi(client);
// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
// These three accounts are for testing purposes

// Shown for demonstration purposes. NEVER reveal secret mnemonics in practice.
// These three accounts are for testing purposes
// Paste in phrase from Step 1A
final String account1_mnemonic = "PASTE your phrase for account 1 from Step 1A";
final String account2_mnemonic = "PASTE your phrase for account 2 from Step 1A";
final String account3_mnemonic = "PASTE your phrase for account 3 from Step 1A";

Account acct1 = new Account(account1_mnemonic);
Account acct2 = new Account(account2_mnemonic);
Account acct3 = new Account(account3_mnemonic);

// get last round and suggested tx fee
// We use these to get the latest round and tx fees
// These parameters will be required before every
// Transactio
// We will account for changing transaction parameters
// before every transaction in this example

ChangingBlockParms cp = null;
try {
cp = getChangingParms(algodApiInstance);
} catch (ApiException e) {
e.printStackTrace();
return;
}
System.out.println(acct1.getAddress().toString())
System.out.println(acct2.getAddress().toString());
System.out.println(acct3.getAddress().toString());
}
}

// Terminal should look similar to this:
// THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
// AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
// 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU

```


Step 2 Create a New Asset
---------------------------

The ability to create asserts directly on the blockcahin is an exciting capability of the Algorand Blockchain. Possible uses include currency, game leader boards, points in a loyalty system, shares of an asset, and securities such as stocks, bonds, and derivaties.  


**Info:**
    The decimals value determines the placement of the decimal. For example, when `decimals = 2`, and the `amount = 1000`, the acatual amount is 10.00. So, when a  transfer of amount of 10 is made, the actual transfer is .10


**Note:**

    With the **Java** code solution, paste this snippet at the end of the `main` function, before the final two curly braces `}}`

**Task 2-1** Account 1 creates an asset called latinum and sets Account 2 as the manager, reserve, freeze, and clawback address.


```java tab="Java"
        // Java
        // Insert after Step 1 code
        // Create a new asset

        // The following parameters are asset specific
        // and will be re-used throughout the example.

        // Total number of this asset available for circulation
        BigInteger assetTotal = BigInteger.valueOf(10000);
        // Whether user accounts will need to be unfrozen before transacting
        boolean defaultFrozen = false;
        // Decimals specifies the number of digits to display after the decimal
        // place when displaying this asset. A value of 0 represents an asset
        // that is not divisible, a value of 1 represents an asset divisible
        // into tenths, and so on. This value must be between 0 and 19
        Integer assetDecimals = 0;
        // Used to display asset units to user
        String unitName = "LATINUM";
        // Friendly name of the asset
        String assetName = "latinum";
        // Optional string pointing to a URL relating to the asset
        String url = "http://this.test.com";
        // Optional hash commitment of some sort relating to the asset. 32 character
        // length.
        String assetMetadataHash = "16efaa3924a6fd9d3a4824799a4ac65d";
        // The following parameters are the only ones
        // that can be changed, and they have to be changed
        // by the current manager
        // Specified address can change reserve, freeze, clawback, and manager
        Address manager = acct2.getAddress();
        // Specified address is considered the asset reserve
        // (it has no special privileges, this is only informational)
        Address reserve = acct2.getAddress();
        // Specified address can freeze or unfreeze user asset holdings
        Address freeze = acct2.getAddress();
        // Specified address can revoke user asset holdings and send
        // them to other addresses
        Address clawback = acct2.getAddress();

        Transaction tx = Transaction.createAssetCreateTransaction(acct1.getAddress(), BigInteger.valueOf(1000),
                cp.firstRound, cp.lastRound, null, cp.genID, cp.genHash, assetTotal, assetDecimals, defaultFrozen,
                unitName, assetName, url, assetMetadataHash.getBytes(), manager, reserve, freeze, clawback);
        // Update the fee as per what the BlockChain is suggesting
        Account.setFeeByFeePerByte(tx, cp.fee);

        // Sign the Transaction
        SignedTransaction signedTx = acct1.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        BigInteger assetID = null;
        try {
            TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
            // Now that the transaction is confirmed we can get the assetID
            com.algorand.algosdk.algod.client.model.Transaction ptx = algodApiInstance
                    .pendingTransactionInformation(id.getTxId());
            assetID = ptx.getTxresults().getCreatedasset();
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        System.out.println("AssetID = " + assetID);
        AssetParams assetInfo = algodApiInstance.assetInformation(assetID);
        System.out.println(assetInfo);
        // terminal out put should look similar to this:
        // Transaction ID: class TransactionID {
        // txId: HZYZW2FRF6YXJZOJTI2YTDOKNHMTPQR5DBGZMIMGCWXWDC36N3RA
        // }
        // Transaction HZYZW2FRF6YXJZOJTI2YTDOKNHMTPQR5DBGZMIMGCWXWDC36N3RA confirmed in
        // round 4288214
        // AssetID = 150575
        // class AssetParams {
        // assetname: latinum
        // clawbackaddr: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // decimals: 0
        // defaultfrozen: false
        // freezeaddr: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // managerkey: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // metadatahash: [B@41ab013
        // reserveaddr: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // total: 10000
        // unitname: LATINUM
        // url: http://this.test.com
        // }
```


Step 3 Configure Asset Manager
------------------------------
Assets can be managed as to which accounts have roles for overall manager, reserve, freeze, and clawback functions. By default all of these roles are set to the creator account. 

Asset reconfiguration allows the address specified as manager to change any of the special addresses for the asset, such as the reserve address. To keep an address the same, it must be re-specified in each new configuration transaction. Supplying an empty address is the same as turning the associated feature off for this asset. Once a special address is set to the empty address, it can never change again. For example, if an asset configuration transaction specifying clawback="" were issued, the associated asset could never be revoked from asset holders, and clawback="" would be true for all time. The  strictEmptyAddressChecking argument can help with this behavior: when set to its default true, makeAssetConfigTxn will throw an error if any undefined management addresses are passed.

**Task 3-1** Here, the current manager (Account 2) issues an asset configuration transaction that assigns Account 1 as the new manager.


```java tab="Java"
         // Java
         // Insert after Step 2's code

        // Change Asset Configuration:
        // Next we will change the asset configuration
        // Note that configuration changes must be done by
        // The manager account, which is currently acct2
        // Note in this transaction we are re-using the asset
        // creation parameters and only changing the manager
        // and transaction parameters like first and last round

        // First we update standard Transaction parameters
        // To account for changes in the state of the blockchain
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (final ApiException e) {
            e.printStackTrace();
            return;
        }

        tx = Transaction.createAssetConfigureTransaction(acct2.getAddress(), BigInteger.valueOf(1000), cp.firstRound,
                cp.lastRound, null, cp.genID, cp.genHash, assetID, acct1.getAddress(), reserve, freeze, clawback,
                false);
        // Update the fee as per what the BlockChain is suggesting
        Account.setFeeByFeePerByte(tx, cp.fee);
        // The transaction must be signed by the current manager account
        // We are reusing the signedTx variable from the first transaction in the
        // example
        signedTx = acct2.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        try {
            final TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
        } catch (final Exception e) {
            e.printStackTrace();
            return;
        }

        // Next we will list the newly created asset
        // Get the asset information for the newly changed asset

        assetInfo = algodApiInstance.assetInformation(assetID);
        // The manager should now be the same as the creator
        System.out.println(assetInfo);
        System.out.println("Creator = " + assetInfo.getCreator());

        // terminal outout should look similar to this...
        // Transaction A4NJTLFO5UHCNAAIKYASCJSO2O6V3TNKEEJFQO3S3RY6IKD2F3NQ confirmed in
        // round 4290988
        // class AssetParams {
        // assetname: latinum
        // clawbackaddr: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // decimals: 0
        // defaultfrozen: false
        // freezeaddr: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // managerkey: THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
        // metadatahash: [B@41ab013
        // reserveaddr: AJNNFQN7DSR7QEY766V7JDG35OPM53ZSNF7CU264AWOOUGSZBMLMSKCRIU
        // total: 10000
        // unitname: LATINUM
        // url: http://this.test.com
        // }
        // Creator = THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
 
```

Step 4 Opt-in to Receive Asset
------------------------------

Once the asset has been created, the next thing to do is send assets to other accounts. 

Before a user can begin transacting with an asset, the user must first issue an asset acceptance transaction. This is a special case of the asset transfer transaction, where the user sends 0 assets to themself. After issuing this transaction, the user can begin transacting with the asset. Each new accepted asset increases the user's minimum balance.

**Task 4-1** Account 3 opts-in to receive the new asset by sending a 0 amount transfer of the asset to itself.

```java tab="Java"
        // Java
        // insert after code from Step 3's code

        // Opt-in to Receive Asset
        // All accounts that want recieve the new asset have to opt in.
        // To do this they send an asset transfer of the new asset to themseleves with
        // an ammount of 0
        //
        // In this example we are setting up the 3rd recovered account to
        // receive the new asset
        // First we update standard Transaction parameters
        // To account for changes in the state of the blockchain
        // com.algorand.algosdk.algod.client.model.Account act = new
        // com.algorand.algosdk.algod.client.model.Account();
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (final ApiException e) {
            e.printStackTrace();
            return;
        }
        tx = Transaction.createAssetAcceptTransaction(acct3.getAddress(), BigInteger.valueOf(1000), cp.firstRound,
                cp.lastRound, null, cp.genID, cp.genHash, assetID);
        // Update the fee based on the network suggested fee
        Account.setFeeByFeePerByte(tx, cp.fee);
        // The transaction must be signed by the current manager account
        // We are reusing the signedTx variable from the first transaction in the
        // example
        signedTx = acct3.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        com.algorand.algosdk.algod.client.model.Account act = algodApiInstance
                .accountInformation(acct3.getAddress().toString());
        try {
            final TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
            // We can now list the account information for acct3
            // and see that it can accept the new asseet
            assetInfo = algodApiInstance.assetInformation(assetID);
            System.out.println("AssetID: " + assetID);
            System.out.println("Creator = " + assetInfo.getCreator());

            act = algodApiInstance.accountInformation(acct3.getAddress().toString());
            System.out.println("Account opt in for Asset transfer " + acct3.getAddress().toString());
            System.out.println("Account Asset Holding Amount = " + act.getHolding(assetID).getAmount());

        } catch (final Exception e) {
            e.printStackTrace();
            return;
        }
        // terminal output should be similar to this...

        // Transaction BQYA3R7L3M63VGLQBTWZURLWSBISM62KFRUGCK2RGBTQTBNNULUA confirmed in round 4296935
        // AssetID: 151126
        // Creator = THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
        // Account opt in for Asset transfer 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
        // Account Asset Holding Amount = 0
```

Step 5 Transfer an Asset
------------------------

Transfer an asset allows users to transact with assets, after they have issued asset acceptance transactions. The optional closeRemainderTo argument can be used to stop transacting with a particular asset. Now that the opt-in has been done on a potential receiving accout, assets can be transferred.

**Note:**
    A frozen account can always close out to the asset creator.

**Task 5-1** This code has Account 1 sending 10 latinum to Account 3.


```java tab="Java"
        // Java
        // Insert code after Step 4's code

        // Transfer the Asset:
        // Now that account3 can recieve the new asset we can tranfer assets
        // from the creator (account 1) to account3

        // First we update standard Transaction parameters
        // To account for changes in the state of the blockchain
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (ApiException e) {
            e.printStackTrace();
            return;
        }
        // Next we set asset xfer specific parameters
        // We set the assetCloseTo to null so we do not close the asset out
        Address assetCloseTo = new Address();
        BigInteger assetAmount = BigInteger.valueOf(10);
        tx = Transaction.createAssetTransferTransaction(acct1.getAddress(), acct3.getAddress(), assetCloseTo,
                assetAmount, BigInteger.valueOf(1000), cp.firstRound, cp.lastRound, null, cp.genID, cp.genHash,
                assetID);
        // Update the fee based on the network suggested fee
        Account.setFeeByFeePerByte(tx, cp.fee);
        // The transaction must be signed by the sender account
        // We are reusing the signedTx variable from the first transaction in the
        // example
        signedTx = acct1.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        try {
            TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
            // We can now list the account information for acct3
            // and see that it now has 10 of the new asset
            assetInfo = algodApiInstance.assetInformation(assetID);
            System.out.println("AssetID: " + assetID);
            System.out.println("Creator = " + assetInfo.getCreator());

            act = algodApiInstance.accountInformation(acct3.getAddress().toString());
            System.out.println("Account w Asset transfered to " + acct3.getAddress().toString());
            System.out.println("Account Asset Holding Amount = " + act.getHolding(assetID).getAmount());

        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        // termianl output should look similar to this
        // Transaction HOYHCDQD3ZMHW5TOA2NBSDGVB4GPRQ26R6AQLNZPC64QJKL5RUHA confirmed in round 4296943
        // AssetID: 151126
        // Creator = THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
        // Account w Asset transfered to 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
        // Account Asset Holding Amount = 10

```

Step 6 Freeze an Asset
-----------------------

To freeze or unfreeze an asset, this transaction must be sent from the account specified as the freeze manager for the asset.

**Task 6-1** The freeze address (Account 2) freezes Account 3's latinum holdings.



```java tab="Java"
         // Java
        // Freeze the Asset:
        // The asset was created and configured to allow freezing an account
        // If the freeze address is blank, it will no longer be possible to do this.
        // In this example we will now freeze account3 from transacting with the
        // The newly created asset.
        // The freeze transaction is sent from the freeze acount
        // Which in this example is account2
        // First we update standard Transaction parameters
        // To account for changes in the state of the blockchain
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (ApiException e) {
            e.printStackTrace();
            return;
        }
        // Next we set asset xfer specific parameters
        boolean freezeState = true;
        // The sender should be freeze account acct2
        // Theaccount to freeze should be set to acct3
        tx = Transaction.createAssetFreezeTransaction(acct2.getAddress(), acct3.getAddress(), freezeState,
                BigInteger.valueOf(1000), cp.firstRound, cp.lastRound, null, cp.genHash, assetID);
        // Update the fee based on the network suggested fee
        Account.setFeeByFeePerByte(tx, cp.fee);
        // The transaction must be signed by the freeze account acct2
        // We are reusing the signedTx variable from the first transaction in the
        // example
        signedTx = acct2.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        try {
            TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
            // We can now list the account information for acct3
            // and see that it now frozen
            // Note--currently no getter method for frozen state
            act = algodApiInstance.accountInformation(acct3.getAddress().toString());
            System.out.println(act.getHolding(assetID).toString());
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        assetInfo = algodApiInstance.assetInformation(assetID);
        System.out.println("AssetID: " + assetID);
        System.out.println("Creator = " + assetInfo.getCreator());

        act = algodApiInstance.accountInformation(acct3.getAddress().toString());
        System.out.println("Account w Asset frozen " + acct3.getAddress().toString());
        System.out.println("Account Asset Holding Info = " + act);

        // termianl out put should show something similar to this
        // AssetID: 151234
        // Creator = THQHGD4HEESOPSJJYYF34MWKOI57HXBX4XR63EPBKCWPOJG5KUPDJ7QJCM
        // Account w Asset frozen
        // 3ZQ3SHCYIKSGK7MTZ7PE7S6EDOFWLKDQ6RYYVMT7OHNQ4UJ774LE52AQCU
```


Step 7 Revoke an Asset
----------------------

Revoking an asset allows an asset's revocation manager to transfer assets on behalf of another user. It will only work when issued by the asset's revocation manager.

**Task 7-1** The clawback address (Account 2) revokes 10 latinum from Account 3 and places it back with Account 1.

```java tab="Java"
        // Java
        // Insert after Step 6's code

        // Revoke the asset:
        // The asset was also created with the ability for it to be revoked by
        // clawbackaddress. If the asset was created or configured by the manager
        // not allow this by setting the clawbackaddress to a blank address
        // then this would not be possible.
        // We will now clawback the 10 assets in account3. Account2
        // is the clawbackaccount and must sign the transaction
        // The sender will be be the clawback adress.
        // the recipient will also be be the creator acct1 in this case
        // First we update standard Transaction parameters
        // To account for changes in the state of the blockchain
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (ApiException e) {
            e.printStackTrace();
            return;
        }
        // Next we set asset xfer specific parameters
        assetAmount = BigInteger.valueOf(10);
        tx = Transaction.createAssetRevokeTransaction(acct2.getAddress(), acct3.getAddress(), acct1.getAddress(),
                assetAmount, BigInteger.valueOf(1000), cp.firstRound, cp.lastRound, null, cp.genID, cp.genHash,
                assetID);
        // Update the fee based on the network suggested fee
        Account.setFeeByFeePerByte(tx, cp.fee);
        // The transaction must be signed by the clawback account
        // We are reusing the signedTx variable from the first transaction in the
        // example
        signedTx = acct2.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        try {
            TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
            // We can now list the account information for acct3
            // and see that it now has 0 of the new asset
            act = algodApiInstance.accountInformation(acct3.getAddress().toString());
            System.out.println("Account 3 Asset amount = " + act.getHolding(assetID).getAmount());
            act = algodApiInstance.accountInformation(acct1.getAddress().toString());
            System.out.println("Account 1 Asset amount = " + act.getHolding(assetID).getAmount());
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        // your terminal output should look something like this
        // Transaction 6ZC6DBYAY4DI7HFZPDHRINXZQNDZX3WHHTM6OOYLPJ6REHUSRYOQ confirmed in round 4298849
        // Account 3 Asset amount = 0
        // Account 1 Asset amount = 10000
```


Step 8 Destroy an Asset
-----------------------

Asset destruction allows the creator to remove the asset from the ledger, if all outstanding assets are held by the creator. 

**Task 8-1** With all assets back in the creator's account, the manager (Account 1) destroys the asset.


```java tab="Java"
        // Java
        // Insert after Step 7's code

        // Destroy the Asset:
        // All of the created assets should now be back in the creators
        // Account so we can delete the asset.
        // If this is not the case the asset deletion will fail
        // The address for the from field must be the creator
        // First we update standard Transaction parameters
        // To account for changes in the state of the blockchain
        try {
            cp = getChangingParms(algodApiInstance);
        } catch (ApiException e) {
            e.printStackTrace();
            return;
        }
        // Next we set asset xfer specific parameters
        // The manager must sign and submit the transaction
        // This is currently set to acct1
        tx = Transaction.createAssetDestroyTransaction(acct1.getAddress(), BigInteger.valueOf(1000), cp.firstRound,
                cp.lastRound, null, cp.genHash, assetID);
        // Update the fee based on the network suggested fee
        Account.setFeeByFeePerByte(tx, cp.fee);
        // The transaction must be signed by the manager account
        // We are reusing the signedTx variable from the first transaction in the
        // example
        signedTx = acct1.signTransaction(tx);
        // send the transaction to the network and
        // wait for the transaction to be confirmed
        try {
            TransactionID id = submitTransaction(algodApiInstance, signedTx);
            System.out.println("Transaction ID: " + id);
            waitForTransactionToComplete(algodApiInstance, signedTx.transactionID);
            // We can now list the account information for acct1
            // and see that the asset is no longer there
            act = algodApiInstance.accountInformation(acct1.getAddress().toString());
            System.out.println("Does AssetID: " + assetID + " exist? " + act.getThisassettotal().containsKey(assetID));
        } catch (Exception e) {
            e.printStackTrace();
            return;
        }
        // your terminal output should look similar to this
        // Transaction SV5UPKAG7CI3BUZ74SL5VDFHRS2SGUVKAJVKNHSUZWIPM6HSRFEA confirmed in
        // round 4299035
        // Does AssetID: 151265 exist? false
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
