package v1

import (
	"net/http"
	"sync"

	"github.com/RDelg/compare-encrypt-time/src/models"
	"github.com/RDelg/compare-encrypt-time/src/services"
	setting "github.com/RDelg/compare-encrypt-time/src/settings"
	"github.com/gin-gonic/gin"
)

func parseMessage(c *gin.Context) models.Message {
	var message models.Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return message
}

// RemoteFunctionAdapter
func RemoteFunctionAdapter(c *gin.Context) {
	message := parseMessage(c)
	if message.Context["action"] == "encrypt" {
		encrypt(message, c)
	} else if message.Context["action"] == "decrypt" {
		decrypt(message, c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "action not supported"})
	}

}

func encrypt(message models.Message, c *gin.Context) {

	svc := services.NewCypherService(setting.AppSetting.CypherSetting.Key, setting.AppSetting.CypherSetting.Iv)

	var wg sync.WaitGroup
	results := make([]string, len(message.Calls))

	for i, call := range message.Calls {
		wg.Add(1)
		go func(i int, call []string) {
			defer wg.Done()
			encrypted, err := svc.Encrypt(call[0])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			results[i] = encrypted
		}(i, call)
	}
	wg.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results})
}

func decrypt(message models.Message, c *gin.Context) {
	svc := services.NewCypherService(setting.AppSetting.CypherSetting.Key, setting.AppSetting.CypherSetting.Iv)

	var wg sync.WaitGroup
	results := make([]string, len(message.Calls))

	for i, call := range message.Calls {
		wg.Add(1)
		go func(i int, call []string) {
			defer wg.Done()
			decrypted, err := svc.Decrypt(call[0])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			results[i] = decrypted
		}(i, call)
	}
	wg.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results})

}

// Decrypt decrypts a message
func Decrypt(c *gin.Context) {
	message := parseMessage(c)
	decrypt(message, c)

}

// Encrypt
func Encrypt(c *gin.Context) {
	message := parseMessage(c)
	encrypt(message, c)
}
