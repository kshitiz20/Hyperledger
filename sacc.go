package main

import (
  "fmt"

    "encoding/json"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type ChaincodeForGettingDetails struct {

}

type Costumer struct {
SSN   string `json:"ssn"`
FirstName  string `json:"firstname"`
LastName string `json:"lastname"`
MobileNumber  string `json:"mobileno"`
  DOB string `json:"dob"`
  Email string `json:"email"`
}

func (t *ChaincodeForGettingDetails) Init(stub shim.ChaincodeStubInterface) peer.Response {
return shim.Success(nil)
}

func (t *ChaincodeForGettingDetails) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

    var fn,args = stub.GetFunctionAndParameters()
    //var result string
   //  var err error
    if fn == "update"{
      return t.update(stub, args)
    }else if fn == "query"{
      return t.query(stub, args)
    }

    return shim.Error("Invalid Smart Contract function name.")
}

func (t *ChaincodeForGettingDetails) update(stub shim.ChaincodeStubInterface, args []string) peer.Response{
  if len(args)!=6{
    return shim.Error("Incorrect args. Give some more arguments")
  }
  var newcost= Costumer{FirstName:args[1], LastName:args[2],MobileNumber:args[3],DOB:args[4], Email:args[5]}
  newcostasBytes,_ :=json.Marshal(newcost)
  stub.PutState(args[0], newcostasBytes)
  return shim.Success(nil)
}


func (t *ChaincodeForGettingDetails) query(stub shim.ChaincodeStubInterface, args []string) peer.Response{

if len(args) != 1 {
return shim.Error("Incorrect number of arguments. Expecting 1")
}

carAsBytes, _ := stub.GetState(args[0])
return shim.Success(carAsBytes)
}

func main() {

// Create a new Smart Contract
err := shim.Start(new(ChaincodeForGettingDetails))
if err != nil {
fmt.Printf("Error creating new Smart Contract: %s", err)
}
}
