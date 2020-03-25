package validation

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/go-playground/validator/v10"
)

type emailValidationResponse struct {
    Valid      bool   `json:"valid"`
    Block      bool   `json:"block"`
    Disposable bool   `json:"disposable"`
    Domain     string `json:"doamin"`
    Text       string `json:"text"`
    Reason     string `json:"reason"`
    MXHost     string `json:"mx_host"`
    MXInfo     string `json:"mx_info"`
    MXIP       string `json:"mx_ip"`
}

func emailValidation(v *validate) validator.Func {
    baseURL := "https://mailcheck.p.rapidapi.com?domain=%s"
    return func(fl validator.FieldLevel) bool {
        url := fmt.Sprintf(baseURL, fl.Field().String())

        req, _ := http.NewRequest("GET", url, nil)

        req.Header.Add("x-rapidapi-host", "mailcheck.p.rapidapi.com")
        req.Header.Add("x-rapidapi-key", v.Config.EmailAPIKey)

        res, _ := http.DefaultClient.Do(req)

        defer res.Body.Close()

        body, _ := ioutil.ReadAll(res.Body)
        var payload emailValidationResponse
        if err := json.Unmarshal(body, &payload); err != nil {
            log.Error(err)
            return false
        }

        if payload.Valid && !payload.Block && !payload.Disposable {
            return true
        }

        return false
    }
}
