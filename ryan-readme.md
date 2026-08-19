- not done Go before so trying to keep it inline with the Online (go.dev) documentation/examples
- there's probably much nicer ways to doing the things I've done

# Online
- UI: https://gymfish.ryansouthgate.com/
- API:
      - (GET) GetProducts: https://gymfish.ryansouthgate.com/api/products/1 (id=404 - returns error)
      - (POST) CreateOrder: https://gymfish.ryansouthgate.com/api/orders
- Bruno Request Suite committed to folder `bruno` (Bruno > Postman)

# Pacakges installed
- "go.uber.org/dig@v1" // Just DI - nothing else
- "gin" // Seems to be a popular API framework
- "decimal" // To handle money numbers
- "testify" // To test

# Docker
- run local: `docker compose up`
- used a simple `nginx.conf` I use for other self-hosted apps (with CDN cache headers to cloudflare)
- deployed on a "server" in my loft: https://ryansouthgate.com/my-old-laptop-is-my-new-web-server/
- Using CloudFlared (by Cloudflare) for HTTPS certs and secure tunnel into my home server (without exposing home IP address)

# Testing
- would really do testing at all service and repository layers, and HTTP handlers for different HTTP Statuses
