# githubcomnimblewaynimblego

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentResponse">AgentResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractResponse">ExtractResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MapResponse">MapResponse</a>

Methods:

- <code title="post /v1/agent">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.Agent">Agent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentParams">AgentParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentResponse">AgentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/extract">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.Extract">Extract</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractParams">ExtractParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractResponse">ExtractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/map">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.Map">Map</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MapParams">MapParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MapResponse">MapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agents

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentListResponse">AgentListResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGetResponse">AgentGetResponse</a>

Methods:

- <code title="get /v1/agents">client.Agents.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentListParams">AgentListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentListResponse">AgentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{template_name}">client.Agents.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGetResponse">AgentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Crawl

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlListResponse">CrawlListResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlRunResponse">CrawlRunResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlStatusResponse">CrawlStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlTerminateResponse">CrawlTerminateResponse</a>

Methods:

- <code title="get /v1/crawl">client.Crawl.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlListParams">CrawlListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlListResponse">CrawlListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/crawl">client.Crawl.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlRunParams">CrawlRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlRunResponse">CrawlRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/crawl/{id}">client.Crawl.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlStatusResponse">CrawlStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/crawl/{id}">client.Crawl.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlService.Terminate">Terminate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#CrawlTerminateResponse">CrawlTerminateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
