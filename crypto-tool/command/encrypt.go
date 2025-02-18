/*
 * Copyright (c)2025. zhengzongwei
 * golang-tools is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *         http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 *
 */

package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt data using a salt",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Data to encrypt is required.")
			return
		}
		data := args[0]
		salt, _ := cmd.Flags().GetString("salt")
		if salt == "" {
			salt = generateMachineCode()
		}
		tool := NewCryptoTool(salt)
		encryptedData, err := tool.encrypt(data)
		if err != nil {
			fmt.Println("Encryption failed:", err)
			return
		}
		fmt.Println("Encrypted data:", encryptedData)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringP("salt", "s", "", "Salt for key derivation")
}
