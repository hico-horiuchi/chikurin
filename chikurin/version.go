package chikurin

import (
	"fmt"
	"time"

	latest "github.com/tcnksm/go-latest"
)

const TIMEOUT_SEC = 1

func verCheck(version string) <-chan *latest.CheckResponse {
	verCheckCh := make(chan *latest.CheckResponse)

	go func() {
		fixFunc := latest.DeleteFrontV()
		githubTag := &latest.GithubTag{
			Owner:             "hico-horiuchi",
			Repository:        "chikurin",
			FixVersionStrFunc: fixFunc,
		}

		res, _ := latest.Check(githubTag, fixFunc(version))
		verCheckCh <- res
	}()

	return verCheckCh
}

func Version(version string) string {
	var print []byte
	print = append(print, fmt.Sprintf("chikurin version %s\n", version)...)
	verCheckCh := verCheck(version)

	for {
		select {
		case res := <-verCheckCh:
			if res != nil && res.Outdated {
				print = append(print, fmt.Sprintf("Latest version of chikurin is %s, please update it\n", res.Current)...)
			}
			return string(print)
		case <-time.After(TIMEOUT_SEC * time.Second):
			return string(print)
		}
	}
}
