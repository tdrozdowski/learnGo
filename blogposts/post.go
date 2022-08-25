package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil
}
