package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	// ModuleName defines the module name
	ModuleName = "escrow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	// GlobalStoreKeyPrefix is a prefix for versioning store
	GlobalStoreKeyPrefix = []byte{0x00}

	ReceiverStoreKeyPrefix = append(GlobalStoreKeyPrefix, []byte("Receiver")...)

	DayReceiverAmountStoreKeyPrefix = append(GlobalStoreKeyPrefix, []byte("DayCoinReceiverAmount")...)

	DayReceiverPaidStoreKeyPrefix = append(GlobalStoreKeyPrefix, []byte("DayReceiverPaid")...)
)

func GetEscrowAddress() sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte("Escrow")))
}



func ReceiverStoreKey(receiver string)[]byte{
	ret := ReceiverStoreKeyPrefix
	ret = append(ret, []byte(receiver)...)
	return ret
}

// Generate key for each block+coin+receiver => total amount in store
func DayReceiverAmountStoreKey(dayId int64, receiver string) []byte {
	ret := DayReceiverAmountStoreKeyPrefix
	ret = append(ret, Int64ToBytes(dayId)...)
	ret = append(ret, []byte(receiver)...)
	return ret
}

func DayReceiverPaidStoreKey(dayId int64, receiver string) []byte {
	ret := DayReceiverPaidStoreKeyPrefix
	ret = append(ret, Int64ToBytes(dayId)...)
	ret = append(ret, []byte(receiver)...)
	return ret
}

func Int64ToBytes(num int64) []byte {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, num)
	return b.Bytes()
}

func BytesToInt64(b []byte) int64 {
	var num int64
	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &num); err != nil {
		Logger.Error(fmt.Sprintf("%x could not be decoded to int64", b))
	}
	return num
}

// Get days since genesis
func GetDayId(blockTime int64) int64 {
	return int64(time.Unix(blockTime, 0).Sub(time.Unix(GetGenesisBlockTime(), 0)).Hours() / 24) // int64 rounds down
}

func GetGenesisBlockTime() int64 {
	return 1585699200 // Wednesday, 1 April 2020 00:00:00 GMT
}

func GetFirstDayId() int64 {
	return GetDayId(GetGenesisBlockTime())
}
