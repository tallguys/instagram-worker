package crawler

import (
	"fmt"
	accountBusiness "instagram-worker/internal/app/model/business/account"
	accountPersist "instagram-worker/internal/app/model/persistence/account"
	"os"

	logger "github.com/sirupsen/logrus"
)

// Crawl crawl user
func Crawl(username string) error {
	account, err := accountPersist.FindAccountByUsername(username)
	if err != nil {
		switch err.(type) {
		case *accountPersist.AccountNotFoundError:
			err = accountPersist.InsertAccount(username)
			if err != nil {
				return err
			}
			account, err = accountPersist.FindAccountByUsername(username)
			if err != nil {
				return err
			}
			break
		default:
			return err
		}
	}

	pageInfo, err := accountBusiness.Request(username)
	if err != nil {
		return err
	}

	profile := pageInfo.SharedData.EntryData.ProfilePage[0].Graphql.User

	if profile.IsPrivate {
		return fmt.Errorf("The account %s is private", username)
	}

	err = fetchShortCodes(pageInfo, *account)
	if err != nil {
		return err
	}

	logger.Info("fetch shortcodes finish")

	// get the updated account info
	account, err = accountPersist.FindAccountByUsername(username)
	if err != nil {
		return err
	}

	dir := fmt.Sprintf("%s/%s", cfg.Download.Folder, username)
	err = os.MkdirAll(dir, 777)
	if err != nil {
		return err
	}

	shortcodes := account.Shortcodes

	maxConncurrency := cfg.Worker.Concurrency

	jobs := make(chan string, maxConncurrency)
	results := make(chan struct{}, len(shortcodes))

	for w := 0; w < maxConncurrency; w++ {
		go worker(username, account.ID, jobs, results)
	}

	for _, shortcode := range shortcodes {
		jobs <- shortcode
	}

	close(jobs)

	for range shortcodes {
		<-results
	}

	return nil
}

func worker(username string, accountID int, jobs <-chan string, results chan<- struct{}) {
	for shortcode := range jobs {
		err := process(username, shortcode, accountID)
		if err != nil {
			logger.Error(fmt.Sprintf("process %s unsuccessfully. Error: %v\n", shortcode, err))
		} else {
			logger.Info(fmt.Sprintf("process %s successfully\n", shortcode))
		}

		results <- struct{}{}
	}
}
