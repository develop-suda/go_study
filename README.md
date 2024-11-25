# go_study

### DI学習用ブランチ
[参考動画　一番最初に見るべき資料](https://www.youtube.com/watch?v=4wS6RwG9W_s)
[参考記事](https://qiita.com/yoshinori_hisakawa/items/a944115eb77ed9247794)


- 依存関係は一方向になるようにする（双方向依存はしない）
- 依存関係は循環しないようにする
- 上位のモジュールは下位のモジュールに依存してはいけない
- 抽象は実装に依存してはいけない、実装は抽象に依存するべき