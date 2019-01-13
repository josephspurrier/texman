package texman_test

import (
	"io/ioutil"
	"testing"

	"github.com/josephspurrier/texman"

	"github.com/stretchr/testify/assert"
)

const (
	testFile1 = "testdata/base.txt"
)

func matchFile(filename string) string {
	b, _ := ioutil.ReadFile(filename)
	return string(b)
}

func TestString(t *testing.T) {
	s := texman.NewFile("not-exist")
	err := s.Load()
	assert.NotNil(t, err)

	s = texman.NewFile(testFile1)
	err = s.Load()
	assert.Nil(t, err)

	b, err := ioutil.ReadFile(testFile1)
	assert.Nil(t, err)

	out := s.String()
	assert.Equal(t, string(b), out)

	outB := s.Byte()
	assert.Equal(t, b, outB)
}

func TestValidate(t *testing.T) {
	s := texman.NewFile(testFile1)
	err := s.Load()
	assert.Nil(t, err)

	for err := range []error{
		s.Insert(0, 1, "test"),
		s.Overwrite(1, -1, "test"),
		s.InsertLine(1, -1),
		s.DeleteLine(-1),
		s.Delete(-1, -1),
	} {
		assert.NotNil(t, err)
	}
}

func TestInsert(t *testing.T) {
	s := texman.NewFile(testFile1)
	err := s.Load()
	assert.Nil(t, err)

	err = s.Insert(1, 1, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/insert1.txt"), s.String())

	s.Load()

	err = s.Insert(2, 6, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/insert2.txt"), s.String())

	s.Load()

	err = s.Insert(2, 1, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/insert3.txt"), s.String())

	s.Load()

	err = s.Insert(7, 2, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/insert4.txt"), s.String())

	//ioutil.WriteFile("testdata/insert4.txt", []byte(s.String()), 0644)
}

func TestOverwrite(t *testing.T) {
	s := texman.NewFile(testFile1)
	err := s.Load()
	assert.Nil(t, err)

	err = s.Overwrite(1, 1, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/overwrite1.txt"), s.String())

	s.Load()

	err = s.Overwrite(2, 6, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/overwrite2.txt"), s.String())

	s.Load()

	err = s.Overwrite(2, 1, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/overwrite3.txt"), s.String())

	s.Load()

	err = s.Overwrite(7, 2, "Ab1")
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/overwrite4.txt"), s.String())

	//ioutil.WriteFile("testdata/overwrite4.txt", []byte(s.String()), 0644)
}

func TestDeleteLine(t *testing.T) {
	s := texman.NewFile(testFile1)
	err := s.Load()
	assert.Nil(t, err)

	err = s.DeleteLine(1)
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/deleteline1.txt"), s.String())

	s.Load()

	err = s.DeleteLine(2)
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/deleteline2.txt"), s.String())

	s.Load()

	err = s.DeleteLine(6)
	assert.NotNil(t, err)

	s.Load()

	err = s.DeleteLine(5)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	s := texman.NewFile(testFile1)
	err := s.Load()
	assert.Nil(t, err)

	err = s.Delete(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/delete1.txt"), s.String())

	s.Load()

	err = s.Delete(2, 6)
	assert.NotNil(t, err)

	s.Load()

	err = s.Delete(2, 1)
	assert.Nil(t, err)
	assert.Equal(t, matchFile("testdata/delete2.txt"), s.String())

	s.Load()

	err = s.Delete(7, 2)
	assert.NotNil(t, err)

	//ioutil.WriteFile("testdata/delete3.txt", []byte(s.String()), 0644)
}
