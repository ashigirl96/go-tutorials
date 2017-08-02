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

### For is Go's "while"

セミコロンを省略して, C言語のwhileはGoでは`for`だけを使う

### Switch

Goの`switch`は他の言語と比較して, 大きく異なるのは, `case`の最後で自動的に`break`する. `break`せずに通したい場合は, `fallthrough`文を`case`の最後に記述する

### Defer

`defer`は, `defer`へ渡した関数の実行を, 呼び出し元の関数の終わり(returnする)まで遅延させるもの.

`defer`へ渡した関数の引数は, すぐに評価されるが, その関数自体は呼び出し元の関数がreturnするまで実行されない

### Pointers

ポインタは変数のメモリアドレスを指す

```
var p *int
i := 42
p = &i
```

### Structs

`structs`(構造体)は, `field`の集まりである

(`type`宣言は型を宣言するためのもの)

### Pointers to structs

`struct`の`Field`は, `struct`のポインタを通してアクセスすることもできる

`Field`Xにアクセスするには`(*p).X`のようにかく. Goでは代わりに`p.X`と書く

### Arrays

`[n]T`型は, 型`T`の`n`個の変数の配列(array)を表す

```
var a [10]int
```

### Slices

配列は固定長です。一方で、スライスは可変長です。より柔軟な配列とみなすことgあできます。スライスは配列よりも一般的です。

型`[]T`は型`T`のスライスを表します。


### Slice length and capacity

スライスにはlengthとcapacityの両方がある

スライスの長さはそれに含まれる要素数

capacityは、スライスのシア書の要素から数えて、元となる配列の要素数

### Creating a slice with make

`make`関数で動的サイズの配列を作れる.

`make`関数はゼロ化された配列を割り当て、その配列を指すスライスを返す

### Maps

`map`はキーと値とを関連つける

マップのゼロ値は`nil`.

`nil`マップはキーを持っておらず, キーを追加することができない

### Function Closures

Goの関数は`Closure`です. クロージャは、それ自身の外部から変数を参照する関数値. 関数は, 参照された変数へアクセスして変えることができ、関数は変数へ`"bind"`されてる

`adder`関数はクロージャを返す. 各クロージャは、それ自身の`sum`変数へbindされる

### Methods

クラス(`class`)がないけど、型にメソッド(`method`)を定義できる

メソッドは、特別なレシーバ(`receiver`)引数を関数に取ります

`reciver`は、`func`キーワードとメソッド名の間に自身の引数リストで表現する

### Pointer Receivers

レシーバの型が, ある型`T`への構文`*T`があることを意味する. `T`は`*int`のようなポインタ自身を取ることができない

例では`*Vertex`に`Scale`メソッドが定義されてる

Pointer Reciverを持つメソッド(ここでは`Scale`)は、レシーバが指す変数を変更できる. Reciver自身を更新することが多いため, Variable ReciverよりもPointer Reciverの方が一般的

`Scale`の宣言から`*`を消して、プログラムがどのように振る舞うか確認する

Variable Receiverでは、`Scale`メソッドの操作は元の`Vertex`変数のコピーを操作する.

つまり, `main`関数で宣言した`Vertex`変数を変更するためには, `Scale`メソッドはPointer Receiverにする必要がある

なんというか、破壊的メソッドのような役割があるような.

### Choosing a value of Pointer Receiver

Pointer Receiverを使う理由

- メソッドがレシーバが指す先の変数を変更するため
- メソッドの呼び出し毎に変数のコピーを避けるため. レシーバが大きな構造体である場合に効率的

### Interfaces

interface型は、methodのsignatureの集まりで定義

methodの集まりを実装した値を, `interface`型の変数へ持たせることができる

### Type assertions

Type Assertionは、interfaceの値の基になる具体的な値を利用する手段

### Stringers

Stringers Interfaceは、`string`として表現できる型

Pythonでいう__str__(self)を書き換えるような感じだろうか

### Reader

io パッケージは、データストリームを読むことを表現する io.Reader インタフェースを規定されてる

Goの標準ライブラリには、インタフェース、ファイル、ネットワーク接続、圧縮、暗号化、などで 多くの実装されてる





