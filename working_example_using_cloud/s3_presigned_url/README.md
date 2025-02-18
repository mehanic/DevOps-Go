# Steps for creating Pre-Signed URLs for Amazon S3 Buckets

- Create an `.env` file and add access keys here.

```
	AWS_ACCESS_KEY_ID="AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY="AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN="token"
```

- Crease `launch.json` file to use IntelliSense and env files.

```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "envFile": "${workspaceFolder}/.env",
            "args": [
 
            ]
        }
    ]
}
```

- `aws-sdk-go-v2` is used here.
- In order to run the code build it first with `go build ./` and run the code `go run ./`.