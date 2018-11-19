package prototype

import (
	"github.com/pkg/errors"
)

func (a *AccountCreateOperation) GetAuthorities(auths *[]Authority) {

}
func (a *AccountCreateOperation) GetRequiredPosting(auths *map[string]bool) {

}

func (a *AccountCreateOperation) GetRequiredOwner(auths *map[string]bool) {

}
func (a *AccountCreateOperation) GetAdmin(*[]AccountAdminPair) {

}
func (a *AccountCreateOperation) IsVirtual() {

}


func (a *AccountCreateOperation) GetRequiredActive(auths *map[string]bool) {
	(*auths)[a.Creator.Value] = true
}


func (a *AccountCreateOperation) Validate() error {

	if a == nil{
		return ErrNpe
	}

	if err := a.Creator.Validate(); err != nil{
		return errors.WithMessage(err, "Creator error" )

	}

	if err := a.NewAccountName.Validate();err != nil{
		return errors.WithMessage(err, "NewAccountName error" )
	}

	if a.MemoKey == nil {
		return errors.New("Memo key cant be empty")
	}
	if err := a.MemoKey.Validate(); err != nil {
		return errors.WithMessage(err, "Memo key error" )
	}

	if a.Posting == nil {
		return errors.New("Posting Key cant be empty")
	}

	if err := a.Posting.Validate(); err != nil {
		return errors.WithMessage(err, "Posting error" )
	}

	if a.Active == nil {
		return errors.New("Posting Key cant be empty")
	}
	if err := a.Active.Validate(); err != nil {
		return errors.WithMessage(err, "Active error" )
	}
	if a.Owner == nil {
		return errors.New("Posting Key cant be empty")
	}
	if err := a.Owner.Validate(); err != nil {
		return errors.WithMessage(err, "Owner error" )
	}

	if a.Fee == nil || a.Fee.Value == 0 {
		return errors.New("Account Create must set Fee")
	}

	return nil
}