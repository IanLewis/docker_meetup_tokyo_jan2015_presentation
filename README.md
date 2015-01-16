#  Docker Meetup Tokyo Jan 2015 Lightning Talk Presentation

## この資料について

僕の資料はGoの[present](https://godoc.org/golang.org/x/tools/cmd/present)というツールでできている。

presentを使うと、資料の中にGoコードを動かすことができる。 
でも、[go-dockerclient](https://github.com/fsouza/go-dockerclient/)と連動するために、
dockerの環境が必要なので、すこし面倒。 そのために、このVagrant Boxを作りました。

# 動的バージョン

プレゼンテーションを動かせる環境を作るために、

    $ vagrant up

を実行するだけ。

サーバーの起動が終わるまで、しばらく時間かかるけど、終わったら、
http://localhost:3999/lt.slide で資料を見ることができる。

# 性的バージョン

すぐ見れるけど、コード動かすことができないバージョンも用意しています。

https://ianlewis.github.io/docker_meetup_tokyo_jan2015_presentation
