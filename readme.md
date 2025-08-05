# traQ ID を google form に自動入力する

## NeoShowcase のビルド設定

- Deploy Type: Runtime
- Build Type: Buildpack
- Context: .
- Use Database: No

## 環境変数

- ENTRY: traQ ID を入力したい質問のID (数字)
- URL: Google Form の URL (https:// から viewform? まで)

ENTRY に入れる数字の取得方法は↓を参考にするといい
https://blog.nakachon.com/2016/12/22/how-to-add-url-parameter-for-google-form/

