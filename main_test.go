package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {

	fmt.Println("run testrun")

	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	fmt.Println("run testrun222")

	in := "message"

	rsp, err := http.Get("http://localhost:18080/" + in)

	fmt.Println("run testrun333")

	if err != nil {
		t.Errorf("Failtd to get: %+v\n", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %+v\n", err)
	}

	// httpサーバーの戻り値を件処する
	want := fmt.Sprintf("Hello, %s!\n", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run関数に終了通知を送る
	cancel()

	// run関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
