// Copyright (c) 2020 Pawel Maslanka
// Contact: pawmas.pawelmaslanka@gmail.com
// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.

package main

import (
	"log"
	"os"
	"os/signal"
)

func watchSignal(done chan struct{}) {
	ch := make(chan os.Signal, 1)
	defer close(ch)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Printf("Interrupt signal...")
	close(done)
}

func main() {
	// flag.Parse()
	go HandleMgmtRequest()

	done := make(chan struct{})
	go watchSignal(done)
	<-done
}
