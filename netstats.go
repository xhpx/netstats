/*
 *=============================================================================
 *
 *		File Name:	netstats.go
 *
 *	  Description:	statistic network packets to a file or output.
 *	                card.
 *
 *		  Version:	1.0
 *		  Created:  18/2/2019
 *		 Compiler:  go build
 *
 *    	   Author:  XuHongping
 *    	   E-Mail:  xuhongping108@gmail.com
 *   	  Company:  IDSS
 *=============================================================================
 */
package main

import (
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

/*
type CmdSt struct {
	config string
	dev    string
}

var cmdSt CmdSt
*/

var (
	handle     *pcap.Handle
	configfile               = flag.String("c", "/etc/netstat.conf", "Configure file name to read from")
	device                   = flag.String("i", "eth0", "Interface to read packets from .")
	snaplen                  = flag.Int("s", 65536, "Snap length, number of bytes max to read per packet")
	promisc                  = flag.Bool("p", true, " Set interface mode")
	timeout    time.Duration = 30 * time.Second
	err        error
)

/*
* This struct stores  count of statistics
 */
type pc_st struct {
	icmp  int
	tcp   int
	udp   int
	http  int
	recv  int
	send  int
	total int
}

var pc pc_st

func usage() {

	flag.Parse()

	fmt.Println("c", *configfile)
	fmt.Println("i", *device)
}

func decodeIPv4Pkt(pkt gopacket.Packet) {
	if tcp := pkt.Layer(layers.LayerTypeTCP); tcp != nil {
		//to do

	}

	if udp := pkt.Layer(layers.LayerTypeUDP); udp != nil {
		//to do
	}

}

func decodeIPv6Pkt(pkt gopacket.Packet) {
	/*reserve*/
}
func processPacket(pkt gopacket.Packet) {

	if ethernetLayer := pkt.Layer(layers.LayerTypeEthernet); ethernetLayer == nil {
		fmt.Println("Not ethernet packet!")
		return
	}

	if ipv4 := pkt.Layer(layers.LayerTypeIPv4); ipv4 != nil {
		decodeIPv4Pkt(pkt)
	} else if ipv6 := pkt.Layer(layers.LayerTypeIPv6); ipv6 != nil {
		decodeIPv6Pkt(pkt)
	}

	if err := pkt.ErrorLayer(); err != nil {
		fmt.Println("Error decoding  packet:", err)
	}
}

func main() {
	usage()

	inact, err := pcap.NewInactiveHandle(*device)

	if err != nil {
		log.Fatal(err)
	}

	if err = inact.SetSnapLen(*snaplen); err != nil {
		log.Fatal("Can't set snap length :%v", err)
	} else if err = inact.SetPromisc(*promisc); err != nil {
		log.Fatal("can't set promisc mode :%v", err)
	} else if err = inact.SetTimeout(timeout); err != nil {
		log.Fatal("can't set timeout :%v", err)
	}

	if handle, err = inact.Activate(); err != nil {
		log.Fatal("PCAP Activate error:", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		fmt.Printf("%T", packet)
		processPacket(packet)
	}

}
