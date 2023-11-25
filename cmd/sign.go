/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Command to generate pluslink signature",
	Long:  `Command to generate pluslink signature`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// getSignCmd represents the getSign command
var getSignCmd = &cobra.Command{
	Use:   "get",
	Short: "Command to generate pluslink signature with GET http method",
	Long:  `Command to generate pluslink signature with GET http method`,
	Run: func(cmd *cobra.Command, args []string) {
		userCodeVal, _ := cmd.Flags().GetString("userCode")
		timeVal, _ := cmd.Flags().GetString("time")
		pwdVal, _ := cmd.Flags().GetString("password")

		md5 := md5.New()
		md5.Write([]byte(strings.ToLower(userCodeVal)))
		userCodeHashBytes := md5.Sum(nil)

		userCodeHash := hex.EncodeToString(userCodeHashBytes)

		md5.Reset()
		md5.Write([]byte(pwdVal))
		pwdHashBytes := md5.Sum(nil)

		pwdHash := hex.EncodeToString(pwdHashBytes)

		signPlain := userCodeHash + "#" + pwdHash + "#" + timeVal

		md5.Reset()
		md5.Write([]byte(signPlain))
		signHashBytes := md5.Sum(nil)

		signHash := hex.EncodeToString(signHashBytes)

		fmt.Printf("userCode : %s\n", userCodeVal)
		fmt.Printf("time : %s\n", timeVal)
		fmt.Printf("signature : %s\n", signHash)
	},
}

// postSignCmd represents the postSign command
var postSignCmd = &cobra.Command{
	Use:   "post",
	Short: "Command to generate pluslink signature with POST http method",
	Long:  `Command to generate pluslink signature with POST http method`,
	Run: func(cmd *cobra.Command, args []string) {
		commandVal, _ := cmd.Flags().GetString("command")
		timeVal, _ := cmd.Flags().GetString("time")
		pwdVal, _ := cmd.Flags().GetString("password")

		md5 := md5.New()
		md5.Write([]byte(pwdVal))
		pwdHashBytes := md5.Sum(nil)

		pwdHash := hex.EncodeToString(pwdHashBytes)

		signPlain := strings.ToUpper(commandVal) + timeVal + pwdHash

		md5.Reset()
		md5.Write([]byte(signPlain))
		signHashBytes := md5.Sum(nil)

		signHash := hex.EncodeToString(signHashBytes)

		fmt.Printf("command : %s\n", strings.ToUpper(commandVal))
		fmt.Printf("time : %s\n", timeVal)
		fmt.Printf("signature : %s\n", signHash)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.AddCommand(getSignCmd)
	signCmd.AddCommand(postSignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getSignCmd.PersistentFlags().StringP("userCode", "u", "", "pluslink client's code (HHxxxxx)")
	getSignCmd.PersistentFlags().StringP("time", "t", "", "time parameter used in transaction request")
	getSignCmd.PersistentFlags().StringP("password", "p", "", "pluslink client's transaction password")

	postSignCmd.PersistentFlags().StringP("command", "c", "", "command parameter used in transaction request (ex: PAY.PULSA)")
	postSignCmd.PersistentFlags().StringP("time", "t", "", "time parameter used in transaction request")
	postSignCmd.PersistentFlags().StringP("password", "p", "", "pluslink client's transaction password")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getSignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
