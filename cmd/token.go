package cmd

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

var o = &Options{}

type transporter struct {
	*http.Transport
	FakeTLSTermination bool
}

type Options struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	tokenCmd.Flags().StringVarP(&o.TokenURL, "tokenurl", "u", "", "token url")
	tokenCmd.Flags().StringVarP(&o.ClientID, "clientid", "i", "", "client id")
	tokenCmd.Flags().StringVarP(&o.ClientSecret, "clientsecret", "s", "", "secret")
}

func Run() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transporter{
			FakeTLSTermination: false,
			Transport:          &http.Transport{},
		},
	})
	oauthConfig := clientcredentials.Config{
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		TokenURL:     o.TokenURL,
	}
	t, err := oauthConfig.Token(ctx)
	if err != nil {
		panic(err)
	}
	pp.Println(t)

	log.Println("------ END --", time.Since(start))
}
