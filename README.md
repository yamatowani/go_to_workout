# Go to Workout CLI

Go to Workout CLIは、筋トレのログを記録するためのCLIアプリケーションです。種目、重さ、セット数、レップ数、日付を入力し、それらの情報をNotionデータベースに保存できます。

## 目次
- [Go to Workout CLI](#go-to-workout-cli)
  - [目次](#目次)
  - [インストール](#インストール)
  - [使い方](#使い方)
    - [コマンド](#コマンド)
    - [フラグ](#フラグ)
    - [コマンド例](#コマンド例)
  - [環境変数の設定](#環境変数の設定)

## インストール

このアプリケーションを使用するには、Go言語がインストールされている必要があります。以下のコマンドでリポジトリをクローンして、依存関係をインストールします。

```bash
git clone https://github.com/yourusername/workout-tracker-cli.git
cd workout-tracker-cli
go mod tidy
```

## 使い方
1. Notionにログインし、インテグレーションを作成し、APIトークンを取得する。
2. Notion上に保存先のDBを作成する。プロパティに、[Exercise(テキスト型), Weight(数値型), Sets(数値型), Reps(テキスト型), Date(日付型)]を設定する。先ほど追加したインテグレーションをDBに追加する。
3. 下記を参考にアプリケーションに環境変数を設定する。

### コマンド
`log`コマンドで新しいトレーニングログを記録します。

### フラグ

`--exercise` : 種目名（必須）
`--weight` : 使用した重さ（kg）（必須）
`--sets` : セット数（必須）
`--reps` : セットごとのレップ数（カンマ区切り）（必須）
`--date` : トレーニングの日付（YYYY-MM-DD形式）（必須）

### コマンド例
```bash
go run . log --exercise "Bench Press" --weight 45 --sets 4 --reps "10,10,10,5" --date "2023-10-29"
```

## 環境変数の設定
Notion APIを使用するために、以下の環境変数を設定する必要があります。

API_TOKEN : Notion APIトークン。Notion上でインテグレーションを作成して、API Tokenを取得してください。
DATABASE_ID : データベースID。保存したIDBにインテグレーションを追加し、データベースの共有リンクから取得してください。

`https://www.notion.so/<DATABASE_ID>?v=<VIEW_ID>`

```bash
export API_TOKEN="your_notion_api_token"
export DATABASE_ID="your_database_id"
```
