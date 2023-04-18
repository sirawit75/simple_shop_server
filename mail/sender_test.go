package mail

import (
	"sirawit/shop/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	config, err := config.LoadWorkerConfig("../cmd/worker")
	assert.NoError(t, err)
	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	subject := "test"
	content := `
		<h1>Hello from Sirawit</h1>
		<p>This message from test server</p>
	`
	to := []string{"hibrari_boss@hotmail.com"}
	err = sender.SendEmail(subject, content, to)
	assert.NoError(t, err)
}
