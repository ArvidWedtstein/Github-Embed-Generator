#### Go API project for generating svg for github embed

# Generator 
https://arvidgithubembed.herokuapp.com/static/

## Better description & docs will come soon. (When i have time left)


![](https://img.shields.io/github/go-mod/go-version/arvidwedtstein/github-embed-generator?style=for-the-badge)

## Routes
Languages are comma seperated
| Routes | Query |
|---|---|
| /skills | ?languages=lang1,lang2,lang3 |
| /rankList | ?users=user1,user2 |
| /mostactivity | ?org=devco-morkjebla |
| /commitactivity | ?user=arvidwedtstein&repo=github-embed-generator |
| /project | ?user=arvidwedtstein&repo=github-embed-generator |

Example: `https://arvidgithubembed.herokuapp.com/skills?languages=php,mysql,javascript,typescript`
### Customization

- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Body text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_
- `title` - Card's custom title _(string)_
- `boxcolor` - Color of the boxes with the logos inside _(hex color)_



 users are comma seperated
> /rankList/:users

#### Common Options:
All hex colors without '#' please
- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Body text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_ 
- `boxcolor` - Card's languages color _(hex color)_
- `title` - Card's custom title _(string)_

Example: 
`/ranklist?users=lartrax,arvidwedtstein&bordercolor=black&titlecolor=red&textcolor=green&backgroundcolor=yellow&title=test`


> /mostactivity?org=devco-morkjebla


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


# Copilot

https://github.com/github/copilot-docs/blob/main/docs/visualstudiocode/gettingstarted.md#getting-started-with-github-copilot-in-visual-studio-code