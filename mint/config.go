package mint

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/elnosh/gonuts/cashu/nuts/nut06"
)

type Config struct {
	PrivateKey     string
	DerivationPath string
}

func GetConfig() Config {
	return Config{
		PrivateKey:     os.Getenv("MINT_PRIVATE_KEY"),
		DerivationPath: os.Getenv("MINT_DERIVATION_PATH"),
	}
}

func getMintInfo() (*nut06.MintInfo, error) {
	mintInfo := nut06.MintInfo{
		Name:        os.Getenv("MINT_NAME"),
		Version:     "gonuts/0.0.1",
		Description: os.Getenv("MINT_DESCRIPTION"),
	}

	mintInfo.LongDescription = os.Getenv("MINT_DESCRIPTION_LONG")
	mintInfo.Motd = os.Getenv("MINT_MOTD")

	privateKey := secp256k1.PrivKeyFromBytes([]byte(os.Getenv("MINT_PRIVATE_KEY")))
	mintInfo.Pubkey = hex.EncodeToString(privateKey.PubKey().SerializeCompressed())

	contact := os.Getenv("MINT_CONTACT_INFO")
	var mintContactInfo [][]string
	if len(contact) > 0 {
		err := json.Unmarshal([]byte(contact), &mintContactInfo)
		if err != nil {
			return nil, fmt.Errorf("error parsing contact info: %v", err)
		}
	}
	mintInfo.Contact = mintContactInfo

	nuts := nut06.NutsMap{
		4: nut06.NutSetting{
			Methods: []nut06.MethodSetting{
				{Method: "bolt11", Unit: "sat"},
			},
			Disabled: false,
		},
		5: nut06.NutSetting{
			Methods: []nut06.MethodSetting{
				{Method: "bolt11", Unit: "sat"},
			},
			Disabled: false,
		},
		7:  map[string]bool{"supported": false},
		8:  map[string]bool{"supported": false},
		9:  map[string]bool{"supported": false},
		10: map[string]bool{"supported": false},
		11: map[string]bool{"supported": false},
		12: map[string]bool{"supported": false},
	}

	mintInfo.Nuts = nuts
	return &mintInfo, nil
}
