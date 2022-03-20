package main

import (
	"bytes"
	"fmt"
	// "os"
	"strings"

	"github.com/mragungsetiaji/learn-go/protobuf/model"
	"github.com/golang/protobuf/jsonpb"
)

func main() {

	var star1 = &model.Star{
		Id: "x1",
		Name: "Robert Pattison",
		Gender: model.StarGender_MALE,
	}

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", star1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", star1.String())

	var movie1 = &model.Movie{
		Id: "m1",
		Title: "The Batman",
		Year: 2022,
	}

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", movie1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", movie1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, movie1)
	if err1 != nil {
		fmt.Printf("# ==== Error 1: %v \n", err1)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON\n       %v \n", jsonString)

	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.Movie)

	err := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err != nil {
		fmt.Printf("# ==== Error 2: %v \n", err)
	}
	fmt.Printf("# ==== As Proto Object\n       %#v \n", protoObject)

}

