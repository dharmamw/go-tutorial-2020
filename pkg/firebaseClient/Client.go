package firebaseclient

import (
	"context"
	"encoding/json"

	"tugas-arif/internal/config"
	"tugas-arif/pkg/errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	sharedClient = &firestore.Client{}
	credentials  = map[string]string{
		
			"type": "service_account",
			"project_id": "testarif-f4faa",
			"private_key_id": "07743b80a872c0d7fd235df0c6efd0420246857c",
			"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCHVeWI1xIm/Ud6\nKd4kVOl1G1D28QPnWXY5PqrdD9Mp3QJV5/cOd6bBVCyhR882lFJ7ZZLd7x/YaYTm\nNRt2DWJddfDwnQF1W4IuWqpSTqHK2KfMiqAKV/0roJZTlQ765yliYZ0zybDr0szk\nIGavIs1akJt0j/zOp99KX8PvAUPV/9+m44BrBrCKMWrrYYpquycYrQB+3lQssXHV\nwZyl1RZQZfx/VwmOhKv++CN+1xRpM54BUoJItnE9D3doFGIysbJINZP/fdR+dAMs\nUrHlcY3rlRer4riakdcUzSAa6F2Hn2YtddJNUOayOwkZ4rJlbA+eAUV2h/AiEhMl\nXKi+dPc5AgMBAAECggEABAzjEL33Ehza12IkMgwq7ISDxZ0GLXwoJXBygELP/BAp\nC+8kUSK6Zk2hnelE31cZdnReX+JsBPfesJEYU4EoPeNtiy+RWQGctMxyQOjCvz80\nTmyiFSbEmFpD0ktcwALMSOB8NvRei8BEV40Zfzy+zp4TUpRYnMq5zENvdsC5igPS\nRYwt1ha0h0NYEYd3+gCNO/F9dhTLZTj0S0mL9ff/faYYhJmzhO4+lZ3GW4zn+g1L\nNaZphYnJ7D6B+Vyh2HXa5en5GmoGxCCywL2sOWOELT+RZddfRqMiQUDxdE5z/RCv\nIeXd5H//ZhZLc/4Dp1fBtmruFqat193176EvuznrsQKBgQC/ELXQvnKNmfZJesJe\ns8EESGCBciauJ9IDvPXtVm/kjvzXZzJMNhh2DUaLU/Kc4nfho3RgEijObYZRQbuW\nkIp4ChmYEoVf7ultmFw9vcSk8uDUEklLUFQrFEKD8VTtWLv0HiGEk4lm4YoWnPKi\np1BtycSxt831++FFPMyoKuWMpQKBgQC1VIVRnQCYdj0Lji3x3Aca1vBmxuHmi6CO\nkIdJXBZNByK/TZNTaXkkp8yW0eVxNvNu7Ixap+8MIp7CthmE0+uIlYw6a3qPKp2a\nevpftSpynxjWA9iBJ11tnZu4UmoTz/3c57QvIunkJaxITlVqpEOJoljBYH74YZOH\nTIYqjLbYBQKBgFiFCmdCxNnb2eIjMMglaahtS+DNHSSUqFU5B4tE/6QZpwS49/Gd\nImoXLnbAlueeeMIeM32LDELPNWqSFLHmF3ET5NWyxv4yNw2iiCHGuMNfD1DRhAmT\nlts6kLKGbb1k3fd0ujytCfyTQ6HEZxl6gOXMlAduS8rKPo0QZRUIgr9NAoGAd5R0\nkvy5ztFysnMh43TZjp6eTPjtMo9z43B2dy9uWX/SL1xmQsS0qjKqXe+vori9UrJW\nYNaMc3FFR1y1eX4Tvq/4mPIWEeHlq2FcSc98XbiDtWc12P5vw4EDl0tqPwSUAqEe\nl1Mr8VPSyKA3/iqzi0lvxJ7xPLWEh940QE0pq00CgYEAlm8FyDEgvIapE1liB/fI\nTmCQu4hpTXHGOkYkXQXTiitF/UgwBI2j5/JymiUf+pMcq8eR/1kXbTJL+xpHM+9E\nJHQMX2WEMWj7I8iVu4LXcHQG0vwl63n2rdxwAnCTdbtDMv/AsEFqzW3djgOMV6mg\nMLSD0ipuEBco2SM9K2uhTms=\n-----END PRIVATE KEY-----\n",
			"client_email": "firebase-adminsdk-pnlki@testarif-f4faa.iam.gserviceaccount.com",
			"client_id": "117187113971759922683",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-pnlki%40testarif-f4faa.iam.gserviceaccount.com",
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
