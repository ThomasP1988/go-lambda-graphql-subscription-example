package config

import "os"

type Table int
type SecondaryIndex int

type SecondaryIndexes = map[SecondaryIndex]string

type TableConfig struct {
	Name           string
	SecondaryIndex SecondaryIndexes
}
type TableRegistry map[Table]TableConfig

const (
	USER Table = iota
	GRAPHQL_SUB_CONNECTION
	GRAPHQL_SUB_SUBSCRIPTION
	GRAPHQL_SUB_EVENT
)

const (
	EmailIndex SecondaryIndex = iota
	OperationIndex
)

func GetTables(env *Stage) TableRegistry {

	SetEnv(env)

	if env == nil {
		envS := Stage(os.Getenv("env"))
		env = &envS
	}

	return TableRegistry{
		USER: TableConfig{
			Name: string(*env) + "-user",
			SecondaryIndex: SecondaryIndexes{
				EmailIndex: "email-index",
			},
		},
		GRAPHQL_SUB_CONNECTION: TableConfig{
			Name:           string(*env) + "-gql-sub-connections",
			SecondaryIndex: SecondaryIndexes{},
		},
		GRAPHQL_SUB_SUBSCRIPTION: TableConfig{
			Name: string(*env) + "-gql-sub-subscriptions",
			SecondaryIndex: SecondaryIndexes{
				OperationIndex: "operation-index",
			},
		},
		GRAPHQL_SUB_EVENT: TableConfig{
			Name:           string(*env) + "-gql-sub-events",
			SecondaryIndex: SecondaryIndexes{},
		},
	}

}
