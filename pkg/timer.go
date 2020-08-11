package pkg

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"sync/atomic"
	"time"
)

type TimerConf struct {
	Interval    time.Duration
	TotalTimes  int
	Destination string
	Filename    string
	Block       int
	TextType    string
}

type UdpTimer struct {
	data    []byte
	conf    *TimerConf
	udpAddr *net.UDPAddr
}

func NewUdpTimer(conf *TimerConf) (*UdpTimer, error) {
	str, err := ioutil.ReadFile(conf.Filename)
	if err != nil {
		panic(fmt.Sprintf("cannot read file: %v", err))
	}

	fmt.Println(string(str))

	data := []byte{}
	if conf.TextType == "binary" {
		data, err = hex.DecodeString(strings.TrimSuffix(string(str), "\n"))
		if err != nil {
			panic(fmt.Sprintf("file is not in hex, %v", err))
		}
	} else {
		data = []byte(strings.TrimSuffix(string(str), "\n"))
	}

	dst, err := net.ResolveUDPAddr("udp", conf.Destination)
	if err != nil {
		panic(fmt.Sprintf("resolve udp address failed: %v", err))
	}

	return &UdpTimer{
		data:    data,
		conf:    conf,
		udpAddr: dst,
	}, err
}

func (t *UdpTimer) Exec() {
	execFunc := func(b *int32) {
		if err := t.SendUdpPacket(); err != nil {
			fmt.Printf("send udp failed: %v", err)
		}
		atomic.StoreInt32(b, 1)
	}

	for times := 0; times < t.conf.TotalTimes; times++ {
		tr := time.NewTimer(t.conf.Interval)
		flagInt := int32(0)

		go execFunc(&flagInt)
		select {
		case <-tr.C:
			if atomic.LoadInt32(&flagInt) != 1 {
				panic("command is not done in timeslot")
			}
		}
	}
}

func (t *UdpTimer) SendUdpPacket() error {
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		fmt.Printf("create listen failed: %v\n", err)
		return err
	}
	defer conn.Close()

	for i := 0; i < t.conf.Block; i++ {
		_, err = conn.WriteTo(t.data, t.udpAddr)
		if err != nil {
			fmt.Printf("send udp failed: %v\n", err)
			return err
		}
	}

	return nil
}
