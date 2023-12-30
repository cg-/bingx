package standard

type Standard struct {
	apiKey    string
	secretKey string
}

func Create(api, secret string) *Standard {
	return &Standard{
		apiKey:    api,
		secretKey: secret,
	}
}
