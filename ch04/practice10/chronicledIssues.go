package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	current := time.Now()
	chronicledConst := make(map[int][]github.Issue)

	for _, item := range result.Items {
		diff := current.Sub(item.CreatedAt)
		diffHours := diff.Hours()

		if diffHours < 731 { //現在時刻との差分が1ヶ月（約730時間以内）の場合。もっといい方法はあるのだろうけどとりあえずこれ。。
			pattern := 0
			chronicledConst[pattern] = append(chronicledConst[pattern], *item)
		} else if diffHours < 8761 { //現在時刻との差分が1年（8760時間以内）の場合/
			pattern := 1
			chronicledConst[pattern] = append(chronicledConst[pattern], *item)
		} else {
			pattern := 2
			chronicledConst[pattern] = append(chronicledConst[pattern], *item)
		}
	}

	for k, v := range chronicledConst {
		if k == 0 {
			fmt.Println("Within 1month since cases bacame open")
			for _, i := range v {
				fmt.Printf("%s #%-5d %9.9s %.55s\n", i.CreatedAt.Format("2006-01-02"), i.Number, i.User.Login, i.Title)
			}
		} else if k == 1 {
			fmt.Println("Within 1year since cases bacame open")
			for _, i := range v {
				fmt.Printf("%s #%-5d %9.9s %.55s\n", i.CreatedAt.Format("2006-01-02"), i.Number, i.User.Login, i.Title)
			}
		} else {
			fmt.Println("Over 1year since cases bacame open")
			for _, i := range v {
				fmt.Printf("%s #%-5d %9.9s %.55s\n", i.CreatedAt.Format("2006-01-02"), i.Number, i.User.Login, i.Title)
			}
		}
	}
}
