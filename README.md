# lxbot-kip-counter

1. Discordのbotを作成する https://discordapp.com/developers/applications
2. 1で作成したbotの `Privileged Gateway Intents` > `MESSAGE CONTENT INTENT` を有効にする
3. 1で作成したbotのトークンを取得する
4.

```bash
cp sample.env .env
# .envにDiscordのbotトークンを設定
docker compose up
```
