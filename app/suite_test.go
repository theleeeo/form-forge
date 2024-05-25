package app

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/response"
)

const (
	dbName = "testdb"
)

type TestSuiteRepo struct {
	suite.Suite
	db       *sql.DB
	stopFunc func() error

	app *App
}

func (t *TestSuiteRepo) SetupSuite() {
	log.SetOutput(os.Stderr)
	t.T().Log("setting up the suite")

	db, stopFunc, err := GetTestMariadb(dbName)
	if err != nil {
		t.T().Fatal(err)
	}

	t.db = db
	t.stopFunc = stopFunc

	formRepo, err := form.NewMySql(nil, db)
	if err != nil {
		t.T().Fatal(err)
	}

	responseRepo, err := response.NewMySql(nil, db)
	if err != nil {
		t.T().Fatal(err)
	}

	formService := form.NewService(formRepo)
	responseService := response.NewService(responseRepo)
	t.app = New(formService, responseService)
}

func (t *TestSuiteRepo) TearDownAllSuite() {
	if t.stopFunc != nil {
		if err := t.stopFunc(); err != nil {
			t.T().Errorf("Failed to stop the container: %v", err)
		}
	}
}

func (t *TestSuiteRepo) SetupTest() {
}

func (t *TestSuiteRepo) BeforeTest(suiteName, testName string) {

}

func (t *TestSuiteRepo) AfterTest(suiteName, testName string) {
	if err := t.ClearDatabase(); err != nil {
		t.T().Errorf("Failed to clear database: %v", err)
	}
}

func (t *TestSuiteRepo) ClearDatabase() error {
	_, err := t.db.Exec("DELETE FROM answers")
	if err != nil {
		return err
	}

	_, err = t.db.Exec("DELETE FROM responses")
	if err != nil {
		return err
	}

	_, err = t.db.Exec("DELETE FROM options")
	if err != nil {
		return err
	}

	_, err = t.db.Exec("DELETE FROM questions")
	if err != nil {
		return err
	}

	_, err = t.db.Exec("DELETE FROM forms")
	if err != nil {
		return err
	}

	return nil
}

// run the test suite
func Test_TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteRepo))
}
