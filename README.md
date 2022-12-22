# telegram-to-openai

# prepare

Before start, you must create a OPENAI api token and a telegram bot token.

# build

Use goland to build the project

    go mod download
    go build .
    
# deploy on VPS

Copy the binary file to where you want

    copy telegram-to-openai /usr/bin/
    
Copy .service file to systemd folder

    copy telegram-to-openai.service /etc/systemd/system/
    
Change parameters

    vi /etc/systemd/system/telegram-to-openai.service
    
Enable && start service

    systemctl daemon-reload
    systemctl enable telegram-to-openai --now
