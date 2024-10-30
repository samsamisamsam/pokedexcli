package main

import "os"

func exitCallback(cfg *config, s ...string) error {
	os.Exit(0)
	return nil
}
