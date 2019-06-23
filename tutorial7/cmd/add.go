/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"encoding/binary"
	"encoding/gob"
	"bytes"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task.",
	Long: `Add a new task. For example:
		task add this is a new task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		db, err := bolt.Open("task.db", os.ModeAppend, nil)

		if err != nil {
			panic(err)
		}

		defer db.Close()

		db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
			
			if err != nil {
				return fmt.Errorf("Create bucket: %s", err)
			}
			
			buf := &bytes.Buffer{}
			gob.NewEncoder(buf).Encode(args)

			id, _ := bucket.NextSequence()
			return bucket.Put(itob(int(id)), buf.Bytes())
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}