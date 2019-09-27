package chrome

import (
	"context"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func Run(task chromedp.Tasks) error {
	opts := make([]chromedp.ExecAllocatorOption, 0)
	opts = append(opts, chromedp.Flag("headless", true))
	opts = append(opts, chromedp.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"))
	allocatorCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocatorCtx, chromedp.WithLogf(log.Printf))
	err := chromedp.Run(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
func GetText(doc *goquery.Document, path string) string {
	return strings.TrimSpace(strings.Replace(doc.Find(path).Text(), "\n", "", -1))
}
