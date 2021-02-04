import model

import (
	"errors"
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(accounnt *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base `valid:"required"`
	Kind string `json:"kind" valid:"notnull"`
	Key string `json:"key" valid:"notnull"`
	AccountID string `json:"account_id" valid:"notnull"`
	Accout *Account `valid:"-"`
	Status string `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New(text: "invalid type of key")
	}
	
	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New(text: "invalid status")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey {
		Kind: kind,
		Account: account,
		Key: key
		Status: "active"
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}