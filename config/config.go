package config

type Config struct {
	OutDir string
	Port   string

	Mail bool

	DestMail string
	SendMail string
	Smtp     string
	SmtpName string
	SmtpPwd  string

	Timeout int
}

// Constructor for the config struct
func NewConfig(
	outDir string,
	port string,
	mail bool,
	destMail string,
	sendMail string,
	smtp string,
	smtpName string,
	smtpPwd string,
	timeout int,
) *Config {
	c := new(Config)
	c.OutDir = outDir
	c.Port = port
	c.Mail = mail
	c.DestMail = destMail
	c.SendMail = sendMail
	c.Smtp = smtp
	c.SmtpName = smtpName
	c.SmtpPwd = smtpPwd
	c.Timeout = timeout
	return c
}
