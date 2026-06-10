# Shared Params Types

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#AutoScrollActionParam">AutoScrollActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#ClickActionParam">ClickActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#EvalActionParam">EvalActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#FetchActionParam">FetchActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#FillActionParam">FillActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#GetCookiesActionParam">GetCookiesActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#GotoActionParam">GotoActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#PressActionParam">PressActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#ScreenshotActionParam">ScreenshotActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#ScrollActionParam">ScrollActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#WaitActionParam">WaitActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#WaitForElementActionParam">WaitForElementActionParam</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go/shared#WaitForNavigationActionParam">WaitForNavigationActionParam</a>

# githubcomnimblewaynimblego

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractResponse">ExtractResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractAsyncResponse">ExtractAsyncResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractBatchResponse">ExtractBatchResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MapResponse">MapResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /v1/extract">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.Extract">Extract</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractParams">ExtractParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractResponse">ExtractResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/extract/async">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.ExtractAsync">ExtractAsync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractAsyncParams">ExtractAsyncParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractAsyncResponse">ExtractAsyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/extract/batch">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.ExtractBatch">ExtractBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractBatchParams">ExtractBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#ExtractBatchResponse">ExtractBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/map">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.Map">Map</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MapParams">MapParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MapResponse">MapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search">client.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#GithubcomnimblewaynimblegoService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SearchParams">SearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agent

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentListResponse">AgentListResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGenerateResponse">AgentGenerateResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGetResponse">AgentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGetGenerationResponse">AgentGetGenerationResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunResponse">AgentRunResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunAsyncResponse">AgentRunAsyncResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunBatchResponse">AgentRunBatchResponse</a>

Methods:

- <code title="get /v1/agents">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentListParams">AgentListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentListResponse">AgentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/generations">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGenerateParams">AgentGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGenerateResponse">AgentGenerateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/{template_name}">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateName <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGetResponse">AgentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/agents/generations/{generation_id}">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.GetGeneration">GetGeneration</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, generationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentGetGenerationResponse">AgentGetGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/run">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunParams">AgentRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunResponse">AgentRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/async">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.RunAsync">RunAsync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunAsyncParams">AgentRunAsyncParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunAsyncResponse">AgentRunAsyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agents/batch">client.Agent.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentService.RunBatch">RunBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunBatchParams">AgentRunBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#AgentRunBatchResponse">AgentRunBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# Tasks

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskListResponse">TaskListResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskGetResponse">TaskGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskResultsResponse">TaskResultsResponse</a>

Methods:

- <code title="get /v1/tasks">client.Tasks.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskListParams">TaskListParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskListResponse">TaskListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tasks/{task_id}">client.Tasks.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskGetResponse">TaskGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tasks/{task_id}/results">client.Tasks.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskService.Results">Results</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#TaskResultsResponse">TaskResultsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchGetResponse">BatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchProgressResponse">BatchProgressResponse</a>

Methods:

- <code title="get /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/batches/{batch_id}">client.Batches.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchGetResponse">BatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches/{batch_id}/progress">client.Batches.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchService.Progress">Progress</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#BatchProgressResponse">BatchProgressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# DomainKnowledge

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#DomainKnowledgeGetDriverResponse">DomainKnowledgeGetDriverResponse</a>

Methods:

- <code title="get /v1/domain-knowledge/driver">client.DomainKnowledge.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#DomainKnowledgeService.GetDriver">GetDriver</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#DomainKnowledgeGetDriverParams">DomainKnowledgeGetDriverParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#DomainKnowledgeGetDriverResponse">DomainKnowledgeGetDriverResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Media

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaRunResponse">MediaRunResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaRunAsyncResponse">MediaRunAsyncResponse</a>

Methods:

- <code title="post /v1/media">client.Media.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaRunParams">MediaRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaRunResponse">MediaRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/media/async">client.Media.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaService.RunAsync">RunAsync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaRunAsyncParams">MediaRunAsyncParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#MediaRunAsyncResponse">MediaRunAsyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Serp

Response Types:

- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunResponse">SerpRunResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunAsyncResponse">SerpRunAsyncResponse</a>
- <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunBatchResponse">SerpRunBatchResponse</a>

Methods:

- <code title="post /v1/serp">client.Serp.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunParams">SerpRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunResponse">SerpRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/serp/async">client.Serp.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpService.RunAsync">RunAsync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunAsyncParams">SerpRunAsyncParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunAsyncResponse">SerpRunAsyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/serp/batch">client.Serp.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpService.RunBatch">RunBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunBatchParams">SerpRunBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go">githubcomnimblewaynimblego</a>.<a href="https://pkg.go.dev/github.com/Nimbleway/nimble-go#SerpRunBatchResponse">SerpRunBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
