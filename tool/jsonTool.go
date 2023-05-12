package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// toJSON 将任何结构体转换为 JSON 格式并写入文件
func ToJSON(obj interface{}, filename string, dirPath string) error {
	// 将对象转换为 JSON 格式
	jsonBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal object to JSON: %s", err)
	}

	// 确保目录存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %s", err)
		}
	}

	// 将 JSON 写入文件
	filePath := fmt.Sprintf("%s/%s", dirPath, filename)
	err = ioutil.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %s", err)
	}

	log.Println("JSON file written successfully")

	return nil
}

func DeleteKey(filename string, dirPath string, ketToDelete string) {
	filepath := dirPath + filename
	// 读取JSON文件
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	var obj interface{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	deleteKey(obj, ketToDelete)

	ToJSON(obj, filename, dirPath)
}

func deleteKey(obj interface{}, key string) {
	switch v := obj.(type) {
	case map[string]interface{}:
		// 如果是map类型，则递归删除子节点
		for k := range v {
			if k == key {
				delete(v, k)
			} else {
				deleteKey(v[k], key)
			}
		}
	case []interface{}:
		// 如果是数组类型，则递归删除子节点
		for i := range v {
			deleteKey(v[i], key)
		}
	}
}
