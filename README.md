# Prepare

Before start, you must create an OPENAI api token and a telegram bot token.

# Build

Use golang to build the project

    go mod download
    go build .
    
# Deploy on VPS

Copy the binary file to where you want

    copy telegram-to-openai /usr/bin/
    
Copy .service file to systemd folder

    copy telegram-to-openai.service /etc/systemd/system/
    
Change parameters

    vi /etc/systemd/system/telegram-to-openai.service
    
Enable && start service

    systemctl daemon-reload
    systemctl enable telegram-to-openai --now

# Other

You can also write a docker file to run it