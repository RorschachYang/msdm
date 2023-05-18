package tool

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(url, destination string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func DownloadImages() {
	// 读取 JSON 文件
	file, err := os.Open("./data/source/cardsdata.json") // 替换为你的 JSON 文件路径
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// 解码 JSON 数据
	var cardsdata []CardSource
	err = json.NewDecoder(file).Decode(&cardsdata)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 创建 cards 文件夹（如果不存在）
	err = os.MkdirAll("cardsimg", 0755)
	if err != nil {
		fmt.Println("Error creating 'cards' folder:", err)
		return
	}

	// 下载 src 字段对应的文件到 cards 文件夹
	for _, carddata := range cardsdata {
		err = downloadFile(carddata.Src, filepath.Join("cardsimg", filepath.Base(carddata.Src)))
		if err != nil {
			fmt.Println("Error downloading src file:", err)
			return
		}
	}

	// 创建 variants 文件夹（如果不存在）
	err = os.MkdirAll("variantsimg", 0755)
	if err != nil {
		fmt.Println("Error creating 'variants' folder:", err)
		return
	}

	// 下载 variants 字段对应的文件到 variants 文件夹
	for _, carddata := range cardsdata {
		for _, variant := range carddata.Variants {
			err = downloadFile(variant, filepath.Join("variantsimg", filepath.Base(variant)))
			if err != nil {
				fmt.Println("Error downloading variant file:", err)
				return
			}
		}
		if err != nil {
			fmt.Println("Error downloading src file:", err)
			return
		}
	}

	fmt.Println("Files downloaded successfully.")
}
