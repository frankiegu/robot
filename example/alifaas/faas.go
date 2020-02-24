package main

import (
	"encoding/json"
	"fmt"
	"github.com/sqeven/robot/issue"
	"github.com/sqeven/robot/processor/drone_promote"
	"github.com/sqeven/robot/processor/say_chat"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	value, err := url.ParseQuery(string(b))
	git := value.Get("payload")
	if strings.Contains(git, "github.com") {
		event := &issue.IssueCommentEvent{}
		eventstr := []byte(git)
		if len(eventstr) == 0 {
			return
		}
		err = json.Unmarshal(eventstr, &event.GitHub)
		if err != nil {
			fmt.Printf("decode event error : %s", err)
			return
		}
		fmt.Printf("repo name is : %s body: %s/n", *event.GitHub.Repo.FullName, *event.GitHub.Comment.Body)
		config := issue.NewConfig("", "", "github")
		issue.Regist("/say", &say_chat.SayChat{"", ""})
		issue.Regist("/promote", &drone_promote.DronePromote{"", ""})
		err = issue.Process(config, *event)
		if err != nil {
			fmt.Printf("promote error %s", err)
		}
	} else {
		event := &issue.IssueCommentEvent{}
		eventstr := []byte(git)
		if len(eventstr) == 0 {
			return
		}
		err = json.Unmarshal(eventstr, &event.GoGs)
		if err != nil {
			fmt.Printf("decode event error : %s", err)
			return
		}
		fmt.Printf("repo name is : %s body: %s/n", event.GoGs.Repository.FullName, event.GoGs.Comment.Body)
		config := issue.NewConfig("", "", "gogs")
		issue.Regist("/say", &say_chat.SayChat{"http://openapi.tuling123.com/openapi/api/v2", "sqeven"})
		err = issue.Process(config, *event)
		if err != nil {
			fmt.Printf("promote error %s", err)
		}
	}
}

func promoteEvent(body []byte) *issue.IssueCommentEvent {
	event := &issue.IssueCommentEvent{}
	if strings.Contains(string(body), "github.com") {
		json.Unmarshal(body, event)
		config := issue.NewConfig("", "", "github")
		issue.Regist("/say", &say_chat.SayChat{"", ""})
		issue.Regist("/promote", &drone_promote.DronePromote{"", ""})
		err := issue.Process(config, *event)
		if err != nil {
			fmt.Printf("promote error %s", err)
			return nil
		}
	} else {
		json.Unmarshal(body, event)
		config := issue.NewConfig("", "", "gogs")
		issue.Regist("/say", &say_chat.SayChat{"", ""})
		err := issue.Process(config, *event)
		if err != nil {
			fmt.Printf("promote error %s", err)
			return nil
		}
	}

	return event
}

func main() {
	fmt.Println("FunctionCompute go runtime inited.")
	http.HandleFunc("/", handler)
	port := os.Getenv("FC_SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":"+port, nil)
}
