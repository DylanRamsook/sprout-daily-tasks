# Sprout Activity Recording
Creates a Sprout task for you given your email, sprout url, and password for sprout.  Meant to be used with a scheduler
like cron because I always forget to submit Sprout stuff.  You can probably just use task manager for Windows.   

### Prerequisites

Something that can run Go.  https://golang.org/doc/install

### Installing

Get + build the project 

```
go get -u github.com/DylanRamsook/sprout-daily-tasks/cmd
```

Test by running "cmd"

```
/Users/dylanr/go/bin/cmd -url https://SOMECOMPANY.sproutatwork.com -email Dylanator@something -password D03440efjalskh
```

Set up your cronjob 

```
crontab -e 
Type "i" to begin editing.
0 10 * * * /Users/dylanr/go/bin/cmd -url https://SOMECOMPANY.sproutatwork.com -email Dylanator@something -password D03440efjalskh
"esc" + Type ":wq" to Save and Quit.
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

##Notes I used  GoLang version 1.14beta1


