module github.com/makerdao/vdb-mcd-transformers

go 1.15

require (
	github.com/AlecAivazis/survey/v2 v2.1.1
	github.com/BurntSushi/toml v0.3.1
	github.com/dave/jennifer v1.3.0
	github.com/ethereum/go-ethereum v1.9.22
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.0.0
	github.com/makerdao/vdb-transformer-utilities v0.0.0-20201021185605-2af87288ff1b
	github.com/makerdao/vulcanizedb v0.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.7.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)

replace gopkg.in/urfave/cli.v1 => gopkg.in/urfave/cli.v1 v1.20.0

replace github.com/ethereum/go-ethereum => github.com/makerdao/go-ethereum v1.9.21-rc1
