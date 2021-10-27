package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"named-mps-gen/gen"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Option struct {
	Mode     string
	PathList []string
}

func InitCl() Option {
	opt := Option{}

	flag.StringVar(
		&opt.Mode,
		"m",
		"ignore",
		"fixを指定すると多目的スライダーに対応していない差分設定ファイルを修正します\n（全てのファイルを修正できるわけではありません）")
	flag.Parse()
	opt.PathList = flag.Args()

	return opt
}

func main() {
	opt := InitCl()

	exepath, _ := os.Executable()
	tmpdata, err := ioutil.ReadFile(filepath.Dir(exepath) + "/template.lua")
	if err != nil {
		println("テンプレートファイルを読み込めませんでした")
		return
	}
	tmptext := string(tmpdata)

	for _, path := range opt.PathList {
		fmt.Printf("%s\n", filepath.Base(path))

		// 読み込み
		data, err := ioutil.ReadFile(path)
		if err != nil {
			println("差分設定ファイルを読み込めませんでした")
			continue
		}
		text, _, _ := transform.String(japanese.ShiftJIS.NewDecoder(), string(data))

		// 解析
		deff := gen.GetDeffInfo(text)

		// 生成
		header := deff.ToTracksHeaderStr()
		body := strings.ReplaceAll(tmptext, "track_list", deff.ToTracksAccessStr())

		// 出力
		outfile, _ := os.Create(filepath.Dir(path) + "/mps-" + filepath.Base(path))
		defer outfile.Close()
		writer := bufio.NewWriter(transform.NewWriter(outfile, japanese.ShiftJIS.NewEncoder()))
		writer.WriteString(header + "\n\n" + body)
		writer.Flush()
	}
}
