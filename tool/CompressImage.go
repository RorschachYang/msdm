package tool

// func CompressImages(dirPath string) {

// 	files, err := ioutil.ReadDir(dirPath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, file := range files {
// 		if file.IsDir() {
// 			continue // 忽略子目录
// 		}

// 		fileName := file.Name()
// 		filePath := filepath.Join(dirPath, fileName)

// 		compressImage(dirPath, fileName)

// 		if err != nil {
// 			fmt.Printf("Error compressing %s: %v\n", filePath, err)
// 		} else {
// 			fmt.Printf("%s compressed successfully\n", filePath)
// 		}
// 	}
// }

// func compressImage(dirPath string, fileName string) {
// 	filePath := dirPath + fileName

// 	// 打开 WebP 图片
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Println("图片打开失败" + filePath)
// 	}
// 	defer file.Close()

// 	// 解码 WebP 图片
// 	img, err := webp.Decode(file)
// 	if err != nil {
// 		fmt.Println("图片解码失败" + filePath)
// 	}

// 	// 如果宽度小于等于 200 像素，则不进行处理
// 	if img.Bounds().Dx() <= 150 {
// 		return
// 	}

// 	// 降低图片的分辨率
// 	newImg := resize.Resize(150, 0, img, resize.Lanczos3)

// 	// 确保目录存在
// 	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
// 		err := os.MkdirAll(dirPath, 0755)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	// 创建输出文件
// 	outputFile, err := os.Create(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer outputFile.Close()

// 	// 将处理后的图片保存为 WebP 格式
// 	var b bytes.Buffer
// 	if err := webp.Encode(&b, newImg, &webp.Options{Quality: 85}); err != nil {
// 		fmt.Println("图片保存失败" + filePath)
// 	}
// 	if _, err := outputFile.Write(b.Bytes()); err != nil {
// 		panic(err)
// 	}
// }
