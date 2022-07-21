package application_gateway

import (
	"context"
	"strconv"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var applications = []string{"positive-authorisation", "negative-authorisation", "path-related-error-handling", "missing-resources-error-handling", "proxy-cases", "proxy-errors"}

func (gs *GatewaySuite) TestGetRequest() {

	for _, app := range applications {
		app, err := gs.cli.ApplicationconnectorV1alpha1().Applications().Get(context.Background(), app, v1.GetOptions{})
		gs.Nil(err)

		gs.Run(app.Spec.Description, func() {
			for _, service := range app.Spec.Services {
				gs.Run(service.Description, func() {
					http := NewHttpCli(gs.T())

					for _, entry := range service.Entries {
						if entry.Type != "API" {
							gs.T().Log("Skipping event entry")
							continue
						}

						expectedCode, err := getExpectedHTTPCode(service)
						if err != nil {
							gs.T().Log("Error during getting the error code from description -> applicationCRD")
							gs.T().Fail()
						}

						res, _, err := http.Get(entry.CentralGatewayUrl)
						gs.Nil(err, "Request failed")
						gs.Equal(expectedCode, res.StatusCode, "Incorrect response code")
					}
				})
			}
		})
	}
}

func (gs *GatewaySuite) TestResponseBody() {
	app, err := gs.cli.ApplicationconnectorV1alpha1().Applications().Get(context.Background(), "proxy-cases", v1.GetOptions{})
	gs.Nil(err)
	for _, service := range app.Spec.Services {
		gs.Run(service.Description, func() {
			http := NewHttpCli(gs.T())

			for _, entry := range service.Entries {
				if entry.Type != "API" {
					gs.T().Log("Skipping event entry")
					continue
				}

				expectedCode, err := getExpectedHTTPCode(service)
				if err != nil {
					gs.T().Log("Error during getting the error code from description -> applicationCRD")
					gs.T().Fail()
				}

				_, body, err := http.Get(entry.CentralGatewayUrl)
				gs.Nil(err, "Request failed")

				codeStr := strconv.Itoa(expectedCode)

				gs.Equal(codeStr, string(body), "Incorrect body")
			}
		})
	}
}
