package function_playground

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%s:%v\n", k, v)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))

	html := `
<html>
<head>
<meta name="google-site-verification" content="uvH6Ab-FlxjwRMyKCn4chx6Pm8PpicF-LCpBRBv4fq0" />
<meta name="google-site-verification" content="WT9pzr9BT14fxp4J0FEAOcu7GsAz3-7yXhFCo_046to" />
</head>
<body>
</body>
</html>
`

	w.Header().Set("Content-Type", "html/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(html))
	if err != nil {
		fmt.Println(err.Error())
	}
}
