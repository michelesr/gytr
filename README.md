# Go YAML Template Renderer

Write [templates](https://godoc.org/text/template)
and render them using data from a [YAML](https://en.wikipedia.org/wiki/YAML)
file.

## Installation

    $ go get 'github.com/michelesr/gytr'

## Usage

```
Usage of gytr:
  -d string
        Data YAML file (default "data.yml")
  -t string
        Template file (default "template")
```

### Examples

#### sshconfig

This is the `template` file:

```
{{range .hosts}}
Host {{.name}}
   User {{.user}}
   Port {{.port}}
{{end}}
```

This is the `data.yml` file:

```yaml
hosts:
-
  name: example.org
  user: example
  port: 22
-
  name: foo.bar
  user: foobar
  port: 2222
```

Launch `gytr`:

```
$ gytr

Host example.org
   User example
   Port 22

Host foo.bar
   User foobar
   Port 2222
```

#### HTML

`template`:

```html
<!DOCTYPE html>
<html>
    <head>
        <title>{{.title}}</title>
    </head>
    <body>{{range .paragraphs}}
        <p>
          {{.}}
        </p>{{end}}
    </body>
</html>
```

`data.yml`:

```yaml
title: Example
paragraphs:
  - A short paragraph
  - >-
    The path of the righteous man is beset on all sides by the iniquities of the
    selfish and the tyranny of evil men. Blessed is he who, in the name of
    charity and good will, shepherds the weak through the valley of darkness,
    for he is truly his brother's keeper and the finder of lost children. And I
    will strike down upon thee with great vengeance and furious anger those who
    would attempt to poison and destroy My brothers. And you will know My name
    is the Lord when I lay My vengeance upon thee.
  - >-
    Now that we know who you are, I know who I am. I'm not a mistake! It all
    makes sense! In a comic, you know how you can tell who the arch-villain's
    going to be? He's the exact opposite of the hero. And most times they're
    friends, like you and me! I should've known way back when... You know why,
    David? Because of the kids. They called me Mr Glass.
```

Output:

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Example</title>
    </head>
    <body>
        <p>
          A short paragraph
        </p>
        <p>
          The path of the righteous man is beset on all sides by the iniquities of the selfish and the tyranny of evil men. Blessed is he who, in the name of charity and good will, shepherds the weak through the valley of darkness, for he is truly his brother's keeper and the finder of lost children. And I will strike down upon thee with great vengeance and furious anger those who would attempt to poison and destroy My brothers. And you will know My name is the Lord when I lay My vengeance upon thee.
        </p>
        <p>
          Now that we know who you are, I know who I am. I'm not a mistake! It all makes sense! In a comic, you know how you can tell who the arch-villain's going to be? He's the exact opposite of the hero. And most times they're friends, like you and me! I should've known way back when... You know why, David? Because of the kids. They called me Mr Glass.
        </p>
    </body>
</html>
```
