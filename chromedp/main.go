package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "time"

    cdp "github.com/knq/chromedp"
    cdptypes "github.com/knq/chromedp/cdp"
)

func main() {
    var err error

    // create context
    ctxt, cancel := context.WithCancel(context.Background())
    defer cancel()

    // create chrome instance
    c, err := cdp.New(ctxt)
    if err != nil {
        log.Fatal(err)
    }

    // run task list
    var site, res string
    err = c.Run(ctxt, googleSearch("site:plod.tv", "Year of the runs", &site, &res))
    if err != nil {
        log.Fatal(err)
    }

    // shutdown chrome
    err = c.Shutdown(ctxt)
    if err != nil {
        log.Fatal(err)
    }

    // wait for chrome to finish
    err = c.Wait()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("saved screenshot of #testimonials from search result listing `%s` (%s)", res, site)
}

func googleSearch(q, text string, site, res *string) cdp.Tasks {
    var buf []byte
    sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
    return cdp.Tasks{
        cdp.Navigate(`https://www.google.com`),
        cdp.Sleep(2 * time.Second),
        cdp.WaitVisible(`#hplogo`, cdp.ByID),
        cdp.SendKeys(`#lst-ib`, q, cdp.ByID),
        cdp.Click(`input[name="btnK"]`, cdp.ByQuery),
        cdp.WaitNotVisible(`input[name="btnI"]`, cdp.ByQuery),
        cdp.Text(sel, res),
        cdp.Click(sel),
        cdp.Sleep(2 * time.Second),
        cdp.WaitVisible(`#footer`, cdp.ByQuery),
        cdp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, cdp.ByQuery),
        cdp.Location(site),
        cdp.Screenshot(`#testimonials`, &buf, cdp.ByID),
        cdp.ActionFunc(func(context.Context, cdptypes.FrameHandler) error {
            return ioutil.WriteFile("testimonials.png", buf, 0644)
        }),
    }
}
