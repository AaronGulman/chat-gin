package chat

type Message struct{
	Username string `json:"username"`
	Content string `json:"content"`
}


type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"` 
	Email    string `json:"email"`   
    }