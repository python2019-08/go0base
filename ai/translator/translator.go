package translator

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func TransLator_main() {
	fmt.Println("TransLator_main:1")
	r := gin.Default()

	v1 := r.Group("/api/v1")

	fmt.Println("TransLator_main:2")

	// http://localhost:8088/api/v1/translate
	v1.POST("/translate", translator)

	r.Run(":8088")
}

func translator(c *gin.Context) {
	var requestData struct {
		OutputLang string `"json:"outputLang"`
		Text       string `json:"text"`
	}
	fmt.Println("ai/translator/translator.go:1")
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json."})
		return
	}

	fmt.Println("ai/translator/translator.go:2")
	//创建prompt
	prompt := prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		prompts.NewSystemMessagePromptTemplate("你是一个只能翻译文本的翻译引擎，不需要进行解释。", nil),
		prompts.NewHumanMessagePromptTemplate(`翻译这段文字到{{.outputLang}}：{{.text}}`,
			[]string{"outputLang", "text"}),
	})

	// 填充prompt
	vals := map[string]any{
		"outputLang": requestData.OutputLang,
		"text":       requestData.Text,
	}

	messages, err := prompt.FormatMessages(vals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//连接olLama
	llm, err := ollama.New(ollama.WithModel("modelscope.cn/unsloth/DeepSeek-R1-Distill-Qwen-1.5B-GGUF"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := []llms.MessageContent{
		llms.TextParts(messages[0].GetType(), messages[0].GetContent()),
		llms.TextParts(messages[1].GetType(), messages[1].GetContent()),
	}

	response, err := llm.GenerateContent(context.Background(), content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response.Choices[0].Content})
}
