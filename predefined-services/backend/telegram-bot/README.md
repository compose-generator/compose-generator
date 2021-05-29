## Telegram Bot
Telegram is an instant messenger, that supports chatting with bots. It paves the way for easily creating own bot applications for every purpose you can imagine. This service will give you the option to write your own Telegram bot, using Python.

### Setup
Telegram bot is considered as backend service and can therefore be found in backends collection, when generating the compose configuration with Compose Generator.

### Usage
1. Download Telegram onto your device
    - [Android](https://play.google.com/store/apps/details?id=org.telegram.messenger)
    - [iOS](https://apps.apple.com/de/app/telegram-messenger/id686449807)
    - [Desktop](https://desktop.telegram.org)
2. Search a chat called `BotFather` and open the command list by clicking on the `/` button at the bottom
3. Select `/newbot create a new bot` and follow the instructions from BotFather
4. Copy the API token BotFather sent to you and paste it, when Compose Generator asks you for the API token
5. Finish the setup of your bot in Compose Generator
6. Click on the link to your bot, BotFather sent you
7. Finished! Now you can chat with your demo bot, Compose Generator generated for you! Send the message `/roll` to the bot to make it roll the dice for you!

You can customize `bot.py` in your source directory on the server to customize the bots answers and functionalities. Have fun!