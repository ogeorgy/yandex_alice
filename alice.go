package alice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler обрабатывает POST-запрос на webhook URL(path) и генерирует ответ
func Handler(path string, Propagate func(APIRequest) Response) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading the body: %v\n", err)
		}

		var req APIRequest
		err = json.Unmarshal(data, &req)
		if err != nil {
			fmt.Printf("Error unmarshal: %v\n", err)
		}

		res := APIResponse{
			Version:  req.Version,
			Session:  req.Session,
			Response: Propagate(req),
		}

		data, err = json.Marshal(res)
		if err != nil {
			fmt.Printf("Error marshal: %v\n", err)
		}
		fmt.Fprint(w, string(data))
	})
}
