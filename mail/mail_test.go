package mail

import "testing"

func TestSendEmail(t *testing.T) {
	p := SendMailParams{
		From:       "Vito <17863130159@163.com>",
		To:         []string{"732185332@qq.com"},
		BCC:        nil,
		CC:         nil,
		Subject:    "test",
		Addr:       "smtp.163.com:465",
		Username:   "17863130159@163.com",
		Password:   "ZZSGXMWJZQOPGAZP",
		Host:       "smtp.163.com",
		ServerName: "smtp.163.com",
	}
	err := SendEmail(p, "test")
	if err != nil {
		t.Logf("error %s", err)
	}
}
