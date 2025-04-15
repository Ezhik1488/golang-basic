package api

import (
	"3-struct/config"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/fatih/color"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type JsonBinAPI struct {
	url    string
	client *http.Client
	Headers
}

type Headers struct {
	contentType string
	xMasterKey  string // Api ключ
	XBinPrivate bool   // Сделать Bin приватным или нет, по умолчанию приватный
	XBinMeta    bool   // Отображать мета информацию при GET запросах или нет
}

type CreateResponse struct {
	MetaData Metadata `json:"metadata"`
}

type Metadata struct {
	ID string `json:"id"`
}

func NewJsonBinAPI(cfg *config.Config) *JsonBinAPI {
	return &JsonBinAPI{
		url:    cfg.ApiUrl,
		client: &http.Client{},
		Headers: Headers{
			contentType: "application/json",
			xMasterKey:  cfg.ApiKey,
		},
	}
}

func (api *JsonBinAPI) Get(binID string) (string, error) {
	// Конструирование URL
	baseUrl := api.buildUrl(binID)

	// Создание GET запроса
	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		color.Red("Ошибка при формировании запроса. %v", err.Error())
		return "", err
	}
	req.Header.Add("Content-Type", api.contentType)
	req.Header.Add("X-Master-Key", api.xMasterKey)
	req.Header.Add("X-Bin-Meta", strconv.FormatBool(api.XBinMeta))

	// Выполнение GET запроса
	resp, err := api.makeRequest(req, "DELETE")
	if err != nil {
		return "", err
	}

	// Чтение ответа
	var result []byte
	result, err = io.ReadAll(resp.Body)

	return string(result), nil
}

func (api *JsonBinAPI) Create(data []byte) (*CreateResponse, error) {
	req, err := http.NewRequest("POST", api.url, bytes.NewBuffer(data))
	if err != nil {
		color.Red("Ошибка при формировании запроса. %v", err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", api.contentType)
	req.Header.Add("X-Master-Key", api.xMasterKey)
	req.Header.Add("X-Bin-Private", strconv.FormatBool(api.XBinPrivate))

	resp, err := api.makeRequest(req, "POST")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			color.Red("Ошибка при закрытии Response Body ")
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		color.Red("Ошибка при чтении тела ответа")
		return nil, err
	}

	var result CreateResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		color.Red("Ошибка при анмаршале ответа")
		return nil, err
	}

	return &result, nil
}

func (api *JsonBinAPI) Update(data []byte, binID string) error {
	// Конструирование URL
	baseUrl := api.buildUrl(binID)

	req, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewBuffer(data))
	if err != nil {
		color.Red("Ошибка при формировании запроса. %v", err.Error())
		return err
	}

	req.Header.Add("Content-Type", api.contentType)
	req.Header.Add("X-Master-Key", api.xMasterKey)

	resp, err := api.makeRequest(req, "PUT")
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		color.Red("Не удалось обновить данные. Проверьте BinID. StatusCode: %v", resp.StatusCode)
	}
	return nil
}

func (api *JsonBinAPI) Delete(binID string) error {
	// Конструирование URL
	baseUrl := api.buildUrl(binID)

	// Создание GET запроса
	req, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		color.Red("Ошибка при формировании запроса. %v", err.Error())
		return err
	}

	req.Header.Add("X-Master-Key", api.xMasterKey)

	// Выполнение DELETE запроса
	resp, err := api.makeRequest(req, "DELETE")
	if err != nil {
		return err
	}

	if resp.StatusCode == 400 {
		return errors.New("invalid BinID")
	}
	return nil
}

func (api *JsonBinAPI) makeRequest(request *http.Request, method string) (*http.Response, error) {
	resp, err := api.client.Do(request)
	if err != nil {
		color.Red("При выполнении %s запроса, произошла ошибка %v", method, err.Error())
		return nil, err
	}

	return resp, nil
}

func (api *JsonBinAPI) buildUrl(pathQuery ...string) *url.URL {
	baseUrl, err := url.Parse(api.url)
	if err != nil {
		color.Red("URL адрес некорректный. %v", err.Error())
		return nil
	}
	for _, v := range pathQuery {
		baseUrl.Path = path.Join(baseUrl.Path, v)
	}
	return baseUrl
}
