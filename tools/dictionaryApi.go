package tools

import (
	inputforms "DictionaryENtoENBackend/controllers/_inputforms"
	"DictionaryENtoENBackend/entity/dictionary"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
)

func SendToDictionaryApi(inputForm inputforms.Search) (*dictionary.DocOutPut, error) {
	response, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + inputForm.Word)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	output := new(dictionary.DocOutPut)
	err = jsoniter.Unmarshal(body, output)
	if err != nil {
		panic(err)
	}

	return output, nil
}
