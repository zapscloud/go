package main

import (
	"fmt"
	"log"

	"github.com/kr/pretty"
	"github.com/zapscloud/golib/database"
	"github.com/zapscloud/golib/zaps"
)

func init() {
}

func main() {
	zapscloud, err := zaps.NewZapsCloud("applicationname", "key", "secret")
	zapscloud.SetStage("dev")

	log.Println("Error ", err)

	// zapscloud.SetTenant("tenant")
	zapsdb, err := database.NewZapsDB(zapscloud)
	log.Println("Error ", err)

	res, err := zapsdb.GetOne("student_class", "34", "")
	fmt.Println("Result values ", err)
	pretty.Println(res)

	txnid := ""

	// txndata, err := zapsdb.StartTransaction()
	// fmt.Println("Txn  Result ", err)
	// pretty.Println(txndata)

	// txnid = txndata["transaction_id"].(string)

	reqbody := map[string]interface{}{
		"class_id":    41,
		"class_name":  "class II41",
		"class_total": 41,
	}

	res, err = zapsdb.Insert("student_class", reqbody, txnid)
	fmt.Println("Insert response ", res, err)

	reqbody = map[string]interface{}{
		"class_id":   41,
		"class_name": "class II 41 updated",
	}
	res, err = zapsdb.UpdateOne("student_class", "41", reqbody, txnid)
	fmt.Println("Update response ", res, err)

	// reqbody = []byte(`{
	// 	"class_name": "class VI updated"
	// }`)
	// filter := `{"class_id":"6"}`
	// res, err = zapsdb.UpdateMany("student_class", filter, reqbody)
	// fmt.Println("Update response ", res, err)

	// filter := `{"class_total":{"$gte":30}}`
	// res, err = zapsdb.GetMany("student_class", filter, "", 0, 0)
	// fmt.Println("Result values ", err)

	// pretty.Println(res)

	// aggvalue := `{"_id":{"student_class":"student_class"},  "count":{"$sum":1}}`

	// res, err = zapsdb.Aggregate("student_class", "", aggvalue, "", 0, 0)
	// fmt.Println("Result values ", res, err)

	// res, err = zapsdb.DeleteOne("student_class", "38", txnid)
	// fmt.Println("Delete response ", res, err)

	// filter := `{"class_total":{"$gte":40}}`
	// res, err = zapsdb.DeleteMany("student_class", filter, txnid)
	// fmt.Println("Delete response ", res, err)

	// duration := time.Duration(10) * time.Second
	// time.Sleep(duration)

	// if err != nil {
	// 	txndata, err = zapsdb.RollbackTransaction(txnid)
	// 	fmt.Println("Txn  Result ", err)
	// 	pretty.Println(txndata)
	// } else {
	// 	txndata, err = zapsdb.CommitTransaction(txnid)
	// 	fmt.Println("Txn  Result ", err)
	// 	pretty.Println(txndata)
	// }

	// res, err := zapsdb.CreateCollection("testcollection", "id", "This is test collection using go lang")
	// fmt.Println("Collection ", res, err)

	// res, err := zapsdb.GetCollectionList("", "", 0, 0)
	// resjson, err := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Collection list :: ", string(resjson), err)

	// // res, err = zapsdb.RemoveCollection("testcollection")
	// // fmt.Println("Remove collection ", res, err)

	// res, err = zapsdb.GetCollection("student")
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Collection Get Resonse ::", string(resjson), err)

	// zapsapp, err := app.NewZapsApp(zapscloud)

	// res, err := zapsapp.GetTenantList("", "", 0, 5)
	// resjson, err := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Collection list :: ", string(resjson), err)

	// res, err = zapsapp.GetTenant("tenant4")
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Tenant Detail :: ", string(resjson), err)

	// zapsauth, err := auth.NewZapsAuth(zapscloud)
	// fmt.Println("Client Auth :: ", err)

	// res, err = zapsauth.GetClientList("", "", 0, 0)
	// fmt.Println("Client List :: ", err)

	// resjson, err := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client List JSON:: ", string(resjson), err)

	// res, err = zapsauth.GetClient("book001")
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

	// res, err := zapsauth.ValidateClient("book00112", "book00112")
	// resjson, err := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Validate Client detail :: ", string(resjson), err)

	// reqbody := []byte(`{
	// 	"client_key": "book00112",
	// 	"client_secret": "book00112",
	// 	"grants": [
	// 		"client_credentials"
	// 	]
	// }`)

	// res, err = zapsauth.CreateClient(reqbody)
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Create Client detail :: ", string(resjson), err)

	// res, err = zapsauth.RemoveClient("book00112")
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

	// res, err := zapsauth.GetUserList(`user_id="91999999999"`, "", 0, 0)
	// resjson, err := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Validate User List: ", string(resjson), err)

	// res, err = zapsauth.GetUser("91999999999")
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

	// reqbody := []byte(`{
	// 	"user_id": "91999999999",
	// 	"user_secret": "password",
	// 	"user_name":"gokultest",
	// 	"user_status":"A"
	// }`)

	// res, err = zapsauth.CreateUser(reqbody)
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

	// res, err = zapsauth.ValidateUser("91999999999", "password")
	// resjson, err = json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

	// reqbody := []byte(`{
	// 	"auth":"Basic ZGJwZGV2MDAyOmRicGRldjAwMg==",
	// 	"tokenrequest":{
	// 		"grant_type":"password",
	// 		"username": "91999999999",
	// 		"password": "password"
	// 	}
	// }`)

	// res, err := zapsauth.AccessToken(reqbody)
	// resjson, err := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

	// res, err := zapsauth.ValidateToken("fde2d7eb665200b9a39b34106c815d7beef778d01")
	// resjson, _ := json.MarshalIndent(res, "", "  ")
	// fmt.Println("Client detail :: ", string(resjson), err)

}
