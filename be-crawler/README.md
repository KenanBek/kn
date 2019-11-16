# Current workflow

```

# Execute main application
go run cmd/main.go

# Get collection names using mongocli
go run mongocli.go colls -d kn

# Get document in collection using mongocli
go run mongocli.go list links -d kn

# Run a mongodb Docker container
docker run --name mongodb -p 27017:27017 mongo

# Start an existing mongodb in attached mode
docker start -a mongodb

```

# Type of entities

- Links (uri, hash)
- Posts (uri, hash, title)

When we start crawling source urls, as a result, we get article links saved as a `posts` entities. URIs itself will be saved as a `links` entity. It's important to note that in `links` we will also store uris for intermediate pages, likes categories, search pages and etc. In this case we should have some mechanism for re-parse case: intermediate pages might need to be parsed based on schedule.

# Work Milestones

Initial version:

- Get links from sources (do not crawl inner pages)
- Recognize post links and save them
- Hash table for the links to avoid reparse and recheck for the post links
- API to expose articles
- List page with paging to list post links
- Details page with web view to view selected link

Needed infrastructure and core features:

- Versioning and packaging
- Architectural decisions (be, api, fe, user interactions)
- Deployment pipeline
- Containerization

Upcoming features for future versions:

- Crawl inner pages intelligently (with schedule for category and paging links, at least)
- Index for post link's content
- Search based on index
- User database and user interaction features (save, like, comment)
