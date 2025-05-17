To handle 10,000 concurrent gRPC requests per second across multiple data centers, I would:

    Horizontal Scaling:
    Deploy multiple instances of the service behind a load balancer in each data center to distribute the load evenly.

    gRPC Load Balancing:
    Use a gRPC-aware load balancer (like Envoy or built-in gRPC client-side load balancing) to efficiently route requests to healthy service instances.

    Report Storage Persistence:
    Replace the in-memory map with a distributed database or key-value store (like Redis, Cassandra, or a cloud-managed DB) to persist reports reliably and allow access from any instance or data center.

    Reliability and Fault Tolerance:
    Use health checks and retries in the client and load balancer. Implement circuit breakers and graceful shutdowns to handle failures smoothly.

    Deployment Strategy:
    Use container orchestration (Kubernetes) to manage scaling, rolling updates, and failover across multiple data centers.

This approach ensures the system can handle high concurrency, provides data persistence, and is reliable and scalable.
