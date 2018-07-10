/*
データをtemplateに渡す方法に関する
*/

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

/* 初期化関数 */
func init() {

	/* index.gohtmlファイルをtemplateタイプに変換し、Must関数でwrapsする*/
	tpl = template.Must(template.ParseFiles("./01_template/index.gohtml"))
}

/* MAIN */
func main() {
	/* 指定したtemplateにデータを渡す*/
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "MAI")

	/* エラーの場合 */
	if err != nil {
		/* LOG出力 */
		log.Fatalln(err)
	}
}
