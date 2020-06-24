package entry

type RequestParams struct {
	Q    string `json:"q"`
	From string `json:"from"`
	To   string `json:"to"`
}

type AppConfig struct {
	AppKey string `json:"appKey"`
	AppPwd string `json:"appPwd"`
}

type TranslationParams struct {
	RequestParams
	AppConfig
	Salt     string `json:"salt"`
	Sign     string `json:"sign"`
	SignType string `json:"signType"`
	CurTime  string `json:"curtime"`
}

type Result struct {
	ErrorCode    string                 `json:"errorCode"`
	Query        string                 `json:"query"`
	Translation  []string               `json:"translation"`
	//Basic        DictBasic              `json:"basic"`
	Web          []DictWeb              `json:"web,omitempty"`
	//Lang         string                 `json:"l"`
	//Dict         map[string]interface{} `json:"dict,omitempty"`
	//Webdict      map[string]interface{} `json:"webdict,omitempty"`
	TSpeakUrl    string                 `json:"tSpeakUrl,omitempty"`
	SpeakUrl     string                 `json:"speakUrl,omitempty"`
	ReturnPhrase []string               `json:"returnPhrase,omitempty"`
}

type DictWeb struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type DictBasic struct {
	UsPhonetic string   `json:"us-phonetic"`
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	UkSpeech   string   `json:"uk-speech"`
	UsSpeech   string   `json:"us-speech"`
	Explains   []string `json:"explains"`
}
