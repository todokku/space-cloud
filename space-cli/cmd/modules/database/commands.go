package database

import (
	"github.com/spaceuptech/space-cli/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GenerateSubCommands is the list of commands the database module exposes
func GenerateSubCommands() []*cobra.Command {

	var generaterule = &cobra.Command{
		Use:  "db-rules",
		RunE: actionGenerateDBRule,
	}

	var generateconfig = &cobra.Command{
		Use:  "db-config",
		RunE: actionGenerateDBConfig,
	}

	var generateschema = &cobra.Command{
		Use:  "db-schema",
		RunE: actionGenerateDBSchema,
	}
	return []*cobra.Command{generaterule, generateconfig, generateschema}
}

// GetSubCommands is the list of commands the database module exposes
func GetSubCommands() []*cobra.Command {

	var getrule = &cobra.Command{
		Use:  "db-rules",
		RunE: actionGetDbRules,
	}

	var getconfig = &cobra.Command{
		Use:  "db-config",
		RunE: actionGetDbConfig,
	}

	var getschema = &cobra.Command{
		Use:  "db-schema",
		RunE: actionGetDbSchema,
	}

	return []*cobra.Command{getrule, getconfig, getschema}
}

func actionGetDbRules(cmd *cobra.Command, args []string) error {
	// Get the project and url parameters
	project := viper.GetString("project")
	commandName := cmd.Use

	params := map[string]string{}
	switch len(args) {
	case 1:
		params["dbAlias"] = args[0]
	case 2:
		params["dbAlias"] = args[0]
		params["col"] = args[1]
	}
	objs, err := GetDbRule(project, commandName, params)
	if err != nil {
		return nil
	}

	if err := utils.PrintYaml(objs); err != nil {
		return nil
	}
	return nil
}

func actionGetDbConfig(cmd *cobra.Command, args []string) error {
	// Get the project and url parameters
	project := viper.GetString("project")
	commandName := cmd.Use

	params := map[string]string{}
	if len(args) != 0 {
		params["dbAlias"] = args[0]
	}
	objs, err := GetDbConfig(project, commandName, params)
	if err != nil {
		return nil
	}

	if err := utils.PrintYaml(objs); err != nil {
		return nil
	}
	return nil
}

func actionGetDbSchema(cmd *cobra.Command, args []string) error {
	// Get the project and url parameters
	project := viper.GetString("project")
	commandName := cmd.Use

	params := map[string]string{}
	switch len(args) {
	case 1:
		params["dbAlias"] = args[0]
	case 2:
		params["dbAlias"] = args[0]
		params["col"] = args[1]
	}

	objs, err := GetDbSchema(project, commandName, params)
	if err != nil {
		return nil
	}
	if err := utils.PrintYaml(objs); err != nil {
		return nil
	}
	return nil
}

func actionGenerateDBRule(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		_ = utils.LogError("incorrect number of arguments", nil)
		return nil
	}
	dbruleConfigFile := args[0]
	dbrule, err := generateDBRule()
	if err != nil {
		return nil
	}

	_ = utils.AppendConfigToDisk(dbrule, dbruleConfigFile)
	return nil
}

func actionGenerateDBConfig(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		_ = utils.LogError("incorrect number of arguments", nil)
		return nil
	}
	dbruleConfigFile := args[0]
	dbrule, err := generateDBConfig()
	if err != nil {
		return nil
	}

	_ = utils.AppendConfigToDisk(dbrule, dbruleConfigFile)
	return nil
}

func actionGenerateDBSchema(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		_ = utils.LogError("incorrect number of arguments", nil)
		return nil
	}
	dbruleConfigFile := args[0]
	dbrule, err := generateDBSchema()
	if err != nil {
		return nil
	}

	_ = utils.AppendConfigToDisk(dbrule, dbruleConfigFile)
	return nil
}
