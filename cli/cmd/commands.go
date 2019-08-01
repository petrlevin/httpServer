package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"

	"os"
)


type CommandBuilder struct{

}


func (cb CommandBuilder) Build()(cm *cobra.Command){
	rootCmd := &cobra.Command{}



	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}
	copyCmd := &cobra.Command{
		Use:   "copy",
		Short: "Copy files",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("We are about copying files %s", strings.Join(args,";"))
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(copyCmd)
	return rootCmd;
}









func Execute() {
	rootCmd :=CommandBuilder{}.Build()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}