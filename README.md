# Go-QRLogo

This is golang app to generate QR code with circular logo in the middle, you can adjust qr content, qr size, logo url and percentage logo size. The default output is an image.Image, but you can encode it to jpeg or png file.

<p align="center">
    <img src="https://github.com/nsrvel/go-qrlogo/blob/master/example/qrsample.jpeg"  width="30%" height="30%">
</p>
---

## Buy me a coffee

Whether you use this project, have learned something from it, or just like it, please consider supporting it by buying me a coffee, so I can dedicate more time on open-source projects like this :)

<a href="https://www.buymeacoffee.com/igorantun" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

---

## Features

-   Material Design
-   Emoji support
-   User @mentioning
-   Private messaging
-   Message deleting (for admins)
-   Ability to kick/ban users (for admins)
-   See other user's IPs (for admins)
-   Other awesome features yet to be implemented

![User Features](http://i.imgur.com/WbF1fi2.png)

![Admin Features](http://i.imgur.com/xQFaadt.png)

#### There are 3 admin levels:

-   **Helper:** Can delete chat messages
-   **Moderator:** The above plus the ability to kick and ban users
-   **Administrator:** All the above plus send global alerts and promote/demote users

---

## Setup

Clone this repo to your desktop and run `npm install` to install all the dependencies.

You might want to look into `config.json` to make change the port you want to use and set up a SSL certificate.

---

## Usage

After you clone this repo to your desktop, go to its root directory and run `npm install` to install its dependencies.

Once the dependencies are installed, you can run `npm start` to start the application. You will then be able to access it at localhost:3000

To give yourself administrator permissions on the chat, you will have to type `/role [your-name]` in the app console.

---

## License

> You can check out the full license [here](https://github.com/IgorAntun/node-chat/blob/master/LICENSE)

This project is licensed under the terms of the **MIT** license.
