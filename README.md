#### Go API project for generating svg for github embed

# Generator 
https://arvidgithubembed.herokuapp.com/static/

## Better description & docs will come soon. (When i have time left)


# Features:

- [GitHub Stats Card](#github-stats-card)
- [Top Languages Card](#top-languages-card)
- [Project Card](#project-card)
- [Skills Card](#skills-card)
- [Organization Activity Card](#org-activity-card)
- [Repository Commit Activity Card](#repo-commit-activity-card)
- [Streak Card](#streak-card)
- [Icon](#icon-card)
- [Themes](#themes)
- [Customization](#customization)
  - [Common Options](#common-options)
  - [Stats Card Exclusive Options](#stats-card-exclusive-options)
  - [Repo Card Exclusive Options](#repo-card-exclusive-options)
  - [Language Card Exclusive Options](#language-card-exclusive-options)
  - [Wakatime Card Exclusive Option](#wakatime-card-exclusive-options)
- [Deploy Yourself](#deploy-on-your-own-vercel-instance)


![](https://img.shields.io/github/go-mod/go-version/arvidwedtstein/github-embed-generator?style=for-the-badge)


## GitHub Stats Card

Change the `?user=` to your own Github username.

```md
[![Arvid's Github Stats](https://arvidgithubembed.herokuapp.com/stats?user=arvidwedtstein)]
```

### Hide your stats

This API allows you to hide individual stats with the query parameter `?hide=`. Multiple stats to hide have to be comma-seperated. 

> Options: `contributions,milestones,packages,forks,releases,watchers,stars,disk,pull,issues,repocontributions,orgcontributions`
   
```md
[![Arvid's Github Stats](https://arvidgithubembed.herokuapp.com/stats?user=arvidwedtstein&hide=contributions)]
```

### Themes

Yes. Da themes have arrived. You can now customize your card without having these long urls

To use a theme add the parameter `&theme=THEME_NAME` to the url

```md
[![Arvid's Github Stats](https://arvidgithubembed.herokuapp.com/stats?user=arvidwedtstein&theme=retro)]
```

#### All themes :-

light, dark, rgb, lartrax, retro, vue-dark, ig9te, github, red, toringe

### Customization

In addition to themes you can also customize your card with URL parameters.

#### Common Options:
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

Example:
`https://arvidgithubembed.herokuapp.com/commitactivity?user=arvidwedtstein&repo=github-embed-generator&titlecolor=000000&textColor=000000&backgroundcolor=ffffff&hide_week=false`


## Hide Individual Stats

`&hide=stat1,stat2...`

<table>
   <thead>
      <tr>
         <th>Options</th>
         <th>Query</th>
      </tr>
   </thead>
   <tbody>
      <tr>
         <td>Contributions</td>
         <td>&hide=contributions</td>
      </tr>
      <tr>
         <td>Milestones</td>
         <td>&hide=milestones</td>
      </tr>
      <tr>
         <td>Packages</td>
         <td>&hide=packages</td>
      </tr>
      <tr>
         <td>Forks</td>
         <td>&hide=forks</td>
      </tr>
      <tr>
         <td>Releases</td>
         <td>&hide=releases</td>
      </tr>
      <tr>
         <td>Watchers</td>
         <td>&hide=watchers</td>
      </tr>
      <tr>
         <td>Stars Earned</td>
         <td>&hide=stars</td>
      </tr>
      <tr>
         <td>Disk Usage</td>
         <td>&hide=disk</td>
      </tr>
      <tr>
         <td>Pull Requests</td>
         <td>&hide=pull</td>
      </tr>
      <tr>
         <td>Issues</td>
         <td>&hide=issues</td>
      </tr>
      <tr>
         <td>Repositories Contributed To</td>
         <td>&hide=repocontributions</td>
      </tr>
      <tr>
         <td>Organizations Contributed To</td>
         <td>&hide=orgcontributions</td>
      </tr>
   </tbody>
</table>

# Contribution

Contributions are welcome!
Feel free to open a pull request or an issue

Make sure your request is meaningful and you have tested the app locally before submitting a pull request.




# TODO

- Repository Timeline?


