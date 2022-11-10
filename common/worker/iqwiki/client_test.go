package iqwiki

var (
	client = NewClient()
)

// func TestGetUserActivityList(t *testing.T) {
// 	address := "0x2fE6aCD015384E1ee5138eF79fe1a434dA8FA12e"
// 	resp, err := client.GetUserActivityList(context.Background(), strings.ToLower(address))

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(len(resp))
// }

// func TestGetUserList(t *testing.T) {
// 	resp, err := client.GetUserList(context.Background())

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Log(len(resp))
// }
