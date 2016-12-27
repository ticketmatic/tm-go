package accountparameters

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// Get all configured account parameters
func Getlist(client *ticketmatic.Client) ([]*ticketmatic.AccountParameter, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/accountparameters")

	var obj []*ticketmatic.AccountParameter
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get an account parameter
func Get(client *ticketmatic.Client, name string) (*ticketmatic.AccountParameter, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/accountparameters/{name}")
	r.UrlParameters(map[string]interface{}{
		"name": name,
	})

	var obj *ticketmatic.AccountParameter
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Set an account parameter
func Set(client *ticketmatic.Client, data *ticketmatic.AccountParameter) (*ticketmatic.AccountParameter, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/accountparameters")
	r.Body(data)

	var obj *ticketmatic.AccountParameter
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
