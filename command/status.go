package command

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"text/tabwriter"
)

func CmdStatus(c *cli.Context) error {
	host := c.String("host")
	port := c.Uint("port")

	resp, err := http.Get(fmt.Sprintf("http://%s:%d/JSON|42", host, port))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response StatusResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	writer := tabwriter.NewWriter(c.App.Writer, 4, 4, 1, ' ', 0)

	fmt.Fprintln(writer, "ID\tModel\tSkin\tDriver")

	for i, v := range response.Cars {
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\n", i + 1, v.Model, v.Skin, v.DriverName)
	}

	writer.Flush()

	return nil
}

type StatusResponse struct {
	Cars []*Car `json:"Cars"`
}

type Car struct {
	Model           string `json:"Model"`
	Skin            string `json:"Skin"`
	DriverName      string `json:"DriverName"`
	DriverTeam      string `json:"DriverTeam"`
	IsConnected     bool   `json:"IsConnected"`
	IsRequestedGUID bool   `json:"IsRequestedGUID"`
	IsEntryList     bool   `json:"IsEntryList"`
}
