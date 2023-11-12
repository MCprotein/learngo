vscode format on save 설정 방법

<br/>

1. command + shift + p 눌러서 settings.json 접속

2. 아래 코드 추가

```json
"[go]": {
    "editor.defaultFormatter": "golang.go"
}
```

<br/>
참조: https://code.visualstudio.com/docs/languages/go#_formatting
