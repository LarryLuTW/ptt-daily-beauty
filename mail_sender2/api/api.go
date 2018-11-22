	// // "main/mail"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"

	// "main/model"

	// url := "https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// body, _ := ioutil.ReadAll(resp.Body)

	// beauties := make([]model.Beauty, 3)
	// err = json.Unmarshal(body, &beauties)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", beauties)