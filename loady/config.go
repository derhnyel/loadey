// Config is a configuration.
// type Config struct {
// 	Proxy    Proxy     `json:"proxy"`
// 	Backends []Backend `json:"backends"`
// }

// // Proxy is a reverse proxy, and means load balancer.
// type Proxy struct {
// 	Port string `json:"port"`
// }

// Backend is servers which load balancer is transferred.
// type Resource struct {
// 	URL    string `json:"url"`
// 	IsDead bool
// 	mu     *sync.RWMutex
// }