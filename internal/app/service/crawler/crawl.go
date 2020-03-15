package crawler

import (
	"fmt"
	accountBusiness "instagram-worker/internal/app/model/business/account"
	accountPersist "instagram-worker/internal/app/model/persistence/account"
)

// Crawl crawl user
func Crawl(username string) error {
	account, err := accountPersist.FindAccountByUsername(username)
	if err != nil {
		return err
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

	fmt.Println("fetch shortcodes finsih")

	// get the updated account info
	account, err = accountPersist.FindAccountByUsername(username)

	shortcodes := account.Shortcodes

	for _, shortcode := range shortcodes {
		err := process(username, shortcode, account.ID)
		if err != nil {
			fmt.Printf("process %s unsuccessfully. Error: %v\n", shortcode, err)
		} else {
			fmt.Printf("process %s successfully\n", shortcode)
		}
	}

	return nil
}
