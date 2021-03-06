package utils

import (
	"context"
	"fmt"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/prototype"
	"github.com/coschain/contentos-go/rpc/pb"
	"hash/crc32"
	"math/rand"
	"syscall"
	"time"
)


func GenerateSignedTxAndValidate2(client grpcpb.ApiServiceClient, ops []interface{}, signers *wallet.PrivAccount) (*prototype.SignedTransaction, error) {
	privKey := &prototype.PrivateKeyType{}
	pk, err := prototype.PrivateKeyFromWIF(signers.PrivKey)
	if err != nil {
		return nil, err
	}
	privKey = pk
	return GenerateSignedTxAndValidate3(client, ops, privKey)
}

func GenerateSignedTxAndValidate3(client grpcpb.ApiServiceClient, ops []interface{}, privKey *prototype.PrivateKeyType) (*prototype.SignedTransaction, error) {
	chainState, err := GetChainState(client)
	if err != nil {
		return nil, err
	}
	return GenerateSignedTxAndValidate4(chainState.Dgpo, 30, ops, privKey)
}

func GetChainState(client grpcpb.ApiServiceClient) (*grpcpb.ChainState, error) {
	req := &grpcpb.NonParamsRequest{}
	resp, err := client.GetStatisticsInfo(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp.State, nil
}

func GenerateSignedTxAndValidate4(dgp *prototype.DynamicProperties, expiration uint32, ops []interface{}, privKey *prototype.PrivateKeyType) (*prototype.SignedTransaction, error) {
	refBlockPrefix := common.TaposRefBlockPrefix(dgp.HeadBlockId.Hash)
	// occupant implement
	refBlockNum := common.TaposRefBlockNum(dgp.HeadBlockNumber)
	tx := &prototype.Transaction{RefBlockNum: refBlockNum, RefBlockPrefix: refBlockPrefix, Expiration: &prototype.TimePointSec{UtcSeconds: dgp.Time.UtcSeconds + expiration}}
	for _, op := range ops {
		tx.AddOperation(op)
	}

	signTx := prototype.SignedTransaction{Trx: tx}

	res := signTx.Sign(privKey, prototype.ChainId{Value: 0})
	signTx.Signature = &prototype.SignatureType{Sig: res}

	if err := signTx.Validate(); err != nil {
		return nil, err
	}

	return &signTx, nil
}

func GenerateUUID(content string) uint64 {
	crc32q := crc32.MakeTable(0xD5828281)
	randContent := content + string(rand.Intn(1e5))
	return uint64(time.Now().Unix()*1e9) + uint64(crc32.Checksum([]byte(randContent), crc32q))
}

func GetPassphrase(reader PasswordReader) (string, error) {
	fmt.Print("Enter passphrase > ")
	bytePassphrase, err := reader.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	passphrase := string(bytePassphrase)
	return passphrase, nil
}
