package settings

type DefinitionType string

const (
	AwsRds DefinitionType = "aws-rds"
)

// AWSRDS configuration structure
type AWSRDS struct {
	Host     string
	User     string
	Password string
	DBName   string
}
