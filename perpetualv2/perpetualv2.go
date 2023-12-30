package perpetualv2

type PerpetualV2 struct {
	apiKey    string
	secretKey string
}

func Create(api, secret string) *PerpetualV2 {
	return &PerpetualV2{
		apiKey:    api,
		secretKey: secret,
	}
}
