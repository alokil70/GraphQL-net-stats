package graph

import (
	"fmt"
	"test/graph/model"
	"test/statgo"
)

func getSpeed(netIO *statgo.NetIOStats) *model.Iface {

    iface := &model.Iface{
        Name:  netIO.IntName,
        TxSec: fmt.Sprintf("%.2f", float64(netIO.TX)/1024),
        RxSec: fmt.Sprintf("%.2f", float64(netIO.RX)/1024),
        Tx:    fmt.Sprintf("%d", netIO.TX),
        Rx:    fmt.Sprintf("%d", netIO.RX),
    }

    return iface
}

func getInterfaces() []*model.Iface {

    var ifaces []*model.Iface

    s := statgo.NewStat().NetIOStats()

	for _, i := range s {
        // fmt.Println("RX ", i.RX)
        // fmt.Println("TX ", i.TX)
        // fmt.Println("IntName ", i.IntName)
        ifaces = append(ifaces, getSpeed(i))
	}

    return ifaces
}
