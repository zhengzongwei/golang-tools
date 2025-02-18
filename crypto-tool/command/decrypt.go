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

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt data using a salt",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Data to decrypt is required.")
			return
		}
		data := args[0]
		salt, _ := cmd.Flags().GetString("salt")
		tool := NewCryptoTool(salt)
		decryptedData, err := tool.decrypt(data)
		if err != nil {
			fmt.Println("Decryption failed:", err)
			return
		}
		fmt.Println("Decrypted data:", decryptedData)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringP("salt", "s", "", "Salt for key derivation")
}
