package account

import (
	"errors"

	"github.com/asdine/storm"
	"github.com/google/uuid"
)

type AccountAttributes struct {
	Country               string                `json:"country"`
	BaseCurrency          string                `json:"base_currency"`
	AccountNumber         string                `json:"account_number"`
	BankID                string                `json:"bank_id"`
	BankIDCode            string                `json:"bank_id_code"`
	Bic                   string                `json:"bic"`
	IBAN                  string                `json:"iban"`
	Title                 string                `json:"title"`
	Name                  string                `json:"name"`
	AccountClassification AccountClassification `json:"account_classification"`
	JointAccount          bool                  `json:"joint_account"`
	Status                string                `json:"status"`
}

type Account struct {
	Type           string            `json:"type"`
	ID             uuid.UUID         `json:"id" storm:"id"`
	OrganisationID uuid.UUID         `json:"organisation_id,omitempty"`
	Version        int               `json:"version,omitempty"`
	Attributes     AccountAttributes `json:"attributes"`
}

type AccountClassification string

const (
	dbPath                        string                = "account.db"
	AccountBusinessClassification AccountClassification = "Business"
	AccountPersonalClassification AccountClassification = "Personal"
)

// Errors
var (
	ErrorRecordInvalid = errors.New("account record is invalid")
)

// All return all accounts from the database
func All(pageNumber, pageSize int) ([]Account, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	accounts := []Account{}
	err = db.All(&accounts, storm.Limit(pageSize), storm.Skip(pageNumber*pageSize))
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Find an account
func Find(id uuid.UUID) (*Account, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	account := new(Account)
	err = db.One("ID", id, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// Delete an account from database
func Delete(id uuid.UUID) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	account := new(Account)
	err = db.One("ID", id, account)
	if err != nil {
		return err
	}
	return db.DeleteStruct(account)
}

// Save updates or creates a given record to database
func (account *Account) Save() error {
	if err := account.validate(); err != nil {
		return err
	}

	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(account)
}

func (account *Account) validate() error {
	if account.Type == "" {
		return ErrorRecordInvalid
	}
	return nil
}
