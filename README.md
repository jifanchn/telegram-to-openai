# What is this

This project links a Telegram bot and the OpenAI API together. 

Instead of using an online browser-based token, this project uses the OpenAI official RESTful API. 

The result is that users will not be annoyed with endless waiting and red alerts, but the downside is that you may be charged for this service by OpenAI.

![](doc/d1bb1b0103f448af7bc3cc2a69eb916.jpg)


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