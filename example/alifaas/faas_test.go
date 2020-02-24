package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_promoteEvent(t *testing.T) {
	body := []byte(`{
  "action": "created",
  "issue": {
    "url": "https://api.github.com/repos/sqeven/testgit/issues/1",
    "repository_url": "https://api.github.com/repos/sqeven/testgit",
    "labels_url": "https://api.github.com/repos/sqeven/testgit/issues/1/labels{/name}",
    "comments_url": "https://api.github.com/repos/sqeven/testgit/issues/1/comments",
    "events_url": "https://api.github.com/repos/sqeven/testgit/issues/1/events",
    "html_url": "https://github.com/sqeven/testgit/issues/1",
    "id": 568760388,
    "node_id": "MDU6SXNzdWU1Njg3NjAzODg=",
    "number": 1,
    "title": "test",
    "user": {
      "login": "sqeven",
      "id": 18329103,
      "node_id": "MDQ6VXNlcjE4MzI5MTAz",
      "avatar_url": "https://avatars0.githubusercontent.com/u/18329103?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/sqeven",
      "html_url": "https://github.com/sqeven",
      "followers_url": "https://api.github.com/users/sqeven/followers",
      "following_url": "https://api.github.com/users/sqeven/following{/other_user}",
      "gists_url": "https://api.github.com/users/sqeven/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/sqeven/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/sqeven/subscriptions",
      "organizations_url": "https://api.github.com/users/sqeven/orgs",
      "repos_url": "https://api.github.com/users/sqeven/repos",
      "events_url": "https://api.github.com/users/sqeven/events{/privacy}",
      "received_events_url": "https://api.github.com/users/sqeven/received_events",
      "type": "User",
      "site_admin": false
    },
    "labels": [

    ],
    "state": "open",
    "locked": false,
    "assignee": null,
    "assignees": [

    ],
    "milestone": null,
    "comments": 2,
    "created_at": "2020-02-21T06:23:55Z",
    "updated_at": "2020-02-21T07:21:29Z",
    "closed_at": null,
    "author_association": "OWNER",
    "body": ""
  },
  "comment": {
    "url": "https://api.github.com/repos/sqeven/testgit/issues/comments/589530950",
    "html_url": "https://github.com/sqeven/testgit/issues/1#issuecomment-589530950",
    "issue_url": "https://api.github.com/repos/sqeven/testgit/issues/1",
    "id": 589530950,
    "node_id": "MDEyOklzc3VlQ29tbWVudDU4OTUzMDk1MA==",
    "user": {
      "login": "sqeven",
      "id": 18329103,
      "node_id": "MDQ6VXNlcjE4MzI5MTAz",
      "avatar_url": "https://avatars0.githubusercontent.com/u/18329103?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/sqeven",
      "html_url": "https://github.com/sqeven",
      "followers_url": "https://api.github.com/users/sqeven/followers",
      "following_url": "https://api.github.com/users/sqeven/following{/other_user}",
      "gists_url": "https://api.github.com/users/sqeven/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/sqeven/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/sqeven/subscriptions",
      "organizations_url": "https://api.github.com/users/sqeven/orgs",
      "repos_url": "https://api.github.com/users/sqeven/repos",
      "events_url": "https://api.github.com/users/sqeven/events{/privacy}",
      "received_events_url": "https://api.github.com/users/sqeven/received_events",
      "type": "User",
      "site_admin": false
    },
    "created_at": "2020-02-21T07:21:29Z",
    "updated_at": "2020-02-21T07:21:29Z",
    "author_association": "OWNER",
    "body": "/say test"
  },
  "repository": {
    "id": 62297584,
    "node_id": "MDEwOlJlcG9zaXRvcnk2MjI5NzU4NA==",
    "name": "testgit",
    "full_name": "sqeven/testgit",
    "private": false,
    "owner": {
      "login": "sqeven",
      "id": 18329103,
      "node_id": "MDQ6VXNlcjE4MzI5MTAz",
      "avatar_url": "https://avatars0.githubusercontent.com/u/18329103?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/sqeven",
      "html_url": "https://github.com/sqeven",
      "followers_url": "https://api.github.com/users/sqeven/followers",
      "following_url": "https://api.github.com/users/sqeven/following{/other_user}",
      "gists_url": "https://api.github.com/users/sqeven/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/sqeven/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/sqeven/subscriptions",
      "organizations_url": "https://api.github.com/users/sqeven/orgs",
      "repos_url": "https://api.github.com/users/sqeven/repos",
      "events_url": "https://api.github.com/users/sqeven/events{/privacy}",
      "received_events_url": "https://api.github.com/users/sqeven/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/sqeven/testgit",
    "description": null,
    "fork": false,
    "url": "https://api.github.com/repos/sqeven/testgit",
    "forks_url": "https://api.github.com/repos/sqeven/testgit/forks",
    "keys_url": "https://api.github.com/repos/sqeven/testgit/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/sqeven/testgit/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/sqeven/testgit/teams",
    "hooks_url": "https://api.github.com/repos/sqeven/testgit/hooks",
    "issue_events_url": "https://api.github.com/repos/sqeven/testgit/issues/events{/number}",
    "events_url": "https://api.github.com/repos/sqeven/testgit/events",
    "assignees_url": "https://api.github.com/repos/sqeven/testgit/assignees{/user}",
    "branches_url": "https://api.github.com/repos/sqeven/testgit/branches{/branch}",
    "tags_url": "https://api.github.com/repos/sqeven/testgit/tags",
    "blobs_url": "https://api.github.com/repos/sqeven/testgit/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/sqeven/testgit/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/sqeven/testgit/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/sqeven/testgit/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/sqeven/testgit/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/sqeven/testgit/languages",
    "stargazers_url": "https://api.github.com/repos/sqeven/testgit/stargazers",
    "contributors_url": "https://api.github.com/repos/sqeven/testgit/contributors",
    "subscribers_url": "https://api.github.com/repos/sqeven/testgit/subscribers",
    "subscription_url": "https://api.github.com/repos/sqeven/testgit/subscription",
    "commits_url": "https://api.github.com/repos/sqeven/testgit/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/sqeven/testgit/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/sqeven/testgit/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/sqeven/testgit/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/sqeven/testgit/contents/{+path}",
    "compare_url": "https://api.github.com/repos/sqeven/testgit/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/sqeven/testgit/merges",
    "archive_url": "https://api.github.com/repos/sqeven/testgit/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/sqeven/testgit/downloads",
    "issues_url": "https://api.github.com/repos/sqeven/testgit/issues{/number}",
    "pulls_url": "https://api.github.com/repos/sqeven/testgit/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/sqeven/testgit/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/sqeven/testgit/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/sqeven/testgit/labels{/name}",
    "releases_url": "https://api.github.com/repos/sqeven/testgit/releases{/id}",
    "deployments_url": "https://api.github.com/repos/sqeven/testgit/deployments",
    "created_at": "2016-06-30T09:21:32Z",
    "updated_at": "2020-02-21T06:09:27Z",
    "pushed_at": "2020-02-21T06:09:25Z",
    "git_url": "git://github.com/sqeven/testgit.git",
    "ssh_url": "git@github.com:sqeven/testgit.git",
    "clone_url": "https://github.com/sqeven/testgit.git",
    "svn_url": "https://github.com/sqeven/testgit",
    "homepage": null,
    "size": 0,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 1,
    "license": null,
    "forks": 0,
    "open_issues": 1,
    "watchers": 0,
    "default_branch": "master"
  },
  "sender": {
    "login": "sqeven",
    "id": 18329103,
    "node_id": "MDQ6VXNlcjE4MzI5MTAz",
    "avatar_url": "https://avatars0.githubusercontent.com/u/18329103?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/sqeven",
    "html_url": "https://github.com/sqeven",
    "followers_url": "https://api.github.com/users/sqeven/followers",
    "following_url": "https://api.github.com/users/sqeven/following{/other_user}",
    "gists_url": "https://api.github.com/users/sqeven/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/sqeven/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/sqeven/subscriptions",
    "organizations_url": "https://api.github.com/users/sqeven/orgs",
    "repos_url": "https://api.github.com/users/sqeven/repos",
    "events_url": "https://api.github.com/users/sqeven/events{/privacy}",
    "received_events_url": "https://api.github.com/users/sqeven/received_events",
    "type": "User",
    "site_admin": false
  }
}`)
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test event decode json",
			args{body},
		},
	}
	for _, tt := range tests {
		fmt.Println(os.Getenv("GITHUB_USER"))
		e := promoteEvent(tt.args.body)
		if *e.GitHub.Repo.FullName != "sqeven/testgit" {
			t.Errorf("%s", *e.GitHub.Repo.FullName)
		}
	}
}
