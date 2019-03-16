package userSys

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eoscanada/eos-go/token"

	"github.com/eoscanada/eos-go"

	"github.com/gin-gonic/gin"
)

type Turn struct {
	Fuck string `json:"fuck"`
}

func (h *RestHandler) Register(c *gin.Context) {

	//var userPk = "5KQwrPbwdL6PhXujxW37FSSQZ1JiwsST4cqQzDeyXtP79zkvFD3"
	var userPk = "5HwDQMp1Ljax42ZNP6ZsJUFhYER5QCGmqWbuM84znVUMn24M9ij"
	//var userPublic =
	var userAccount = "qwertasdfzxc"
	var accountTo = "qwertyasdfgz"
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(userPk)
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	h.srvcContext.BCServer.EOSClint.SetSigner(keyBag)
	from := eos.AccountName(userAccount)
	to := eos.AccountName(accountTo)
	quantity, err := eos.NewEOSAssetFromString("1.0000 EOS")
	memo := ""
	if err != nil {
		panic(fmt.Errorf("invalid quantity: %s", err))
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(h.srvcContext.BCServer.EOSClint); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	tx := eos.NewTransaction([]*eos.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)
	aaa := tx.TransactionHeader.Expiration.Time.Unix()
	logger.Debugf("%+v", aaa)
	signedTx, packedTx, err := h.srvcContext.BCServer.EOSClint.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %s", err))
	}
	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %s", err))
	}

	fmt.Println(string(content))
	fmt.Println()

	response, err := h.srvcContext.BCServer.EOSClint.PushTransaction(packedTx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %s", err))
	}

	payload := &Turn{
		Fuck: hex.EncodeToString(response.Processed.ID),
	}
	c.JSON(http.StatusOK, payload)
}
