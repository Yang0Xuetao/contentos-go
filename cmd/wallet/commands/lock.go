package commands

import (
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/cmd/wallet/wallet"
)

var LockCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock",
		Short: "lock a account",
		Run:   lock,
	}
	return cmd
}

func lock(cmd *cobra.Command, args []string) {
	o := cmd.Context["wallet"]
	w := o.(wallet.Wallet)
	name := args[0]
	err := w.Lock(name)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v\n", err))
		return
	}
	fmt.Println(fmt.Sprintf("unlock account %s success", name))
}