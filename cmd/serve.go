/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	adminv1 "whalepower.tech/api/v1/admin"
	nodev1 "whalepower.tech/api/v1/node"
	testv1 "whalepower.tech/api/v1/test"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "cloudnative computing platform",
	Long:  `a cloudnatvive computing platform, support nivia card.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		testv1.Entry.InitTestRouter(r.Group("/test"))
		nodev1.Entry.InitNodeRouter(r.Group("/node"))
		adminv1.Entry.InitAdminRouter(r.Group("/admin"))
		r.Run(fmt.Sprintf(":%d", globalCfg.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
