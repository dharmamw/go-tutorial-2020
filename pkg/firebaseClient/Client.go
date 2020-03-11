package firebaseclient

import (
	"context"
	"encoding/json"

	"go-tutorial-2020/internal/config"
	"go-tutorial-2020/pkg/errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	sharedClient = &firestore.Client{}
	credentials  = map[string]string{
		"type":                        "service_account",
		"project_id":                  "testv1-e51a9",
		"private_key_id":              "d07f7a53eb253f44bbedc2c5f90e39028eb7ef4c",
		"private_key":                 "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzFgQ498g0IOL8\nQZ8klVhfPi7wj6e5C/8tzO8yDkiq6bRn44dT9gxgSuL263a4qjgmq1Z8CDMkYbMy\nlWp9Z3P79H9MuYG9g7fkkMUqBsy2XSTTgnfM20eJaocKlo82yBl8rI0Wu5TUCL93\nqwE8iphM6PFVMELv7yXquYeQmc2zW0fQmRsBzVjXLVWWeITVtm3rtWcaXsasFcWj\n7OYltrKDTjNcNhfe1C380MEwu9s3TbfwIuynXUli2POz+G4BaCINDZw2OUrd7qxR\nlUpo2m3aOU4DEC7GWYmhldEQN/lZNSS5ofCSv/6UgOhuuMFzei9Nc/I/pF4Ms4wz\nn+5zdutTAgMBAAECggEAElvi8e4g5Ysrid22UVtIFHd8+2dWs+cxcUINkW2acH8F\n4hKWuHpEWYjDzLll0deNaxVsh8mVaJjldH2RzapR3xX5COYJkWKT8wgOVlkdGmLh\nGBfLbUJbipBTqLe3lc+coUXVLuvq/XOqITv0I/83TscgmnGYox06n14GskG6LzUz\nyOdPCsX4/1ylsc7XjP/ATWYj3mb2+kudAPOJ/a35bSIMiO6+csZqJd/16WNpJP6e\ngmBGTh4qXSvedId7pa3p0CJD7yEi8TnGX0SR4M1GHn0aGdTdWwR/7Oh3o0QazOd2\nVzJk3ZpzaceiJSZ/nz9ZhCgpiheV2ohw5Iis11P0AQKBgQDqudTChUwT8ozChZZp\nSudd78QrwTUlOu5wozrYJIg4h+qwfMVfVmocCehhhLXwqrGoBUw++dcDw6xC7+4s\n3qZwI1tpY8qwEW2um7lx0WQUC/g0YmGVruTyW6NP+GDoTEsP6DCxN7RiP/YniM4M\n0KgqfDW7d1LEguXEbxkTC7lkAQKBgQDDUTbs+TfLdna8ApfKp26q46JijjC49g5X\n+Ou1Nu2esNvWIVSCcaIuSXmr4p0NrzxiLvnvwPpCrMmsT7/yBrW/7unO7+IwjUMu\n3NJyq0U8834GSuSLuMEMWboqiBWhYxhUduFARS5I0S11mCRAPR9g4DEjDdpzkDtt\nO/D7Eb9/UwKBgQC4Itf5UUCBVYF/B/Ua3JvsYS9vc74RWs26pxJ+hQon7tf5Y6gi\nRlQvcsZN2iIwjneX67AIp045scLtL/OUV4YR6mrHnnVe8g0tMRSlaTItV8Z6scMD\n/ZO1XJYdIihDk7Y+4FNyctSbTrn4AaZC/10tFwu/6LeRWW6OTulIu6XQAQKBgQCM\nWKtTFLFW7kTbTDGuWSlYekGQ+ANipMwhwqf8iv+r7AHLmB+Vq/mRsRJQxIF999E3\n6/GEqfIqPuabfqK6Ur/+rrorDIxHvvnrjplZ9F/IMF76Po6DJ7rwGPmA3lBMq1ws\nNVAeUpezkztLKIvD7SfDANXODoJOT/GjyFKc1l/4KQKBgQCSRmDxmkUECTupeOH1\na+yt5yITr7jyiZRvIYyTyIKFZ69rm59bxnEog0vcUQqp0rTC282gEf2UVU+wbWPU\n3OA7kjoOlxefDP05MDKmO4O+pV1ckrth6nNzIWl/LNRHUSmZQUf31mYZslppyAM8\nzoTpccVO2NwYA0CVU0PIPcHnRg==\n-----END PRIVATE KEY-----\n",
		"client_email":                "firebase-adminsdk-q8uew@testv1-e51a9.iam.gserviceaccount.com",
		"client_id":                   "110691299263637436058",
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-q8uew%40testv1-e51a9.iam.gserviceaccount.com",
	}
)

// Client ...
type Client struct {
	Client *firestore.Client
}

// NewClient ...
func NewClient(cfg *config.Config) (*Client, error) {
	var c Client
	cb, err := json.Marshal(credentials)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to marshal credentials!")
	}
	option := option.WithCredentialsJSON(cb)
	c.Client, err = firestore.NewClient(context.Background(), cfg.Firebase.ProjectID, option)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firebase client!")
	}
	return &c, err
}
