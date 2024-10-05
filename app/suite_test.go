package app

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/response"
)

type TestSuiteRepo struct {
	suite.Suite
	testDB *TestDB

	app *App
}

func (t *TestSuiteRepo) SetupSuite() {
	log.SetOutput(os.Stderr)
	t.T().Log("setting up the suite")

	schema, err := os.ReadFile(filepath.Join("..", "schema.sql"))
	if err != nil {
		t.T().Fatal(err)
	}

	testDB, err := SetupTestPostgresql("test_formforge", string(schema))
	if err != nil {
		t.T().Fatal(err)
	}

	t.testDB = testDB

	formRepo := form.NewPgRepo(testDB.Pool)

	responseRepo := response.NewPgRepo(testDB.Pool)
	if err != nil {
		t.T().Fatal(err)
	}

	formService := form.NewService(formRepo)
	responseService := response.NewService(responseRepo)
	t.app = New(formService, responseService)
}

func (t *TestSuiteRepo) TearDownAllSuite() {
	if err := t.testDB.Shutdown(); err != nil {
		t.T().Errorf("Failed to close test DB: %v", err)
	}
}

func (t *TestSuiteRepo) SetupTest() {
}

func (t *TestSuiteRepo) BeforeTest(suiteName, testName string) {
	if err := t.testDB.Reset(); err != nil {
		t.T().Fatal(err)
	}
}

func (t *TestSuiteRepo) AfterTest(suiteName, testName string) {

}

// run the test suite
func Test_TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteRepo))
}
