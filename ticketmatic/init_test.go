package ticketmatic

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	accountcode := os.Getenv("TM_TEST_ACCOUNTCODE")
	accesskey := os.Getenv("TM_TEST_ACCESSKEY")
	secretkey := os.Getenv("TM_TEST_SECRETKEY")

	if accountcode == "" || accesskey == "" || secretkey == "" {
		log.Println("No test variables found. Make sure you set the")
		log.Println("TM_TEST_ACCOUNTCODE, TM_TEST_ACCESSKEY and TM_TEST_SECRETKEY")
		log.Println("environment variables.")
		os.Exit(1)
	}

	server = "https://qa.ticketmatic.com/api"

	os.Exit(m.Run())
}
