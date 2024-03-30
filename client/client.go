package client

// CS 161 Project 2

// Only the following imports are allowed! ANY additional imports
// may break the autograder!
// - bytes
// - encoding/hex
// - encoding/json
// - errors
// - fmt
// - github.com/cs161-staff/project2-userlib
// - github.com/google/uuid
// - strconv
// - strings

import (
	"encoding/json"

	userlib "github.com/cs161-staff/project2-userlib"
	"github.com/google/uuid"

	// hex.EncodeToString(...) is useful for converting []byte to string

	// Useful for string manipulation

	// Useful for formatting strings (e.g. `fmt.Sprintf`).
	"fmt"

	// Useful for creating new error messages to return using errors.New("...")
	"errors"

	// Optional.
	_ "strconv"
)

// // This serves two purposes: it shows you a few useful primitives,
// // and suppresses warnings for imports not being used. It can be
// // safely deleted!
// func someUsefulThings() {

// 	// Creates a random UUID.
// 	randomUUID := uuid.New()

// 	// Prints the UUID as a string. %v prints the value in a default format.
// 	// See https://pkg.go.dev/fmt#hdr-Printing for all Golang format string flags.
// 	userlib.DebugMsg("Random UUID: %v", randomUUID.String())

// 	// Creates a UUID deterministically, from a sequence of bytes.
// 	hash := userlib.Hash([]byte("user-structs/alice"))
// 	deterministicUUID, err := uuid.FromBytes(hash[:16])
// 	if err != nil {
// 		// Normally, we would `return err` here. But, since this function doesn't return anything,
// 		// we can just panic to terminate execution. ALWAYS, ALWAYS, ALWAYS check for errors! Your
// 		// code should have hundreds of "if err != nil { return err }" statements by the end of this
// 		// project. You probably want to avoid using panic statements in your own code.
// 		panic(errors.New("An error occurred while generating a UUID: " + err.Error()))
// 	}
// 	userlib.DebugMsg("Deterministic UUID: %v", deterministicUUID.String())

// 	// Declares a Course struct type, creates an instance of it, and marshals it into JSON.
// 	type Course struct {
// 		name      string
// 		professor []byte
// 	}

// 	course := Course{"CS 161", []byte("Nicholas Weaver")}
// 	courseBytes, err := json.Marshal(course)
// 	if err != nil {
// 		panic(err)
// 	}

// 	userlib.DebugMsg("Struct: %v", course)
// 	userlib.DebugMsg("JSON Data: %v", courseBytes)

// 	// Generate a random private/public keypair.
// 	// The "_" indicates that we don't check for the error case here.
// 	var pk userlib.PKEEncKey
// 	var sk userlib.PKEDecKey
// 	pk, sk, _ = userlib.PKEKeyGen()
// 	userlib.DebugMsg("PKE Key Pair: (%v, %v)", pk, sk)

// 	// Here's an example of how to use HBKDF to generate a new key from an input key.
// 	// Tip: generate a new key everywhere you possibly can! It's easier to generate new keys on the fly
// 	// instead of trying to think about all of the ways a key reuse attack could be performed. It's also easier to
// 	// store one key and derive multiple keys from that one key, rather than
// 	originalKey := userlib.RandomBytes(16)
// 	derivedKey, err := userlib.HashKDF(originalKey, []byte("mac-key"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	userlib.DebugMsg("Original Key: %v", originalKey)
// 	userlib.DebugMsg("Derived Key: %v", derivedKey)

// 	// A couple of tips on converting between string and []byte:
// 	// To convert from string to []byte, use []byte("some-string-here")
// 	// To convert from []byte to string for debugging, use fmt.Sprintf("hello world: %s", some_byte_arr).
// 	// To convert from []byte to string for use in a hashmap, use hex.EncodeToString(some_byte_arr).
// 	// When frequently converting between []byte and string, just marshal and unmarshal the data.
// 	//
// 	// Read more: https://go.dev/blog/strings

// 	// Here's an example of string interpolation!
// 	_ = fmt.Sprintf("%s_%d", "file", 1)
// }

// // This is the type definition for the User struct.
// // A Go struct is like a Python or Java class - it can have attributes
// // (e.g. like the Username attribute) and methods (e.g. like the StoreFile method below).
// // 这里来定义一下结构体
// type TreeNode struct {
// 	Value    int
// 	Children []*TreeNode
// }
// type User struct {
// 	// Username string
// 	// Password string
// 	// // Private_key是用来对用户进行解密的
// 	// // Signature_key是用来对数据进行解密的
// 	// Private_key   userlib.PKEDecKey
// 	// Signature_key userlib.DSSignKey
// 	// // 这里对于Intermediate Id想以树的方式来定义
// 	// IntermediateUUIDmap TreeNode

// 	Username string
// 	Password string
// 	// ADDD LATER ---------------------------------
// 	PrivateKey       userlib.PKEDecKey
// 	SignatureKey     userlib.DSSignKey
// 	IntermediateUUID map[string]map[string]Invitation

// 	// You can add other attributes here if you want! But note that in order for attributes to
// 	// be included when this struct is serialized to/from JSON, they must be capitalized.
// 	// On the flipside, if you have an attribute that you want to be able to access from
// 	// this struct's methods, but you DON'T want that value to be included in the serialized value
// 	// of this struct that's stored in datastore, then you can use a "private" variable (e.g. one that
// 	// begins with a lowercase letter).
// }
// type FileNode struct {
// 	Contents []byte
// 	PrevUUID uuid.UUID
// 	NextUUID uuid.UUID
// }

// // 包含文件对应的 FileNode 地址
// type FileLocator struct {
// 	FirstFileNodeUUID uuid.UUID
// 	LastFileNodeUUID  uuid.UUID
// 	SymKeyFN          []byte
// 	MacKeyFN          []byte
// }

// // 文件分享接收者通过 Intermediate 获取 fileLocator 的解密密钥
// type Intermediate struct {
// 	FileLocatorUUID   uuid.UUID
// 	SymKeyFileLocator []byte
// 	MacKeyFileLocator []byte
// }

// // 每个用户通过 keyFile 来打开 file
// type KeyFile struct {
// 	isFileOwner bool
// 	FileUUID    uuid.UUID
// 	SymKeyFile  []byte
// 	MacKeyFile  []byte
// }

// // 文件分享邀请
// type Invitation struct {
// 	IntermediateUUID uuid.UUID
// 	SymKeyInter      []byte
// 	MacKeyInter      []byte
// }

// This serves two purposes: it shows you a few useful primitives,
// and suppresses warnings for imports not being used. It can be
// safely deleted!
func someUsefulThings() {

	// Creates a random UUID.
	randomUUID := uuid.New()

	// Prints the UUID as a string. %v prints the value in a default format.
	// See https://pkg.go.dev/fmt#hdr-Printing for all Golang format string flags.
	userlib.DebugMsg("Random UUID: %v", randomUUID.String())

	// Creates a UUID deterministically, from a sequence of bytes.
	hash := userlib.Hash([]byte("user-structs/alice"))
	deterministicUUID, err := uuid.FromBytes(hash[:16])
	if err != nil {
		// Normally, we would `return err` here. But, since this function doesn't return anything,
		// we can just panic to terminate execution. ALWAYS, ALWAYS, ALWAYS check for errors! Your
		// code should have hundreds of "if err != nil { return err }" statements by the end of this
		// project. You probably want to avoid using panic statements in your own code.
		panic(errors.New("An error occurred while generating a UUID: " + err.Error()))
	}
	userlib.DebugMsg("Deterministic UUID: %v", deterministicUUID.String())

	// Declares a Course struct type, creates an instance of it, and marshals it into JSON.
	type Course struct {
		name      string
		professor []byte
	}

	course := Course{"CS 161", []byte("Nicholas Weaver")}
	courseBytes, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}

	userlib.DebugMsg("Struct: %v", course)
	userlib.DebugMsg("JSON Data: %v", courseBytes)

	// Generate a random private/public keypair.
	// The "_" indicates that we don't check for the error case here.
	var pk userlib.PKEEncKey
	var sk userlib.PKEDecKey
	pk, sk, _ = userlib.PKEKeyGen()
	userlib.DebugMsg("PKE Key Pair: (%v, %v)", pk, sk)

	// Here's an example of how to use HBKDF to generate a new key from an input key.
	// Tip: generate a new key everywhere you possibly can! It's easier to generate new keys on the fly
	// instead of trying to think about all of the ways a key reuse attack could be performed. It's also easier to
	// store one key and derive multiple keys from that one key, rather than
	originalKey := userlib.RandomBytes(16)
	derivedKey, err := userlib.HashKDF(originalKey, []byte("mac-key"))
	if err != nil {
		panic(err)
	}
	userlib.DebugMsg("Original Key: %v", originalKey)
	userlib.DebugMsg("Derived Key: %v", derivedKey)

	// A couple of tips on converting between string and []byte:
	// To convert from string to []byte, use []byte("some-string-here")
	// To convert from []byte to string for debugging, use fmt.Sprintf("hello world: %s", some_byte_arr).
	// To convert from []byte to string for use in a hashmap, use hex.EncodeToString(some_byte_arr).
	// When frequently converting between []byte and string, just marshal and unmarshal the data.
	//
	// Read more: https://go.dev/blog/strings

	// Here's an example of string interpolation!
	_ = fmt.Sprintf("%s_%d", "file", 1)
}

// This is the type definition for the User struct.
// A Go struct is like a Python or Java class - it can have attributes
// (e.g. like the Username attribute) and methods (e.g. like the StoreFile method below).
type User struct {
	Username string
	Password string
	// ADDD LATER ---------------------------------
	PrivateKey       userlib.PKEDecKey
	SignatureKey     userlib.DSSignKey
	IntermediateUUID map[string]map[string]Invitation

	// You can add other attributes here if you want! But note that in order for attributes to
	// be included when this struct is serialized to/from JSON, they must be capitalized.
	// On the flipside, if you have an attribute that you want to be able to access from
	// this struct's methods, but you DON'T want that value to be included in the serialized value
	// of this struct that's stored in datastore, then you can use a "private" variable (e.g. one that
	// begins with a lowercase letter).
}

// Where contain the address fileNode
type FileLocator struct {
	FirstFileNodeUUID uuid.UUID
	LastFileNodeUUID  uuid.UUID
	SymKeyFN          []byte
	MacKeyFN          []byte
}

// Where to read content of file
type FileNode struct {
	Contents []byte
	PrevUUID uuid.UUID
	NextUUID uuid.UUID
}

// Where for recipient get the key for fileLocator
type Intermediate struct {
	FileLocatorUUID   uuid.UUID
	SymKeyFileLocator []byte
	MacKeyFileLocator []byte
}

// Every user has keyFile to open the file
type KeyFile struct {
	IsFileOwner bool
	FileUUID    uuid.UUID
	SymKeyFile  []byte
	MacKeyFile  []byte
}

type Invitation struct {
	IntermediateUUID uuid.UUID
	SymKeyInter      []byte
	MacKeyInter      []byte
}

// NOTE: The following methods have toy (insecure!) implementations.

// #########################################################################            INITUSER         ############################################################################################
func InitUser(username string, password string) (userdataptr *User, err error) {
	//----HANDLE ERRORS-----------------------
	//Error Case1: empty username
	if len(username) == 0 {
		return nil, errors.New("Username must be greater than 0 characters")
	}
	hashUsername := userlib.Hash([]byte(username))
	userUUID, err := uuid.FromBytes(hashUsername[:16])
	//Error Case2: Short short
	if err != nil {
		return nil, errors.New("Hash(username) must be at least 16 bytes")
	}
	//Error Duplicate username, use other name
	if _, ok := userlib.DatastoreGet(userUUID); ok == true {
		return nil, errors.New("UUID already exists in DataStore")
	}
	//----------------------------------------

	// Generating Public-Key Encryption and Digital Signatures
	publicKey, privateKey, _ := userlib.PKEKeyGen()
	signatureKey, verifyKey, _ := userlib.DSKeyGen()

	// Store public-key and verify-key on KeyStore
	err = userlib.KeystoreSet(username+"publicKey", publicKey)
	if err != nil {
		return nil, errors.New("CANNOT set user's publicKey on keystore")
	}
	err = userlib.KeystoreSet(username+"verifyKey", verifyKey)
	if err != nil {
		return nil, errors.New("CANNOT set user's verifyKey on keystore")
	}

	// Initilize new user
	newUser := User{username, password, privateKey, signatureKey, map[string]map[string]Invitation{}}
	newUserBytes, errMarshal := json.Marshal(newUser)
	if errMarshal != nil {
		return nil, errors.New("Marshal error, cannot convert User Struct to bytes")
	}

	// Generate symmetricKey and macKey from username and password to Encrypt then Tag
	symEncKey, macKey := GenerateKeys(username, password)
	iv := userlib.RandomBytes(16)
	newUserEncrypted := userlib.SymEnc(symEncKey, iv, newUserBytes)
	hmacTag, hmacError := userlib.HMACEval(macKey, newUserEncrypted)
	if hmacError != nil {
		return nil, errors.New("input as key for hmac inshould be a 16-byte key")
	}

	// Store the new userUUID to datastore
	userlib.DatastoreSet(userUUID, append(newUserEncrypted, hmacTag...))

	return &newUser, nil
}

// #########################################################################            GETUSER         ############################################################################################
func GetUser(username string, password string) (userdataptr *User, err error) {
	var userdata User
	hashUsername := userlib.Hash([]byte(username))
	userUUID, err := uuid.FromBytes(hashUsername[:16])

	//------HANDLE ERRORS---------------------------------------------------
	//Error: Short short
	if err != nil {
		return nil, errors.New("Hash(username) must be at least 16 bytes")
	}
	dataValue, ok := userlib.DatastoreGet(userUUID)
	//Error: no username exist
	if ok == false {
		return nil, errors.New("UUID(the user) does not exists in DataStore")
	}
	//----------------------------------------------------------------------

	// Retrieve newUserEncrypted
	newUserEncrypted := dataValue[:len(dataValue)-64]
	// Retrieve hmacTag
	hmacTag := dataValue[len(dataValue)-64:]

	// Recreate the hmacTag from username and pw to verify authentic of tag
	symEncKey, macKey := GenerateKeys(username, password)
	hmacTagVerify, hmacError := userlib.HMACEval(macKey, newUserEncrypted)
	if hmacError != nil {
		return nil, errors.New("input as key for hmac in InitUser should be a 16-byte key")
	}

	// Confirm authenticity using HMACEqual()
	if !userlib.HMACEqual(hmacTagVerify, hmacTag) {
		return nil, errors.New("Data has been modified or wrong Password")
	}

	// Decrypt to get User
	newUserBytes := userlib.SymDec(symEncKey, newUserEncrypted)
	errMarshal := json.Unmarshal(newUserBytes, &userdata)
	if errMarshal != nil {
		return nil, errors.New("Unmarshal error, cannot convert bytes to User Struct")
	}
	return &userdata, nil
}

// #########################################################################            STOREFILE         ############################################################################################
func (userdata *User) StoreFile(filename string, content []byte) (err error) {

	keyFile, err := userdata.VerifyThenDecKeyFile(filename)
	if err != nil && err.Error() != "keyFileUUID does not exists in DataStore" {
		return err
	}

	//Filename doesnot existed in userdata => Create a new file
	if keyFile == nil {
		fileLocatorUUID := uuid.New() // Create a random uuid for the file
		// store file content in filenode
		FirstFileNodeUUID := uuid.New()
		LastFileNodeUUID := uuid.New()

		// Generate a new symKey and macKey
		symKeyFN, macKeyFN, ok := GenerateSymAndMacKey("enc-mac-filenode")
		if ok != nil {
			return ok
		}
		err = EncFileNode(symKeyFN, macKeyFN, content, FirstFileNodeUUID, uuid.Nil, LastFileNodeUUID)
		if err != nil {
			return err
		}
		err = EncFileNode(symKeyFN, macKeyFN, nil, LastFileNodeUUID, FirstFileNodeUUID, uuid.Nil)
		if err != nil {
			return err
		}

		// Generate symmetricKey and macKey from owner to Encrypt then Tag the FileLocator
		symKeyFL, macKeyFL, ok := GenerateSymAndMacKey("enc-mac-filelocator")
		if ok != nil {
			return ok
		}
		err = EncFileLocator(symKeyFL, macKeyFL, fileLocatorUUID, FirstFileNodeUUID, LastFileNodeUUID, symKeyFN, macKeyFN)
		if err != nil {
			return err
		}

		// Create and store the keyfile struct for file
		err = userdata.EncKeyFile(filename, true, fileLocatorUUID, symKeyFL, macKeyFL)
		if err != nil {
			return err
		}
		return nil

	} else {
		// Filename existed => Overwrite it
		// Generate new first and last node for fileLocator
		FirstFileNodeUUID := uuid.New()
		LastFileNodeUUID := uuid.New()

		// Generate a new symKey and macKey
		symKeyFN, macKeyFN, ok := GenerateSymAndMacKey("enc-mac-filenode")
		if ok != nil {
			return ok
		}
		err = EncFileNode(symKeyFN, macKeyFN, content, FirstFileNodeUUID, uuid.Nil, LastFileNodeUUID)
		if err != nil {
			return err
		}
		err = EncFileNode(symKeyFN, macKeyFN, nil, LastFileNodeUUID, FirstFileNodeUUID, uuid.Nil)
		if err != nil {
			return err
		}

		// Re-Encrypt fileLocator with new First and Last Node
		fileLocatorUUID, symFileLocKey, macFileLocKey, err := userdata.GetFileLocator(filename)
		if err != nil {
			return err
		}
		err = EncFileLocator(symFileLocKey, macFileLocKey, fileLocatorUUID, FirstFileNodeUUID, LastFileNodeUUID, symKeyFN, macKeyFN)
		if err != nil {
			return err
		}
	}

	return nil

}

// #########################################################################            APPENDTOFILE         ############################################################################################
func (userdata *User) AppendToFile(filename string, content []byte) error {
	// Get UUID of data located file
	fileLocatorUUID, symFileLocKey, macFileLocKey, err := userdata.GetFileLocator(filename)
	//----Handle error--------
	if err != nil {
		return err
	}

	fileLocator, err := VerifyThenDecFileLocator(symFileLocKey, macFileLocKey, fileLocatorUUID)
	if err != nil {
		return err
	}

	oldLastFileNode, err := VerifyThenDecFileNode(fileLocator.SymKeyFN, fileLocator.MacKeyFN, fileLocator.LastFileNodeUUID)
	if err != nil {
		return err
	}

	// Add new last node and Encrypt it
	newLastFileNodeUUID := uuid.New()
	err = EncFileNode(fileLocator.SymKeyFN, fileLocator.MacKeyFN, nil, newLastFileNodeUUID, fileLocator.LastFileNodeUUID, uuid.Nil)
	if err != nil {
		return err
	}

	// Update the oldLastFileNode.NextUUID
	err = EncFileNode(fileLocator.SymKeyFN, fileLocator.MacKeyFN, content, fileLocator.LastFileNodeUUID, oldLastFileNode.PrevUUID, newLastFileNodeUUID)
	if err != nil {
		return err
	}

	// Update last node UUID of fileLocator
	fileLocator.LastFileNodeUUID = newLastFileNodeUUID

	// Re-Encrypt fileLocator
	err = EncFileLocator(symFileLocKey, macFileLocKey, fileLocatorUUID, fileLocator.FirstFileNodeUUID, fileLocator.LastFileNodeUUID, fileLocator.SymKeyFN, fileLocator.MacKeyFN)
	if err != nil {
		return err
	}

	return nil
}

// #########################################################################            LOADFILE         ############################################################################################
func (userdata *User) LoadFile(filename string) (content []byte, err error) {
	var contents []byte
	// Get UUID of data located file
	fileLocatorUUID, symFileLocKey, macFileLocKey, err := userdata.GetFileLocator(filename)
	//----Handle error--------
	if err != nil {
		return nil, err
	}
	// Get data
	fmt.Println("LoadFile file Locator")
	fmt.Println(fileLocatorUUID)
	fileLocator, err := VerifyThenDecFileLocator(symFileLocKey, macFileLocKey, fileLocatorUUID)
	if err != nil {
		return nil, err
	}

	curNodeUUID := fileLocator.FirstFileNodeUUID
	for curNodeUUID != uuid.Nil {
		curNodeDec, err := VerifyThenDecFileNode(fileLocator.SymKeyFN, fileLocator.MacKeyFN, curNodeUUID)
		if err != nil {
			return nil, err
		}
		contents = append(contents, curNodeDec.Contents...)
		curNodeUUID = curNodeDec.NextUUID
	}
	return contents, nil
}

// #########################################################################            CREATEINVITATION         ############################################################################################
func (userdata *User) CreateInvitation(filename string, recipientUsername string) (invitationPtr uuid.UUID, err error) {
	//--------HANDLE ERROR--------------------------------------------------------------
	keyFile, err := userdata.VerifyThenDecKeyFile(filename)
	//Error: The given filename does not exist in the personal file namespace of the caller
	if err != nil {
		return uuid.Nil, err
	}
	//Error: recipientUsername does not exist
	recipientUUID, err := uuid.FromBytes(userlib.Hash([]byte(recipientUsername))[:16])
	_, ok := userlib.DatastoreGet(recipientUUID)
	if !ok {
		return uuid.Nil, errors.New("UUID(recipient) does not exists in DataStore")
	}
	// Case: user can't access the data
	_, err = userdata.LoadFile(filename)
	if err != nil {
		return uuid.Nil, errors.New("User cannot access the data")
	}
	//-----------------------------------------------------------------------------------

	var invitation Invitation
	var invitationUUID uuid.UUID

	if keyFile.IsFileOwner { // Invitation send by owner
		// Create a new intermediate then encrypt and mac it
		symKeyIntermediate, macKeyIntermediate, err := GenerateSymAndMacKey("enc-mac-intermediate-struct")
		if err != nil {
			return uuid.Nil, err
		}

		intermediateUUID := uuid.New()
		err = EncIntermediate(symKeyIntermediate, macKeyIntermediate, intermediateUUID, keyFile.FileUUID, keyFile.SymKeyFile, keyFile.MacKeyFile)
		if err != nil {
			return uuid.Nil, err
		}

		// Create an invitation hold the IntermediateUUID and 2 interKeys
		invitationUUID = uuid.New()
		invitation = Invitation{intermediateUUID, symKeyIntermediate, macKeyIntermediate}

		// Update File's owner IntermediateUUID
		if userdata.IntermediateUUID[filename] == nil {
			userdata.IntermediateUUID[filename] = map[string]Invitation{recipientUsername: invitation}
		} else {
			userdata.IntermediateUUID[filename][recipientUsername] = invitation
		}
		// //######## RE ENCRYPT USER ########################## ----  I COMMENT IT OUT BECAUSE I'M NOT SURE WE NEED THIS, SO FAR THERE IS NO BUGS
		hashUsername := userlib.Hash([]byte(userdata.Username))
		userUUID, err := uuid.FromBytes(hashUsername[:16])

		newUser := User{userdata.Username, userdata.Password, userdata.PrivateKey, userdata.SignatureKey, userdata.IntermediateUUID}
		newUserBytes, errMarshal := json.Marshal(newUser)
		if errMarshal != nil {
			return uuid.Nil, errors.New("Marshal error, cannot convert User Struct to bytes")
		}

		// Generate symmetricKey and macKey from username and password to Encrypt then Tag
		symEncKey, macKey := GenerateKeys(userdata.Username, userdata.Password)
		iv := userlib.RandomBytes(16)
		newUserEncrypted := userlib.SymEnc(symEncKey, iv, newUserBytes)
		hmacTag, hmacError := userlib.HMACEval(macKey, newUserEncrypted)
		if hmacError != nil {
			return uuid.Nil, errors.New("input as key for hmac inshould be a 16-byte key")
		}

		// Store the new userUUID to datastore
		userlib.DatastoreSet(userUUID, append(newUserEncrypted, hmacTag...))
		// //####################################################
	} else { // Invitation send by recipient
		// Create an invitation hold the IntermediateUUID and 2 interKeys
		invitationUUID = uuid.New()
		invitation = Invitation{keyFile.FileUUID, keyFile.SymKeyFile, keyFile.MacKeyFile}
	}

	invitationBytes, errMarshal := json.Marshal(invitation)
	if errMarshal != nil {
		return uuid.Nil, errors.New("Marshal error, cannot convert invitation Struct to bytes")
	}

	// Use recipient's public key to encrypt the invitation
	recipientPublicKey, ok := userlib.KeystoreGet(recipientUsername + "publicKey")
	if ok == false {
		return uuid.Nil, errors.New("CANNOT get user's publicKey on keystore")
	}
	symKeyEncrypted, invitationEncrypted, err := HybridEncryption(recipientPublicKey, invitationBytes, "encrypt-invitation")
	// Create list to hold symKeyEncrypted, invitationEncrypted and send them over server
	data := [][]byte{symKeyEncrypted, invitationEncrypted}
	dataBytes, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		return uuid.Nil, errors.New("Marshal error, cannot convert Struct to bytes")
	}
	// ###################################################

	// Use sender's signature key to sign the invitation
	signature, err := userlib.DSSign(userdata.SignatureKey, dataBytes)
	if err != nil {
		return uuid.Nil, errors.New("CANNOT sign the invitation by sender's signature key")
	}

	userlib.DatastoreSet(invitationUUID, append(dataBytes, signature...))

	return invitationUUID, nil
}

// #########################################################################            ACCEPTINVITATION         ############################################################################################
func (userdata *User) AcceptInvitation(senderUsername string, invitationPtr uuid.UUID, filename string) error {
	// Check if User already has a file with the chosen filename in their personal file
	keyFile, err := userdata.VerifyThenDecKeyFile(filename)
	if err != nil && err.Error() != "keyFileUUID does not exists in DataStore" {
		return err
	}
	if keyFile != nil {
		return errors.New("User already has a file with the chosen filename in their personal file")
	}

	// Get invitation and signature
	dataValue, ok := userlib.DatastoreGet(invitationPtr)
	if ok == false {
		return errors.New("Something about the invitationUUID is wrong")
	}
	dataBytes, signature := dataValue[:len(dataValue)-256], dataValue[len(dataValue)-256:]

	// Get the sender verify key to verify the invitation
	verifyKey, ok := userlib.KeystoreGet(senderUsername + "verifyKey")
	if ok == false {
		return errors.New("CANNOT get sender's verifyKey on keystore")
	}
	// Verify the invitation
	err = userlib.DSVerify(verifyKey, dataBytes, signature)
	if err != nil {
		return errors.New("The invitation was tampered")
	}
	// Decrypted the invitation
	var data [][]byte
	errMarshal := json.Unmarshal(dataBytes, &data)
	if errMarshal != nil {
		return errors.New("Unmarshal error, cannot convert bytes to [][]byte")
	}
	symKeyEncrypted, invitationEncrypted := data[0], data[1]
	invitationBytes, err := HybridDecryption(userdata.PrivateKey, symKeyEncrypted, invitationEncrypted)
	if err != nil {
		return err
	}
	//unmarshal
	var invite Invitation
	errMarshal = json.Unmarshal(invitationBytes, &invite)
	if errMarshal != nil {
		return errors.New("Unmarshal error, cannot convert bytes to invitation")
	}

	// Verify whether we has been revoked
	_, err = VerifyThenDecIntermediate(invite.SymKeyInter, invite.MacKeyInter, invite.IntermediateUUID)
	if err != nil {
		return errors.New("The invitation is no longer valid due to revocation")
	}

	// Create the keyfile struct for intermediateUUID
	err = userdata.EncKeyFile(filename, false, invite.IntermediateUUID, invite.SymKeyInter, invite.MacKeyInter)
	if err != nil {
		return err
	}

	return nil
}

// #########################################################################            REVOKEACCESS         ############################################################################################
func (userdata *User) RevokeAccess(filename string, recipientUsername string) error {
	//Check if The given filename does not exist in the personal file namespace of the caller
	keyFile, err := userdata.VerifyThenDecKeyFile(filename)
	if err != nil {
		return err
	}
	//Check if The given filename is not currently shared with recipientUsername
	invi, ok := userdata.IntermediateUUID[filename][recipientUsername]
	if ok == false {
		return errors.New("The given filename is not currently shared with recipientUsername")
	}

	//------------------ENCRYPT NEW FILENODE, NEW FILELOCATOR AND NEW KEYFILE-----------------

	//Download the whole content of the file and encrypt the new uuid
	content, err := userdata.LoadFile(filename)
	if err != nil {
		return err
	}

	//Delete old fileLocator
	userlib.DatastoreDelete(keyFile.FileUUID)

	// Create a new-random uuid for fileNode
	FirstFileNodeUUID := uuid.New()
	LastFileNodeUUID := uuid.New()

	// Generate a new symKey and macKey
	symKeyFN, macKeyFN, err := GenerateSymAndMacKey("enc-mac-filenode")
	if err != nil {
		return err
	}
	err = EncFileNode(symKeyFN, macKeyFN, content, FirstFileNodeUUID, uuid.Nil, LastFileNodeUUID)
	if err != nil {
		return err
	}
	err = EncFileNode(symKeyFN, macKeyFN, nil, LastFileNodeUUID, FirstFileNodeUUID, uuid.Nil)
	if err != nil {
		return err
	}

	// Create a new-random uuid for the fileLocator
	fileLocatorUUID := uuid.New()
	// Generate symmetricKey and macKey from owner to Encrypt then Tag the FileLocator
	symKeyFL, macKeyFL, err := GenerateSymAndMacKey("enc-mac-filelocator")
	if err != nil {
		return err
	}

	// Encrypt the fileLocator with new keys and new uuid
	err = EncFileLocator(symKeyFL, macKeyFL, fileLocatorUUID, FirstFileNodeUUID, LastFileNodeUUID, symKeyFN, macKeyFN)
	if err != nil {
		return err
	}

	// Update new keyfile struct for fileLocator
	err = userdata.EncKeyFile(filename, true, fileLocatorUUID, symKeyFL, macKeyFL)
	if err != nil {
		return err
	}

	//------DELETE THE INTERMEDIATE SHARED WITH REVOKE DIRECT-RECIPIENT---------------
	userlib.DatastoreDelete(invi.IntermediateUUID)
	delete(userdata.IntermediateUUID[filename], recipientUsername)

	//------UPDATE THE INTERMEDIATE FOR OTHER NON-REVOKED USERS-----------------------
	for key, users := range userdata.IntermediateUUID {
		if key == filename {
			// Get the sym and mac key for re-encrypt intermediateUUID
			for _, invite := range users {
				err = EncIntermediate(invite.SymKeyInter, invite.MacKeyInter, invite.IntermediateUUID, fileLocatorUUID, symKeyFL, macKeyFL)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
