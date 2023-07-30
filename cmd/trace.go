
package cmd

import (
	"fmt"
	"net/http"
	"net"
	"io/ioutil"
	"github.com/spf13/cobra"
	"encoding/json"
	"os/exec"
	"time"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "trace the IP or DNS domain & show its location",
	Long:  `trace the IP or DNS domain & show its location`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)>0{
			for _ , arg := range args{

				// checking if given argument is an IP or a DNS domain
				check := isIP(arg)
				if check {
					showData(arg)
				}
				if !check {
					ip, err := getIPAddress(arg)
					if err != nil {
						fmt.Println(err)
						continue
					}
					if len(ip) == 0 {
						fmt.Printf("no IP address found for domain: %s", arg)
						continue
					}
					showData(ip)
			}
		}
	}else{
		fmt.Println("NO IP ADDRESS OR DNS DOMAIN ENTERED")
	}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

// Check if the input is an IP address

func isIP(args string) bool{
		ip := net.ParseIP(args)
		if ip != nil {
			if ip!=nil{
				return true
			}
		}
		return false 
}

// Get the IP address of a domain
func getIPAddress(domain string) (string, error) {
	ip, err := net.LookupIP(domain)
	if err != nil {
		return "", err
	}
	return ip[0].String(), nil
}

// trackingResult represents the structure of the data received from the API

type trackingResult struct{
	IP	     string `json:"ip"`
	City	 string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	LOC      string `json:"loc"`
	Postal   string `json:"postal"`
	TimeZone string `json:"timezone"`
}

// showData retrieves IP geolocation data and opens the location on Google Maps

func showData(ip string){
	url := "http://ipinfo.io/" + ip + "/json"
	responseByte := getData(url)

	var responseByteData trackingResult

	checkValid := json.Valid(responseByte)

	if checkValid {
		err := json.Unmarshal(responseByte, &responseByteData) // unmarshalling json
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println("DATA FOUND :")

		fmt.Printf(" IP : %s\n City : %s\n Region : %s\n Country : %s\n LOC : %s\n Postal : %s\n TimeZone : %s\n\n",
				responseByteData.IP,responseByteData.City,responseByteData.Region,responseByteData.Country,
				responseByteData.LOC,responseByteData.Postal,responseByteData.TimeZone)

		openBrowser(responseByteData.LOC)
	}
	
}

// openBrowser opens the user's default web browser with the Google Maps location

func openBrowser(co_ordinates string){
	url := fmt.Sprintf("http://maps.google.com/?q=%s", co_ordinates)

	fmt.Println("------Google map is loading------") 
	time.Sleep(5 * time.Second)              // adding a time delay of 5 seconds
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

// getData fetches data from the given URL and returns the response as bytes

func getData(url string) []byte {
	response , err := http.Get(url)
 	checkFetchError(err)
	responseByte , err := ioutil.ReadAll(response.Body)
 	checkFetchError(err)
	return responseByte
}

// checking if any error in fetching data 

func checkFetchError(err error){
	if err != nil{
		fmt.Println("unable to fetch data")
	}
}

