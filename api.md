# githubcomplazafyiplazago

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeometryUnionParam">GeometryUnionParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#LineStringGeometryParam">LineStringGeometryParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MultiLineStringGeometryParam">MultiLineStringGeometryParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MultiPointGeometryParam">MultiPointGeometryParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MultiPolygonGeometryParam">MultiPolygonGeometryParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#PointGeometryParam">PointGeometryParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#PolygonGeometryParam">PolygonGeometryParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeoJsonFeature">GeoJsonFeature</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Geometry">Geometry</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#LineStringGeometry">LineStringGeometry</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MultiLineStringGeometry">MultiLineStringGeometry</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MultiPointGeometry">MultiPointGeometry</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MultiPolygonGeometry">MultiPolygonGeometry</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#PointGeometry">PointGeometry</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#PolygonGeometry">PolygonGeometry</a>

# Features

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#BatchRequestParam">BatchRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SpatialPredicateParam">SpatialPredicateParam</a>

Methods:

- <code title="get /api/v1/features/{type}/{id}">client.Features.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, type\_ <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeoJsonFeature">GeoJsonFeature</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/features/batch">client.Features.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureBatchParams">FeatureBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/features">client.Features.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureQueryParams">FeatureQueryParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Datasets

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Dataset">Dataset</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetList">DatasetList</a>

Methods:

- <code title="post /api/v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetNewParams">DatasetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/datasets/{id}">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#Dataset">Dataset</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/datasets">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetListParams">DatasetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetList">DatasetList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/datasets/{id}">client.Datasets.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#DatasetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Geocode

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#AutocompleteRequestParam">AutocompleteRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeForwardRequestParam">GeocodeForwardRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeReverseRequestParam">GeocodeReverseRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#AutocompleteResult">AutocompleteResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeResult">GeocodeResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodingFeature">GeocodingFeature</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ReverseGeocodeResult">ReverseGeocodeResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeBatchResponse">GeocodeBatchResponse</a>

Methods:

- <code title="post /api/v1/geocode/autocomplete">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Autocomplete">Autocomplete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeAutocompleteParams">GeocodeAutocompleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#AutocompleteResult">AutocompleteResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode/batch">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeBatchParams">GeocodeBatchParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeBatchResponse">GeocodeBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Forward">Forward</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeForwardParams">GeocodeForwardParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeResult">GeocodeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/geocode/reverse">client.Geocode.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeService.Reverse">Reverse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#GeocodeReverseParams">GeocodeReverseParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ReverseGeocodeResult">ReverseGeocodeResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Methods:

- <code title="post /api/v1/search">client.Search.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SearchService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#SearchQueryParams">SearchQueryParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Routing

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#IsochroneRequestParam">IsochroneRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MatrixRequestParam">MatrixRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#NearestRequestParam">NearestRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RouteRequestParam">RouteRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MatrixResult">MatrixResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#NearestResult">NearestResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RouteResult">RouteResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochroneResponse">RoutingIsochroneResponse</a>

Methods:

- <code title="post /api/v1/isochrone">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Isochrone">Isochrone</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochroneParams">RoutingIsochroneParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingIsochroneResponse">RoutingIsochroneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/matrix">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Matrix">Matrix</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingMatrixParams">RoutingMatrixParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#MatrixResult">MatrixResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/nearest">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Nearest">Nearest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingNearestParams">RoutingNearestParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#NearestResult">NearestResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/route">client.Routing.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingService.Route">Route</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RoutingRouteParams">RoutingRouteParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#RouteResult">RouteResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Elevation

Params Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupRequestParam">ElevationLookupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationProfileRequestParam">ElevationProfileRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupResult">ElevationLookupResult</a>
- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationProfileResult">ElevationProfileResult</a>

Methods:

- <code title="post /api/v1/elevation">client.Elevation.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupParams">ElevationLookupParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#ElevationLookupResult">ElevationLookupResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
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

- <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#PlazaqlQueryParam">PlazaqlQueryParam</a>

Methods:

- <code title="post /api/v1/query">client.Query.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#QueryExecuteParams">QueryExecuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go">githubcomplazafyiplazago</a>.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#FeatureCollection">FeatureCollection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tiles

Methods:

- <code title="get /api/v1/tiles/{z}/{x}/{y}">client.Tiles.<a href="https://pkg.go.dev/github.com/plazafyi/plaza-go#TileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, z <a href="https://pkg.go.dev/builtin#int64">int64</a>, x <a href="https://pkg.go.dev/builtin#int64">int64</a>, y <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
