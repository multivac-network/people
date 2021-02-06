package main

import "fmt"
import "regexp"
import "context"
import "github.com/vartanbeno/go-reddit/reddit"

func main() {
  ctx := context.Background()
  posts, _, err := reddit.DefaultClient().Subreddit.RisingPosts(ctx, "wallstreetbets", &reddit.ListOptions{
			Limit: 100,
    })
	if err != nil {
		return
	}
  for _, post := range posts {
    regex := regexp.MustCompile(`\s([\$A-Z]{2,5})\s`)
    matched := regex.MatchString(post.Body)
    if matched && post.Score >= 100 {
      vals := regex.FindAllString(post.Body, -1)
      for _, val := range vals {
        fmt.Printf("Match: %s\tScore: %d\t%s\n", val, post.Score, post.ID)
      }
    }
  }
	
}