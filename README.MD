# GitHub Notification Fetcher

## Install
Use `brew install shohag121/tap/letmeknow`
Or
```bash
brew tap shohag121/tap
brew install letmeknow
```

## Usage

### Login
Login with GitHub [Personal Token](https://github.com/settings/tokens).
```bash
letmeknow auth login -t <PERSONALGITHUBTOKEN>
```
### Logout
Logout of the script.
```bash
letmeknow auth logout -f
```
### Check auth status
Display currently authenticated user information.
```bash
letmeknow auth
```

### Check auth status
Show as List table.
```bash
letmeknow list
```

or 
Show as notification.
```bash
letmeknow list -d=false
```

