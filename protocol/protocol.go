// Create a message, marshall into a byte buffer, and de-marshall it out. Check identical.

package main

import (
	proto "code.google.com/p/goprotobuf/proto"
	fmt "fmt"
	message "github.com/cf-guardian/prototype/protocol/message"
	log "log"
)

func main() {
	testMsg := &message.SlaveMessage{
		Msg:   proto.String("hello"),
		Reply: proto.String("hi, y'all")}
	data, err := proto.Marshal(testMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newTestMsg := &message.SlaveMessage{}

	err = proto.Unmarshal(data, newTestMsg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// Now test and newTest contain the same data.
	if testMsg.GetMsg() != newTestMsg.GetMsg() {
		log.Fatalf("data mismatch %q != %q", testMsg.GetMsg(), newTestMsg.GetMsg())
	} else {
		fmt.Printf("Msg component matches\n")
	}

	if testMsg.GetReply() != "" && newTestMsg.GetReply() != "" {
		if testMsg.GetReply() != newTestMsg.GetReply() {
			log.Fatalf("data mismatch %q != %q", testMsg.GetReply(), newTestMsg.GetReply())
		} else {
			fmt.Printf("Reply component matches\n")
		}
	} else {
		log.Fatalf("don't both have Reply")
	}
	fmt.Printf("Finished comparison.")

}
