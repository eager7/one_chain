package btc

import (
	"encoding/hex"
	"fmt"
	"github.com/gcash/bchd/bchec"
	"github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip32"
	"testing"
)

const (
	url  = "39.108.13.219:8332"
	usr  = "btc"
	pass = "blcokchain"
)

func TestDecodeTx(t *testing.T) {
	btc, err := Initialize(url, usr, pass)
	if err != nil {
		t.Fatal(err)
	}
	raw := `0200000007648b3ed7adeb993885858b6adf333fc3478a1cbcc0f3a8c7558a408630a557ee010000006a4730440220284eea93d650c86f03a9327e1f7e9753cde03624c7a88799acd8837240c8d21e022023cac444acc8408c9f789d7cc462be2f30a8ff766b912e42edc0ace79ee0aa1f01210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff93d70d6455c952bebe32fbff8dc2faec82c36572abf6e76eed7760afeb576ec9000000006a47304402205b937b65d3be195229d02902d6bd34e02b837038565d902f24ae4706543a8fe702205d873010e78fe7db8e864563dc3cd6cbb4fb98e3cb96155b185d95b879a421fc01210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff1078e638f30aac29f5ceccc878ac990e5aec207da957d07748bb90c3eeaca0ae000000006b483045022100fe0d65ee265cf58e471c91668b8cb3665c9f55024e471de38c953391fc40501f02204b4af17ac9a19f6486448f061421abec53b690a984891537d6e295312173481801210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff9fdf4fce5380ed68cdc9de83cafb0c6b4845183e921029a370f2b76f167cb80b000000006b4830450221009a28c64fc8208df08db3b6fb3892f3feb0b0ec514b59136f00c4706ec1fddbf0022055899d5056dcddfb24235902aaaa03fe351e45346995a95645736cfee503fe0001210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff3ca6efdaa863407b2849117922e4e53956561832b634c24d2cf4ed408907312a000000006b483045022100dfea5d1e1e9f5054dfff7c33669bb1f82f6ac1e8a6873c04a4fbdfa115ed4f9d0220282b66e491efb19ad80f7d8648012c4fcd6fb9d8fa18c8bf3d413dde622a18f901210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff0424e6febb0fb748cdc64ebf2b770f94d411a791da85e06186524c3761335266000000006a4730440220531c16fd81512741d9b5c9d43f61e003039f713ac65b58617cdcde72bd3c53ff02201dd9a43fd778634d31dd0d311afa42a20bd75734d1816b7c9a7dc545efd1489101210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff4b8102bb10c1716eac9ca68832e837122e87c01a023b1f2bf6e2af95da22d149000000006a47304402206b2bc09991862fc607a781ae8bb54700864a6e32fa650a572a51f2a9d3000aa502206d565ab3c4407e2d6f80c5d96acf87364ece5a7afe9739354bc473c0f2a86a3b01210339d5b09fd17ff969c9e40c4546c616528c1ec3f36d7f079561590f4c98190660ffffffff03eb080000000000001976a914cdc19923868f659bb0d435dc2bb8a6dc6a934b9d88ac20020000000000001976a914f3dea51c73d75ece8e0c89198cd8f6a139fba34688ac0000000000000000536a4c5002000467696674010001010000300e6764cb238e570c504265a095d1ea528a6cde53c8d371b1366548f848075e02042d9efdadaf833d5801120d8cbfb36601000100010004000000000100040000000100000000`
	fmt.Println(btc.CheckTransactionStandard(raw))
}

func TestSeed(t *testing.T) {
	seed := "9964c397116eeb285c86b73f0a5eee5f87f54897"
	by, err := hex.DecodeString(seed)
	if err != nil {
		t.Fatal(err)
	}
	key, err := bip32.NewMasterKey(by)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key.B58Serialize())
}

func TestWallet(t *testing.T) {
	param := chaincfg.MainNetParams
	seed := "9804342764c9d136146854eb0e19919f637a3c9e"
	by, err := hex.DecodeString(seed)
	if err != nil {
		t.Fatal(err)
	}
	wallet, err := NewFromSeed(by)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("m/44'/%d'/0'/0/%d",param.HDCoinType, i)
		pri, err := wallet.derivePrivateKey(hdwallet.MustParseDerivationPath(path))
		if err != nil {
			t.Fatal(err)
		}
		wif, err := bchutil.NewWIF((*bchec.PrivateKey)(pri), &param, true)
		if err != nil {
			t.Fatal(err)
		}
		pubKeyHash := bchutil.Hash160(wif.PrivKey.PubKey().SerializeCompressed())
		addressPubKeyHash, err := bchutil.NewAddressPubKeyHash(pubKeyHash, &param)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(param.HDCoinType, i, path, addressPubKeyHash.String(), hex.EncodeToString(wif.SerializePubKey()), wif.String())
	}
}
