package urlconstants

const (
	API                string = "/api"
	AUTH               string = API + "/auth"
	AUTH_SIGN_IN       string = AUTH + "/signin"
	AUTH_SIGN_UP       string = AUTH + "/signup"
	AUTH_REFRESH_TOKEN string = AUTH + "/refresh-token"
)

func AUTH_ALLOWED_NO_JWT() []string {
	return []string{AUTH_SIGN_IN, AUTH_SIGN_UP}
}

func AUTH_ALLOWED_REFRESH_JWT() []string {
	return []string{AUTH_REFRESH_TOKEN}
}

func NO_JWT() []string {
	//change []string{} with no jwt urls
	return append(AUTH_ALLOWED_NO_JWT(), []string{}...)
}
