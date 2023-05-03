package web

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"syscall"

	filex "github.com/liuxiaobopro/gobox/file"
)

func name(createProName string) {
	// 检查git是否安装
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("未安装git")
		return
	}

	gitProName := "qsgo-web-templete"
	fmt.Printf("开始创建项目: %s (如果失败,请多次尝试)\n", createProName)

	// 检查文件夹是否存在
	if _, err := os.Stat(gitProName); err == nil {
		// 存在,删除
		fmt.Println("删除文件夹: ", gitProName)
		os.RemoveAll(gitProName)
	}
	if _, err := os.Stat(createProName); err == nil {
		// 存在,删除
		fmt.Println("删除文件夹: ", createProName)
		os.RemoveAll(createProName)
	}

	// gitPath := fmt.Sprintf("http://github.com/liuxiaobopro/%s.git", gitProName)
	gitPath := fmt.Sprintf("http://gitee.com/liuxiaobopro/%s.git", gitProName)
	cmd := exec.Command("git", "clone", gitPath)

	fmt.Println("执行命令：", cmd.Args)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("执行命令时出错：", err)
		fmt.Println("标准错误信息：", stderr.String())
		return
	}

	// 修改clone下来的项目名
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("修改项目名: ", createProName)
		if err := syscall.Rename(gitProName, createProName); err != nil {
			fmt.Println("修改文件夹名时出错：", err)
			return
		}

		// 删除.git
		fmt.Println("删除.git文件夹")
		if err := os.RemoveAll(fmt.Sprintf("%s/.git", createProName)); err != nil {
			fmt.Println("删除.git文件夹时出错：", err)
			return
		}

		// 修改项目下所有文件内容包含gitProName改成arg
		fmt.Printf("修改项目下所有文件内容包含%s改成%s\n", gitProName, createProName)
		if err := filex.ReplaceInDir(createProName, gitProName, createProName); err != nil {
			fmt.Printf("修改项目下所有文件内容包含%s改成%s时出错：%s\n", gitProName, createProName, err)
			return
		}

		// git init && git add . && git commit -m "init"
		fmt.Println("git init && git add . && git commit -m \"init\"")
		cmd := exec.Command("git", "init")
		cmd.Dir = createProName
		if err := cmd.Run(); err != nil {
			fmt.Println("git init时出错：", err)
			return
		}

		cmd = exec.Command("git", "add", ".")
		cmd.Dir = createProName
		if err := cmd.Run(); err != nil {
			fmt.Println("git add .时出错：", err)
			return
		}

		cmd = exec.Command("git", "commit", "-m", "init")
		cmd.Dir = createProName
		if err := cmd.Run(); err != nil {
			fmt.Println("git commit -m \"init\"时出错：", err)
			return
		}

		fmt.Printf("项目%s创建成功\n", createProName)

		// go mod tidy
		fmt.Println("go mod tidy")
		cmd = exec.Command("go", "mod", "tidy")
		cmd.Dir = createProName
		if err := cmd.Run(); err != nil {
			fmt.Println("go mod tidy时出错：", err)
			return
		}
		fmt.Println("Done!")
	}()
	wg.Wait()
}
