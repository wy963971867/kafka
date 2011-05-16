package kafka

import (
	"testing"
	//"fmt"
	"bytes"
	"container/list"
)

func TestMessageCreation(t *testing.T) {
	payload := []byte("testing")
	msg := NewMessage(payload)
	if msg.magic != 0 {
		t.Errorf("magic incorrect")
		t.Fail()
	}


	// generated by kafka-rb: e8 f3 5a 06
	expected := []byte { 0xe8, 0xf3, 0x5a, 0x06 }
	if !bytes.Equal(expected, msg.checksum[:]) {
		t.Fail()
	}
}


func TestMessageEncoding(t *testing.T) {
	payload := []byte("testing")
	msg := NewMessage(payload)
	
	// generated by kafka-rb:
	expected := []byte { 0x00,0x00,0x00,0x0c,0x00,0xe8,0xf3,0x5a,0x06,0x74,0x65,0x73,0x74,0x69,0x6e,0x67 }
	if !bytes.Equal(expected, msg.Encode()) {
	  t.Fail()
	}
	
	// verify round trip
	msgDecoded := Decode(msg.Encode())
	if !bytes.Equal(msgDecoded.payload, payload) {
	  t.Fail()
	}
	if !bytes.Equal(msgDecoded.payload, payload) {
	  t.Fail()
	}
  chksum := []byte { 0xE8, 0xF3, 0x5A, 0x06 }
  if !bytes.Equal(msgDecoded.checksum[:], chksum) {
	  t.Fail()
	}
  if msgDecoded.magic != 0 {
	  t.Fail()
	}
}

func TestPublishRequestEncoding(t *testing.T) {
	payload := []byte("testing")
	msg := NewMessage(payload)
	
	messages := list.New()
	messages.PushBack(msg)
	pubBroker := NewBrokerPublisher("localhost:9092", "test", 0)
	request := pubBroker.broker.EncodePublishRequest(REQUEST_PRODUCE, messages)
	
	// generated by kafka-rb:
	expected := []byte { 0x00,0x00,0x00,0x20,0x00,0x00,0x00,0x04,0x74,0x65,0x73,0x74,
	                     0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x10,0x00,0x00,0x00,0x0c,
	                     0x00,0xe8,0xf3,0x5a,0x06,0x74,0x65,0x73,0x74,0x69,0x6e,0x67 }
	
	if !bytes.Equal(expected, request) {
	  t.Errorf("expected length: %d but got: %d", len(expected), len(request))
	  t.Errorf("expected: %X\n but got: %X", expected, request)
	  t.Fail()
	}
}

func TestConsumeRequestEncoding(t *testing.T) {
  
  pubBroker := NewBrokerPublisher("localhost:9092", "test", 0)
	request := pubBroker.broker.EncodeConsumeRequest(REQUEST_FETCH, 0, 1048576)
	
  // generated by kafka-rb, encode_request_size + encode_request
  expected := []byte { 0x00,0x00,0x00,0x18,0x00,0x01,0x00,0x04,0x74,
                       0x65,0x73,0x74,0x00,0x00,0x00,0x00,0x00,0x00,
                       0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x10,0x00,0x00 }

 	if !bytes.Equal(expected, request) {
 	  t.Errorf("expected length: %d but got: %d", len(expected), len(request))
 	  t.Errorf("expected: %X\n but got: %X", expected, request)
 	  t.Fail()
 	}
}

