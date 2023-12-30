package spot

type Spot struct {
	apiKey    string
	secretKey string
}

func Create(api, secret string) *Spot {
	return &Spot{
		apiKey:    api,
		secretKey: secret,
	}
}
