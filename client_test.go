package core

import (
	"bytes"
	"context"
	"log"
	"testing"
)

func getResource() *Resource {
	cli := NewClient()
	rs, err := cli.GetResource(context.Background(), "127.0.0.1:40004", true, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	return rs
}

// go test -v -test.run=Test_List ./
func Test_List(t *testing.T) {

	rs := getResource()
	t.Log(rs.List(""))

	t.Log(rs.List("proto.Account"))
}

// go test -v -test.run=Test_Desc ./
func Test_Desc(t *testing.T) {
	rs := getResource()
	t.Log(rs.Describe("proto.Account.DescUserInfo"))
}

// go test -v -test.run=Test_Invoke ./
func Test_Invoke(t *testing.T) {
	rs := getResource()

	bs := []byte(`{"UserID":1532}`)
	buf := bytes.NewBuffer(bs)

	t.Log(rs.Invoke(context.Background(), "proto.Account.DescUserInfo", buf))
}

func Test_Invoke2(t *testing.T) {
	cli := NewClient()
	t.Log(cli.Invoke(context.Background(),
		"10.40.212.34:40004",
		"Account",
		"DescUserInfo",
		`{"UserID":1532}`,
	))
}
