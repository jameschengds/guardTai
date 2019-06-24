package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go/token"
	"github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("eos-cilent")
)

type Eos struct {
	EOSClint *eos.API
}

func EOSCilentInit(url string) (e *Eos) {
	api := eos.New(url)
	return &Eos{
		EOSClint: api,
	}
}

func (e *Eos) PushTX(userAccount, accountTo, userPk, memo string) (string, error) {
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(userPk)
	if err != nil {
		logger.Errorf("import private key: %s", err)
		return "", nil
	}
	e.EOSClint.SetSigner(keyBag)
	from := eos.AccountName(userAccount)
	to := eos.AccountName(accountTo)
	quantity, err := eos.NewEOSAssetFromString("1.0000 EOS")
	if err != nil {
		logger.Errorf("invalid quantity: %s", err)
		return "", nil
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(e.EOSClint); err != nil {
		logger.Errorf("filling tx opts: %s", err)
		return "", nil
	}

	tx := eos.NewTransaction([]*eos.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)
	aaa := tx.TransactionHeader.Expiration.Time.Unix()
	logger.Debugf("%+v", aaa)
	signedTx, packedTx, err := e.EOSClint.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		logger.Errorf("sign transaction: %s", err)
		return "", nil
	}
	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		logger.Errorf("json marshalling transaction: %s", err)
		return "", nil
	}
	packetContent, err := json.MarshalIndent(packedTx, "", "  ")
	if err != nil {
		logger.Errorf("json marshalling transaction: %s", err)
		return "", nil
	}

	logger.Debug(string(content))
	logger.Debug(string(packetContent))

	response, err := e.EOSClint.PushTransaction(packedTx)
	if err != nil {
		logger.Errorf("push transaction: %s", err)
		return "", nil
	}

	return hex.EncodeToString(response.Processed.ID), err
}

func (e *Eos) NewAccount(creator, newAccount eos.AccountName, pubKey ecc.PublicKey, buyRAMAmount, cpuStake, netStake eos.Asset, doTransfer bool) {
	actions := []*eos.Action{}
	actions = append(actions, system.NewNewAccount(creator, newAccount, pubKey))
	actions = append(actions, system.NewDelegateBW(creator, newAccount, cpuStake, netStake, doTransfer))
	actions = append(actions, system.NewBuyRAM(creator, newAccount, uint64(buyRAMAmount.Amount)))
	resp, err := e.EOSClint.SignPushActions(actions...)
	if err == nil {
		data, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}
