package keymethod

import (
	"encoding/hex" //16進数
	"os"

	"github.com/Fuses-Garage/UnityGo/util"
)

func GetPubKey() string {
	bytes, err := os.ReadFile("derrsapubkey.key")
	util.CheckErr(err)
	hexstr := hex.EncodeToString(bytes)
	return hexstr
}

func GetPrvKey() string {
	bytes, err := os.ReadFile("derrsaprvkey.key")
	util.CheckErr(err)
	hexstr := hex.EncodeToString(bytes)
	return hexstr
}
