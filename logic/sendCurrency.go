package logic

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"fmt"
	"github.com/amstee/blockchain/classes"
	"github.com/amstee/blockchain/models"
	"github.com/amstee/blockchain/utils"
)

func SendCurrency(db *gorm.DB, args []string) {
	blockchain := classes.GetBlockChainFromGenesis(db)
	wallets := models.GetWallets(db)
	from := args[0]
	to := args[1]
	amount, err := strconv.Atoi(args[2]); if err != nil {
		fmt.Println("Invalid amount")
	}
	txs := classes.NewTransaction(from, to, amount, blockchain, wallets)
	reward := classes.NewCoinBaseTX(from, utils.RandStringBytesRmndr(16))
	blockchain.MineBlock([]*models.TransactionModel{reward, txs})

	fmt.Printf("Block mined by sender, reward = %d\n", reward.Vout[0].Value)
}