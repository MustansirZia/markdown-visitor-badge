#  👑 markdown-visitor-badge

![Badge](/static/1.png)

![Example](/static/2.png)

> Get this embeddable visitor badge instantly for your GitHub profile or any markdown file.

[![Documentation](https://godoc.org/github.com/mustansirzia/markdown-visitor-badge?status.svg)](http://godoc.org/github.com/mustansirzia/markdown-visitor-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/MustansirZia/markdown-visitor-badge)](https://goreportcard.com/report/github.com/MustansirZia/markdown-visitor-badge)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)


# Prerequisities
1. GitHub Account. (https://github.com)
2. Vercel Account. (https://vercel.com)

# Instructions.
1. Create a free distributed Redis cache on Vercel using [this](https://vercel.com/storage/kv) link.
2. Press the **Deploy** button to take your badge to the cloud!

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FMustansirZia%2Fmarkdown-visitor-badge&env=REDIS_HOST,REDIS_PORT,REDIS_USERNAME,REDIS_PASSWORD,REDIS_DATABASE,REDIS_USE_TLS&envDescription=Redis%20connection%20variables.)

3. In Vercel create project screen, you need to fill in all environment variables related to Redis. If you did step 1, you must already have them. You can bring your own Redis as well since the client used in the badge code is a generic Redis client which can connect to any datastore that supports Redis protocol.
*If you're using Vercel KV, make sure to set `REDIS_USE_TLS` to `true` and `REDIS_DATABASE` to `0`.*

4. After the deployment is complete, you should have a url of the deployment. It must be something like this 
`https://<your-slug>.vercel.app`. Note this down for the next step.

5. Paste the following markdown snippet inside your README.md file of your GitHub profile or any markdown file where you want to track visitors.
```markdown

 ![visitors](<url-from-step-4>/api/count)

```
6. And Voila! You should be able to see the badge appear instantly. *Since this is a global vercel function, your users should also see the badge appear almost instantly irrespective of their location on the globe*.

# Customization.
The look and feel of the badge can be customised by providing a set of query params inside the markdown snippet as described in the below picture.

![Customization](/static/3.jpeg)

# License.
MIT.
