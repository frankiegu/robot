package issue

import (
	"fmt"
	"os"
	"strings"

	"github.com/gogs/go-gogs-client"
	"github.com/google/go-github/github"
)

type IssueCommentEvent struct {
	GitHub github.IssueCommentEvent
	GoGs   gogs.IssueCommentPayload
}

type GitClient struct {
	GitHub *github.Client
	GoGs   *gogs.Client
}

type Command struct {
	Type    string
	Command string
}

type IssueEvent struct {
	*IssueCommentEvent
	Command *Command
	Client  *GitClient
}

var robot map[string]Robot

type Robot interface {
	Process(config Config, event IssueEvent) error
}

type Config struct {
	UserName string
	Password string
	Token    string
	Git      string
}

func NewConfig(user string, passwd string, git string) Config {
	if user == "" {
		if git == "github" {
			user = os.Getenv("GITHUB_USER")
		} else {
			user = os.Getenv("GOGS_URL")
		}
	}
	if passwd == "" {
		if git == "github" {
			passwd = os.Getenv("GITHUB_PASSWD")
		} else {
			passwd = os.Getenv("GOGS_TOKEN")
		}
	}
	return Config{UserName: user, Password: passwd, Git: git}
}

func NewGitClient(config Config) *GitClient {
	gitClient := new(GitClient)
	//GitHub
	tp := github.BasicAuthTransport{
		Username: config.UserName,
		Password: config.Password,
	}
	gitClient.GitHub = github.NewClient(tp.Client())
	// GoGs
	gitClient.GoGs = gogs.NewClient(config.UserName, config.Password)
	return gitClient
}

func Process(config Config, event IssueCommentEvent) error {
	commands := []*Command{}
	client := NewGitClient(config)
	//decode commands
	if config.Git == "github" {
		commands = decodeFromBody(event.GitHub.Comment.Body)
	} else {
		commands = decodeFromBody(&event.GoGs.Comment.Body)
	}

	fmt.Println("commands from body:", commands)

	for _, command := range commands {
		issueEvent := IssueEvent{
			&event,
			command,
			client,
		}
		fmt.Println("process command", command.Type, command.Command)
		if v, ok := robot[command.Type]; ok {
			v.Process(config, issueEvent)
		}
	}

	return nil
}

// Regist user robot
func Regist(command string, r Robot) {
	if robot == nil {
		robot = make(map[string]Robot)
	}
	robot[command] = r
}

func decodeFromBody(body *string) []*Command {
	var res []*Command
	lines := strings.Split(*body, "\n")
	for _, line := range lines {
		if !validCommand(line) {
			continue
		}
		res = append(res, decodeCommand(line))
	}
	return res
}

func validCommand(s string) bool {
	for _, b := range s {
		t := byte(b)
		if t != ' ' && t != '/' {
			return false
		}
		if t == '/' {
			return true
		}
		if t == ' ' {
			continue
		}
	}
	return false
}

// decode command
func decodeCommand(s string) *Command {
	command := &Command{}
	var i, j int
	fmt.Printf("decode cmd: %s\n", s)
	for i = range s {
		if byte(s[i]) == '/' {
			break
		}
	}
	var flag bool
	for j = i; j < len(s); j++ {
		if byte(s[j]) == ' ' {
			flag = true
			command.Type = s[i:j]
		}
		if flag && byte(s[j]) != ' ' {
			command.Command = s[j:]
			break
		}
	}
	return command
}
