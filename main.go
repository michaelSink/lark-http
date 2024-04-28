package lark

func main() {

	server := new(Server)
	server.Address = "localhost:42069"

	server.ListenAndServe()

	return
}
