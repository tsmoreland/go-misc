package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"sync"
)

func BadWaitGroup(files []string) {
	var wg sync.WaitGroup

	for _, file := range files {
		path := file
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := os.ReadFile(path)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(len(data))
			}
		}()
	}

	wg.Wait()
}

func ErrGroup(files []string) {
	var eg errgroup.Group
	for _, file := range files {
		path := file
		eg.Go(func() error {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			log.Println(len(data))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Print(err)
	}
}

func ErrGroupWithContext(ctx context.Context, files []string) {
	eg, ctx := errgroup.WithContext(ctx)
	for _, file := range files {
		path := file
		eg.Go(func() error {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			select {
			case <-ctx.Done():
				log.Print(ctx.Err())
				return ctx.Err()
			default:
				log.Println(len(data))
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Print(err)
	}
}
