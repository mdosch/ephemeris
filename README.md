[![GoDoc](https://godoc.org/github.com/skx/ephemeris?status.svg)](http://godoc.org/github.com/skx/ephemeris)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/ephemeris)](https://goreportcard.com/report/github.com/skx/ephemeris)
[![license](https://img.shields.io/github/license/skx/ephemeris.svg)](https://github.com/skx/ephemeris/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/ephemeris.svg)](https://github.com/skx/ephemeris/releases/latest)


* [Ephemeris](#ephemeris)
* [Installation](#installation)
* [Blog Generation](#blog-generation)
   * [Blog Format](#blog-format)
* [Demo Blog](#demo-blog)
* [Hacking](#hacking)
* [Feedback](#feedback)




# Ephemeris

`ephemeris` is a golang CLI-application which will generate a blog from a collection of static text-files, complete with:

* An archive-view.
  * Showing posts by year, and month.
* Comment support.
  * See [COMMENTS.md](COMMENTS.md) for more details on the setup required.
* A tag-cloud.
  * Containing all tags, and a list of posts using a specified tag.
* An RSS feed.
  * Containing the most recent ten posts.
  * Full text is included in the feed.

The project was primarily written to generate [my own blog](https://blog.steve.fi/), which was previously generated with the perl-based [chronicle blog compiler](https://steve.fi/Software/chronicle/) - if you've used `chronicle` you may consult the [brief notes on migration](MIGRATION.md).



# Installation

You can install from source, by cloning the repository and running:

    $ cd ephemeris/cmd/ephemeris
    $ go build .
    $ go install .

Or if you just wish to install the binary:

    $ go install github.com/skx/ephemeris/cmd/ephemeris@latest

Alternatively you may find precompiled binaries available for many systems upon the [release page](https://github.com/skx/evalfilter/releases).




# Blog Generation

A blog is generated from two things:

* A series of blog-posts, stored beneath a given directory.
* An optional set of comments, which are plaintext files associated with a given blog-post.

To build/generate/create your blog you need to create a configuration file that contains the appropriate directories.  The configuration file is assumed to be named `ephemeris.json` in the current-directory, and a sample configuration file would look like this:

        {
          "CommentsPath":  "./comments/",
          "OutputPath":    "./output/",
          "PostsPath":     "./posts/",
          "Prefix":        "http://blog.steve.fi/"
        }

Once you have a configuration file simply run the command to compile and generate your blog:

    $ ephemeris

As expected the generated output will be placed beneath the `output/` directory.  The possible configuration-keys in the JSON file are:

* `PostsPath` - **Mandatory**
  * This is the path to the directory containing your blog-posts.
  * This directory will be searched recursively for content.
* `CommentAPI`
  * The URL of the CGI script to receive comments, this is used in the add-comment form.
    * See [COMMENTS.md](COMMENTS.md) for a discussion of comments.
* `CommentsPath`
  * This is the path to the directory containing your comments.
  * If this is empty then no comments will be read/inserted into your output
  * See [COMMENTS.md](COMMENTS.md) for a discussion of comments.
* `OutputPath`
  * The path beneath which all output content should be written.
  * This defaults to `output/` if not specified.
* `Prefix` - **Mandatory**
  * This is the URL-prefix used to generate all links.
* `ThemePath`
  * This is the path to a local theme you're using, if you don't wish to use the default theme embedded within the binary.
  * See the [theming](#theming) section in this document for more details.


There is a command-line flag which lets you specify an alternative configuration-file, if you do not wish to use the default.  Run `ephemeris -help` to see details.




## Blog Format

The input to this program is a directory tree containing a series of blog-posts.  Each post will be stored in a single file, with the entry being prefixed by a simple header containing meta-data.

A sample post would look like this:

```
Tags: compilers, assembly, golang, brainfuck
Date: 14/06/2020 19:00
Subject: Writing a brainfuck compiler.
Format: markdown

So last night I had the idea..
```

There are a few things to note here:

* The header and the content are separated by a single blank line.
* The date **MUST** be in the specified format.
* If there is no `format: markdown` header then the body will be assumed to be HTML.
  * All my early posts were written in HTML.
  * Later I switched to markdown.

As noted the input directory will be processed recursively, which allows you to group posts by topic, year, or in any other way you might prefer.  I personally file my entries by year, like so:

```
data/
  ├── 2005
  | ├── 1.txt
  | └── 2.txt
  ├── 2006
  | ├── 3.txt
  | └── 4.txt
  ..
  ├── 2018
  | ├── 5.txt
  | └── 6.txt
  ├── 2019
  | ├── 7.txt
  | └── 8.txt
  └── 2020
    ├── 9.txt
    ├── 10.txt
    └── 11.txt
```


# Demo Blog

There is a demo-blog contained within this repository, along with a lightly-modified theme.  To compile the blog into a set of HTML output-pages simple change into the appropriate directory and run the command:

```
$ cd _demo
$ ephemeris
```

This will generate `./output/`.  As you can see from the [configuration file](_demo/ephemeris.json) the blog will have an URL-prefix of `http://localhost:8000` so you can try serving it with a local webserver on that port to view it in your browser:

```
cd output/
python -m SimpleHTTPServer 8000
```

Once the simple HTTP-server is running open http://localhost:8000/ with your browser to see the compiled/generated result.




# Theming

The main binary contains an embedded set of `text/template` resources which are used to generate the output blog.  If you wished to update those static-resources you'd need to edit the application-source and rebuild it which is a bit cumbersome, and future releases would almost certainly overwrite your changes.

Instead of having to rebuild the application to change the generated output you can use a local directory of templates instead of the embedded resources.

To get started you should export the default templates to a local directory:

```
ephemeris -export-theme=./blog-theme/
```

This will give you the following contents:

```
blog-theme/
├── archive_page.tmpl
├── archive.tmpl
├── entry.tmpl
├── inc
│   ├── add_comment_form.tmpl
│   ├── blog_post.tmpl
│   ├── comments_on_blog_post.tmpl
│   ├── css.tmpl
│   ├── recent_posts.tmpl
│   └── rss.tmpl
├── index.rss
├── index.tmpl
├── tag_page.tmpl
└── tags.tmpl

1 directory, 13 files
```

Now that you have the local templates available you can edit them, changing the text and layout as you wish, and specify that local directory as the `ThemePath` in your `ephemeris.json` configuration file.

* **NOTE:** The templates are processed using the standard [golang text/template](https://golang.org/pkg/text/template/) package.




# Feedback

This project is not documented to my usual and preferred standard, no doubt it will improve over time.

However there are many blog packages out there, so I expect this project will only be of interest to those wishing to switch from `chronicle`.

Steve
