package say_chat

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/sqeven/robot/issue"
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

func (d *SayChat) Process(event issue.IssueEvent) error {
	config := issue.NewConfig("", "")
	git := issue.NewGitClient(config)
	input := &github.IssueComment{Body: github.String(":robot:Hi, i am robot!")}
	cmd := decodeCmd(event.Command.Command)
	name := *event.IssueCommentEvent.Repo.Name
	lnamespace := strings.Split(*event.IssueCommentEvent.Repo.FullName, "/")
	if len(lnamespace) != 2 {
		return fmt.Errorf("get repo name failed:%s", *event.IssueCommentEvent.Repo.FullName)
	}
	namespace := lnamespace[0]
	git.Issues.CreateComment(context.Background(), namespace, name, *event.IssueCommentEvent.Issue.Number, input)
	fmt.Println("promte info : ", namespace, name, cmd.Content, cmd.Params)

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
