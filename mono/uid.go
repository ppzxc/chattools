package mono

import (
	"github.com/aohorodnyk/uid"
	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"
	"math/rand"
)

var mono monoton.Monoton
var generator = uid.NewProviderCustom(8, uid.NewRand(), uid.NewEncoderBase62())

func GetUuid() string {
	return generator.MustGenerate().String()
}

func GetIntUid() int {
	return rand.Intn(99999999-10000000) + 10000000
}

func Init() {
	node := uint64(1)
	initialTime := uint64(1577865600000) // set 2020-01-01 PST as initial time
	m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
	if err != nil {
		panic(err)
	}
	mono = m
}

func GetMONOID() string {
	return mono.Next()
}