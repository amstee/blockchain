package config

type BlockchainConf struct {
	Version byte
	CheckSumLen int
}

var BlockchainConfig = new(BlockchainConf)