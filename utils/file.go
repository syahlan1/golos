package utils

import (
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

func UploadFile(file *multipart.FileHeader, dir string) (filename, filepath string, err error) {

	uploadedFile, err := file.Open()
	if err != nil {
		log.Println("tes1")
		return "", "", err
	}
	defer uploadedFile.Close()

	filename = file.Filename
	trueFilename := CreateFilename(0, GetFileExtension(filename))
	filepath = dir + trueFilename

	uploadDir := "./" + dir
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Println("tes2")
		return "", "", err
	}

	destinationFile, err := os.Create(filepath)
	if err != nil {
		log.Println("tes3")
		return "", "", err
	}
	defer destinationFile.Close()

	if _, err := uploadedFile.Seek(0, 0); err != nil {
		log.Println("tes4")
		return "", "", err
	}
	if _, err := io.Copy(destinationFile, uploadedFile); err != nil {
		log.Println("tes5")
		return "", "", err
	}

	return trueFilename, "/"+filepath, nil
}

func GetFileExtension(filename string) (ext string) {
	extension := path.Ext(filename)
	extension = extension[1:]
	return extension
}

func RandomNum(min, max, i int) string {
	rand.Seed(time.Now().UnixNano() + int64(i))
	return strconv.Itoa(rand.Intn(max-min) + min)
}

func CreateFilename(i int, ext string) string {
	time := time.Now().Format("20060102150405")
	return time + RandomNum(1000000000000, 9999999999999, i) + "." + ext
}
