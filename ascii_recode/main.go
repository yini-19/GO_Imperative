package main

func main() {
	input := "Hello, World!"
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		panic(err)
	}

	art := GenerateArt(input, banner)
	println(art)
}
