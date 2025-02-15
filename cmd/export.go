package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/daichirata/hammer/internal/hammer"
)

var (
	exportExample = `
* Export spanner schema
  hammer export spanner://projects/projectId/instances/instanceId/databases/databaseName > schema.sql`

	exportCmd = &cobra.Command{
		Use:     "export SOURCE",
		Short:   "Export schema",
		Example: exportExample,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("must specify 1 argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			sourceURI := args[0]

			source, err := hammer.NewSource(ctx, sourceURI)
			if err != nil {
				return err
			}
			ddl, err := source.DDL(ctx)
			if err != nil {
				return err
			}
			for _, stmt := range ddl.List {
				fmt.Println(stmt.SQL() + ";\n")
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(exportCmd)
}
