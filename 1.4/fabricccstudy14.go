/**

  Copyright xuehuiit Corp. 2018 All Rights Reserved.

  http://www.xuehuiit.com

  QQ 411321681

 */

package main


import (

	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"github.com/hyperledger/fabric/common/util"
	"encoding/json"
	//"strings"
	"time"
)

//定义一个机构体，作为chaincode的主对象，可以是任何符合go语言规范的命名方式
type fabriccc struct {

}

/** 

	系统初始化方法， 在部署chaincode的过程中当执行命令

    peer chaincode instantiate -o orderer.robertfabrictest.com:7050 -C
         roberttestchannel -n r_test_cc6 -v 1.0 -c '{"Args":["init","a","100","b","200"]}'
         -P "OR	('Org1MSP.member','Org2MSP.member')"

    的时候会调用该方法


	https://github.com/hyperledger/fabric/blob/release/core/chaincode/shim/interfaces.go  所有的注释这里

*/
func (t *fabriccc) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println(" <<  ========  success init it is view in docker  ==========  >> ")
	return shim.Success([]byte("success init "))
}

/**

  主业务逻辑，在执行命令
  peer chaincode invoke -o 192.168.23.212:7050 -C roberttestchannel -n r_test_cc6 -c '{"Args":["invoke","a","b","1"]}'

  的时候系统会调用该方法并传入相关的参数，注意 "invoke" 之后的参数是需要传入的参数

*/

func (t *fabriccc) Invoke(stub shim.ChaincodeStubInterface) pb.Response {


	function,args := stub.GetFunctionAndParameters()

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is big invoke   ========== ")
	fmt.Printf(" parm is  %s  %s  %s  %s  \n " , a_parm , b_parm , c_parm ,function)


	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")


}


// test invoke method
func (t *fabriccc) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {



	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is small invoke   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  %s  \n " , a_parm , b_parm , c_parm )



	// 设定值
	if a_parm == "set"{

		stub.PutState(b_parm,[]byte(c_parm))
		return shim.Success( []byte( "success invok " + c_parm  )  )

	}else if a_parm == "get"{   //取单个值

		return t.query(stub, args)


	}else if a_parm == "delete" { //删除某个值


		fmt.Println("  ========  delete   ========== %s ",b_parm)

		err := stub.DelState(b_parm)

		if err != nil {
			return shim.Error(" 删除出现错误！！！！！")
		}

		return shim.Success([]byte(" 删除正确！！！！！  "))


	}else if a_parm == "GetHistoryForKey" { //取单个值的历史记录

		keysIter, err := stub.GetHistoryForKey(b_parm);


		if err != nil {
			return shim.Error(fmt.Sprintf("GetHistoryForKey failed. Error accessing state: %s", err))
		}
		defer keysIter.Close()

		var keys []string

		for keysIter.HasNext() {

			response, iterErr := keysIter.Next()
			if iterErr != nil {
				return shim.Error(fmt.Sprintf("GetHistoryForKey operation failed. Error accessing state: %s", err))
			}

			//交易编号
			txid := response.TxId
			//交易的值
			txvalue := response.Value
			//当前交易的状态
			txstatus := response.IsDelete
			//交易发生的时间戳
			txtimesamp :=response.Timestamp

			tm := time.Unix(txtimesamp.Seconds, 0)
			datestr := tm.Format("2006-01-02 03:04:05 PM")


			fmt.Printf(" Tx info -   txid : %s   value :  %s  if delete: %t   datetime : %s \n ", txid , string(txvalue) , txstatus , datestr )

			keys = append( keys , txid)

		}


		jsonKeys, err := json.Marshal(keys)
		if err != nil {
			return shim.Error(fmt.Sprintf("query operation failed. Error marshaling JSON: %s", err))
		}

		return shim.Success(jsonKeys)


	}else{

		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))

	}



	return shim.Success(  []byte( "curr method is invoke" )  )



}


// test query method
func (t *fabriccc) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {



	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is query  ========== ")

	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


	// 设定值
	if a_parm == "set"{

		stub.PutState(b_parm,[]byte(c_parm))
		return shim.Success( []byte( "success invok " + c_parm  )  )

	}else if a_parm == "get"{   //取单个值

		return t.query(stub, args)


	}else if a_parm == "delete" { //删除某个值


		fmt.Println("  ========  delete   ========== %s ",b_parm)

		err := stub.DelState(b_parm)

		if err != nil {
			return shim.Error(" 删除出现错误！！！！！")
		}

		return shim.Success([]byte(" 删除正确！！！！！  "))


	}else if a_parm == "GetHistoryForKey" { //取单个值的历史记录

		keysIter, err := stub.GetHistoryForKey(b_parm);


		if err != nil {
			return shim.Error(fmt.Sprintf("GetHistoryForKey failed. Error accessing state: %s", err))
		}
		defer keysIter.Close()

		var keys []string

		for keysIter.HasNext() {

			response, iterErr := keysIter.Next()
			if iterErr != nil {
				return shim.Error(fmt.Sprintf("GetHistoryForKey operation failed. Error accessing state: %s", err))
			}

			//交易编号
			txid := response.TxId
			//交易的值
			txvalue := response.Value
			//当前交易的状态
			txstatus := response.IsDelete
			//交易发生的时间戳
			txtimesamp :=response.Timestamp

			tm := time.Unix(txtimesamp.Seconds, 0)
			datestr := tm.Format("2006-01-02 03:04:05 PM")


			fmt.Printf(" Tx info -   txid : %s   value :  %s  if delete: %t   datetime : %s \n ", txid , string(txvalue) , txstatus , datestr )

			keys = append( keys , txid)

		}


		jsonKeys, err := json.Marshal(keys)
		if err != nil {
			return shim.Error(fmt.Sprintf("query operation failed. Error marshaling JSON: %s", err))
		}

		return shim.Success(jsonKeys)


	}else{

		return shim.Success([]byte("success invok  and Not opter !!!!!!!! "))

	}



	return shim.Success( []byte("curr method is query") )
}



func main() {
	err := shim.Start(new(fabriccc))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

