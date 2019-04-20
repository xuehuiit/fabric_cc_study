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
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/msp"
	//"github.com/hyperledger/fabric/common/util"
	"encoding/json"
	//"strings"
	//"time"
	"bytes"
	"time"
	"encoding/gob"
	"github.com/hyperledger/fabric/common/tools/protolator"
)


type blockchain struct {
	Bctype string `json:"bctype"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Consensus      string `json:"consensus"`
	Storage       string    `json:"storage"`
}


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
	}else if function == "get" {
		return t.get(stub, args)
	}else if function == "set" {
		return t.set(stub, args)
	}else if function == "delete" {
		return t.delete(stub, args)
	}else if function == "GetHistoryForKey" {
		return t.GetHistoryForKey(stub, args)
	}else if function == "addNewBc" {
		return t.addNewBc(stub,args)
	}else if function == "querySearch" {
		return t.querySearch(stub,args)
	}else if function == "querySearch4Composite" {
		return t.querySearch4Composite(stub,args)
	}else if function == "set4Composite" {
		return t.set4Composite(stub,args)
	}else if function == "rangeQuery"{
		return t.rangeQuery(stub,args)
	}else if function == "querySearch4Pageing"{
		return t.querySearch4Pageing(stub,args)
	}else if function == "querySearch4Composite4Pageing"{
		return t.querySearch4Composite4Pageing(stub,args)
	}else if function == "rangeQuery4Pageing"{

		return t.rangeQuery4Pageing(stub,args)

	}else if function == "setPrivate" {

		return t.setPrivate(stub,args)

	}else if function == "getPrivate" {

		return t.getPrivate(stub,args)

	}else if function == "getChannelID" {

		return t.getChannelID(stub,args)

	}else if function == "getBinding" {

		return t.getBinding(stub,args)

	}else if function == "getTransient" {

		return t.getTransient(stub)

	}else if function == "getCreator" {

		return t.getCreator(stub)

	}else if function == "getSignedProposal" {

		return t.getSignedProposal(stub)

	}else if function == "getTxTimeStamp" {

		return t.getTxTimeStamp( stub )

	}else if function == "setStateValidationParameter" {

		return t.setStateValidationParameter( stub )

	}else if function == "getStateValidationParameter" {

		return t.getStateValidationParameter( stub ,args )

	}else if function == "InvokeChaincode"{

		return t.InvokeChaincode( stub ,args )

	}else if function == "GetTxID" {
		return t.getTxID(stub)

	}


	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"query\"")


}


////  ============================  chaincode 之间互相调用 ============================================


func (t *fabriccc) InvokeChaincode( stub shim.ChaincodeStubInterface , args []string ) pb.Response {


	var a_parm_channelname = args[0]
	var b_parm_chaincodename = args[1]
	var c_parm_callmethod = args[2]
	var c_parm1 = args[3]
	var c_parm2 = args[4]
	var c_parm3 = args[5]

	cow_parms := []string{c_parm_callmethod,c_parm1,c_parm2,c_parm3}
	queryArgs1 := make([][]byte, len(cow_parms))
	for i, arg := range cow_parms {
		queryArgs1[i] = []byte(arg)
	}

	cow_response := stub.InvokeChaincode(b_parm_chaincodename,queryArgs1,a_parm_channelname)


	if cow_response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", cow_response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}

	stub.PutState("1_"+c_parm2,[]byte(c_parm3))
	stub.PutState("2_"+c_parm2,[]byte(c_parm3))
	stub.PutState("3_"+c_parm2,[]byte(c_parm3))
	stub.PutState("4_"+c_parm2,[]byte(c_parm3))


	cow_result := string(cow_response.Payload)

	fmt.Printf(" cow info -   result : %s   \n  ", cow_result )

	return shim.Success([]byte(cow_result))


}






////  =============================  chaincode 中设置 KEY 级别的背书策略  ===================================


/**
  给某个key设置规则
 */
func (t *fabriccc) setStateValidationParameter(stub shim.ChaincodeStubInterface) pb.Response {


	err := stub.SetStateValidationParameter("key_org1",[]byte("OR ('Org1MSP.member','Org2MSP.member')"))
	if err != nil {
		fmt.Printf("Returning error ******************\n")
		return shim.Error(err.Error())
	}



	return shim.Success(nil)

}



/**
  给某个key设置规则
 */
func (t *fabriccc) getStateValidationParameter( stub shim.ChaincodeStubInterface , args []string ) pb.Response {



	//var a_parm = args[0]
	var b_parm = args[1]
	//var c_parm = args[2]


	policy , err := stub.GetStateValidationParameter(b_parm)
	if err != nil {
		fmt.Printf("Returning error ******************\n")
		return shim.Error(err.Error())
	}

	fmt.Printf(" %s ",string(policy))

	return shim.Success( policy )


}



////   ====================================================================================================

/**
  获取当前交易时间
 */
func (t *fabriccc) getTxTimeStamp(stub shim.ChaincodeStubInterface) pb.Response {


	fmt.Printf("\nBegin*** getTxTimeStamp \n")
	txTimeAsPtr, err := stub.GetTxTimestamp()
	if err != nil {
		fmt.Printf("Returning error ******************\n")
		return shim.Error(err.Error())
	}
	fmt.Printf("\t returned value from stub: %v\n", txTimeAsPtr)
	fmt.Printf("\t After converting time to Unix format %s \n", time.Unix(txTimeAsPtr.Seconds, int64(txTimeAsPtr.Nanos)).String())
	fmt.Printf("\nEnd*** getTxTimeStamp \n")
	//return shim.Success([]byte(txTimeAsPtr))
	return shim.Success(nil)

}


/**
  获取当前交易的提案信息
 */
func (t *fabriccc) getSignedProposal(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Printf("\nBegin*** getSignedProposal \n")
	signedProposal, err := stub.GetSignedProposal()

	if err != nil {
		fmt.Printf("Returning error ****************** ")
		return shim.Error(err.Error())
	}
	fmt.Printf("\t returned value from stub: %v", signedProposal)
	fmt.Printf("\nEnd*** getSignedProposal \n")
	buf := &bytes.Buffer{}
	protolator.DeepMarshalJSON(buf, signedProposal)
	fmt.Printf(string(buf.Bytes()))

	return shim.Success([]byte(buf.Bytes()))

}

/**
获取当前方法调用者的身份信息
 */
func (t *fabriccc) getCreator(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Printf("\nBegin*** getCreator \n")
	creator, err := stub.GetCreator()

	if err != nil {
		fmt.Printf("GetCreator Error")
		return shim.Error(err.Error())
	}

	si := &msp.SerializedIdentity{}
	err2 := proto.Unmarshal(creator, si)

	if err2 != nil {
		fmt.Printf("Proto Unmarshal Error")
		return shim.Error(err2.Error())
	}


	buf := &bytes.Buffer{}
	protolator.DeepMarshalJSON(buf, si)
	fmt.Printf("End*** getCreator \n")
	fmt.Printf(string(buf.Bytes()))



	//return shim.Success([]byte(buf.Bytes()))


	return shim.Success(creator)
}



/**
 * 客户端在发起提案的时候可以放置客户要的私有信息

peer chaincode invoke -o orderer.qklszzn.com:7050 -C mytestchannel -n fb14_cc_study  -c '{"Args":["getTransient","a","eos","b4","blockchain"]}' --transient "{\"name\":\"IHJvYmVydGZlbmcgaGVyZQ==\"}"


目前还没有验证是否信息会被提交到区块中 TODO ??????


 */
func (t *fabriccc) getTransient(stub shim.ChaincodeStubInterface) pb.Response {


	fmt.Printf("\nBegin*** getTransient \n")
	payload, err := stub.GetTransient()
	fmt.Printf(" payload from chaincode : %v", payload)
	if err != nil {
		return shim.Error(err.Error())
	}
	for key, currArg := range payload {
		fmt.Printf("Inside ... Loop")
		fmt.Printf("payload[%d] := %s\n", key, currArg)
	}
	b, err2 := GetBytes(payload)
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	fmt.Printf("\nEnd*** getTransient \n")
	return shim.Success([]byte(b))
}

// 还不知道做什么用的
func (t *fabriccc) getBinding(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	/*bindtype , err := stub.GetBinding()

	if err != nil {
		fmt.Printf("Error getting transaction timestamp: %s", err)
		return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err))
	}

	fmt.Printf("GetBinding : %v \n ",bindtype  )

	return shim.Success(bindtype)*/


	binding, err := stub.GetBinding()
	if err != nil {
		fmt.Printf("Returning error ****************** ")
		return shim.Error(err.Error())
	} else if binding == nil {
		fmt.Printf("###### No Transaction Binding is generated ###### ")
		return shim.Error("###### No Transaction Binding is generated ###### ")
	}
	fmt.Printf("\t returned value from stub : %v\n", binding)
	fmt.Printf("End*** getBinding \n")
	return shim.Success(binding)

}



// 获取当前的交易时间
func (t *fabriccc) etTxTimestamp(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	txtime,err:= stub.GetTxTimestamp()
	if err != nil {
		fmt.Printf("Error getting transaction timestamp: %s", err)
		return shim.Error(fmt.Sprintf("Error getting transaction timestamp: %s", err))
	}


	tm := time.Unix(txtime.Seconds, 0)

	fmt.Printf("Transaction Time: %v \n ", tm.Format("2006-01-02 03:04:05 PM"))

	return shim.Success([]byte(fmt.Sprint("  time is :   %s   ",tm.Format("2006-01-02 15:04:05"))))

}




// 获取当前的交易编号
func (t *fabriccc) getTxID(stub shim.ChaincodeStubInterface ) pb.Response {

	txid := stub.GetTxID();
	fmt.Println("  ========  GetTxID   ==========  %s  ",txid)
	return shim.Success([]byte(txid))


}





// 获取当前的channel的名称
func (t *fabriccc) getChannelID(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is setPrivate   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm )


	channelid := stub.GetChannelID();
	fmt.Println("  ========  GetTxID   ==========  %s  ",channelid)

	return shim.Success( []byte( "success invok " + channelid  )  )


}




////   ===================   private data =====================================   /////


//PutPrivateData(collection,key,value)
//GetPrivateData(collection,key)
//GetPrivateDataByRange(collection, startKey, endKey string)

// 上述三个方法和普通的key处理方法是一样，么有什么太多的特色，具体倒是详细说明即可


// test setPrivate method
func (t *fabriccc) setPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is setPrivate   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm )

	//stub.PutState(b_parm,[]byte(c_parm))

	stub.PutPrivateData(a_parm,b_parm,[]byte(c_parm));


	return shim.Success( []byte( "success invok " + c_parm  )  )


}


// test getPrivate method
func (t *fabriccc) getPrivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]

	fmt.Println("  ========  curr method is getPrivate   ========== ")
	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )

	//keyvalue,err := stub.GetState(b_parm)
	keyvalue, err := stub.GetPrivateData( a_parm , b_parm )

	if( err != nil  ){

		return shim.Error(" finad error! ")
	}

	return shim.Success( keyvalue )

}


////   ============   =================   /////



/**
 区间查询
 */
func (t *fabriccc) rangeQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var a_parm= args[0]
	var b_parm= args[1]
	var c_parm= args[2]


	fmt.Println("  ========  curr method is rangeQuery   ==========  ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm  )


	resultsIterator, err := stub.GetStateByRange(a_parm, b_parm) //这里应为传入参数，但为了简化这里直接Hard code 为 car1 、 car3
	if err != nil {
		return shim.Error("Query by Range failed")
	}
	defer resultsIterator.Close() //释放迭代器

	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString(`{"result":[`)

	for resultsIterator.HasNext() {

		queryResponse, err := resultsIterator.Next() //获取迭代器中的每一个值
		key := queryResponse.Key

		if err != nil {
			return shim.Error("Fail")
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}


		buffer.WriteString(key+" -> "+string(queryResponse.Value)) //将查询结果放入Buffer中
		bArrayMemberAlreadyWritten = true

	}

	buffer.WriteString(`]}`)
	fmt.Print("Query result: %s", buffer.String())

	return shim.Success(buffer.Bytes())


}


/**
 带有分页的区间查询
 */
func (t *fabriccc) rangeQuery4Pageing(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var a_parm= args[0]
	var b_parm= args[1]
	var c_parm= args[2]


	fmt.Println("  ========  curr method is rangeQuery   ==========  ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm  )


	resultsIterator, err := stub.GetStateByRange(a_parm, b_parm)


	resultsIterator, responseMetadata, err := stub.GetStateByRangeWithPagination(a_parm, b_parm ,3 ,"" )


	defer resultsIterator.Close()


	if err != nil {
		return shim.Error(err.Error())
	}

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, responseMetadata)

	fmt.Printf("- rangeQuery4Pageing queryResult:\n%s\n", bufferWithPaginationInfo.String())


	//return buffer.Bytes(), nil


	return shim.Success( []byte( "success querySearch "  )  )


}





//======================================     组合键存储和查询   ===================================


/**
 基于符合键的查询
 */
func (t *fabriccc) querySearch4Composite(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]



	fmt.Println("  ========  curr method is querySearch4Composite   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm  )


	colorIdResultsIterator, err := stub.GetStateByPartialCompositeKey ("bctype#name#conose", []string{a_parm})


	//colorIdResultsIterator, err := stub.GetStateByPartialCompositeKeyWithPagination ("bctype#name#conose", []string{a_parm},3,"");



	if err != nil {
		return shim.Error(err.Error())
	}

	defer colorIdResultsIterator.Close()

	for colorIdResultsIterator.HasNext() {

		responseRange, err := colorIdResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}


		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedColor := compositeKeyParts[0]
		returnedMarbleName := compositeKeyParts[1]

		fmt.Printf("- found a marble from index:%s bctype:%s name:%s\n", objectType, returnedColor, returnedMarbleName)

	}


	return shim.Success( []byte( "success querySearch "  )  )

}



/**
 基于符合键的查询
 */
func (t *fabriccc) querySearch4Composite4Pageing(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]



	fmt.Println("  ========  curr method is querySearch4Composite   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm  )


	//返回包含给出颜色的组合键的迭代器
	//colorIdResultsIterator, err := stub.GetStateByPartialCompositeKey ("bctype#name#conose", []string{a_parm})

	resultsIterator, responseMetadata, err  := stub.GetStateByPartialCompositeKeyWithPagination ("bctype#name#conose", []string{a_parm},3,"");

	defer resultsIterator.Close()


	if err != nil {
		return shim.Error(err.Error())
	}

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, responseMetadata)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", bufferWithPaginationInfo.String())


	//return buffer.Bytes(), nil


	return shim.Success( []byte( "success querySearch "  )  )


}






/**
 基于符合键的查询
 */
func (t *fabriccc) set4Composite(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]
	var d_parm = args[3]


	fmt.Println("  ========  curr method is set4Composite   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm )


	//基于复合键的存储模式
	indexName := "bctype#name#conose"
	colorNameIndexKey, err := stub.CreateCompositeKey( indexName , []string{ a_parm , b_parm } )

	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	stub.PutState(colorNameIndexKey, value)


	return shim.Success( []byte(d_parm) )


}



//================================================================================================================




/**
 通过复杂的条件查询，适用于coushdb的存放模型
 */
func (t *fabriccc) querySearch(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]



	fmt.Println("  ========  curr method is set   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm  )


	queryString := fmt.Sprintf(`{"selector":{"bctype":"%s"}}`, "pow")
	resultsIterator, err := stub.GetQueryResult(queryString)

	if err != nil {
		return shim.Error("Rich query failed")
	}
	defer resultsIterator.Close() //释放迭代器

	var buffer bytes.Buffer
	bArrayMemberAlreadyWritten := false
	buffer.WriteString(`{"result":[`)


	for resultsIterator.HasNext() {

		queryResponse, err := resultsIterator.Next() //获取迭代器中的每一个值
		if err != nil {
			return shim.Error("Fail")
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value)) //将查询结果放入Buffer中
		bArrayMemberAlreadyWritten = true

	}

	buffer.WriteString(`]}`)
	fmt.Print("Query result: %s", buffer.String())


	return shim.Success( []byte( "success querySearch "  )  )

}

/**
 通过复杂的条件查询，适用于coushdb的存放模型( 带有扉页模式 )
 */
func (t *fabriccc) querySearch4Pageing(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]



	fmt.Println("  ========  curr method is set   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm  )


	queryString := fmt.Sprintf(`{"selector":{"bctype":"%s"}}`, "pow")

	resultsIterator, responseMetadata, err := stub.GetQueryResultWithPagination( queryString ,3 ,"" )


	defer resultsIterator.Close()


	if err != nil {
		return shim.Error(err.Error())
	}

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	bufferWithPaginationInfo := addPaginationMetadataToQueryResults(buffer, responseMetadata)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", bufferWithPaginationInfo.String())


	//return buffer.Bytes(), nil


	return shim.Success( []byte( "success querySearch "  )  )

}



/**
 初始化
 */
func (t *fabriccc) addNewBc(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]
	var d_parm = args[3]


	fmt.Println("  ========  curr method is set   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm )

	bc := &blockchain{a_parm, b_parm, c_parm, d_parm}
	bcJSONasBytes, err := json.Marshal(bc)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(b_parm, bcJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}



	//基于复合键的存储模式
	indexName := "bctype~name"
	colorNameIndexKey, err := stub.CreateCompositeKey( indexName , []string{bc.Bctype,bc.Name } )

	if err != nil {
		return shim.Error(err.Error())
	}

	value := []byte{0x00}
	stub.PutState(colorNameIndexKey, value)


	return shim.Success( []byte(d_parm) )


	}






// test set method
func (t *fabriccc) set(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is set   ========== ")

	fmt.Printf(" parm is  %s  %s  %s  \n " , a_parm , b_parm , c_parm )

	stub.PutState(b_parm,[]byte(c_parm))
	return shim.Success( []byte( "success invok " + c_parm  )  )


}


// test get method
func (t *fabriccc) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]

	fmt.Println("  ========  curr method is get   ========== ")
	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )

	keyvalue,err := stub.GetState(b_parm)

	if( err != nil  ){

		return shim.Error(" finad error! ")
	}

	return shim.Success( keyvalue )

}

// test get method
func (t *fabriccc) getprivate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]

	fmt.Println("  ========  curr method is get   ========== ")
	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


	//keyvalue,err := stub.GetPrivateData()

	keyvalue,err := stub.GetState(b_parm)


	if( err != nil  ){

		return shim.Error(" finad error! ")
	}

	return shim.Success( keyvalue )

}

// test delete method
func (t *fabriccc) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]

	fmt.Println("  ========  curr method is delete   ========== ")
	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


	err := stub.DelState(b_parm)

	if err != nil {
		return shim.Error(" 删除出现错误！！！！！")
	}

	return shim.Success([]byte(" 删除正确！！！！！  "))


}



// test GetHistoryForKey method
func (t *fabriccc) GetHistoryForKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]

	fmt.Println("  ========  curr method is GetHistoryForKey   ========== ")
	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


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

		/*txtimesamp :=response.Timestamp
		tm := time.Unix(txtimesamp.Seconds, 0)
		datestr := tm.Format("2006-01-02 03:04:05 PM")
*/

		fmt.Printf(" Tx info -   txid : %s   value :  %s  if delete: %t   datetime : %s \n ", txid , string(txvalue) , txstatus , " " )

		keys = append( keys , txid)

	}


	jsonKeys, err := json.Marshal(keys)
	if err != nil {
		return shim.Error(fmt.Sprintf("query operation failed. Error marshaling JSON: %s", err))
	}

	return shim.Success(jsonKeys)

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

		var keyvalue []byte
		var err error
		keyvalue,err = stub.GetState(b_parm)

		if( err != nil  ){

			return shim.Error(" finad error! ")
		}


		return shim.Success( keyvalue )


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

			/*txtimesamp :=response.Timestamp
			tm := time.Unix(txtimesamp.Seconds, 0)
			datestr := tm.Format("2006-01-02 03:04:05 PM")*/


			fmt.Printf(" Tx info -   txid : %s   value :  %s  if delete: %t   datetime : %s \n ", txid , string(txvalue) , txstatus , "  " )

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


func (t *fabriccc) Query( stub shim.ChaincodeStubInterface) pb.Response {


	_,args := stub.GetFunctionAndParameters()

	var a_parm= args[0]
	var b_parm= args[1]
	var c_parm= args[2]

	fmt.Println("  ========  curr method big  query  ========== ")

	fmt.Printf(" parm is  %s  %s  %s   \n ", a_parm, b_parm, c_parm)

	return shim.Success( []byte("return for big query") );
}

// test query method
func (t *fabriccc) query( stub shim.ChaincodeStubInterface , args []string) pb.Response {



	var a_parm = args[0]
	var b_parm = args[1]
	var c_parm = args[2]


	fmt.Println("  ========  curr method is small query  ========== ")

	fmt.Printf(" parm is  %s  %s  %s   \n " , a_parm , b_parm , c_parm )


	// 设定值
	if a_parm == "set"{

		stub.PutState(b_parm,[]byte(c_parm))
		return shim.Success( []byte( "success invok " + c_parm  )  )

	}else if a_parm == "get"{   //取单个值

		var keyvalue []byte
		var err error
		keyvalue,err = stub.GetState(b_parm)

		if( err != nil  ){

			return shim.Error(" finad error! ")
		}


		return shim.Success( keyvalue )


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
			//txtimesamp :=response.Timestamp

			//tm := time.Unix(txtimesamp., 0)
			//datestr := tm.Format("2006-01-02 03:04:05 PM")


			fmt.Printf(" Tx info -   txid : %s   value :  %s  if delete: %t   datetime : %s \n ", txid , string(txvalue) , txstatus , "d" )

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



///////////   ===  mgethod ====


func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}


func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}


func addPaginationMetadataToQueryResults(buffer *bytes.Buffer, responseMetadata *pb.QueryResponseMetadata) *bytes.Buffer {

	buffer.WriteString("[{\"ResponseMetadata\":{\"RecordsCount\":")
	buffer.WriteString("\"")
	buffer.WriteString(fmt.Sprintf("%v", responseMetadata.FetchedRecordsCount))
	buffer.WriteString("\"")
	buffer.WriteString(", \"Bookmark\":")
	buffer.WriteString("\"")
	buffer.WriteString(responseMetadata.Bookmark)
	buffer.WriteString("\"}}]")

	return buffer
}






func main() {
	err := shim.Start(new(fabriccc))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

