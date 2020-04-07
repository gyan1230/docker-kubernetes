package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Home :
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "In Home"})
}

//SendInfo :
func SendInfo(c *gin.Context) {

	//Open a new hash interface to write to
	h := md5.New()
	//Copy the file in the hash interface and check for any error

	//Get the 16 bytes hash
	hashInBytes := h.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String := hex.EncodeToString(hashInBytes)

	fmt.Println("The hash is: ", returnMD5String)
	c.JSON(http.StatusOK, gin.H{"hash": returnMD5String})
}
