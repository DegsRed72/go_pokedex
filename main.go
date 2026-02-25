package main

func main() {
	current_cfg := &config{Next: nil, Previous: nil}
	startRepl(current_cfg)
}
