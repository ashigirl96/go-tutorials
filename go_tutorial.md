# Go

## [Goコードの書き方](http://golang-jp.org/doc/code.html#PackageNames)

### ワークスペース

- `src`はパッケージに構成された(1ディレクトリに1パッケージ)Goのソースコードを含み,
- `pkg`はパッケージオブジェクトを含み,
- `bin`は実行形式コマンドを含む.

`go`ツールはソースパッケージをビルとし, その結果生成されるバイナリを`pkg`および`bin`にインストールする.

```go
bin/
    streak                         # コマンド実行形式ファイル
    todo                           # コマンド実行形式ファイル
pkg/
    linux_amd64/
        code.google.com/p/goauth2/
            oauth.a                # パッケージオブジェクト
        github.com/nf/todo/
            task.a                 # パッケージオブジェクト
src/
    code.google.com/p/goauth2/
        .hg/                       # mercurialレポジトリメタデータ
        oauth/
            oauth.go               # パッケージソース
            oauth_test.go          # テストソース
    github.com/nf/
        streak/
            .git/                  # gitレポジトリメタデータ
            oauth.go               # コマンドソース
            streak.go              # コマンドソース
        todo/
            .git/                  # gitレポジトリメタデータ
            task/
                task.go            # パッケージソース
            todo.go                # コマンドソース
```
            
は, 2つのコマンド(`streak & todo`)            と2つのライブラリ(`oauth`&`task`)からなる3つのレポジトリ(`goauth2`, `streak`, `todo`)が入ってます。

### GOPATH環境変数

ワークスペースの場所を示す.

始めるにあたって, ワークスペースディレクトリを作成し, GOPATHをセットする.

`$ export PATH=$PATH:$GOPATH/bin`

### パッケージパス

標準ライブラリ

`"fmt"`, `"net/http"`のような短いパス

コードをソースレポジトリに置く場合は, ベースパスとしてソースリポジトリのルールを利用する必要がある.

例えば, `github.com/user`にあるGitHubアカウントを利用する場合には, おれがベースパス

ソースコードを置く場合は, 

`$ mkdir -p $GOPATH/src/github.com/user`

### 初めてのプログラム

パッケージを選ぶ

`mkdir $GOPATH/src/github.com/ashigirl96/hello`

に$hello.go$を書く

そこに`go`ツールでビルドして, インストールできるようになる.

`$ go install github.com/ashigirl96/hello`

### 初めてのライブラリ

ライブラリを書く.

` $go build github.com/ashigirl96/newmath`

`build`コマンドはファイルの出力はない. ファイル出力ならば`go install`を使う. ワークスペースの`pkg`ディレクトリにパッケージオブジェクトが生成される.

newmathパッケージがビルドされたことを確認したら, 修正して使ってみよう.

### テスティング

Goはgo testコマンドとtestingコマンドで構成される軽量なテストフレームワークを提供する

テストをかくときは, `_test.go`で終わる名前のファイルを作成し, `func (t *testing.T)`というシグネチャの`TestXXX`という名前の関数を作成する. テストフレームワークはこれらの関数を1つ1つ実行して, `t.Errorやt.Fiail`といったFail関数が呼ばれると, テスおtが失敗したとみなす.



