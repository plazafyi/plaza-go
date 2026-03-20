# githubcomplazafyiplazago

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeoJsonFeature">GeoJsonFeature</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeoJsonGeometry">GeoJsonGeometry</a>

# Elements

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#BatchRequestParam">BatchRequestParam</a>

Methods:

- <code title="get /api/v1/features/{type}/{id}">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, type\_ <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeoJsonFeature">GeoJsonFeature</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/features/batch">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementBatchParams">ElementBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/features/lookup">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeoJsonFeature">GeoJsonFeature</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/features/nearby">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.Nearby">Nearby</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementNearbyParams">ElementNearbyParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/features/nearby">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.NearbyPost">NearbyPost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementNearbyPostParams">ElementNearbyPostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/features">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementQueryParams">ElementQueryParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/features">client.Elements.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementService.QueryPost">QueryPost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElementQueryPostParams">ElementQueryPostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Datasets

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Dataset">Dataset</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetList">DatasetList</a>

Methods:

- <code title="post /api/v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetNewParams">DatasetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/datasets/{id}">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetList">DatasetList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/datasets/{id}">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v1/datasets/{id}/features">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.Features">Features</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetFeaturesParams">DatasetFeaturesParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Geocode

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#AutocompleteResult">AutocompleteResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeResult">GeocodeResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodingFeature">GeocodingFeature</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ReverseGeocodeResult">ReverseGeocodeResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeBatchResponse">GeocodeBatchResponse</a>

Methods:

- <code title="get /api/v1/geocode/autocomplete">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Autocomplete">Autocomplete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeAutocompleteParams">GeocodeAutocompleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#AutocompleteResult">AutocompleteResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode/autocomplete">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.AutocompletePost">AutocompletePost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeAutocompletePostParams">GeocodeAutocompletePostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#AutocompleteResult">AutocompleteResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode/batch">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeBatchParams">GeocodeBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeBatchResponse">GeocodeBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/geocode">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Forward">Forward</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeForwardParams">GeocodeForwardParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeResult">GeocodeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.ForwardPost">ForwardPost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeForwardPostParams">GeocodeForwardPostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeResult">GeocodeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/geocode/reverse">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Reverse">Reverse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeReverseParams">GeocodeReverseParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ReverseGeocodeResult">ReverseGeocodeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode/reverse">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.ReversePost">ReversePost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeReversePostParams">GeocodeReversePostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ReverseGeocodeResult">ReverseGeocodeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Methods:

- <code title="get /api/v1/search">client.Search.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SearchService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SearchQueryParams">SearchQueryParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/search">client.Search.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SearchService.QueryPost">QueryPost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SearchQueryPostParams">SearchQueryPostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Routing

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MatrixRequestParam">MatrixRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RouteRequestParam">RouteRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MatrixResult">MatrixResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#NearestResult">NearestResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RouteResult">RouteResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochroneResponse">RoutingIsochroneResponse</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochronePostResponse">RoutingIsochronePostResponse</a>

Methods:

- <code title="get /api/v1/isochrone">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Isochrone">Isochrone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochroneParams">RoutingIsochroneParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochroneResponse">RoutingIsochroneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/isochrone">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.IsochronePost">IsochronePost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochronePostParams">RoutingIsochronePostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochronePostResponse">RoutingIsochronePostResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/matrix">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Matrix">Matrix</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingMatrixParams">RoutingMatrixParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MatrixResult">MatrixResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/nearest">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Nearest">Nearest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingNearestParams">RoutingNearestParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#NearestResult">NearestResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/nearest">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.NearestPost">NearestPost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingNearestPostParams">RoutingNearestPostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#NearestResult">NearestResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/route">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Route">Route</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingRouteParams">RoutingRouteParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RouteResult">RouteResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Elevation

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationProfileRequestParam">ElevationProfileRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationBatchResult">ElevationBatchResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupResult">ElevationLookupResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationProfileResult">ElevationProfileResult</a>

Methods:

- <code title="post /api/v1/elevation/batch">client.Elevation.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationBatchParams">ElevationBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationBatchResult">ElevationBatchResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/elevation">client.Elevation.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupParams">ElevationLookupParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupResult">ElevationLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/elevation">client.Elevation.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationService.LookupPost">LookupPost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupPostParams">ElevationLookupPostParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupResult">ElevationLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/elevation/profile">client.Elevation.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationService.Profile">Profile</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationProfileParams">ElevationProfileParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationProfileResult">ElevationProfileResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# MapMatch

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MapMatchRequestParam">MapMatchRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MapMatchResult">MapMatchResult</a>

Methods:

- <code title="post /api/v1/map-match">client.MapMatch.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MapMatchService.Match">Match</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MapMatchMatchParams">MapMatchMatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MapMatchResult">MapMatchResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Optimize

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeRequestParam">OptimizeRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeCompletedResult">OptimizeCompletedResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeJobStatus">OptimizeJobStatus</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeProcessingResult">OptimizeProcessingResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeResult">OptimizeResult</a>

Methods:

- <code title="post /api/v1/optimize">client.Optimize.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeNewParams">OptimizeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeResult">OptimizeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/optimize/{job_id}">client.Optimize.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OptimizeJobStatus">OptimizeJobStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Query

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#OverpassQueryParam">OverpassQueryParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SparqlQueryParam">SparqlQueryParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SparqlResult">SparqlResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryExecuteResponse">QueryExecuteResponse</a>

Methods:

- <code title="post /api/v1/query">client.Query.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryExecuteParams">QueryExecuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryExecuteResponse">QueryExecuteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/overpass">client.Query.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryService.Overpass">Overpass</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryOverpassParams">QueryOverpassParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/sparql">client.Query.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryService.Sparql">Sparql</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QuerySparqlParams">QuerySparqlParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SparqlResult">SparqlResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tiles

Methods:

- <code title="get /api/v1/tiles/{z}/{x}/{y}">client.Tiles.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#TileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, z <a href="https://pkg.go.dev/builtin#int64">int64</a>, x <a href="https://pkg.go.dev/builtin#int64">int64</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
