package say_chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogs/go-gogs-client"
	"github.com/google/go-github/github"
	"github.com/sqeven/robot/issue"
	"io/ioutil"
	"net/http"
	"strings"
)

type SayChat struct {
	ChatServer string
	ChatToken  string
}

type SayChatCmd struct {
	Content string
	Params  map[string]string
}

type AiReq struct {
	ReqType    int `json:"reqType"`
	Perception struct {
		InputText struct {
			Text string `json:"text"`
		} `json:"inputText"`
		InputImage struct {
			URL string `json:"url"`
		} `json:"inputImage"`
		SelfInfo struct {
			Location struct {
				City     string `json:"city"`
				Province string `json:"province"`
				Street   string `json:"street"`
			} `json:"location"`
		} `json:"selfInfo"`
	} `json:"perception"`
	UserInfo struct {
		APIKey string `json:"apiKey"`
		UserID string `json:"userId"`
	} `json:"userInfo"`
}

type AiRsp struct {
	Intent struct {
		Code       int    `json:"code"`
		IntentName string `json:"intentName"`
		ActionName string `json:"actionName"`
		Parameters struct {
			NearbyPlace string `json:"nearby_place"`
		} `json:"parameters"`
	} `json:"intent"`
	Results []struct {
		GroupType  int    `json:"groupType"`
		ResultType string `json:"resultType"`
		Values     struct {
			URL  string `json:"url"`
			Text string `json:"text"`
		} `json:"values,omitempty"`
	} `json:"results"`
}

func (d *SayChat) AiChat(what string) (string, error) {
	client := &http.Client{}
	aiReq := new(AiReq)
	aiReq.ReqType = 0
	aiReq.UserInfo.APIKey = "0aa1f77c569b46a3ac5c03430fbc62fa"
	aiReq.UserInfo.UserID = d.ChatToken
	aiReq.Perception.InputText.Text = what
	info, err := json.Marshal(aiReq)
	if err != nil {
		return "Hi, i am robot!", err
	}
	reader := string(info)
	req, err := http.NewRequest("POST", d.ChatServer, strings.NewReader(reader))
	if err != nil {
		return "Hi, i am robot!", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "Hi, i am robot!", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Hi, i am robot!", err
	}

	rsp := new(AiRsp)
	bodystr := string(body)
	if err := json.Unmarshal([]byte(bodystr), rsp); err != nil {
		return "Hi, i am robot!", err
	}
	return rsp.Results[0].Values.Text, nil
}

func (d *SayChat) Process(config issue.Config, event issue.IssueEvent) error {
	if config.Git == "github" {
		git := issue.NewGitClient(config)
		input := &github.IssueComment{Body: github.String(":robot:Hi, i am robot!")}
		cmd := decodeCmd(event.Command.Command)
		name := *event.IssueCommentEvent.GitHub.Repo.Name
		lnamespace := strings.Split(*event.IssueCommentEvent.GitHub.Repo.FullName, "/")
		if len(lnamespace) != 2 {
			return fmt.Errorf("get repo name failed:%s", *event.IssueCommentEvent.GitHub.Repo.FullName)
		}
		namespace := lnamespace[0]
		git.GitHub.Issues.CreateComment(context.Background(), namespace, name, *event.IssueCommentEvent.GitHub.Issue.Number, input)
		fmt.Println("promte info : ", namespace, name, cmd.Content, cmd.Params)
	} else {
		git := issue.NewGitClient(config)
		cmd := decodeCmd(event.Command.Command)
		name := event.IssueCommentEvent.GoGs.Repository.Name
		lnamespace := strings.Split(event.IssueCommentEvent.GoGs.Repository.FullName, "/")
		if len(lnamespace) != 2 {
			return fmt.Errorf("get repo name failed:%s", event.IssueCommentEvent.GoGs.Repository.FullName)
		}
		namespace := lnamespace[0]
		what, _ := d.AiChat(cmd.Content)
		input := gogs.CreateIssueCommentOption{Body: what}
		git.GoGs.CreateIssueComment(namespace, name, event.IssueCommentEvent.GoGs.Issue.Index, input)
		fmt.Println("promte info : ", namespace, name, cmd.Content, cmd.Params)
	}

	return nil
}

func decodeCmd(s string) *SayChatCmd {
	cmd := &SayChatCmd{}
	split := splitMultiBlank(s)
	if len(split) < 1 {
		return nil
	}

	cmd.Content = split[0]
	for _, p := range split[1:] {
		t := strings.Split(p, "=")
		if len(t) == 2 {
			if cmd.Params == nil {
				cmd.Params = make(map[string]string)
			}
			cmd.Params[t[0]] = t[1]
		}
	}
	return cmd
}

func splitMultiBlank(s string) []string {
	var res []string
	i, j := 0, 0
	for ; i < len(s) && j < len(s); {
		if s[i] == ' ' {
			i++
			j++
		}
		if j < len(s) && s[j] != ' ' {
			j++
			if j >= len(s) {
				res = append(res, s[i:j])
				break
			}
		} else if j > i {
			res = append(res, s[i:j])
			j++
			i = j
		}
	}
	return res
}
