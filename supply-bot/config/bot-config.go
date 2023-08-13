package config

import "github.com/spf13/viper"

type BotConfig struct {
	RPC             string
	Contract        string
	PrivateKeyOwner string
	PublicKeyOwner  string
}

func ReadBotConfig() BotConfig {
	return BotConfig{
		RPC:             viper.GetString("watcher.rpc"),
		Contract:        viper.GetString("watcher.contract"),
		PrivateKeyOwner: viper.GetString("watcher.private_key_owner"),
		PublicKeyOwner:  viper.GetString("watcher.public_key_owner"),
	}
}
