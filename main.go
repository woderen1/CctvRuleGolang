package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("使用方法: ./main <输入文件> [输出文件]")
        fmt.Println("示例: ./main input.txt")
        fmt.Println("示例: ./main input.txt output.txt")
        return
    }
    
    // 输入和输出文件名
    inputFile := os.Args[1]
    
    // 如果没有指定输出文件，则使用默认命名规则
    outputFile := "modified_" + inputFile
    if len(os.Args) >= 3 {
        outputFile = os.Args[2]
    }
    // 打开原文件
    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Printf("无法打开文件 %s: %v\n", inputFile, err)
        return
    }
    defer file.Close()

    // 创建输出文件
    outFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Printf("无法创建文件 %s: %v\n", outputFile, err)
        return
    }
    defer outFile.Close()

    // 定义正则表达式模式
    // 非贪婪匹配，只替换到第一个匹配的查询参数
    pattern := regexp.MustCompile(`https://dh5wswx02\.v\.cntv\.cn/asp/h5e/hls/850/0303000a/3/default/[^?]+\.m3u8\?[^"]*`)
    
    // 使用bufio逐行读取和处理
    scanner := bufio.NewScanner(file)
    writer := bufio.NewWriter(outFile)
    lineCount := 0
    modifiedCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        lineCount++
        
        // 使用正则表达式进行替换
        if pattern.MatchString(line) {
            // 使用ReplaceAllStringFunc进行更精确的替换
            newLine := pattern.ReplaceAllStringFunc(line, func(match string) string {
                // 提取m3u8文件名之前的部分
                parts := strings.Split(match, "?")
                if len(parts) > 0 {
                    // 替换域名和路径
                    beforeQuery := parts[0]
                    beforeQuery = strings.Replace(beforeQuery, "dh5wswx02.v.cntv.cn", "hls.cntv.lxdns.com", 1)
                    beforeQuery = strings.Replace(beforeQuery, "/asp/h5e/hls", "/asp/hls", 1)
                    return beforeQuery
                }
                return match
            })
            
            writer.WriteString(newLine + "\n")
            modifiedCount++
        } else {
            writer.WriteString(line + "\n")
        }
    }

    writer.Flush()

    if err := scanner.Err(); err != nil {
        fmt.Printf("读取文件时出错: %v\n", err)
        return
    }

    fmt.Printf("处理完成！\n")
    fmt.Printf("总行数: %d\n", lineCount)
    fmt.Printf("修改行数: %d\n", modifiedCount)
    fmt.Printf("原文件: %s (未改动)\n", inputFile)
    fmt.Printf("新文件: %s\n", outputFile)
}
