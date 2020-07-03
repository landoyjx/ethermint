package main

import (
	"fmt"

	"github.com/cosmos/ethermint/version"
	"github.com/spf13/cobra"
)

// var (
// 	VERSION     string
// 	BUILD_TIME  string
// 	GO_VERSION  string
// 	GIT_BRANCH  string
// 	COMMIT_SHA1 string
// )

var versionCmd = &cobra.Command{

	Use: "v",

	Short: "Print the version number of ethermint",
	Long:  `This is ethermint's version`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("version:\t%s \nbuild time:\t%s\ngit branch:\t%s\ngit commit:\t%s\ngo version:\t%s\n", version.VERSION, version.BUILD_TIME, version.GIT_BRANCH, version.COMMIT_SHA1, version.GO_VERSION)
	},
}
