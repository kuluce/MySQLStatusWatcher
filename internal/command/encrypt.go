package command

import (
	"MySQLStatusWatcher/internal/aesencrypt"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Code string

	EncryptCmd = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt the configuration",
		Run: func(cmd *cobra.Command, args []string) {
			_encryped := aesencrypt.Encrypt([]byte(Code))
			fmt.Println("Encrypted code:", _encryped)
		},
	}
)

func init() {
	EncryptCmd.Flags().StringVarP(&Code, "code", "c", "", "Code for encryption")
}

// func decrypt(encrypted string) string {
// 	_encryped_byte, err := hex.DecodeString(encrypted)
// 	if err != nil {
// 		fmt.Println("hex.DecodeString error:", err)
// 		return ""
// 	}
// 	_descryped, err := aesencrypt.AesDecrypt(_encryped_byte)
// 	if err != nil {
// 		fmt.Println("decrypt error:", err)
// 		return ""
// 	}
// 	return string(_descryped)
// }
