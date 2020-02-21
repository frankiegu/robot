package drone_promote

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/sqeven/robot/issue"
	"golang.org/x/oauth2"
)

/*
QSp93SmhZVpJAmb7tWPuWIO
Example API Usage:
curl -i https://cloud.drone.io/api/user \
-H "Authorization: Bearer QSp93SmhZVpJAmb7tWP"
Example CLI Usage:
export DRONE_SERVER=https://cloud.drone.io
export DRONE_TOKEN=QSp93SmhZVpJAmb7tWPuW
drone info
*/
type DronePromote struct {
	DroneServer string
	DroneToken  string
}

type DronePromoteCmd struct {
	Build int
	Target string
	Params map[string]string
}

// process command like /promote 42 test key=value
func (d *DronePromote) Process(event issue.IssueEvent) error {
	d.env()
	// new drone client
	client := *d.client()
	// decode command
	cmd := decodeCmd(event.Command.Command)
	name := *event.IssueCommentEvent.Repo.Name
	lnamespace := strings.Split(*event.IssueCommentEvent.Repo.FullName,"/")
	if len(lnamespace) != 2 {
		return fmt.Errorf("get repo name failed:%s",*event.IssueCommentEvent.Repo.FullName)
	}
	namespace := lnamespace[0]
	fmt.Println("promte info : ",namespace,name,cmd.Build,cmd.Target,cmd.Params)
	// drone promote
	_,err := client.Promote(namespace,name,cmd.Build,cmd.Target,cmd.Params)
	if err != nil {
		fmt.Errorf("promote failed : %s\n",err)
		return err
	}
	return nil
}

// 42 test key=value
// Build:42
// Target:test
// Para:key:value
// Para is optional
func decodeCmd(s string) *DronePromoteCmd{
	var err error
	cmd := &DronePromoteCmd{}
	split := splitMultiBlank(s)
	if len(split) < 2 {
		return nil
	}
	cmd.Build,err = strconv.Atoi(split[0])
	if err != nil {
		return nil
	}
	cmd.Target = split[1]
	for _,p := range split[2:] {
		t := strings.Split(p,"=")
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
	i,j := 0,0
	for ;i<len(s) && j <len(s); {
		if s[i] == ' ' {
			i++
			j++
		}
		if j <len(s) && s[j] != ' ' {
			j++
			if j >= len(s) {
				res = append(res, s[i:j])
				break
			}
		} else if j > i {
			res = append(res, s[i:j])
			j++
			i=j
		}
	}
	return res
}

func (d *DronePromote) client() *drone.Client {
	config := new(oauth2.Config)
	auther := config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: d.DroneToken,
		},
	)

	// create the drone client with authenticator
	client := drone.NewClient(d.DroneServer, auther)
	return &client
}

func (d *DronePromote) env() {
	if d.DroneServer == "" {
		d.DroneServer = os.Getenv("DRONE_SERVER")
	}
	if d.DroneToken == "" {
		d.DroneToken = os.Getenv("DRONE_TOKEN")
	}
}
