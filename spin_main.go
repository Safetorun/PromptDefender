package main

//
//import (
//	"encoding/json"
//	"fmt"
//	spinconfig "github.com/fermyon/spin/sdk/go/config"
//	spinhttp "github.com/fermyon/spin/sdk/go/http"
//	"github.com/prompt_protect_go/aiprompt"
//	"github.com/prompt_protect_go/app"
//	"net/http"
//)
//
//type AppResponse struct {
//	AiScore float32 `json:"aiScore"`
//}
//
//func init() {
//	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
//
//		openAIKey, err := spinconfig.Get("open_ai_api_key")
//
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		answer, _ := app.New(aiprompt.NewOpenAI(openAIKey)).CheckAI(r.URL.Query().Get("prompt"))
//
//		w.Header().Set("Content-Type", "application/json")
//
//		response := AppResponse{AiScore: answer}
//		output, err := json.Marshal(response)
//
//		if err != nil {
//			return
//		}
//
//		_, err = fmt.Fprintln(w, string(output))
//		if err != nil {
//			return
//		}
//	})
//}
//
//func main() {
//}
