
package cmd

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/spf13/cobra"
	"encoding/json"
	"os/exec"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "trace the IP",
	Long: `trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)>0{
			for _ ,ip := range args{
				fmt.Println(ip)
				showData(ip)
			}
		}else{
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type trackingResult struct{
	IP	     string `json:"ip"`
	City	 string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	LOC      string `json:"loc"`
	Postal   string `json:"postal"`
	TimeZone string `json:"timezone"`
}
func showData(ip string){
	url := "http://ipinfo.io/" + ip + "/json"
	responseByte := getData(url)

	var responseByteData trackingResult

	checkValid := json.Valid(responseByte)

	if checkValid {
		err := json.Unmarshal(responseByte, &responseByteData)
		if err != nil{
			fmt.Println("Unable to unmarshal")
		}
		fmt.Println("DATA FOUND :")
		fmt.Println(responseByteData)
		openBrowser(responseByteData.LOC)
	}
	
}

func openBrowser(co_ordinates string){
	url := fmt.Sprintf("https://www.google.com/maps/@%s,15z?entry=ttu",co_ordinates)
	cmd := exec.Command("google-chrome",url)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Encountered error when starting command: %s\n", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command execution error: %s\n", err)
	}
}

func getData(url string) []byte {
	response , err := http.Get(url)
 	checkFetchError(err)
	responseByte , err := ioutil.ReadAll(response.Body)
 	checkFetchError(err)
	return responseByte
}

func checkFetchError(err error){
	if err != nil{
		fmt.Println("unable to fetch data")
	}
}
