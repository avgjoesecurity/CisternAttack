// Written by Ben Glass. Use to do cool stuff.

package main

import (
	"log"
	"os"

	"github.com/k-sone/snmpgo"
)

type TrapListener struct {
	logger *log.Logger
}

func (l *TrapListener) OnTRAP(trap *snmpgo.TrapRequest) {
	// simple logging
	if trap.Error == nil {
		l.logger.Printf("%v %v", trap.Source, trap.Pdu)
	} else {
		l.logger.Printf("%v %v", trap.Source, trap.Error)
	}
}

func NewTrapListener() *TrapListener {
	return &TrapListener{logger: log.New(os.Stdout, "", log.LstdFlags)}
}

func main() {
	server, err := snmpgo.NewTrapServer(snmpgo.ServerArguments{
		LocalAddr: "0.0.0.0:4445",
	})
	if err != nil {
		log.Fatal(err)
	}

	// V2c
	err = server.AddSecurity(&snmpgo.SecurityEntry{
		Version:   snmpgo.V2c,
		Community: "public",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve(NewTrapListener())
	if err != nil {
		log.Fatal(err)
	}
}
