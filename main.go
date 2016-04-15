package main

import (
	"flag"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const envPrefix = "DAVE_"

func strflag(name string, value string, usage string) *string {
	if v, ok := os.LookupEnv(envPrefix + strings.ToUpper(name)); ok {
		return flag.String(name, v, usage)
	}
	return flag.String(name, value, usage)
}

var (
	prefix *string
	owner  *string
	repo   *string
	token  *string
)

func main() {
	log.SetFlags(0)

	prefix = strflag("prefix", "release-", "branch name prefix")
	owner = strflag("owner", "VoIPGRID", "github owner to find repo under")
	repo = strflag("repo", "voipgrid", "github repository to bump version of")
	token = strflag("token", "", "github access token")
	flag.Parse()
	client := github.NewClient(oauth2.NewClient(oauth2.NoContext,
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: *token})))
	branchNames, err := currentBranchNames(client, *owner, *repo)
	if err != nil {
		log.Fatalf("error fetching branchnames: %v", err)
	}
	next := nextBranch(*prefix, branchNames...)
	if next == "" {
		log.Fatalf("no release branches found in %q", branchNames)
	}
	develop, _, err := client.Git.GetRef(*owner, *repo, "refs/heads/develop")
	if err != nil {
		log.Fatalf("error finding refs/heads/develop: %v", err)
	}
	r := "refs/heads/" + next
	ref := github.Reference{
		Ref:    &r,
		Object: develop.Object,
	}
	_, _, err = client.Git.CreateRef(*owner, *repo, &ref)
	if err != nil {
		log.Fatalf("error creating ref %q: %v", r, err)
	}
}

func currentBranchNames(client *github.Client, owner, repo string) (branchNames []string, err error) {
	opt := &github.ListOptions{}
	for {
		var (
			branches []github.Branch
			resp     *github.Response
		)
		branches, resp, err = client.Repositories.ListBranches(owner, repo, opt)
		if err != nil {
			return
		}
		for _, branch := range branches {
			branchNames = append(branchNames, *branch.Name)
		}
		if resp.NextPage == 0 {
			return
		}
		opt.Page = resp.NextPage
	}
}

func nextBranch(prefix string, branchNames ...string) string {
	var versions []*semver.Version
	for _, name := range branchNames {
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		v, err := semver.NewVersion(strings.TrimPrefix(name, prefix))
		if err != nil {
			continue
		}
		versions = append(versions, v)
	}
	if len(versions) < 1 {
		return ""
	}
	sort.Sort(semver.Versions(versions))
	v := versions[len(versions)-1]
	v.BumpMinor()
	return prefix + v.String()
}
