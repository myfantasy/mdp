package mdp

const connStringTest = `{"server":"http://localhost:9170/", "token":"abc"}`

// func TestItemsGetQuery(t *testing.T) {
// 	ti := ItemsGetQuery{TableName: "ttbl_n", IKeys: &[]int64{5, 7, 10}}

// 	b, err := json.Marshal(ti)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if string(b) != `{"table":"ttbl_n","iks":[5,7,10]}` {
// 		t.Fatal("Error convert to json", string(b))
// 	}
// }

// func TestAllDataIGet(t *testing.T) {
// 	c, err := ConnectionGetFromJSON([]byte(connStringTest))

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ig, statusCode, err := c.AllDataIGet("tst1")

// 	if err != nil {
// 		t.Fatal("query fail: ", err, string(ig.ParamsErr), string(ig.InternalErr))
// 	}

// 	if statusCode != 200 {
// 		t.Fatal("statusCode: ", statusCode, "!=", 200)
// 	}

// 	if len(ig.Items) == 0 {
// 		t.Fatal("len(ig.Items) ==", 0)
// 	}
// }
