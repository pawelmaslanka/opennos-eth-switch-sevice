// Copyright (c) 2020 Pawel Maslanka
// Contact: pawmas.pawelmaslanka@gmail.com
// This file is subject to the terms and conditions defined in
// file 'LICENSE.txt', which is part of this source code package.

package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	log "github.com/golang/glog"

	"google.golang.org/grpc"

	mgmt "opennos-eth-switch-service/mgmt"
	serv_param "opennos-eth-switch-service/serv-param"
)

type ethSwitchMgmtServerT struct {
	// TODO: Change it to stream server
	mgmt.UnimplementedEthSwitchMgmtServer
}

func newEthSwitchMgmtServerT() *ethSwitchMgmtServerT {
	return &ethSwitchMgmtServerT{}
}

///////////////////////////////////////// Port breakout /////////////////////////////////////////
func (srv *ethSwitchMgmtServerT) SetPortBreakout(ctx context.Context, req *mgmt.PortBreakoutRequest) (*mgmt.Result, error) {
	ethIntf := req.GetEthIntf()
	if ethIntf == nil || (len(ethIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface")
	}

	chanSpeed := req.GetChannelSpeed()
	if chanSpeed == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No channel speed")
	}

	log.Infof("Set number of channels %s with channel speed %s on interface %s",
		req.GetNumChannels().String(), chanSpeed.GetMode().String(), ethIntf.Ifname)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

///////////////////////////////////////// Aggregate /////////////////////////////////////////
func (srv *ethSwitchMgmtServerT) CreateAggregateIntf(ctx context.Context, req *mgmt.CreateAggregateIntfRequest) (*mgmt.Result, error) {
	aggIntf := req.GetAggIntf()
	if aggIntf == nil || (len(aggIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interfaces")
	}

	log.Infof("Create aggregate interface %s", aggIntf.Ifname)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) DeleteAggregateIntf(ctx context.Context, req *mgmt.DeleteAggregateIntfRequest) (*mgmt.Result, error) {
	aggIntf := req.GetAggIntf()
	if aggIntf == nil || (len(aggIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interfaces")
	}

	log.Infof("Delete aggregate interface %s", aggIntf.Ifname)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) AddEthernetIntfToAggregateIntf(ctx context.Context, req *mgmt.AddEthernetIntfToAggregateIntfRequest) (*mgmt.Result, error) {
	aggIntf := req.GetAggIntf()
	if aggIntf == nil || (len(aggIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interface")
	}

	ethIntfs := req.GetEthIntfs()
	if ethIntfs == nil || (len(ethIntfs) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interfaces")
	}

	strBuilder := strings.Builder{}
	for i, ifname := range ethIntfs {
		if len(ifname.GetIfname()) > 0 {
			str := fmt.Sprintf(" %s", ifname.Ifname)
			if _, err := strBuilder.WriteString(str); err != nil {
				return &mgmt.Result{Status: mgmt.Result_FAILED}, err
			}
		} else {
			return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface at idx %d", i)
		}
	}

	log.Infof("Add %d Ethernet interfaces to aggregate interface %s:%s",
		len(ethIntfs), aggIntf.Ifname, strBuilder.String())
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) RemoveEthernetIntfFromAggregateIntf(ctx context.Context, req *mgmt.RemoveEthernetIntfFromAggregateIntfRequest) (*mgmt.Result, error) {
	aggIntf := req.GetAggIntf()
	if aggIntf == nil || (len(aggIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interface")
	}

	ethIntfs := req.GetEthIntfs()
	if ethIntfs == nil || (len(ethIntfs) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interfaces")
	}

	strBuilder := strings.Builder{}
	for i, ifname := range ethIntfs {
		if len(ifname.GetIfname()) > 0 {
			str := fmt.Sprintf(" %s", ifname.GetIfname())
			if _, err := strBuilder.WriteString(str); err != nil {
				return &mgmt.Result{Status: mgmt.Result_FAILED}, err
			}
		} else {
			return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface at idx %d", i)
		}
	}

	log.Infof("Remove %d Ethernet interfaces from aggregate interface %s:%s",
		len(ethIntfs), aggIntf.Ifname, strBuilder.String())
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

///////////////////////////////////////// Ethernet /////////////////////////////////////////
func (srv *ethSwitchMgmtServerT) AddIpv4AddrToEthernetIntf(ctx context.Context, req *mgmt.AddIpv4AddrToEthernetIntfRequest) (*mgmt.Result, error) {
	ethIntf := req.GetEthIntf()
	if ethIntf == nil || (len(ethIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface")
	}

	addr := req.GetAddr()
	if addr == nil || (len(addr.GetIp()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No IPv4 address")
	}

	log.Infof("Add IPv4 address %s/%d to Ethernet interface %s", addr.Ip, addr.PrfxLen, ethIntf.Ifname)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) RemoveIpv4AddrFromEthernetIntfRequest(ctx context.Context, req *mgmt.RemoveIpv4AddrFromEthernetIntfRequest) (*mgmt.Result, error) {
	ethIntf := req.GetEthIntf()
	if ethIntf == nil || (len(ethIntf.GetIfname()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface")
	}

	addr := req.GetAddr()
	if addr == nil || (len(addr.GetIp()) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No IPv4 address")
	}

	log.Infof("Remove IPv4 address %s/%d from Ethernet interface %s", addr.Ip, addr.PrfxLen, ethIntf.Ifname)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

///////////////////////////////////////// VLAN /////////////////////////////////////////
func (srv *ethSwitchMgmtServerT) CreateVlan(ctx context.Context, req *mgmt.CreateVlanRequest) (*mgmt.Result, error) {
	vlan := req.GetVlan()
	if vlan == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No VLAN")
	}

	log.Infof("Create VLAN %d", vlan.Vid)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) DeleteVlan(ctx context.Context, req *mgmt.DeleteVlanRequest) (*mgmt.Result, error) {
	vlan := req.GetVlan()
	if vlan == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No VLAN")
	}

	log.Infof("Create VLAN %d", vlan.Vid)
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) AddAggregateIntfToVlan(ctx context.Context, req *mgmt.AddAggregateIntfToVlanRequest) (*mgmt.Result, error) {
	vlan := req.GetVlan()
	if vlan == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No VLAN")
	}

	aggIntfs := req.GetAggIntfs()
	if aggIntfs == nil || (len(aggIntfs) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interfaces")
	}

	strBuilder := strings.Builder{}
	for i, ifname := range aggIntfs {
		if len(ifname.GetIfname()) > 0 {
			str := fmt.Sprintf(" %s", ifname.Ifname)
			if _, err := strBuilder.WriteString(str); err != nil {
				return &mgmt.Result{Status: mgmt.Result_FAILED}, err
			}
		} else {
			return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interface at idx %d", i)
		}
	}

	log.Infof("Add %d aggregate interfaces to %s VLAN %d:%s",
		len(aggIntfs), vlan.Mode.String(), vlan.Vid, strBuilder.String())
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) RemoveAggregateIntfFromVlan(ctx context.Context, req *mgmt.RemoveAggregateIntfFromVlanRequest) (*mgmt.Result, error) {
	vlan := req.GetVlan()
	if vlan == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No VLAN")
	}

	aggIntfs := req.GetAggIntfs()
	if aggIntfs == nil || (len(aggIntfs) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interfaces")
	}

	strBuilder := strings.Builder{}
	for i, ifname := range aggIntfs {
		if len(ifname.GetIfname()) > 0 {
			str := fmt.Sprintf(" %s", ifname.Ifname)
			if _, err := strBuilder.WriteString(str); err != nil {
				return &mgmt.Result{Status: mgmt.Result_FAILED}, err
			}
		} else {
			return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No aggregate interface at idx %d", i)
		}
	}

	log.Infof("Remove %d aggregate interfaces from %s VLAN %d:%s",
		len(aggIntfs), vlan.Mode.String(), vlan.Vid, strBuilder.String())
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) AddEthernetIntfToVlan(ctx context.Context, req *mgmt.AddEthernetIntfToVlanRequest) (*mgmt.Result, error) {
	vlan := req.GetVlan()
	if vlan == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No VLAN")
	}

	ethIntfs := req.GetEthIntfs()
	if ethIntfs == nil || (len(ethIntfs) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interfaces")
	}

	strBuilder := strings.Builder{}
	for i, ifname := range ethIntfs {
		if len(ifname.GetIfname()) > 0 {
			str := fmt.Sprintf(" %s", ifname.GetIfname())
			if _, err := strBuilder.WriteString(str); err != nil {
				return &mgmt.Result{Status: mgmt.Result_FAILED}, err
			}
		} else {
			return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface at idx %d", i)
		}
	}

	log.Infof("Add %d Ethernet interfaces to %s VLAN %d:%s",
		len(ethIntfs), vlan.Mode.String(), vlan.Vid, strBuilder.String())
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

func (srv *ethSwitchMgmtServerT) RemoveEthernetIntfFromVlan(ctx context.Context, req *mgmt.RemoveEthernetIntfFromVlanRequest) (*mgmt.Result, error) {
	vlan := req.GetVlan()
	if vlan == nil {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No VLAN")
	}

	ethIntfs := req.GetEthIntfs()
	if ethIntfs == nil || (len(ethIntfs) == 0) {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interfaces")
	}

	strBuilder := strings.Builder{}
	for i, ifname := range ethIntfs {
		if len(ifname.GetIfname()) > 0 {
			str := fmt.Sprintf(" %s", ifname.Ifname)
			if _, err := strBuilder.WriteString(str); err != nil {
				return &mgmt.Result{Status: mgmt.Result_FAILED}, err
			}
		} else {
			return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("No Ethernet interface at idx %d", i)
		}
	}

	log.Infof("Remove %d Ethernet interfaces from %s VLAN %d:%s",
		len(ethIntfs), vlan.Mode.String(), vlan.Vid, strBuilder.String())
	if false {
		return &mgmt.Result{Status: mgmt.Result_FAILED}, fmt.Errorf("Unknown")
	}

	return &mgmt.Result{Status: mgmt.Result_SUCCESS}, nil
}

// HandleMgmtRequest handles all gRPC requests
func HandleMgmtRequest() {
	localhost := fmt.Sprintf(":%d", serv_param.MgmtListeningTcpPortC)
	lis, err := net.Listen(serv_param.MgmtCommProtoC, localhost)
	if err != nil {
		log.Fatalf("Failed to start listening on TCP port %d for management request: %v",
			serv_param.MgmtListeningTcpPortC, err)
	}
	s := grpc.NewServer()
	mgmt.RegisterEthSwitchMgmtServer(s, newEthSwitchMgmtServerT())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve ethernet switch management request: %v", err)
	}
}
