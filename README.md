# ROT
go製のコマンドラインで動くRot13変換器です

## Install
`git clone` して `go install` or `go build`

## option
`-n N`

値を設定することでアルファベットをずらす数を変更することができます

Example:
    rot -n 15 とするとROT15で変換を行います

`-d`
`--decode`

ROT13のときのみ正しく逆変換ができます
        
## TODO
~~-d Decodeの実装~~

標準入力からだけでなく、コマンドライン引数からも文字列を読み込めるようにする

