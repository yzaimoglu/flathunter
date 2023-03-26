<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![BSD-3-Clause LICENSE][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/yzaimoglu/flathunter">
    <img src="images/logo.png" alt="Logo" width=50%>
  </a>

  <p><br></p>

  <p align="center">
    The only program you'll need to find a flat in Germany as a student.
    <br />
    <a href="https://github.com/yzaimoglu/flathunter"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/yzaimoglu/flathunter">Cloud Version</a>
    ·
    <a href="https://github.com/yzaimoglu/flathunter/issues">Report Bug</a>
    ·
    <a href="https://github.com/yzaimoglu/flathunter/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary><strong>Table of Contents</strong></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation-docker">Installation Docker</a></li>
        <li><a href="#installation-bare-metal">Installation Bare-Metal</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <!-- <li><a href="#acknowledgments">Acknowledgments</a></li> -->
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project
Flathunter is an open source software project created by two students that helps students find flats easily by crawling different real estate websites. It works by allowing the user to enter the URLs of the real estate websites they want to monitor, and then crawls these websites to save the flat listings. Whenever a new listing is posted on one of the platforms, the user is notified.

Flathunter is designed to be easy to use and customizable. The user can specify their search criteria, such as location, price, and number of bedrooms on the real estate website of their choice. After that they only need to copy the URL, Flathunter will do the rest. This is one of the features making Flathunter flexible and adaptable.

Flathunter is divided into three: Server, Client & Frontend. It is designed to be scalable, allowing it to handle a large number of users and real estate websites without compromising performance. Flathunter currently works with Ebay-Kleinanzeigen, WG-Gesucht and Immobilienscout24. The current notification methods supported are Discord Webhooks and HTTP Requests. More will be added in the future.

The project is open source, which means that anyone can contribute to it by adding new features, fixing bugs, and improving the code. It is hosted on a public repository on GitHub, where users can view the source code, report issues, and contribute to the project. Flathunter is licensed under the BSD-3-Clause license, which allows anyone to use, modify, and distribute the software for any purpose.

**TL;DR** Flathunter is a valuable tool for students who are searching for affordable housing. Its ability to monitor multiple real estate websites and notify users when new listings are posted makes it a time-saving and efficient tool. Its open source nature also ensures that it will continue to evolve and improve as more people contribute to the project.

[![Flathunter Screenshot][product-screenshot]](https://github.com/yzaimoglu/flathunter)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

**Backend:**
* ![Go][Go]
* ![ArangoDB][ArangoDB]
* ![Redis][Redis]

**Frontend:**
* ![JavaScript][JavaScript]
* ![Vue][VueJS]
* ![Nuxt][NuxtJS]
* ![Tailwind][TailwindCSS]

**General:**
* ![Docker][Docker]
* ![Swagger][Swagger]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started
Getting started with Flathunter is very easy. Make sure to have the prerequisites for your deployment of choice installed. After that follow the installation steps provided by us.

### Prerequisites
**If you want to deploy using the provided Dockerfiles or the Docker Compose file you only need:**
* Make
  *MacOS:* ```xcode-select --install```
  *Linux distros:* should be preinstalled
  *More Info:* [gnu.org](https://www.gnu.org/software/make/)
* Docker & Docker Compose
  *More Info:* [docs.docker.com](https://docs.docker.com/get-docker/)

**If you want to run Flathunter on bare-metal, you'll need:**
* Make
  *MacOS:* ```xcode-select --install```
  *Linux distros:* should be preinstalled
  *More Info:* [gnu.org](https://www.gnu.org/software/make/)
* Node (Tested and functional on v16.15.0)
  *MacOS:* ```brew install node@16```
  *Debian based Linux distros:* ```apt-get install node```
  *More Info:* [nodejs.org](https://nodejs.org)
* npm (Tested and functional on v8.5.5)
  ```sh
  npm install npm@latest -g
  ```
* Go (Tested and functional on v1.20.2)
  *MacOS:* ```brew install go```
  *Debian based Linux distros:* ```apt-get install go```
  *More Info:* [go.dev](https://go.dev/dl/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Installation Docker
1. Clone the git repository at [https://github.com/yzaimoglu/flathunter](https://github.com/yzaimoglu/flathunter)
   ```sh
   git clone https://github.com/yzaimoglu/flathunter.git
   ```
2. Navigate to the base directory and run the following command: *(The command can take a while, be patient)*
   ```sh
   make docker-build
   ```
3. Copy the env.example file into .env and change the environment variables to your liking
   ```sh
   cp env.example .env
   ```
4. Use the following make command to start Flathunter by using docker compose:
   ```sh
   make docker-up
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Installation Bare-Metal
1. Clone the git repository at [https://github.com/yzaimoglu/flathunter](https://github.com/yzaimoglu/flathunter)
   ```sh
   git clone https://github.com/yzaimoglu/flathunter.git
   ```
2. Run the following make command to install all necessary modules and packages:
   ```sh
   make install
   ```
3. Copy the env.example file into .env and change the environment variables to your liking
   ```sh
   cp env.example .env
   ```
4. Try running the different services in development mode using the following make commands:
   ```sh
   make dev-server
   make dev-client
   make dev-frontend
   ```
5. If everything works in development mode, you can build Flathunter for production by using the following make command:
   ```sh
   make prod
   ```
6. You should now see the server and client binaries and the output directory of the nuxt frontend.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage
Work in progress...

_For more examples, please refer to the [Documentation](https://github.com/yzaimoglu/flathunter)_
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- [ ] Server
    - [ ] REST API
- [ ] Client
    - [x] Ebay-Kleinanzeigen
    - [x] WG-Gesucht
    - [ ] Immobilienscout24
    - [ ] immowelt
- [ ] Frontend
    - [ ] User Authentication
    - [ ] Listing Dashboard
    - [ ] Statistics

See the [open issues](https://github.com/yzaimoglu/flathunter/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the BSD-3-Clause License. See `LICENSE` for more information.

#### Used Software with copryight notices:
* github.com/google/uuid (BSD-3-Clause)
Copyright (c) 2009,2014 Google Inc. All rights reserved.
* golang.org/x/crypto (BSD-3-Clause)
Copyright (c) 2009 The Go Authors. All rights reserved.
* github.com/arangodb/go-driver (Apache License 2.0)
Copyright 2017 ArangoDB GmbH
* github.com/gocolly/colly (Apache License 2.0)
Copyright 2018 gocolly
* github.com/alitto/pond (MIT License)
* github.com/go-co-op/gocron (MIT License)
* github.com/gofiber/fiber (MIT License)
* github.com/gofiber/helmet (MIT License)
* github.com/gookit/slog (MIT License)
* github.com/joho/godotenv (MIT License)
* nuxt.com (MIT License)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Yagizalp Zaimoglu - [@yzaimoglu](https://yagizalp.dev) - yagizalp@mitocho.io
Reza Jaber - [@rezajaber](https://rezajaber.dev) - reza@mitocho.io

Project Link: [https://github.com/yzaimoglu/flathunter](https://github.com/yzaimoglu/flathunter)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS
## Acknowledgments
* []()
* []()
* []()

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/yzaimoglu/flathunter.svg?style=flat
[contributors-url]: https://github.com/yzaimoglu/flathunter/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/yzaimoglu/flathunter.svg?style=flat
[forks-url]: https://github.com/yzaimoglu/flathunter/network/members
[stars-shield]: https://img.shields.io/github/stars/yzaimoglu/flathunter.svg?style=flat
[stars-url]: https://github.com/yzaimoglu/flathunter/stargazers
[issues-shield]: https://img.shields.io/github/issues/yzaimoglu/flathunter.svg?style=flat
[issues-url]: https://github.com/yzaimoglu/flathunter/issues
[license-shield]: https://img.shields.io/github/license/yzaimoglu/flathunter.svg?style=flat
[license-url]: https://github.com/yzaimoglu/flathunter/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat&logo=linkedin&colorB=555
[linkedin-url-yz]: https://linkedin.com/in/yagizalp-zaimoglu
[linkedin-url-rj]: https://www.linkedin.com/in/reza-jaber-583b511ba/
[product-screenshot]: images/screenshot.png
[VueJS]: https://img.shields.io/badge/Vue.js-35495E?style=flat&logo=vuedotjs&logoColor=4FC08D
[Go]: https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white
[JavaScript]: https://img.shields.io/badge/javascript-%23323330.svg?style=flat&logo=javascript&logoColor=%23F7DF1E
[Docker]: https://img.shields.io/badge/docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white
[Swagger]: https://img.shields.io/badge/-Swagger-%23Clojure?style=flat&logo=swagger&logoColor=white
[Nginx]: https://img.shields.io/badge/nginx-%23009639.svg?style=flat&logo=nginx&logoColor=white
[TailwindCSS]: https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=flat&logo=tailwind-css&logoColor=white
[NodeJS]: https://img.shields.io/badge/node.js-6DA55F?style=flat&logo=node.js&logoColor=white
[NuxtJS]: https://img.shields.io/badge/Nuxt-002E3B?style=flat&logo=nuxtdotjs&logoColor=#00DC82
[ArangoDB]: https://img.shields.io/badge/ArangoDB-DDE072.svg?style=flat&logo=ArangoDB&logoColor=black
[Redis]: https://img.shields.io/badge/redis-%23DD0031.svg?style=flat&logo=redis&logoColor=white
