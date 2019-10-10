# Type of entities

Now:

- URIHashes
- Posts
- Pages

Change to:

- Links (uri, hash)
- Posts (uri, hash, title)
- 
- 

When we start crawling source urls, as a result, we get article links saved as a `posts` entities. URIs itself will be saved as a `links` entity. It's important to note that in `links` we will also store uris for intermediate pages, likes categories, search pages and etc. In this case we should have some mechanism for re-parse case: intermediate pages might need to be parsed based on schedule.

