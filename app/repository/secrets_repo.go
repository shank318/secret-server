package repository

type secretsRepo struct {
	baseRepo
}

var SecretsRepo ISecretsRepo

func init() {
	SecretsRepo = new(secretsRepo)
}
