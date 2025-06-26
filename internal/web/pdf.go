package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
	"net/http"
)

type Merge2pdfRequest struct {
	URL []string `json:"url"`
}

type Merge2pdfResponse struct {
	Response
}

func merge2pdf(c *gin.Context) {
	fmt.Println("merge2pdf...")

	var data Merge2pdfRequest
	var resposn Merge2pdfResponse
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(data.URL)

	var filePaths []string
	//for _, fileUrl := range data.URL {
	//	localDocxPath := "/Users/lijia/go/src/go-echo/" + strconv.FormatInt(time.Now().Unix(), 10) + "_" + filepath.Base(fileUrl)
	//	err := check.DownloadFile(fileUrl, localDocxPath)
	//	if err != nil {
	//		log.Fatal("Error downloading file:", err)
	//	}
	//	filePaths = append(filePaths, localDocxPath)
	//}

	filePaths = []string{
		"/Users/lijia/go/src/go-echo/1749719598_Manuscript_File_v1.pdf",
		"/Users/lijia/go/src/go-echo/1749719599_JCA-2024-24-Manuscript_File_v1.pdf",
		"/Users/lijia/go/src/go-echo/WECN-2024-17-Manuscript-File.v1.pdf",
	}

	fmt.Println(filePaths)
	outputFile := "/Users/lijia/go/src/go-echo/converted_file.pdf"
	err := mergePDFs(filePaths, outputFile)
	if err != nil {
		log.Fatalf("合并 PDF 文件失败: %v\n", err)
	} else {
		fmt.Println("PDF 文件合并成功:", outputFile)
	}
	resposn.Response.Data = outputFile
	c.IndentedJSON(http.StatusOK, resposn)
}

func mergePDFs(inputFiles []string, outputFile string) error {
	// 合并多个 PDF 文件
	err := api.MergeCreateFile(inputFiles, outputFile, false, nil)
	if err != nil {
		return fmt.Errorf("合并文件失败: %v", err)
	}
	return nil
}
