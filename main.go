package main

func main() {
	// err := godotenv.Load("env/local.env")
	// if err != nil {
	// 	panic("error loading .env file")
	// }
	router := newRouter()
	router.Logger.Fatal(router.Start(":8080"))
}
