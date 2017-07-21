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

テストをかくときは, `_test.go`で終わる名前のファイルを作成し, `func (t *testing.T)`というシグネチャの`TestXXX`という名前の関数を作成する. テストフレームワークはこれらの関数を1つ1つ実行して, `t.Errorやt.Fiail`といったFail関数が呼ばれると, テストが失敗したとみなす.

## A Tour of Go

### Packages

GoのプログラムはPackageで構成される.

プログラムは`main`パッケージから開始される　

### Imports

カッコでpackageのimportをgroup化し, factored import statementとする.

```
import "fmt"
import "math"
```

ともできる.

### Exported names

最初の文字が大文字で始まる名前は、外部のパッケージから参照できるexport name(公開された名前). `Pi`は`math` packageでexportされてる.

### Named return values

戻り値となる変数に名前をつける(named return value). 戻り値に名前をつけると, 関数の最初で定義した変数名としてあ疲れます.

この戻り値の名前は, 戻り値の意味を示す名前とすることで、関数のドキュメントとして表現するようにします。

名前をつけた戻り値の変数を使うと、`return`文に何も書かずに戻せる. これを`"naked" return`と云う

naked returnステートメントは, 短い関数でのみ利用すべき.

### Variable

`var`は変数を宣言する. 

初期化子が与えられている場合, 型を省略することができる. その変数は初期化子が持つ型になる

### Short variable declarations

`var`の代わりに`:=`を使うと暗黙的な型宣言ができる

関数外ではキーワードで始まる宣言(var, func, など)が必要で、`:=`での暗黙宣言ができない

### Type conversions

型変換

変数`v`と型`T`があったとき、`T(v)`は変数`v`を`T`型へ変換する

### Constants

定数は、`const`キーワードを使って変数を同じように宣言する

定数はcharacter、string, boolean, numeric(数値)のみで使える

定数は`:=`を使って宣言できない
