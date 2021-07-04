package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/cobra"
)

var exportCmd *cobra.Command

func export(cmd *cobra.Command, args []string) error {
	CK := os.Getenv("TW_CK")
	CS := os.Getenv("TW_CS")
	AT := os.Getenv("TW_AT")
	AS := os.Getenv("TW_AS")
	if CK == "" || CS == "" || AT == "" || AS == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}
	anaconda.SetConsumerKey(CK)
	anaconda.SetConsumerSecret(CS)
	api := anaconda.NewTwitterApi(AT, AS)

	fmt.Println("Getting block ids...")
	ids, err := getBlockIds(api)
	if err != nil { 
		return fmt.Errorf("Error on getting block ids from Twitter: %s", err)
	}
	fmt.Println("Done!")

	var outFile string
	if len(args) == 0 {
		outFile = "block_ids.txt"
	} else {
		outFile = args[0]
	}

	fmt.Println("Exporting to file...")
	err = exportIdsToFile(ids, outFile)
	if err != nil {
		return fmt.Errorf("Error exporting ids to file: %w", err)
	}
	fmt.Println("Export completed.")
	return nil
}

func exportIdsToFile(ids []int64, outFile string) error {
	file, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(sliceIntToString(ids, " ")))
	return err
}


// https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/mute-block-report-users/api-reference/get-blocks-list
func getBlockIds(api *anaconda.TwitterApi) (ids []int64, err error) {
	ids = make([]int64, 0)
	cur := anaconda.Cursor{Next_cursor: -1, Next_cursor_str: "-1"} // If no cursor is provided, a value of -1 will be assumed, which is the first "page."
	for cur.Next_cursor != 0 {
		v := url.Values{} // https://pkg.go.dev/net/url#Values
		v.Add("curcor", cur.Next_cursor_str)
		cur, err = api.GetBlocksIds(v)

		if err != nil {
			return ids, err
		}

		ids = append(ids, cur.Ids...)
	}
	return ids, err
}

func sliceIntToString(s []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(s), " ", delim, -1), "[]")
}

func init() {
	exportCmd = &cobra.Command{
		Use: "export",
		Short: "Export block ids to file",
		RunE: export,
	}

	rootCmd.AddCommand(exportCmd)
}