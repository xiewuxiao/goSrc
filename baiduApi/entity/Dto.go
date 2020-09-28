package entity

type Token struct {
	REFRESH_TOKEN string `json:refresh_token`
	EXPIRES_IN    int    `json:expires_in`
	SESSION_KEY   string `json:session_key`
	ACCESS_TOKEN  string `json:access_token`
	SCOPE         string `json:scope`
}

type Result struct {
	LOG_ID string `json:log_id`
	WORDS_RESULT_NUM    int    `json:words_result_num`
	WORDS_RESULT []Words `json:words_result`
	
}

type Words struct {
	WORDS string `json:words`
}
