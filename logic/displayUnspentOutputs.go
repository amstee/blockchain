package logic

import (
	"github.com/jinzhu/gorm"
	"github.com/amstee/blockchain/models"
	"log"
)

func DisplayUnspentOutputs(odb *gorm.DB) {
	var unspent []models.TXOutput

	err := odb.Find(&unspent).Error; if err != nil {
		log.Fatalf("An error occured %s ", err)
	}
	for _, out := range unspent {
		out.Display()
	}
}
