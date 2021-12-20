# Technical Interview Exercises - World Clock (TypeScript/React)

これは、[株式会社 ABEJA](https://abejainc.com/ja/) の開発者採用面談で利用するコーディングテストの演習問題のひとつです。

## 制限事項

以下の制限事項を守ってください。

1. `npm install` などで新しいライブラリを追加しないでください
2. React のコンポーネントは、すべて関数コンポーネントで実装してください

## 問題 1

下記スクリーンショットのような世界時計を実装してください。デザインまでは実装されたファイルが `src` ディレクトリ以下に用意していますので、不足している処理を実装する形になります。

![](https://raw.githubusercontent.com/abeja-inc/bizflex-coding-interview-solutions/main/exercises/01-typescript-clock/doc/screen.png?token=AAAEQEAW5UYMW6ZGDQWNHXLBZFN7W)

### 機能要件

世界時計の要件は以下になります。

- 世界各地の都市（東京・シンガポール・ホノルル・ロサンゼルス・オークランド）の現在時刻を表示する
  - 時刻のフォーマットは `(1-24):(00-60)`
- 東京から見た各都市の時差を表示する
- 東京から見た各都市の日付が「昨日」「今日」または「明日」かを表示する

また、世界時計の下にはセッティング表示があり、ここには、

- `src/ClockContext` がタイマーで更新された回数
- 東京の現在時刻を秒まで表示
  - 時刻のフォーマットは `(1-24):(00-59):(00-59)`
- タイマーの発火間隔を以下から選択することができる
  - 停止
  - 100 ミリ秒
  - 500 ミリ秒
  - 1 秒
  - 5 秒

正解となる動作を録画した[動画ファイル](https://github.com/abeja-inc/bizflex-coding-interview-solutions/blob/main/exercises/01-typescript-clock/doc/screen.mov)も参考にしてください。

### 実行方法

このディレクトリで、

```bash
$ npm install
```

することで、依存ライブラリをインストールできます。その後、

```bash
$ npm start
```

で開発サーバが起動します。

### ソースコードの解説

`src` ディレクトリ以下にすべてのファイルが存在します。`src/App.tsx` がエントリーポイントになります。

```tsx
<ClockProvider localTimeZone="Asia/Tokyo" refreshInterval={refreshInterval}>
  <Clock title="東京" timeZone="Asia/Tokyo" />
  <Clock title="シンガポール" timeZone="Asia/Singapore" />
  <Clock title="ホノルル" timeZone="Pacific/Honolulu" />
  <Clock title="ロサンゼルス" timeZone="America/Los_Angeles" />
  <Clock title="オークランド" timeZone="Pacific/Auckland" />
  <Settings onIntervalChange={setRefreshInterval} />
</ClockProvider>
```

世界時計のロジックの実装は `src/ClockContext.tsx` にある `ClockContext` です。このコンテキストは以下のようなステートを持ちます。

```ts
export type ClockState = {
  // The local time zone. The default is the user's time zone.
  localTimeZone: string;

  // The current date and time in the `localTimeZone`.
  now: Dayjs;

  // The total number of updates since the start up.
  updates: number;

  // Update interval in milliseconds. Stop timer if the value is `null`.
  refreshInterval: number | null;
};
```

- **localTimeZone** ユーザーのタイムゾーン。今回は `Asia/Tokyo` にしてください
- **now** 現在時刻を表す `Dayjs` オブジェクトです。今回、日時操作には [Day.js](https://day.js.org/) を使います
- **updates** タイマーで更新された回数です
- **refreshInterval** タイマーの発火間隔です。`null` の場合はタイマーを停止してください

## 問題 2

問題 1 に従って、`ClockContext` を実装すると、タイマーが発火するたびに、`Clock` コンポーネントの render が実行されてしまいます。しかし、今の世界時計の仕様だと分単位で時刻を表示すればいいので、タイマーが発火するたびに `Clock` コンポーネントの render を実行する必要はありません。できるだけ、不必要な render の呼び出しを減らしてください。

**ヒント** 新しいコンテキストを追加しても構いません。
