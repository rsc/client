// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/stellar1/gregor.avdl

package stellar1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type PaymentStatusMsg struct {
	KbTxID KeybaseTransactionID `codec:"kbTxID" json:"kbTxID"`
	TxID   TransactionID        `codec:"txID" json:"txID"`
}

func (o PaymentStatusMsg) DeepCopy() PaymentStatusMsg {
	return PaymentStatusMsg{
		KbTxID: o.KbTxID.DeepCopy(),
		TxID:   o.TxID.DeepCopy(),
	}
}

type GregorInterface interface {
}

func GregorProtocol(i GregorInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "stellar.1.gregor",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type GregorClient struct {
	Cli rpc.GenericClient
}
