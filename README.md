# CLI tool for creating google photos album from your local folder

## Prerequisites
```
Golang Runtime
Google API Client Setup
```

## 1. Clone Repository
```
git clone https://github.com/ksyuk/folder-to-photos-album.git
```

## 2. Create .env
Create ``.env`` file in project root directory.  
Fill the following CLIENT_ID, CLIENT_SECRET, REDIRECT_URI.
```
CLIENT_ID="<fill here>"
AUTH_URI="https://accounts.google.com/o/oauth2/auth"
TOKEN_URI="https://oauth2.googleapis.com/token"
CLIENT_SECRET="<fill here>"
REDIRECT_URI="<fill here>"
SCOPE="https://www.googleapis.com/auth/photoslibrary"
RESPONSE_TYPE="code"
```
## 3. Run Code and Select mode 2 (start router)
```
mode:2
```
Go to "http://localhost:8080/access-token".  
Continue until redirect to the page displaying access token.  
Copy the access token and add the following line to the above ``.env`` file
```
ACCESS_TOKEN="<fill here>"
```

## 4. Locate your folder
Locate your folder in project root directory.

## 5. Run Code and Select mode 1 (update)
```
mode:1.
```
Fill the following info
```
folder name:
album name:
```
example
```
folder name:flower
album name:flower2023/01/01
```