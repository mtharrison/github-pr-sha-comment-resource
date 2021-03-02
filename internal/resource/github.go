package resource

import (
	"context"
	"errors"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

func GetShaFromDir(dir string) (string, error) {

	r, err := git.PlainOpen(dir)
	if err != nil {
		return "", err
	}

	h, err := r.ResolveRevision(plumbing.Revision("HEAD"))
	if err != nil {
		return "", err
	}

	return h.String(), nil
}

func GetPrNumberFromSha(input Input, sha string) (int, error) {

	client, err := githubClient(input.Source.AccessToken, input.Source.V3Endpoint)
	if err != nil {
		return -1, err
	}

	ctx := context.Background()

	pulls, _, err := client.PullRequests.ListPullRequestsWithCommit(ctx, input.Source.Owner(), input.Source.Repo(), sha, nil)
	if err != nil {
		return -1, err
	}

	if len(pulls) < 1 {
		return -1, errors.New("No pulls found for SHA")
	}

	return *pulls[0].Number, nil
}

func FormatComment(comment string) string {

	return safeExpandEnv(comment)
}

func PostComment(input Input, comment string, pr int) (string, error) {

	client, err := githubClient(input.Source.AccessToken, input.Source.V3Endpoint)
	if err != nil {
		return "", err
	}

	ctx := context.Background()

	issueComment := new(github.IssueComment)
	issueComment.Body = &comment
	res, _, err := client.Issues.CreateComment(ctx, input.Source.Owner(), input.Source.Repo(), pr, issueComment)
	if err != nil {
		return "", err
	}

	return *res.HTMLURL, nil
}

func safeExpandEnv(s string) string {
	return os.Expand(s, func(v string) string {
		switch v {
		case "BUILD_ID", "BUILD_NAME", "BUILD_JOB_NAME", "BUILD_PIPELINE_NAME", "BUILD_TEAM_NAME", "ATC_EXTERNAL_URL":
			return os.Getenv(v)
		}
		return "$" + v
	})
}

func githubClient(token, endpoint string) (*github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	if endpoint == "" {
		return github.NewClient(tc), nil
	}

	client, err := github.NewEnterpriseClient(endpoint, endpoint, tc)
	if err != nil {
		return nil, err
	}
	return client, nil
}
