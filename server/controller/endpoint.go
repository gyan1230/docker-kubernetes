package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	err := os.Mkdir("serverdata", 0755)
	if err != nil {
		log.Println(err)
	}
}

//Home :
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "In Home"})
}

//SendInfo :
func SendInfo(c *gin.Context) {
	f, err := os.Create("serverdata/random.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString("new text to append\n"); err != nil {
		log.Println(err)
	}
	//Open a new hash interface to write to
	h := md5.New()
	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//Get the 16 bytes hash
	hashInBytes := h.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String := hex.EncodeToString(hashInBytes)

	fmt.Println("The hash is: ", returnMD5String)
	c.JSON(http.StatusOK, gin.H{"hash": returnMD5String})
}
