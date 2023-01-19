<a name="readme-top"></a>

<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Version][GoVersion]][version-url]

</div>

<br />
<div align="center">
  <a href="https://github.com/ArvidWedtstein/Github-Embed-Generator">
    <img src="https://thetinygopher.com/images/tinygo-logo-small.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Github Embed Generator</h3>

  <p align="center">
    Go API project for generating svg for github embed
    <br />
    <a href="https://github.com/ArvidWedtstein/Github-Embed-Generator"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/ArvidWedtstein/Gtihub-Embed-Generator/issues">Report Bug</a>
    ·
    <a href="https://github.com/ArvidWedtstein/Github-Embed-Generator/issues">Request Feature</a>
  </p>
</div>



<details>
  <summary>Table of Contents</summary>
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
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#github-stats-card">Github Stats Card</a></li>
        <li><a href="#top-languages-card">Top Languages Card</a></li>
        <li><a href="#project-card">Project Card</a></li>
        <li><a href="#skills-card">Skills Card</a></li>
        <li><a href="#org-activity-card">Organization Activity Card</a></li>
        <li><a href="#repo-commit-activity-card">Activity Card for repositories</a></li>
        <li><a href="#streak-card">Streak Card</a></li>
        <li><a href="#icon-card">Icon Card</a></li>
        <li><a href="#themes">Themes</a></li>
        <li><a href="#customization">Customization</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

Go API project for generating svg's, which can be used in github profiles, repository readmes and much more!

***__Project is currently not hosted due to my poorness__***

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

This project is written in Go!

[![Go][Go]][Go-url]


<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

To run this project you'll need to have go installed on your computer. You can download it [here](https://go.dev/dl/)

Check go version

  ```sh
  go version
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/ArvidWedtstein/Github-Embed-Generator.git
   ```
   
2. Create a .env file in the root directory and add your environment variables. 

3. Run
   ```sh
   go run .
   ```


<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Deploying

I had this project deployed on heroku, since netlify doesn't support deploying for GO.

<!-- USAGE EXAMPLES -->

## Usage

This project can be used for adding useful information to your github profile, repository readme and much more!

A generator can to generate the url for you can be found [here](https://arvidgithubembed.herokuapp.com/static/) (current not working)

Here are some useful notes on how to use this project!

### GitHub Stats Card

Change the `?user=` to your own Github username.

```md
[![Arvid's Github Stats](https://arvidgithubembed.herokuapp.com/stats?user=arvidwedtstein)]
```

### Themes

Yes. Da themes have arrived. You can now customize your card without having these long urls 

To use a theme add the parameter `&theme=THEME_NAME` to the url

```md
[![Arvid's Github Stats](https://arvidgithubembed.herokuapp.com/stats?user=arvidwedtstein&theme=retro)]
```
Current themes are:
`
light,
dark,
rgb,
lartrax,
retro,
vue-dark,
ig9te,
github,
red,
toringe
`


### Customization

This API allows you to hide individual stats with the query parameter `?hide=`. Multiple stats to hide have to be comma-seperated. 

> Options: `contributions,milestones,packages,forks,releases,watchers,stars,disk,pull,issues,repocontributions,orgcontributions`
   
```md
[![Arvid's Github Stats](https://arvidgithubembed.herokuapp.com/stats?user=arvidwedtstein&hide=contributions,disk)]
```

#### Some Customization Options:
All hex colors without '#' please
- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Card's text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_ 
- `boxcolor` - Card's box color _(hex color)_
- `title` - Card's custom title _(string)_

#### Repository Commitactivity Card Exclusive Options:

- `hide_week` - Hides the week numbers on the card _(boolean)_
- `repo` - The name of your repository _(github repository name)_

> /commitactivity?:user&:repo

![](https://arvidgithubembed.herokuapp.com/commitactivity?user=arvidwedtstein&repo=github-embed-generator&titlecolor=333333&textColor=000000&backgroundcolor=ffffff&hide_week=false)


<!-- ROADMAP -->

## Roadmap

- [ ] Refactor whole project
- [ ] Implement a better system for embed themes
- [ ] Fix docs


See the [open issues](https://github.com/ArvidWedtstein/Github-Embed-Generator/issues) for a full list of proposed features (and known issues).



<!-- CONTRIBUTING -->

## Contributing

Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this project better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! <3

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.


<!-- CONTACT -->

## Contact

No contact :)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments


<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/ArvidWedtstein/Github-Embed-Generator.svg?style=for-the-badge
[contributors-url]: https://github.com/ArvidWedtstein/Github-Embed-Generator/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ArvidWedtstein/Github-Embed-Generator.svg?style=for-the-badge
[forks-url]: https://github.com/ArvidWedtstein/Github-Embed-Generator/network/members
[stars-shield]: https://img.shields.io/github/stars/ArvidWedtstein/Github-Embed-Generator.svg?style=for-the-badge
[stars-url]: https://github.com/ArvidWedtstein/Github-Embed-Generator/stargazers
[issues-shield]: https://img.shields.io/github/issues/ArvidWedtstein/Github-Embed-Generator.svg?style=for-the-badge
[issues-url]: https://github.com/ArvidWedtstein/Github-Embed-Generator/issues
[license-shield]: https://img.shields.io/github/license/ArvidWedtstein/Github-Embed-Generator.svg?style=for-the-badge
[license-url]: https://github.com/ArvidWedtstein/Github-Embed-Generator/blob/prod/LICENSE.txt
[Go]: https://img.shields.io/badge/Go-454545?style=for-the-badge&logo=Go&logoColor=00A7D0
[Go-url]: https://go.dev
[GoVersion]: https://img.shields.io/github/go-mod/go-version/arvidwedtstein/github-embed-generator?style=for-the-badge
[version-url]: https://go.dev/dl
