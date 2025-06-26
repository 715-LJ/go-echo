package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"github.com/unidoc/unioffice/document"
	"go-echo/internal/check"
	"log"
	"net/http"
)

type Word2pdfRequest struct {
	URL string `json:"url"`
}

type Word2pdfResponse struct {
	Response
}

func word2pdf(c *gin.Context) {
	fmt.Println("word2pdf...")

	var data Word2pdfRequest
	var resposn Word2pdfResponse
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(data)

	fileUrl := data.URL
	localDocxPath := "/Users/lijia/go/src/go-echo/downloaded_file.docx"
	pdfFilePath := "/Users/lijia/go/src/go-echo/converted_file.pdf"
	err := check.DownloadFile(fileUrl, localDocxPath)
	if err != nil {
		log.Fatal("Error downloading file:", err)
	}

	//第二步：将下载的 DOCX 文件转换为 PDF
	err = convertDocxToPDF(localDocxPath, pdfFilePath)
	if err != nil {
		log.Fatal("Error converting DOCX to PDF:", err)
	}

	// 第三步：删除本地 DOCX 文件
	//err = deleteFile(localDocxPath)
	//if err != nil {
	//	log.Fatal("Error deleting DOCX file:", err)
	//}

	resposn.Data = pdfFilePath
	c.IndentedJSON(http.StatusOK, resposn)
}

// 将 DOCX 文件转换为 PDF
func convertDocxToPDF(docxFilePath, pdfFilePath string) error {
	// 提取 .docx 文件中的文本
	_, err := extractTextFromDocx(docxFilePath)
	if err != nil {
		fmt.Println("Error extracting text from DOCX:", err)
		return err
	}

	// 创建 PDF 文件
	//err = createPDF(textContent, pdfFilePath)
	//if err != nil {
	//	fmt.Println("Error creating PDF:", err)
	//	return err
	//}
	return nil
}

// 从 .docx 文件中提取文本内容
func extractTextFromDocx(docxFilePath string) (string, error) {
	//// 打开 .docx 文件
	doc, err := document.Open(docxFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open DOCX file: %v", err)
	}
	defer func(doc *document.Document) {
		err := doc.Close()
		if err != nil {

		}
	}(doc)

	//// 提取所有段落文本
	var textContent string
	for _, para := range doc.Paragraphs() {
		// 提取每个段落中的文本
		for _, run := range para.Runs() {
			textContent += run.Text() // 获取每个run中的文本
		}
		textContent += "\n"
	}

	fmt.Println(textContent)
	return textContent, nil
}

// 使用 gofpdf 创建 PDF 文件
func createPDF(textContent, pdfFilePath string) error {
	// 创建 PDF 文档
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// 设置字体
	pdf.SetFont("Arial", "", 12)

	// 将 .docx 内容写入 PDF 中
	lines := splitTextIntoLines(textContent)
	for _, line := range lines {
		pdf.MultiCell(0, 10, line, "", "", false)
	}

	// 输出 PDF 文件
	err := pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		return fmt.Errorf("failed to create PDF: %v", err)
	}

	fmt.Println("PDF file created successfully:", pdfFilePath)
	return nil
}

// 分割文本内容为行
func splitTextIntoLines(text string) []string {
	var lines []string
	// 你可以根据需求设置文本行的最大长度
	maxLineLength := 80
	for len(text) > maxLineLength {
		// 找到空格位置分割行
		lineEnd := maxLineLength
		for lineEnd > 0 && text[lineEnd] != ' ' {
			lineEnd--
		}
		if lineEnd == 0 {
			// 如果没有空格则按最大长度分割
			lineEnd = maxLineLength
		}

		lines = append(lines, text[:lineEnd])
		text = text[lineEnd:]
	}
	lines = append(lines, text) // 添加剩余文本

	return lines
}

// 删除本地文件
//func deleteFile(filePath string) error {
//	err := os.Remove(filePath)
//	if err != nil {
//		return fmt.Errorf("failed to delete file: %v", err)
//	}
//
//	fmt.Println("File deleted successfully:", filePath)
//	return nil
//}
