# 概要
Go-Seleniumによるブラウザ操作
# 手順 (⚠️chromedriverとchromeのバージョンを一致させる)
`brew install chromedriver`
`go mod init goseleniumtest`
`go get github.com/sclevine/agouti`
`go mod tidy`
`go build ./`