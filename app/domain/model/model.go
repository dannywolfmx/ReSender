package model

type orm struct {
	ID string
	//	CreatedAt string
	//	UpdatedAt string

}

type Order struct {
	orm
	Number  string
	Invoice string
	//	Clients *Client
	//	Mails   []MailDirection
	//	Files   []File
}

type Client struct {
	orm
	Name string
}

type MailDirection struct {
	orm
	Direction string
}

type File struct {
	orm
	Path  string
	Title string
}
