package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/DylanRamsook/sprout-daily-tasks/pkg"
	"log"
	"time"
)

func main() {

	//These are the command line options you have to pass. Only url, email and password are required.
	urlPtr := flag.String("url", "https://<YOUR_COMPANY_HERE_PROBABLY.sproutatwork.com/auth/login_post", "Url for logging in... For me at least")
	emailPtr := flag.String("email", "SomePerson@someemail.com", "Email for logging into Sprout")
	passwordPtr := flag.String("password", "somePassword", "Password for logging into Sprout")
	rememberMePtr := flag.String("rememberMe", "true", "Remember user by default. String, not bool")
	dateTimePtr := flag.String("date", time.Now().Format("2006-01-02"), "Day to set activity to. Format is 2006-01-02")

	//Parse em
	flag.Parse()
	log.Printf(*dateTimePtr)
	//Try logging in with de-referenced pointers.
	resp, err := pkg.LoginToSprout(*emailPtr, *passwordPtr, *urlPtr, *rememberMePtr)
	if err != nil {
		log.Printf("Error authenticating, but I didn't test enough to have good error handling to tell you what went wrong.  Here is the Error message instead")
		log.Printf(err.Error())
	}

	log.Printf("Authenticated successfully.  Probably.")

	//We need sessionId, token, and remember me for the cookie.
	//Ex: sprout_session=Ch1ck3nT3nd1e5; remember_me=1; token=NiceTryGuy
	actResp, err := pkg.CreateSproutActivity(*urlPtr, *dateTimePtr, "90", resp.Cookies(), *rememberMePtr)
	if err != nil {
		log.Printf("Error creating activity, but I didn't test enough to have good error handling to tell you what went wrong.  Here is the Error message instead")
		log.Printf(err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(actResp.Body)
	newStr := buf.String()

	fmt.Printf("Response was" + newStr)

}
