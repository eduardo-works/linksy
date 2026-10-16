# 🔗 Linksy

**Linksy** is a web application for organising and managing bookmarks in a structured way.

The project is being developed primarily as a **learning project**, with the goal of gaining practical experience with Go, backend development, relational databases and REST APIs.

## 🎯 Purpose

Linksy works as a personal bookmark manager.

Bookmarks can be organised using:

- **Pages**, which represent the main sections of the application.
- **Tags**, which can be displayed on multiple pages.
- **Bookmarks**, which can belong to multiple tags.

Pages and tags have a many-to-many relationship, as do tags and bookmarks.

This allows the same bookmark to appear in different contexts without being duplicated within a user's collection.

For example, a Gmail bookmark could belong to both the `Email` and `Personal` tags, while the `Email` tag itself could appear on several pages.

Tags within pages and bookmarks within tags support independent manual ordering.

## 🧱 Stack

The project uses:

- **Go** — backend and REST API
- **PostgreSQL** — relational database
- **Vue.js** — frontend

The project is intended to provide practical experience with:

- Go fundamentals and idiomatic programming
- Go project structure and package organisation
- HTTP servers and REST API development
- Relational database modelling
- SQL and database access
- Many-to-many relationships
- Database migrations
- Data validation and error handling
- Testing
- Authentication and user management at a later stage

The backend will be developed incrementally, prioritising an understanding of Go's standard library and core concepts before introducing additional frameworks or abstractions.

## 🗃️ Data Model

The main entities are:

```text
User
Page
Tag
Bookmark
PageTag
BookmarkTag
```

The main relationships are:

```text
User 1 ─── N Page
User 1 ─── N Tag
User 1 ─── N Bookmark

Page     N ─── N Tag
Tag      N ─── N Bookmark
```

`PageTag` represents the association between a page and a tag, including the tag's position within that page.

`BookmarkTag` represents the association between a bookmark and a tag, including the bookmark's position within that tag.

Each page, tag and bookmark belongs to exactly one user.

Relationships between entities are restricted to the same user.

A bookmark is stored only once per user and can be associated with multiple tags. Different users can independently store bookmarks pointing to the same URL.

## 👤 Users

User authentication and management are not part of the initial implementation.

The data model is nevertheless designed with multiple users in mind. During the first development stage, all data will belong to a predefined default user.

Authentication and user management can therefore be added later without requiring a major redesign of the database structure.

## 🐘 PostgreSQL

To configure PostgreSQL to use fewer resources during development, edit:

```text
/etc/postgresql/16/main/postgresql.conf
```

Change the following values:

```ini
shared_buffers = 64MB
max_connections = 20
work_mem = 2MB
maintenance_work_mem = 32MB
```

Restart the service:

```bash
sudo systemctl restart postgresql
```

These settings are intended for a small local development environment and should not be treated as production recommendations.

## 🖥️ PostgreSQL Client

To use the PostgreSQL client as administrator:

```bash
sudo -u postgres psql
```

To connect using the Linksy development user:

```bash
psql -h localhost -U linksy_dev -d linksy
```

## 🚦 Status

Linksy is currently under development.

Its primary purpose is learning Go through the development of a practical application.

The architecture, implementation details, data model and features may change as the project evolves.

## 🔗 API URLs

For more information about the API routes, see [URLs](URLS.md).
