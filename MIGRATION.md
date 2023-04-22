## Chronicle History

The chronicle blog compiler started life as a simple project, but grew in complexity over time.  Part of the reason the complexity grew was because the project was very flexible:

* It would read a whole bunch flat files, each of which contained a single blog-post.
* Each parsed post would then be inserted into an SQLite database.
* Using this intermediary database a series of plugins would each execute in turn:
  * The plugins would run SQL-queries to extract posts of interest.
    * For example building a tag-cloud.
    * For example building an archive-view.
    * For example outputting the front-page (10 most recent posts) & associated RSS-feed.
* Once complete the SQLite database would be destroyed.

My expectation was that the use of an intermediary SQLite database would allow content to be generated in a very flexible and extensible fashion, however over time it became apparent that I didn't _actually_ need generation to be very flexible!  Most blogs look the same, if you have tags, archives, etc, then that's enough.

In short this project was born to __replace__ chronicle, and perform the things I actually need, rather than what I _suspected_ I might want.



# Migration from chronicle

As with `chronicle` the input to this program is a directory containing a series of blog-posts, along with a path from which to load static-comments.

There are two changes in this project which you will have to adjust to:

* The format of blog-posts became more strict.
* The naming of the comment-files became more strict.


## Blog Posts

Each post will be stored in a single file, with the entry being prefixed by a header containing meta-data. A sample post would look like this:

```
Tags: markdownshare, puppet, marionette, github, oodi, university
Date: 07/04/2020 09:00
Subject: A busy few days
Format: markdown

Over the past few weeks things have been pretty hectic.
```

There are a few things to note here:

* The date **MUST** be in the specified format.
* If there is no `format: markdown` header then the body will be assumed to be HTML.
  * All my early posts were written in HTML.
  * Later I switched to markdown.

A related change is that it is now a __fatal-error__ for a blog-post to have a header-key which is unknown.  To provide a concrete example it was previously possible to write:

```
Subject: I won't write another email client
Tags: golang, email, maildir, maildir-utils
Date: 08/01/2020 19:19
Blah: foo
Publish: later
Format: markdown

Once upon a time I wrote an email client, in a combination of C++ and Lua.
```

Now `Blah`, and `Publish` are explicitly prohibited.



## Comment Files

In the past, comments would be written to files with names such as:

* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-13:18:09`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-13:20:39`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-14:20:40`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.23-November-2008-14:44:15`

We now expect these to be named:

* `${link}.${ctime}`

So these examples would become:

* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227432366`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227439089`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227439239`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227442840`
* `you_ve_had_this_coming_since_the_day_you_arrived.html.1227444255`

See [COMMENTS.md](COMMENTS.md) for a discussion on comment-setup.




## Migration Summary

Start by pointing the configuration file at your existing entries, ignoring comments.  You'll need to fix the `Date:` header of your posts  until the errors go away.

Once you've done that you can now rename the comments, if you use them.
