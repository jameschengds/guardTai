package blockchain

import "github.com/eoscanada/eos-go"

type Eos struct {
	EOSClint *eos.API
}

func EOSCilentInit(url string) (e *Eos) {
	api := eos.New(url)
	return &Eos{
		EOSClint: api,
	}
}
