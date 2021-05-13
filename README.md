# templuxinator
Think helm-charts-templates, but for any file type :) without helmet :) 



# This is... what?


I didnt like helm-chart templating.  
Then I learned it.  
And wanted it for any file, without helm  




templuxinator - simple templating engine to bake `mytemplatefile.anything` + `myvaluesfile.yaml` = `theresultfile.anything`

- `mytemplatefile.anything` should be a text file (UTF-8) in any format (.js, .yaml, .toml, anything), containing text and template-expressions (like *{{ .person2.name }}* )  
  The template can have [golang template functions](https://golang.org/pkg/text/template/#hdr-Functions) + and [sprig functions](http://masterminds.github.io/sprig/) (very much like helm charts)


- `myvaluesfile.yaml` should be a valid yaml file (UTF-8), the values that can be used inside the template
  
- `theresultfile.anything` will be created, containing a copy of the template with the template-expressions parsed


No more, no less, simple.



Demo in Asciinema in [this link](https://asciinema.org/a/413908?autoplay=1) or clicking image
[![templuxinator](https://asciinema.org/a/413908.png)](https://asciinema.org/a/413908?autoplay=1)

____

This was barely tested, but the code is so simple that there is not much to fail :)  
Golang makes it all very straightforward, I just connected the pieces: golang template/sprig lib, cli command arguments, and some file-read/writing :)


Open-source 4 life ;)
